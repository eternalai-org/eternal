package dockercmd

import (
	"testing"
)

func TestListContainers(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test ListContainers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListContainers()
		})
	}
}

func TestListImages(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test ListImages",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListImages()
		})
	}
}

func TestCreateAndStartContainer(t *testing.T) {
	type args struct {
		imageTag      string
		containerName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test CreateAndStartContainer",
			args: args{
				imageTag:      "nginx",
				containerName: "nginx-test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			modelMountDir := "./infer-result"
			if _, err := CreateAndStartContainer(tt.args.imageTag, tt.args.containerName, "5001", modelMountDir, false); (err != nil) != tt.wantErr {
				t.Errorf("CreateContainer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadLocalImageWithCustomName(t *testing.T) {
	type args struct {
		imagePath string
		imageName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test LoadLocalImageWithCustomName",
			args: args{
				imagePath: "./nginx2",
				imageName: "nginx2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadLocalImageWithCustomName(tt.args.imagePath, tt.args.imageName); (err != nil) != tt.wantErr {
				t.Errorf("LoadLocalImageWithCustomName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
