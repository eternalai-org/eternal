package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type UIMessageData struct {
	Section UISectionType
	Text    string
	Color   string
}

func waitForActivity(sub chan UIMessageData) tea.Cmd {
	return func() tea.Msg {
		return UIMessageData(<-sub)
	}
}

type UISectionType string

const (
	UISectionStatusText UISectionType = "status_text"
)

func (m *UIinstance) UpdateSectionText(msg UIMessageData) {
	m.sub <- msg
}
