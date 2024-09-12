package main

import (
	"encoding/json"
	"eternal-infer-worker/apis"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs/dockercmd"
	"eternal-infer-worker/libs/file"
	"eternal-infer-worker/libs/github"
	"eternal-infer-worker/libs/logger"
	_ "eternal-infer-worker/libs/logger"
	"eternal-infer-worker/manager"
	watcher "eternal-infer-worker/task-watcher"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"eternal-infer-worker/tui"

	_ "net/http/pprof"

	tea "github.com/charmbracelet/bubbletea"
)

var VersionTag string

func main() {
	defer func() {
		if r := recover(); r != nil {
			if tui.UI != nil {
				tui.UI.UpdateSectionText(tui.UIMessageData{
					Section: tui.UISectionStatusText,
					Color:   "danger",
					Text:    "Panic attack! 💀 ",
				})
			}
			log.Println("Panic attack", r)
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	logger.DefaultLogger.SetTermPrinter(func(text string) {
		fmt.Println(text)
	})

	var err error
	cfg, cmd, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		panic(err)
	}
	VersionTag = cfg.Version

	releaseInfo, err := github.GetLatestRelease()
	if err != nil {
		fmt.Println("Error getting latest release info: ", err)
	} else {
		if releaseInfo.TagName != VersionTag {
			fmt.Println("New version available: ", releaseInfo.TagName)

			fmt.Println("Release notes: ", releaseInfo.Body)
			willUpdate := false
			if cfg.DisableUpdateOnStart {
				fmt.Println("Update on start is disabled")
				fmt.Println("Please update the program manually (this message will disappear in 5 seconds)")
				time.Sleep(5 * time.Second)
			} else {
				willUpdate = true
			}
			if willUpdate {
				fmt.Println("Removing old binary...")
				err = file.RemoveFile("eternal")
				if err != nil {
					fmt.Println("Error removing old binary: ", err)
				}
				fmt.Println("Downloading latest release...")
				err = github.DownloadLatestRelease("eternal")
				if err != nil {
					fmt.Println("Error downloading latest release: ", err)
					os.Exit(1)
				} else {
					fmt.Println("Downloaded latest release")
					fmt.Println("Please restart the program")
					os.Exit(0)
				}
			}
		} else {
			fmt.Println("You are using the latest version")
		}
	}
	modelManager := manager.NewModelManager(cfg.ModelsDir, cfg.RPC, cfg.NodeMode, cfg.WorkerHub, cfg.DisableGPU)

	newTaskWatcher, err := watcher.NewTaskWatcher(watcher.NetworkConfig{
		RPC: cfg.RPC,
		// WS:  *ws,
	}, VersionTag, cfg.WorkerHub, cfg.Account, cfg.ModelsDir, cfg.LighthouseAPI, cfg.NodeMode, 1, 1, modelManager, nil)
	if err != nil {
		panic(err)
	}

	if cmd != nil {
		// fmt.Println("Command: ", cmd)
		switch cmd.Cmd {
		case "wallet":
			subcmd := cmd.Args[0]
			switch subcmd {
			case "help":
				fmt.Println("wallet available subcommands: ")
				fmt.Println("  help (this message)")
				fmt.Println("  info (show wallet info)")
				fmt.Println("  claim-unstake (claim unstake)")
				fmt.Println("  claim-reward (claim mining reward)")
			case "info":
				info, err := newTaskWatcher.GetWorkerInfo()
				if err != nil {
					panic(err)
				}
				infoBytes, err := json.MarshalIndent(info, "", "  ")
				if err != nil {
					panic(err)
				}
				fmt.Println("Wallet info: ")
				fmt.Println(string(infoBytes))

			case "claim-unstake":
				err := newTaskWatcher.ReclaimStake()
				if err != nil {
					panic(err)
				}

			case "claim-reward":
				err := newTaskWatcher.ClaimMiningReward()
				if err != nil {
					panic(err)
				}
			}
		}
		os.Exit(0)
	} else {

		logger.DefaultLogger.SetTermPrinter(func(text string) {
			fmt.Println(text)
		})
		err = checkRequirement()
		if err != nil {
			panic(err)
		}

		stopChn := make(chan struct{}, 1)
		if !cfg.SilentMode {
			ui := tui.InitialModel(VersionTag, cfg.NodeMode, stopChn, newTaskWatcher, modelManager)
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

			shutdownEmitted := false
			go func() {
				for {
					select {
					case <-stopChn:
						shutdownEmitted = true
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
					}
				}
			}()

			go func() {
				for {
					select {
					case <-stopChn:
						if shutdownEmitted {
							tui.UI.UpdateSectionText(tui.UIMessageData{
								Section: tui.UISectionStatusText,
								Color:   "danger",
								Text:    "Force shutting down..."})
							time.Sleep(1 * time.Second)
							os.Exit(0)
						}
					}
				}
			}()
		}
		// go modelManager.WatchAndPreloadModels()

		go newTaskWatcher.Start()

		go func() {
			err = apis.InitRouter(cfg.Port, newTaskWatcher, VersionTag).StartRouter()
			if err != nil {
				panic(err)
			}
		}()
		select {}
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
