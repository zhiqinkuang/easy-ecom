package service

import "testing"

func TestGetBanner(t *testing.T) {

	tests := []struct {
		name string

		wantErr bool
	}{
		{
			name:    "Valid input",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetBanner()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBanner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
