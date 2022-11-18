package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/nasa70/car-auction/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgStartBidding_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStartBidding
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgStartBidding{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgStartBidding{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
