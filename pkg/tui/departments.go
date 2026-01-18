package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/htekgulds/terminal-rehber/services"
)

// DepartmentsModel represents the departments table model
type DepartmentsModel struct {
	table table.Model
	ready bool
}

// NewDepartmentsModel creates a new departments table model
func NewDepartmentsModel() (*DepartmentsModel, error) {
	// Load departments from services
	departments, err := services.GetDepartments()
	if err != nil {
		return nil, fmt.Errorf("failed to load departments: %w", err)
	}

	// Define table columns
	columns := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Phone", Width: 20},
		{Title: "Manager", Width: 20},
		{Title: "Parent Dept", Width: 20},
	}

	// Build table rows
	rows := make([]table.Row, 0, len(departments))
	for _, dept := range departments {
		manager, err := services.GetPersonById(dept.ManagerId)
		if err != nil {
			return nil, fmt.Errorf("failed to get manager: %w", err)
		}
		parentDept := ""
		if dept.ParentDepartmentId != nil {
			parentDeptObj, err := services.GetDepartmentById(*dept.ParentDepartmentId)
			if err != nil {
				return nil, fmt.Errorf("failed to get parent department: %w", err)
			}

			parentDept = parentDeptObj.Name
		}
		rows = append(rows, table.Row{
			dept.Name,
			dept.Phone,
			fmt.Sprintf("%s %s", manager.FirstName, manager.LastName),
			parentDept,
		})
	}

	// Create table model
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	// Style the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	return &DepartmentsModel{
		table: t,
		ready: false,
	}, nil
}

// Init initializes the model
func (m *DepartmentsModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *DepartmentsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ready {
			m.ready = true
		}
		// Adjust table size based on window size (account for tab bar ~3 lines)
		availableHeight := msg.Height - 8 // Leave space for borders, padding, and tab bar
		availableWidth := msg.Width - 8
		if availableHeight > 0 {
			m.table.SetHeight(availableHeight)
		}
		if availableWidth > 0 {
			m.table.SetWidth(availableWidth)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Handle row selection if needed
			return m, nil
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View renders the UI
func (m *DepartmentsModel) View() string {
	if !m.ready {
		return "Loading departments..."
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)

	content := fmt.Sprintf("%s", m.table.View())

	return style.Render(content)
}
