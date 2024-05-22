package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UIinstance struct {
	version  string
	nodeMode string

	statusText    string
	statusSpinner spinner.Model

	newVersionSpinner spinner.Model

	loadingSpinner spinner.Model

	loadedModel string

	taskManager  TaskManagerI
	modelManager ModelManagerI

	sub    chan UIMessageData
	sublog chan UILogText

	logLines []string
	stopChn  chan struct{}

	expandLog bool
}

var UI *UIinstance

func InitialModel(version, nodeMode string, stopChn chan struct{}, taskMng TaskManagerI, modelMng ModelManagerI) UIinstance {
	m := UIinstance{
		sub:          make(chan UIMessageData, 1),
		version:      version,
		sublog:       make(chan UILogText),
		logLines:     make([]string, 15),
		nodeMode:     nodeMode,
		stopChn:      stopChn,
		taskManager:  taskMng,
		modelManager: modelMng,
	}
	m.resetSpinner()
	return m
}

func (m UIinstance) Init() tea.Cmd {
	return tea.Batch(
		m.loadingSpinner.Tick,
		m.newVersionSpinner.Tick,
		m.statusSpinner.Tick,
		waitForActivity(m.sub),
		waitForActivityLog(m.sublog),
	)
}

func colorToStyle(color string) lipgloss.Style {
	switch color {
	case "danger":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	case "warning":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	case "warning-bg":
		return lipgloss.NewStyle().Background(lipgloss.Color("214"))
	case "normal":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	case "green":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("34"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	}
}

func (m UIinstance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case UILogText:
		m.logLines = append(m.logLines, string(msg))
		if len(m.logLines) > 15 {
			m.logLines = m.logLines[1:]
		}
		return m, tea.Batch(
			waitForActivityLog(m.sublog),
		)
	case UIMessageData:
		switch msg.Section {
		case UISectionStatusText:
			// m.resetStatusSpinner()
			m.SetStatusText(msg.Text)
			m.setStatusSpinnerStyle(colorToStyle(msg.Color))
			return m, tea.Batch(
				waitForActivity(m.sub),
			)
		}

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			go func() {
				m.stopChn <- struct{}{}
			}()
		case "ctrl+u":
			m.SetStatusText("Unstaking and quitting")
			m.setStatusSpinnerStyle(colorToStyle("danger"))
			go func() {
				err := m.taskManager.UnregisterAndQuit()
				if err != nil {
					m.Print("Error: " + err.Error())
					return
				}
				m.stopChn <- struct{}{}
			}()
		case "ctrl+r":
			go func() {
				err := m.taskManager.ReclaimStake()
				if err != nil {
					m.Print("Error: " + err.Error())
					return
				}
			}()
		case "ctrl+e":
			go func() {
				err := m.taskManager.Restake()
				if err != nil {
					m.Print("Error: " + err.Error())
					return
				}
			}()
		case "e":
			go func() {
				err := m.taskManager.ClaimMiningReward()
				if err != nil {
					m.Print("Error: " + err.Error())
					return
				}
			}()
		case "l":
			m.expandLog = !m.expandLog
		}
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		var cmd2 tea.Cmd
		var cmd3 tea.Cmd
		m.loadingSpinner, cmd3 = m.loadingSpinner.Update(msg)
		m.newVersionSpinner, cmd2 = m.newVersionSpinner.Update(msg)
		m.statusSpinner, cmd = m.statusSpinner.Update(msg)
		return m, tea.Batch(cmd,
			cmd2,
			cmd3,
		)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m UIinstance) View() string {
	var s string
	s = fmt.Sprintf("[Eternal] [%v] [%v]\n\n", m.version, m.nodeMode)
	// If the spinner is showing, render it and bail.

	if m.expandLog {
		s += "Logs: ('l' to collapse)------------------ \n\n"
		for _, l := range m.logLines {
			s += logStyle(l) + "\n"
		}
	} else {
		s += "Logs: ('l' to expand)------------------ \n\n"
		s += logStyle(m.logLines[len(m.logLines)-1]) + "\n"
	}

	s += "--------------------------------------- \n"

	if !m.expandLog {
		//status line
		s += fmt.Sprintf("\n %s Status:%s%s\n", m.statusSpinner.View(), " ", textStyle(m.statusText))

		if m.taskManager != nil {

			tasks := m.taskManager.GetCurrentRunningTasks()
			taskIDs := make([]string, 0, len(tasks))
			for _, v := range tasks {
				taskIDs = append(taskIDs, v.TaskID)
			}

			s += fmt.Sprintf("\n %s Current processing (%v):%s%s\n", ">", len(taskIDs), " ", textStyle(fmt.Sprintf("%v", taskIDs)))

			modelLoadingSpinner := ""
			if m.modelManager.GetStatus() != "loaded" {
				modelLoadingSpinner = m.loadingSpinner.View()
			}
			s += fmt.Sprintf("\n %s Assigned Model:%s%s | ModelManager: %s %s\n", ">", " ", textStyle(m.taskManager.GetAssignedModel()), textStyle(m.modelManager.GetStatus()), modelLoadingSpinner)
			s += fmt.Sprintf("\n %s Stake Status:%s%s (%s EAI)\n", ">", " ", textStyle(m.taskManager.StakeStatus()), textStyle(m.taskManager.GetStakedAmount()))
			if m.taskManager.StakeStatus() == "staked" {
				s += bindingShortCutStyle("   Press 'ctrl+u' to unstake and quit\n")
			}

			s += fmt.Sprintf("\n %s Worker Address:%s%s\n", ">", " ", textStyle(m.taskManager.GetWorkerAddress()))
			s += fmt.Sprintf("\n %s Balance:%s%s EAI", ">", " ", textStyle(m.taskManager.GetWorkerBalance()))
			s += fmt.Sprintf(" %s Session Earning:%s%s EAI (%s tasks)\n", "|", " ", textStyle(fmt.Sprintf("%v", m.taskManager.GetSessionEarning())), textStyle(fmt.Sprintf("%v", m.taskManager.GetProcessedTasks())))

			miningAmount := m.taskManager.GetMiningReward()
			if miningAmount != "0" {
				s += fmt.Sprintf("\n %s Mining reward:%s%s EAI\n", ">", " ", textStyle(miningAmount))
				s += bindingShortCutStyle("   Press 'e' to claim mining reward\n")
			}

			unstakeAmount, unstakeTime := m.taskManager.GetUnstakeInfo()
			if unstakeAmount != "0" {
				s += fmt.Sprintf("\n %s Pending unstake:%s%s EAI available to claim at %s\n", ">", " ", textStyle(unstakeAmount), textStyle(unstakeTime.Format("2006-01-02 15:04:05")))
				s += "   Press 'ctrl+r' to reclaim stake\n"
				s += "   Press 'ctrl+e' to restake and start mining again\n"
			}

			hubInfo, err := m.taskManager.GetHubGlobalInfo()
			if err == nil {
				// s += fmt.Sprintf("\n %s Hub Global Info:%s%s\n", ">", " ", textStyle(fmt.Sprintf("%v", hubInfo)))
				s += fmt.Sprintf("\n %s Total models: %s | Total Miners: %s\n", ">", textStyle(fmt.Sprintf("%v", hubInfo.TotalModels)), textStyle(fmt.Sprintf("%v", hubInfo.TotalMiners)))
			}
		}
	}

	s += bindingShortCutStyle("\nPress 'q' or 'ctrl+c' to quit\n")

	if hasNewVer, newVer := m.taskManager.HasNewVersion(); hasNewVer {
		s += "\n "
		s += fmt.Sprintf("%s %s %s", colorToStyle("warning").Render(m.newVersionSpinner.View()), colorToStyle("warning").Render("New version available:"), colorToStyle("warning-bg").Bold(true).Render(newVer))
		s += fmt.Sprintf("\n\n")
		s += colorToStyle("warning-bg").Render(fmt.Sprintf("stop and start the program to update to the new version"))
		// s += colorToStyle("warning-bg").Render(fmt.Sprintf(" %s New version available:%s%s", m.newVersionSpinner.View(), " ", textStyle(newVer)))
		s += "\n"
	}

	return s
}
