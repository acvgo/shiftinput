package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectionModel struct {
	Choices  []item
	Cursor   int
	Selected map[int]struct{}
}

type item struct {
	text    string
	onPress func() tea.Msg
}

func newSelectionModel() SelectionModel {
	return SelectionModel{
		Choices: []item{
			{
				text:    "Show me a progress bar!",
				onPress: func() tea.Msg { return itemSelectionMsg{} },
			},
		},
		Selected: make(map[int]struct{}),
	}
}

func (m SelectionModel) Init() tea.Cmd {
	return nil
}

func (m SelectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
				return m, m.Choices[m.Cursor].onPress
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m SelectionModel) View() string {
	return docStyle.Render(func() string {
		// The header
		s := "Welcome to Addetz' Bubble Tea Tutorial!\n\n"
		s = s + "I'm the selection model.\n\n"

		// Iterate over our choices
		for i, choice := range m.Choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.Cursor == i {
				cursor = ">" // cursor!
			}

			// Is this choice selected?
			checked := " " // not selected
			if _, ok := m.Selected[i]; ok {
				checked = "x" // selected!
			}

			// Render the row
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.text)
		}

		// The footer
		s += helpStyle("\nPress q to quit.\n")

		// Send the UI for rendering
		return s
	}())
}
