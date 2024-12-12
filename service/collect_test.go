package service

import "testing"

func TestCreateCollect(t *testing.T) {
	tests := []struct {
		name    string
		args    CollectService
		wantErr bool
	}{
		{
			name: "Valid input",
			args: CollectService{
				GoodsNum:      "G123456723",
				GoodsName:     "Test Goods",
				GoodsPrice:    99.99,
				CollectStatus: 1,
				GoodsImg:      "https://example.com/image.jpg",
				Tags:          3,
				UserId:        "U12345",
				GoodsDes:      "This is a test goods description.",
			},
			wantErr: false,
		},
		{
			name: "Missing GoodsNum",
			args: CollectService{
				GoodsName:     "Test Goods",
				GoodsPrice:    99.99,
				CollectStatus: 1,
				GoodsImg:      "https://example.com/image.jpg",
				Tags:          3,
				UserId:        "U12345",
				GoodsDes:      "This is a test goods description.",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateCollect(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCollect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateCollect(t *testing.T) {
	type args struct {
		GoodsNum string
		UserId   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				GoodsNum: "G123456723",
				UserId:   "U12345",
			},
			wantErr: false,
		},
		{
			name: "Missing GoodsNum",
			args: args{
				UserId: "U12345",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateCollect(tt.args.UserId, tt.args.GoodsNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCollect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindCollect(t *testing.T) {
	type args struct {
		UserId  string
		GoodNum string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				UserId:  "U12345",
				GoodNum: "G123456723",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FindCollect(tt.args.UserId, tt.args.GoodNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCollect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFindAllCollect(t *testing.T) {
	type args struct {
		UserId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				UserId: "U12345",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FindALLCollect(tt.args.UserId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCollect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
