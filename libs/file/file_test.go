package file

import (
	"testing"
)

func TestUnzipFile(t *testing.T) {
	type args struct {
		zipFilePath string
		dst         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unzip file success",
			args: args{
				zipFilePath: "test.txt.zip",
				dst:         "./unzip_output",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnzipFile(tt.args.zipFilePath, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("UnzipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnzipDir(t *testing.T) {
	type args struct {
		zipFilePath string
		dst         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unzip dir success",
			args: args{
				zipFilePath: "dirzip.zip",
				dst:         "./unzip_output",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnzipFile(tt.args.zipFilePath, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("UnzipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMergeFiles(t *testing.T) {
	type args struct {
		files  []string
		output string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "merge files success",
			args: args{
				files:  []string{"2024-05-02.log.partaa", "2024-05-02.log.partab", "2024-05-02.log.partac", "2024-05-02.log.partad", "2024-05-02.log.partae", "2024-05-02.log.partaf"},
				output: "output.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MergeFiles(tt.args.files, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("MergeFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
