package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the TUI model
type Model struct {
	width  int
	height int
	ready  bool
}

// NewModel creates a new TUI model
func NewModel() *Model {
	return &Model{}
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the UI
func (m *Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	style := lipgloss.NewStyle().
		Width(m.width-2).
		Height(m.height-2).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62"))

	content := lipgloss.NewStyle().
		Padding(1, 2).
		Render("Welcome to Terminal Rehber\n\nPress 'q' to quit")

	return style.Render(content)
}
