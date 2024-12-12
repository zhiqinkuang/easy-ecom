package service

import "testing"

// test createCart
func TestCreateCart(t *testing.T) {
	tests := []struct {
		name    string
		args    CartService
		wantErr bool
	}{
		{
			name: "Valid input",
			args: CartService{
				CartID:     "xferwqrsfs222775",
				UserID:     "xvswqrsdrw245",
				ProductID:  "PROD11223343",
				SelectNum:  2,
				Price:      299.99,
				ImageURL:   "https://example.com/product-image.jpg",
				IsActive:   true,
				Properties: `[{"k": "版本", "v": [{"id": "ZbLielyD9aOJPS5XrhcN1", "name": "私服", "price": 0}], "k_id": "EFZiqxaX6y8sDaMdW5v0W", "is_multiple": false}]`,
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateCartItem(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// test updateCart
func TestUpdateCartItem(t *testing.T) {
	type args struct {
		UserID    string
		CartID    string
		SelectNum int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				CartID:    "xferwqrsfs222775",
				UserID:    "xvswqrsdrw245",
				SelectNum: 4,
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateCartItem(tt.args.UserID, tt.args.CartID, tt.args.SelectNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteCartItem(t *testing.T) {
	type args struct {
		UserID string
		CartID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				CartID: "xferwqrsfs222775",
				UserID: "xvswqrsdrw245",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteCart(tt.args.UserID, tt.args.CartID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCartItem(t *testing.T) {
	type args struct {
		UserID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				UserID: "xvswqrsdrw",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCart(tt.args.UserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
