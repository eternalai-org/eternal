package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Meter,
	}

	textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	boldText  = lipgloss.NewStyle().Bold(true).Render
	logStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("246")).Render

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

func (m *UIinstance) resetStatusSpinner() {
	m.statusSpinner = spinner.New()
	m.statusSpinner.Style = spinnerStyle
	m.statusSpinner.Spinner = spinners[4]
}

func (m *UIinstance) setSpinnerStyle(style lipgloss.Style) {
	m.statusSpinner.Style = style
}

// func (m *UIinstance) ShowSpinner() {
// 	m.spinnerShow = true

// 	if m.spinnerShow {
// 		m.resetSpinner()
// 	}
// }

// func (m *UIinstance) HideSpinner() {
// 	m.spinnerShow = false

// 	if m.spinnerShow {
// 		m.resetSpinner()
// 	}
// }
