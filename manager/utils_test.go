package manager

import "testing"

func Test_getCurrentDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test getCurrentDir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCurrentDir(); got != tt.want {
				t.Errorf("getCurrentDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
