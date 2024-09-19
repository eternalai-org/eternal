package dockercmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docker/cli/opts"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CheckDockerExist() bool {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false
	}
	defer apiClient.Close()

	p, err := apiClient.Ping(context.Background())
	if err != nil {
		return false
	}
	log.Println(p.APIVersion)
	// should be 1.41
	return true
}

func ListContainers() ([]types.Container, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers, nil
}

func ListImages() ([]types.ImageSummary, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers, nil
}

func GetImageInfo(imageName string) (types.ImageSummary, error) {
	imgs, err := ListImages()
	if err != nil {
		return types.ImageSummary{}, err
	}

	for _, img := range imgs {
		for _, tag := range img.RepoTags {
			if tag == imageName {
				return img, nil
			}
		}
	}

	return types.ImageSummary{}, nil
}

func RenameImage(imageName, newImageName string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ImageTag(context.Background(), imageName, newImageName)
	if err != nil {
		return err
	}

	return nil
}

func RemoveImage(imageName string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	_, err = apiClient.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func LoadLocalImageWithCustomName(imagePath, imageName string) error {
	log.Println("[SetupDocker][LoadLocalImageWithCustomName] - imagePath: ", imagePath, " ,imageName: ", imageName)
	loadedImg, err := loadLocalImage(imagePath)
	if err != nil {
		log.Println("[SetupDocker][LoadLocalImageWithCustomName][Err] - imagePath: ", imagePath, " ,imageName: ", imageName, " ,error:", err)
		return err
	}

	log.Println("[SetupDocker][LoadLocalImageWithCustomName] - imagePath: ", imagePath, " ,imageName: ", imageName, " ,loadedImg: ", loadedImg)
	err = RenameImage(loadedImg, imageName)
	if err != nil {
		log.Println("[SetupDocker][LoadLocalImageWithCustomName][Err] - imagePath: ", imagePath, " ,imageName: ", imageName, " ,error:", err)
		return err
	}

	err = RemoveImage(loadedImg)
	if err != nil {
		log.Println("[SetupDocker][LoadLocalImageWithCustomName][Err] - imagePath: ", imagePath, " ,imageName: ", imageName, " ,error:", err)
		return err
	}
	return nil
}

func loadLocalImage(imagePath string) (string, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	defer apiClient.Close()

	imageFile, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}

	a, err := apiClient.ImageLoad(context.Background(), imageFile, false)
	if err != nil {
		return "", err
	}
	defer a.Body.Close()

	body, err := ioutil.ReadAll(a.Body)
	if err != nil {
		return "", err
	}

	// {"stream":"Loaded image: nginx:latest\n"}

	imageName := ""
	// get the image name from the response
	type ImageLoadResponse struct {
		Stream string `json:"stream"`
	}
	// var imageLoadResponse ImageLoadResponse
	// err = json.Unmarshal(body, &imageLoadResponse)
	// if err != nil {
	log.Println(string(body))
	// 	return "", err
	// }

	// 	{"status":"Loading layer","progressDetail":{"current":8192,"total":8192},"progress":"[==================================================\u003e]  8.192kB/8.192kB","id":"f460e43f2ce4"}
	// {"status":"Loading layer","progressDetail":{"current":8192,"total":8192},"progress":"[==================================================\u003e]  8.192kB/8.192kB","id":"f460e43f2ce4"}
	// {"stream":"Loaded image: prompt_generator:latest\n"}

	parts := strings.Split(string(body), "Loaded image: ")
	if len(parts) < 2 {
		return "", errors.New("Failed to load image")
	}

	part2 := parts[1]
	part2 = strings.TrimSpace(part2)

	imageName = part2[:len(part2)-4]

	log.Println("Loaded image: ", imageName)

	return imageName, nil
}

func StopAndRemoveContainer(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		return err
	}

	err = apiClient.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetContainerByName(containerName string) (*types.Container, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return &container, nil
			}
		}
	}

	return nil, nil
}

func PauseContainer(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ContainerPause(context.Background(), containerID)
	if err != nil {
		if strings.Contains(err.Error(), "already paused") {
			return nil
		}
		return err
	}

	return nil
}

func UnpauseContainer(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ContainerUnpause(context.Background(), containerID)
	if err != nil {
		return err
	}

	return nil
}

func StartContainer(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func RemoveContainer(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer apiClient.Close()

	err = apiClient.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetContainerInfo(containerID string) (*types.ContainerJSON, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	info, err := apiClient.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func WaitForContainerToReady(containerID string) error {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	defer apiClient.Close()
	i := 0

	for {
		i++
		if i > 120 {
			return fmt.Errorf("Timeout waiting for container to be ready")
		}
		info, err := apiClient.ContainerInspect(context.Background(), containerID)
		if err != nil {
			return err
		}

		if info.State.Running {
			return nil
		}

		if info.State.Dead || info.State.Status == "exited" {
			var errInfo = struct {
				Message string `json:"message"`
				Code    int    `json:"code"`
			}{
				Message: info.State.Error,
				Code:    info.State.ExitCode,
			}
			return fmt.Errorf("Container exited with code %d: %s", errInfo.Code, errInfo.Message)
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func CreateAndStartContainer(imageTag string, containerName, containerPort, mountFolder string, disableGPU bool) (*container.ContainerCreateCreatedBody, error) {
	existedContainer, err := GetContainerByName(containerName)
	if err != nil {
		return nil, err
	}
	if existedContainer != nil {
		// err = RemoveContainer(existedContainer.ID)
		// if err != nil {
		// 	return nil, errors.Join(fmt.Errorf("Failed to remove existed container %s", containerName), err)
		// }
		a := container.ContainerCreateCreatedBody{
			ID: existedContainer.ID,
		}

		return &a, nil
	}
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer apiClient.Close()

	port := "5000"
	hostPort := containerPort

	newport, err := nat.NewPort("tcp", port)
	if err != nil {
		log.Println("Unable to create docker port")
		return nil, err
	}
	exposedPorts := map[nat.Port]struct{}{
		newport: struct{}{},
	}

	gpuOpts := opts.GpuOpts{}
	if !disableGPU {
		//FOR debug purpose
		//gpuOpts.Set("all")
	}

	hostConfig := &container.HostConfig{
		Resources: container.Resources{DeviceRequests: gpuOpts.Value()},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: mountFolder,
				Target: "/output",
			},
		},
		PortBindings: nat.PortMap{
			newport: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
	}

	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	resp, err := apiClient.ContainerCreate(context.Background(), &container.Config{
		Image:        imageTag,
		ExposedPorts: exposedPorts,
	}, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		return nil, err
	}

	err = apiClient.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}
