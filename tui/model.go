package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelState int

const (
	selectionState modelState = iota
	progressState
)

type Model struct {
	state     modelState
	selection tea.Model
	progress  tea.Model
	text      textinput.Model
	cursor    bool
}

func NewModel() Model {
	return Model{
		selection: newSelectionModel(),
		progress:  newProgress(),
		text:      textinput.New(),
	}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."

	m.text.Placeholder = "Search Terms..."
	m.text.CharLimit = 50
	m.text.Width = 60

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case itemSelectionMsg:
		if !m.cursor {
			m.state = progressState
		}
		cmd = m.progressSelectedCmd()
	case progressSelectedMsg:
		if !m.cursor {
			m.progress, cmd = m.progress.Update(msg)
		}
	case progressIncrementMsg:
		m.progress, cmd = m.progress.Update(msg)
	case progress.FrameMsg:
		m.progress, cmd = m.progress.Update(msg)
	case tea.KeyMsg:
		if m.cursor && msg.String() != "tab" {
			var textCmd tea.Cmd
			m.text, textCmd = m.text.Update(msg) // ? Store the updated textinput model
			return m, tea.Batch(textCmd, cmd)    // ? Return the command
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			fmt.Println("Pressed Tab!")
			m.cursor = !m.cursor
			if m.cursor {
				m.text.Focus()
			} else {
				m.text.Blur()
			}
		default:
			if m.state == selectionState {
				m.selection, cmd = m.selection.Update(msg)
			}
			if m.state == progressState {
				m.progress, cmd = m.progress.Update(msg)
			}
		}
	}

	return m, tea.Batch(cmd)
}

func (m Model) View() string {
	switch m.state {
	case progressState:
		return m.progress.View()
	default:
		test := lipgloss.NewStyle()
		inputView := test.Render(m.text.View())
		selectView := m.selection.View()
		return lipgloss.JoinVertical(lipgloss.Top, selectView, inputView)
	}
}

func (m Model) progressSelectedCmd() tea.Cmd {
	return func() tea.Msg {
		return progressSelectedMsg{}
	}
}
