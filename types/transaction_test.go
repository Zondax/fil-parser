package types

import (
	"math/big"
	"testing"
	"time"
)

func TestTransaction_Equal(t1 *testing.T) {
	tests := []struct {
		name string
		a    Transaction
		b    Transaction
		want bool
	}{
		{
			name: "equal txs",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: true,
		},
		{
			name: "Different BlockData",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    500,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "Different Id",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0001",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: true,
		},
		{
			name: "different ParendId",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "01",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: true,
		},
		{
			name: "different Level",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       1,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different TxTimestamp",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Now(),
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different txCid",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "newtxCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different txFrom",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x001",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different txTo",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x002",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different Amount",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(10000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different GasUsed",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     10,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different Status",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "SysOutOfGas",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different TxType",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Propose",
				TxMetadata:  "{}",
			},
			want: false,
		},
		{
			name: "different metadata",
			a: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{}",
			},
			b: Transaction{
				BasicBlockData: BasicBlockData{
					Height:    1000,
					TipsetCid: "test",
					BlockCid:  "test",
				},
				Id:          "0000",
				ParentId:    "",
				Level:       0,
				TxTimestamp: time.Time{},
				TxCid:       "txCid",
				TxFrom:      "0x000",
				TxTo:        "0x001",
				Amount:      big.NewInt(1000),
				GasUsed:     0,
				Status:      "Ok",
				TxType:      "Send",
				TxMetadata:  "{\"params\": {\"Sectors\":[10,12,34]}}",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t1.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
