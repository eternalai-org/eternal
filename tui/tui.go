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

	loadedModel string

	taskManager  TaskManagerI
	modelManager ModelManagerI

	sub    chan UIMessageData
	sublog chan UILogText

	logLines []string
	stopChn  chan struct{}
}

var UI *UIinstance

func InitialModel(version, nodeMode string, stopChn chan struct{}, taskMng TaskManagerI, modelMng ModelManagerI) UIinstance {
	m := UIinstance{
		sub:          make(chan UIMessageData, 1),
		version:      version,
		sublog:       make(chan UILogText),
		logLines:     make([]string, 5),
		nodeMode:     nodeMode,
		stopChn:      stopChn,
		taskManager:  taskMng,
		modelManager: modelMng,
	}
	return m
}

func (m UIinstance) Init() tea.Cmd {
	return tea.Batch(
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
		return lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	case "waiting":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	case "normal":
		return lipgloss.NewStyle().Foreground(lipgloss.Color("34"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	}
}

func (m UIinstance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case UILogText:
		m.logLines = append(m.logLines, string(msg))
		if len(m.logLines) > 5 {
			m.logLines = m.logLines[1:]
		}
		return m, tea.Batch(
			waitForActivity(m.sub),
			waitForActivityLog(m.sublog),
		)
	case UIMessageData:
		switch msg.Section {
		case UISectionStatusText:
			m.resetStatusSpinner()
			m.SetStatusText(msg.Text)
			m.setSpinnerStyle(colorToStyle(msg.Color))
			return m, tea.Batch(
				m.statusSpinner.Tick,
				waitForActivity(m.sub),
				waitForActivityLog(m.sublog),
			)
		}

		return m, tea.Batch(
			waitForActivity(m.sub),
			waitForActivityLog(m.sublog),
		)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			go func() {
				m.stopChn <- struct{}{}
			}()
		case "ctrl+s":
			m.SetStatusText("Unstaking and quitting")
			m.setSpinnerStyle(colorToStyle("danger"))
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
		}
		return m, tea.Batch(
			waitForActivity(m.sub),
			waitForActivityLog(m.sublog),
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.statusSpinner, cmd = m.statusSpinner.Update(msg)
		return m, tea.Batch(cmd,
			waitForActivity(m.sub),
			waitForActivityLog(m.sublog),
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

	s += "Logs: ---------------------------------- \n\n"

	for _, l := range m.logLines {
		s += logStyle(l) + "\n"
	}

	s += "--------------------------------------- \n"
	//status line
	s += fmt.Sprintf("\n %s Status:%s%s\n", m.statusSpinner.View(), " ", textStyle(m.statusText))

	if m.modelManager != nil {
		s += fmt.Sprintf("\n %s ModelManager Status:%s%s\n", m.statusSpinner.View(), " ", textStyle(m.modelManager.GetStatus()))
	}

	if m.taskManager != nil {
		unstakeAmount, unstakeTime := m.taskManager.GetUnstakeInfo()
		if unstakeAmount != "0" {
			s += fmt.Sprintf("\n %s Pending unstake:%s%s EAI available to claim at %s\n", ">", " ", textStyle(unstakeAmount), textStyle(unstakeTime.Format("2006-01-02 15:04:05")))
			s += "   Press ctrl+r to reclaim stake\n"
		}
		miningAmount := m.taskManager.GetMiningReward()
		if miningAmount != "0" {
			s += fmt.Sprintf("\n %s Mining reward:%s%s EAI\n", ">", " ", textStyle(miningAmount))
		}
		s += fmt.Sprintf("\n %s Assigned Model:%s%s\n", ">", " ", textStyle(m.taskManager.GetAssignedModel()))
		s += fmt.Sprintf("\n %s Stake Status:%s%s (%s EAI)\n", ">", " ", textStyle(m.taskManager.StakeStatus()), textStyle(m.taskManager.GetStakedAmount()))
		s += fmt.Sprintf("\n %s Worker Address:%s%s\n", ">", " ", textStyle(m.taskManager.GetWorkerAddress()))
		s += fmt.Sprintf("\n %s Balance:%s%s EAI", ">", " ", textStyle(m.taskManager.GetWorkerBalance()))
		s += fmt.Sprintf(" %s Session Earning:%s%s EAI (%s tasks)\n", "|", " ", textStyle(fmt.Sprintf("%v", m.taskManager.GetSessionEarning())), textStyle(fmt.Sprintf("%v", m.taskManager.GetProcessedTasks())))

		tasks := m.taskManager.GetCurrentRunningTasks()
		taskIDs := make([]string, 0, len(tasks))
		for _, v := range tasks {
			taskIDs = append(taskIDs, v.TaskID)
		}

		s += fmt.Sprintf("\n %s Current processing (%v):%s%s\n\n", ">", len(taskIDs), " ", textStyle(fmt.Sprintf("%v", taskIDs)))
	}
	s += "\nPress q or ctrl+c to quit\n"
	s += "Press ctrl+s to unstake and quit\n"

	return s
}
