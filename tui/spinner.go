package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	pulseDot = spinner.Spinner{
		Frames: []string{"○", "◎", "●", "◎"},
		FPS:    time.Second / 7,
	}

	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Meter,
		spinner.Hamburger,
		pulseDot,
	}

	textStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	boldText             = lipgloss.NewStyle().Bold(true).Render
	logStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("246")).Render
	bindingShortCutStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0f0f1f")).Render

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))

	loadingSpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

func (m *UIinstance) resetSpinner() {
	m.statusSpinner = spinner.New()
	m.statusSpinner.Style = spinnerStyle
	m.statusSpinner.Spinner = spinners[4]

	m.newVersionSpinner = spinner.New()
	m.newVersionSpinner.Spinner = spinners[7]

	m.loadingSpinner = spinner.New()
	m.loadingSpinner.Style = colorToStyle("warning")
	m.loadingSpinner.Spinner = spinners[6]
}

func (m *UIinstance) setStatusSpinnerStyle(style lipgloss.Style) {
	m.statusSpinner.Style = style
}
