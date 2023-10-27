package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ashishkhuraishy/test-net/app"
	testtypes "github.com/ashishkhuraishy/test-net/x/testnet/types"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	txtype "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/google/uuid"
)

func main() {
	// addr, err := CreateAccount()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(addr)
	addr := "cosmos1308f48v2806cc2da9qmjfuz0npx6gwzs0f7dgy"
	// addr := "cosmos1jqynxg6uyllgqch88wru9ec7nzaam7ya99r0ar"
	CreateTx(addr)

}

/* Create a GRPC client connnection without TLS. */
func CreateGrpcConn() (*grpc.ClientConn, error) {
	var grpcConn *grpc.ClientConn
	var err error
	//  Specify the fully qualified hostname and port of the validator node.
	grpcHost := "127.0.0.1:9090"
	grpcConn, err = grpc.Dial(
		grpcHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return grpcConn, err
}

func NewKeyring() (keyring.Keyring, error) {
	var kr keyring.Keyring
	keyringPassword := ""
	runtimeHomeDir := "~/.keyring-new"

	reader := strings.NewReader("")
	reader.Reset(keyringPassword + "\n")

	encodingConfig := app.MakeEncodingConfig()
	cdc := codec.NewProtoCodec(encodingConfig.InterfaceRegistry)
	kr, err := keyring.New(
		sdktypes.KeyringServiceName(),
		keyring.BackendTest,
		runtimeHomeDir,
		reader,
		cdc,
	)

	return kr, err
}

func CreateAccount() (string, error) {
	keyringPassword := "password"
	kr, err := NewKeyring()
	if err != nil {
		return "failed creating keyring", err
	}

	uid := genUID()
	mnemonicInfo, _, err := kr.NewMnemonic(uid, keyring.English, "", keyringPassword, hd.Secp256k1)
	if err != nil {
		return "", err
	}

	newAddr, err := mnemonicInfo.GetAddress()
	return newAddr.String(), err
}

func genUID() string {
	uid, _ := uuid.NewUUID()
	val := uid.String()

	fmt.Println(val)
	return val
}

func CreateTx(addr string) {
	kr, err := NewKeyring()
	if err != nil {
		log.Fatalln(err)
	}

	accAddr, err := sdktypes.AccAddressFromBech32(addr)
	if err != nil {
		log.Fatalln(err)
	}
	encodingConfig := app.MakeEncodingConfig()
	txConfig := encodingConfig.TxConfig
	txBuilder := txConfig.NewTxBuilder()

	msg := testtypes.NewMsgCreateItem(addr, "New Question", []string{"option a", "option b"})
	if err := txBuilder.SetMsgs(msg); err != nil {
		log.Fatalln(err)
	}

	txBuilder.SetGasLimit(100000)

	// sign the transaction
	keyInfo, err := kr.KeyByAddress(accAddr)
	if err != nil {
		log.Fatalln(err)
	}

	exportedPvtKey, err := kr.ExportPrivKeyArmor(keyInfo.Name, "")
	if err != nil {
		log.Fatalln(err)
	}

	pvtKey, algo, err := crypto.UnarmorDecryptPrivKey(exportedPvtKey, "")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(algo)

	account, err := GetAccount(addr)
	if err != nil {
		log.Fatalln(err)
	}

	accSeq := account.Sequence
	accountNum := account.AccountNumber

	accSeqList := []uint64{accSeq}
	accNumberList := []uint64{accountNum}
	pvtKeyList := []cryptotypes.PrivKey{pvtKey}

	// First round: Gathering signer infos
	var signatureLists []signing.SignatureV2
	for idx, privKey := range pvtKeyList {
		sig := signing.SignatureV2{
			PubKey: privKey.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  encodingConfig.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqList[idx],
		}

		signatureLists = append(signatureLists, sig)
	}

	if err := txBuilder.SetSignatures(signatureLists...); err != nil {
		log.Fatalln(err)
	}

	// Second round signing: Signed by each signer
	var signingList []signing.SignatureV2
	for idx, privKey := range pvtKeyList {
		signerData := authsigning.SignerData{
			AccountNumber: accNumberList[idx],
			Sequence:      accSeqList[idx],
			ChainID:       "testnet",
		}
		signMode := encodingConfig.TxConfig.SignModeHandler().DefaultMode()
		sig, err := tx.SignWithPrivKey(signMode, signerData, txBuilder, privKey, encodingConfig.TxConfig, accSeqList[idx])
		if err != nil {
			log.Fatalln(err)
		}

		signingList = append(signingList, sig)
	}
	if err != nil {
		log.Fatalln(err)
	}

	if err := txBuilder.SetSignatures(signingList...); err != nil {
		log.Fatalln(err)
	}

	txBytes, err := encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		log.Fatalln(err)
	}

	// txData := hex.EncodeToString(txBytes)
	// fmt.Println(base64.NewEncoding(txData))
	// fmt.Println(txBytes)
	hash, height, err := BroadcastSignedTxn(txBytes)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(hash, height)
}

func GetAccount(accountAddress string) (authtypes.BaseAccount, error) {
	var baseAccount authtypes.BaseAccount
	addr, err := sdktypes.AccAddressFromBech32(accountAddress)
	if err != nil {
		return baseAccount, err
	}
	grpcConn, err := CreateGrpcConn()
	if err != nil {
		return baseAccount, err
	}
	defer grpcConn.Close()
	accountClient := authtypes.NewQueryClient(grpcConn)
	accountRes, err := accountClient.Account(
		context.Background(),
		&types.QueryAccountRequest{Address: addr.String()},
	)
	if err != nil {
		return baseAccount, err
	}
	accountData := accountRes.GetAccount().Value
	if err := baseAccount.XXX_Unmarshal(accountData); err != nil {
		return baseAccount, err
	}
	return baseAccount, nil
}

func BroadcastSignedTxn(txBytes []byte) (string, int64, error) {
	grpcConn, err := CreateGrpcConn()
	if err != nil {
		return "", int64(0), err
	}
	defer grpcConn.Close()
	serviceClient := txtype.NewServiceClient(grpcConn)
	broadcastTxRes, err := serviceClient.BroadcastTx(
		context.Background(),
		&txtype.BroadcastTxRequest{
			Mode:    txtype.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return "", int64(0), err
	}
	txResponse := broadcastTxRes.GetTxResponse()
	log.Println(txResponse.String())
	if txResponse.Code != 0 {
		return "", int64(0),
			errors.New("TxResponse.Code: " + strconv.FormatUint(uint64(txResponse.Code), 10))
	}

	fmt.Println(txResponse.TxHash)
	return txResponse.TxHash, txResponse.Height, nil
}
