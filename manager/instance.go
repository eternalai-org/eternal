package manager

import (
	"bytes"
	"encoding/json"
	"eternal-infer-worker/libs/eaimodel"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type ModelInstance struct {
	ModelInfo eaimodel.ModelInfoContract
	ModelPath string

	Port string

	actionLock sync.Mutex

	containerID         string
	verifierContainerID string
	ResultDir           string
	VerifyDir           string

	Ready         bool
	VerifierReady bool

	Loaded         bool
	VerifierLoaded bool
	DisableGPU     bool

	TrainingRequest *TrainingRequest
}

// func setupCondaEnv(name, envFile string) error {
// 	out, err := exec.Command("bash", "./setup_conda.sh", name, envFile).Output()
// 	if err != nil {
// 		log.Println(string(out))
// 		return err
// 	}

// 	log.Println(string(out))
// 	return nil
// }

// func (m *ModelInstance) Setup() error {
// 	if strings.HasPrefix(strings.ToLower(m.ModelInfo.Metadata.ModelURL), "ipfs://") {
// 		// get cid at last part
// 		cid := strings.Replace(m.ModelInfo.Metadata.ModelURL, "ipfs://", "", 1)
// 		filePath, err := lighthouse.DownloadChunkedData(cid, m.ModelPath)
// 		if err != nil {
// 			log.Println("Download file with IPFS gateway got error", err)
// 			return err
// 		}

// 		// unzip the file
// 		err = file.UnzipFile(filePath, m.ModelPath)
// 		if err != nil {
// 			return err
// 		}
// 	} else if strings.HasPrefix(strings.ToLower(m.ModelInfo.Metadata.ModelURL), "https://") {

// 	}

// 	return setupCondaEnv(m.ModelInfo.ModelAddr, m.ModelPath+"/environment.yml")
// }

// func (m *ModelInstance) Start() error {

// 	args := fmt.Sprintf("run -n %v --no-capture-output python %v --port %v ", m.ModelInfo.ModelAddr, m.ModelPath+"/server.py", m.Port)
// 	cmd := exec.Command("conda", strings.Split(args, " ")...)

// 	stderr, err := cmd.StderrPipe()
// 	if err != nil {
// 		return err
// 	}
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		return err
// 	}
// 	cmd.Start()
// 	scanner := bufio.NewScanner(stderr)
// 	scanner.Split(bufio.ScanLines)

// 	for scanner.Scan() {
// 		m := scanner.Text()
// 		log.Println("err", m)
// 		// jobErrLog += fmt.Sprintln(m)
// 	}

// 	scanner2 := bufio.NewScanner(stdout)
// 	scanner2.Split(bufio.ScanLines)
// 	for scanner2.Scan() {
// 		m := scanner2.Text()
// 		log.Println("log", m)
// 		// jobLog += fmt.Sprintln(m)
// 	}
// 	cmd.Wait()

// 	return nil
// }

type InferRequest struct {
	Input struct {
		OutputPath string `json:"output_path"`
		Prompt     string `json:"prompt"`
		Seed       uint64 `json:"seed"`
	} `json:"input"`
}

// type InferResponse struct {
// 	OutputPath string    `json:"output_path"`
// 	Embeddings []float64 `json:"embeddings,omitempty"`
// 	Time       float64   `json:"time,omitempty"`
// }

type InferResponse struct {
	Output struct {
		OutputPath  string `json:"output_path"`
		MagicPrompt string `json:"magic_prompt"`
	} `json:"output"`
}

func (m *ModelInstance) Infer(prompt, outputPath string, seed uint64) (string, error) {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	//call the model server
	infer := InferRequest{}
	infer.Input.OutputPath = outputPath
	infer.Input.Prompt = prompt
	infer.Input.Seed = seed

	url := fmt.Sprintf("http://0.0.0.0:%v/predictions", m.Port)
	//url := "http://172.168.20.114:10757/predictions"

	inferJSON, err := json.Marshal(infer)
	if err != nil {
		return "", err
	}

	log.Println("infer", url, string(inferJSON))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inferJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var inferResp InferResponse
	err = json.Unmarshal(body, &inferResp)
	if err != nil {
		return "", err
	}

	log.Println("inferResp", string(body))

	return inferResp.Output.OutputPath, nil
}

// TODO @liam change later wait for @james
type VerifyRequest struct {
	OrgResult  string `json:"org_result"`
	VerfResult string `json:"verf_result"`
}

func (m *ModelInstance) VerifyResult(orgResult, verfResult string) (string, error) {

	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	//call the model server
	infer := VerifyRequest{
		OrgResult:  orgResult,
		VerfResult: verfResult,
	}

	url := fmt.Sprintf("http://0.0.0.0:%v/predictions", m.Port)

	inferJSON, err := json.Marshal(infer)
	if err != nil {
		return "", err
	}

	log.Println("infer", url, string(inferJSON))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inferJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var inferResp InferResponse
	err = json.Unmarshal(body, &inferResp)
	if err != nil {
		return "", err
	}

	return inferResp.Output.OutputPath, nil
}

func (m *ModelInstance) GetExt() string {
	ext := "png"
	switch m.ModelInfo.Metadata.ModelType {
	case eaimodel.ModelTypeImage:
		ext = "png"
	case eaimodel.ModelTypeText:
		ext = "txt"
	}
	return ext
}
