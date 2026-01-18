package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	tabPeople      = 0
	tabDepartments = 1
)

// Model represents the TUI model with tabs
type Model struct {
	width       int
	height      int
	ready       bool
	activeTab   int
	peopleModel *PeopleModel
	deptModel   *DepartmentsModel
	tabNames    []string
}

// NewModel creates a new TUI model with tabs
func NewModel() (*Model, error) {
	peopleModel := NewPeopleModel()
	deptModel, err := NewDepartmentsModel()
	if err != nil {
		return nil, err
	}

	return &Model{
		activeTab:   tabPeople,
		peopleModel: peopleModel,
		deptModel:   deptModel,
		tabNames:    []string{"People", "Departments"},
	}, nil
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.peopleModel.Init())
	cmds = append(cmds, m.deptModel.Init())
	return tea.Batch(cmds...)
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		// Forward window size to child models (they'll adjust for tab bar)
		peopleModel, cmd1 := m.peopleModel.Update(msg)
		m.peopleModel = peopleModel.(*PeopleModel)
		cmds = append(cmds, cmd1)

		deptModel, cmd2 := m.deptModel.Update(msg)
		m.deptModel = deptModel.(*DepartmentsModel)
		cmds = append(cmds, cmd2)

		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			// Switch to next tab
			m.activeTab = (m.activeTab + 1) % len(m.tabNames)
			return m, nil
		case "shift+tab":
			// Switch to previous tab
			m.activeTab = (m.activeTab - 1 + len(m.tabNames)) % len(m.tabNames)
			return m, nil
		case "1":
			m.activeTab = tabPeople
			return m, nil
		case "2":
			m.activeTab = tabDepartments
			return m, nil
		case "esc":
			// ESC quits only if not in a sub-view
			return m, tea.Quit
		}
	}

	// Forward update to active model
	switch m.activeTab {
	case tabPeople:
		peopleModel, cmd := m.peopleModel.Update(msg)
		m.peopleModel = peopleModel.(*PeopleModel)
		return m, cmd
	case tabDepartments:
		deptModel, cmd := m.deptModel.Update(msg)
		m.deptModel = deptModel.(*DepartmentsModel)
		return m, cmd
	}

	return m, cmd
}

// View renders the UI
func (m *Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	// Render tabs
	tabs := m.renderTabs()

	// Render active view
	var content string
	switch m.activeTab {
	case tabPeople:
		content = m.peopleModel.View()
	case tabDepartments:
		content = m.deptModel.View()
	}

	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		MarginLeft(2).
		MarginTop(1).
		Render("Tab/Shift+Tab: Switch • 1/2: Jump • ↑/↓: Navigate • Enter: Select • q: Quit")

	// Combine tabs and content
	return lipgloss.JoinVertical(lipgloss.Left, tabs, content, helpText)
}

// renderTabs renders the tab bar
func (m *Model) renderTabs() string {
	var tabs []string

	// Active tab style - prominent with bright colors
	activeTabStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true).
		Padding(0, 4)

	// Inactive tab style - subtle but clearly visible
	inactiveTabStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("246")).
		Background(lipgloss.Color("237")).
		Padding(0, 4)

	for i, name := range m.tabNames {
		if i == m.activeTab {
			tabs = append(tabs, activeTabStyle.Render(" "+name+" "))
		} else {
			tabs = append(tabs, inactiveTabStyle.Render(" "+name+" "))
		}
		tabs = append(tabs, "  ")
	}

	// Create tab bar with spacing
	tabBar := lipgloss.JoinHorizontal(lipgloss.Left, tabs...)

	// Add bottom border for the tab bar area
	width := m.width
	if width == 0 {
		width = 80 // Default width if not set
	}

	return lipgloss.NewStyle().
		Width(width).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(2).
		Render(tabBar)
}
