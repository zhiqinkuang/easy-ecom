package service

import "testing"

func TestGetGoodByCategory(t *testing.T) {
	type args struct {
		categoryID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				categoryID: "SfmgAfxHQDxTklNm3Bf4n",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetGoodByCategory(tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoodByCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
