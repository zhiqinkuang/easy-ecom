package service

import "testing"

//

func TestGetCategoriesByParent(t *testing.T) {
	type args struct {
		Parent string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				Parent: "aUEQ8srl1GPZKz6Z0MBvn",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCategoriesByParent(tt.args.Parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCategoriesByActiveAndLevel(t *testing.T) {
	type args struct {
		Level int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				Level: 1,
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCategoriesByActiveAndLevel(tt.args.Level)
			if (err != nil) != tt.wantErr {
				t.Errorf("Createcart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
