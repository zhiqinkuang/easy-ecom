package service

import "testing"

func TestGetGoodAtr(t *testing.T) {
	type args struct {
		GoodsID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				GoodsID: "aUEQ8srl1GPZKz6Z0MBvn",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetGoodAtr(tt.args.GoodsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
