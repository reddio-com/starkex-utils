package asset

import "testing"

func TestGetAssetInfo(t *testing.T) {
	tests := []struct {
		typ     string
		address string
		want    string
	}{
		{"ETH", "0x8322fff2", "0x8322fff2"},
		{"ERC20", "0x07865c6e87b9f70255377e024ace6630c1eaa37f", "0xf47261b000000000000000000000000007865c6e87b9f70255377e024ace6630c1eaa37f"},
	}
	for _, tt := range tests {
		t.Run(tt.typ, func(t *testing.T) {
			got, err := GetAssetInfo(tt.typ, tt.address)
			if err != nil {
				t.Errorf("GetAssetInfo() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("GetAssetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
