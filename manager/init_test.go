package manager

import "testing"

func Test_checkSetupCondaFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test checkSetupCondaFile",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkSetupCondaFile(); (err != nil) != tt.wantErr {
				t.Errorf("checkSetupCondaFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
