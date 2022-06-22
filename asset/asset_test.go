package asset

import (
	"math/big"
	"testing"

	"github.com/reddio-com/starkex-utils/types"
)

func TestGetAssetInfo(t *testing.T) {
	type args struct {
		TokenType types.TokenType
		Address   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ETH",
			args: args{
				TokenType: types.ETH,
				Address:   "0x4240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
			},
			want: "0x8322fff2",
		},
		{
			name: "ERC20",
			args: args{
				TokenType: types.ERC20,
				Address:   "0x4240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
			},
			want: "0xf47261b00000000000000000000000004240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
		},
		{
			name: "ERC721",
			args: args{
				TokenType: types.ERC721,
				Address:   "0x4240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
			},
			want: "0x025717920000000000000000000000004240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
		},
		{
			name: "ERC721M",
			args: args{
				TokenType: types.ERC721M,
				Address:   "0x4240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
			},
			want: "0xb8b866720000000000000000000000004240e8b8c0b6E6464a13F555F6395BbfE1c4bdf1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAssetInfo(tt.args.TokenType, tt.args.Address); got != tt.want {
				t.Errorf("GetAssetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAssetType(t *testing.T) {
	type args struct {
		TokenType types.TokenType
		Address   string
		Quantum   *big.Int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				TokenType: types.ERC20,
				Address:   "0x20a36B174dfb726A33a8416eD2E4894719236140",
				Quantum:   big.NewInt(6),
			},
			want: "24472b3d6de7eb4126da9effb48951f9f397acd92c9b9b528ec9506cc51fbb0",
		},
		{
			name: "test2",
			args: args{
				TokenType: types.ERC721,
				Address:   "0x490a36B174dfb726A33a8416eD2E489471923640",
				Quantum:   big.NewInt(1),
			},
			want: "5681e5f4257b91099832434be78613fe251d1fc3a57b37f98a7c5954828a8f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAssetType(tt.args.TokenType, tt.args.Address, tt.args.Quantum)
			// decoded, _ := hex.DecodeString(fmt.Sprint(got))
			want := new(big.Int)
			want.SetString(tt.want, 16)
			if got.Cmp(want) != 0 {
				t.Errorf("GetAssetType() = %v, want %v", got, want)
			}
		})
	}

}

func TestGetAssetID(t *testing.T) {
	type args struct {
		TokenType types.TokenType
		TokenId   *big.Int
		Address   string
		Quantum   *big.Int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				TokenType: types.ERC20,
				TokenId:   big.NewInt(1),
				Address:   "0x20a36B174dfb726A33a8416eD2E4894719236140",
				Quantum:   big.NewInt(6),
			},
			want: "24472b3d6de7eb4126da9effb48951f9f397acd92c9b9b528ec9506cc51fbb0",
		},
		{
			name: "test2",
			args: args{
				TokenType: types.ERC721,
				TokenId:   big.NewInt(20),
				Address:   "0x490a36B174dfb726A33a8416eD2E489471923640",
				Quantum:   big.NewInt(1),
			},
			want: "34a3d452eec45ff2491610f4261bcd2aa0305730b1eeb9bec373ed679eedd47",
		},
		{
			name: "test3",
			args: args{
				TokenType: types.ERC721M,
				TokenId:   big.NewInt(2),
				Address:   "0x490a36B174dfb726A33a8416eD2E489471923640",
				Quantum:   big.NewInt(1),
			},
			want: "4001092dca3411adf0e2c242bfe94321deea57ffede4acef78af20c270a5348",
		},
		{
			name: "test3",
			args: args{
				TokenType: types.ERC721M,
				TokenId:   big.NewInt(67),
				Address:   "0x490a36B174dfb726A33a8416eD2E489471923640",
				Quantum:   big.NewInt(1),
			},
			want: "400aa6e72093b2960a575428b4627e64257d24a08b8a4e566790937e9e61d54",
		},
		{
			name: "test3",
			args: args{
				TokenType: types.ERC721M,
				TokenId:   big.NewInt(345738),
				Address:   "0x490a36B174dfb726A33a8416eD2E489471923640",
				Quantum:   big.NewInt(1),
			},
			want: "40051d57bde18dddb7e6e149ecc7f449d88af7b54f73990d3162c2992224248",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAssetID(tt.args.TokenType, tt.args.Address, tt.args.Quantum, tt.args.TokenId)
			// decoded, _ := hex.DecodeString(fmt.Sprint(got))
			want := new(big.Int)
			want.SetString(tt.want, 16)
			if got.Cmp(want) != 0 {
				t.Errorf("GetAssetID() = %v, want %v", got, want)
			}
		})
	}
}
