package service

import (
	"github.com/zhiqinkuang/easy-ecom/repository"
	"testing"
)

func TestGetClientByUID(t *testing.T) {
	tests := []struct {
		name    string
		uid     string
		wantErr bool
	}{
		{
			name:    "Valid UID",
			uid:     "12312312412",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := GetClientByUID(tt.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientByUID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && client == nil {
				t.Errorf("GetClientByUID() returned nil, expected client")
			}
		})
	}
}

func TestGetClientByUserName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "13361643493",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := GetClientByUserName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientByUID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && client == nil {
				t.Errorf("GetClientByUID() returned nil, expected client")
			}
		})
	}
}

func TestCreateClient(t *testing.T) {
	tests := []struct {
		name    string
		client  *repository.ClientLogin
		wantErr bool
	}{
		{
			name: "Create new user",
			client: &repository.ClientLogin{
				Username: "13361643493",
				UID:      "12312312412",
				Status:   0,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateClient(tt.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateCaptcha(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    bool
		wantBase64 bool
	}{
		{
			name:       "Generate Captcha with valid data",
			wantErr:    false,
			wantBase64: true, // We expect valid Base64 string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call GenerateCaptcha and capture its output
			masterBase64, thumbBase64, err := GenerateCaptcha()

			// Check if the error matches the expected error
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if masterBase64 and thumbBase64 are not empty (they should be valid Base64 encoded strings)
			if tt.wantBase64 {
				if masterBase64 == "" {
					t.Errorf("GenerateCaptcha() masterBase64 is empty")
				}
				if thumbBase64 == "" {
					t.Errorf("GenerateCaptcha() thumbBase64 is empty")
				}
			}
		})
	}
}
