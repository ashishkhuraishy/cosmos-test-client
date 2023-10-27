package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateItem = "create_item"

var _ sdk.Msg = &MsgCreateItem{}

func NewMsgCreateItem(creator string, question string, options []string) *MsgCreateItem {
	return &MsgCreateItem{
		Creator:  creator,
		Question: question,
		Options:  options,
	}
}

func (msg *MsgCreateItem) Route() string {
	return RouterKey
}

func (msg *MsgCreateItem) Type() string {
	return TypeMsgCreateItem
}

func (msg *MsgCreateItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
