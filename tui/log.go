package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type UILogText string

func waitForActivityLog(sub chan UILogText) tea.Cmd {
	return func() tea.Msg {
		return UILogText(<-sub)
	}
}

// var logLck sync.Mutex

func (m *UIinstance) Print(msg string) {
	// logLck.Lock()
	// defer logLck.Unlock()
	m.sublog <- UILogText(msg)
}
