package main

import (
	"eternal-infer-worker/apis"
	"eternal-infer-worker/libs/dockercmd"
	"eternal-infer-worker/libs/logger"
	_ "eternal-infer-worker/libs/logger"
	"eternal-infer-worker/manager"
	watcher "eternal-infer-worker/task-watcher"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"eternal-infer-worker/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			tui.UI.UpdateSectionText(tui.UIMessageData{
				Section: tui.UISectionStatusText,
				Color:   "danger",
				Text:    "Panic attack! 💀 ",
			})
			log.Println("Panic attack", r)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	var err error
	err = checkRequirement()
	if err != nil {
		panic(err)
	}
	rpc := flag.String("rpc", "https://eternal-ai3.tc.l2aas.com/rpc", "rpc url of the network")
	// ws := flag.String("ws", "", "ws url of the network")
	taskContract := flag.String("task-contract", "0x5E077E883b9b44CC5213140c87d98DAB35E6390e", "task contract address")
	account := flag.String("account", "", "account private key")
	modelsDir := flag.String("models-dir", "./models", "models dir")
	preloadModelStr := flag.String("model", "", "model")
	lighthouseAPI := flag.String("lighthouse-api", "", "lighthouse api")
	port := flag.Int("port", 5656, "port of the server")
	mode := flag.String("mode", "worker", "mode of the server, worker or verifier")

	flag.Parse()

	_ = lighthouseAPI

	if *account == "" || *lighthouseAPI == "" || *port == 0 || *preloadModelStr == "" {
		flag.PrintDefaults()
		return
	}

	modelManager := manager.NewModelManager(*modelsDir, *rpc, *mode)

	newTaskWatcher, err := watcher.NewTaskWatcher(watcher.NetworkConfig{
		RPC: *rpc,
		// WS:  *ws,
	}, *taskContract, *account, *modelsDir, *lighthouseAPI, *mode, 1, 1, modelManager, nil)
	if err != nil {
		panic(err)
	}

	stopChn := make(chan struct{}, 1)

	ui := tui.InitialModel("0.5.0", *mode, stopChn, newTaskWatcher)
	tui.UI = &ui

	logger.DefaultLogger.SetTermPrinter(tui.UI.Print)
	go func() {
		p := tea.NewProgram(ui, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}()
	ui.UpdateSectionText(tui.UIMessageData{
		Section: tui.UISectionStatusText,
		Color:   "waiting",
		Text:    "Starting server...",
	})
	time.Sleep(1 * time.Second)

	err = modelManager.PreloadModels([]string{*preloadModelStr})
	if err != nil {
		panic(err)
	}

	go newTaskWatcher.Start()

	go func() {
		err = apis.InitRouter(*port, newTaskWatcher).StartRouter()
		if err != nil {
			panic(err)
		}
	}()

	for {
		select {
		case <-stopChn:
			tui.UI.UpdateSectionText(tui.UIMessageData{
				Section: tui.UISectionStatusText,
				Color:   "danger",
				Text:    "Shutting down..."})
			err = modelManager.RemoveAllInstanceDocker()
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			os.Exit(0)
		case <-time.After(1 * time.Second):
		}
	}
}

func checkRequirement() error {
	err := checkDockerExist()
	if err != nil {
		return err
	}

	return nil
}

func checkDockerExist() error {
	ok := dockercmd.CheckDockerExist()
	if !ok {
		return fmt.Errorf("docker not found")
	}
	return nil
}

// func checkCondaExist() error {

// 	cmd := exec.Command("conda", "-V")

// 	out, err := cmd.Output()
// 	if err != nil {
// 		return err
// 	}

// 	versionParts := strings.Split(string(out), " ")
// 	if len(versionParts) < 2 {
// 		return fmt.Errorf("conda not found")
// 	}

// 	versions := strings.Split(versionParts[1], ".")
// 	if len(versions) < 2 {
// 		return fmt.Errorf("conda not found")
// 	}

// 	majorVersion, err := strconv.ParseInt(versions[0], 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	minorVersion, err := strconv.ParseInt(versions[1], 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	if majorVersion < 4 || (majorVersion == 4 && minorVersion < 8) {
// 		return fmt.Errorf("conda version is too low")
// 	}

// 	return nil
// }
