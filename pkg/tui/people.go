package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/htekgulds/terminal-rehber/services"
)

// PeopleModel represents the people table model
type PeopleModel struct {
	table  table.Model
	people []services.Person
	width  int
	height int
	ready  bool
}

// NewPeopleModel creates a new people table model
func NewPeopleModel() *PeopleModel {
	// Fetch people data
	people, err := services.GetPeople()
	if err != nil {
		// If there's an error, create empty model
		// In a real scenario, you might want to handle this differently
		people = []services.Person{}
	}

	// Define table columns
	columns := []table.Column{
		{Title: "Name", Width: 25},
		{Title: "Prefix", Width: 12},
		{Title: "Title", Width: 25},
		{Title: "Room", Width: 10},
		{Title: "Phone", Width: 18},
		{Title: "Floor", Width: 6},
	}

	// Convert people to table rows
	rows := make([]table.Row, len(people))
	for i, person := range people {
		prefix := ""
		if person.Prefix != nil {
			prefix = *person.Prefix
		}
		fullName := fmt.Sprintf("%s %s", person.FirstName, person.LastName)
		rows[i] = table.Row{
			fullName,
			prefix,
			person.Title,
			person.Room,
			person.Phone,
			strconv.Itoa(person.Floor),
		}
	}

	// Create table
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
		Bold(false)
	t.SetStyles(s)

	return &PeopleModel{
		table:  t,
		people: people,
	}
}

// Init initializes the model
func (m *PeopleModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *PeopleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		// Update table width to fit window (account for tab bar ~3 lines)
		m.table.SetWidth(msg.Width - 8)
		m.table.SetHeight(msg.Height - 8)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Handle row selection if needed
			return m, nil
		}
	}

	// Update table
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// View renders the UI
func (m *PeopleModel) View() string {
	if !m.ready {
		return "Loading people data..."
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)

	// help := lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("240")).
	// 	MarginTop(1).
	// 	Render("↑/↓: Navigate • Enter: Select")

	content := fmt.Sprintf("%s", m.table.View())

	return style.Render(content)
}
