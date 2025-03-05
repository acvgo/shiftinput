package tui

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

type ProgressModel struct {
	progress progress.Model
}

func newProgress() ProgressModel {
	return ProgressModel{
		progress: progress.New(progress.WithDefaultScaledGradient()),
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progressSelectedMsg:
		return m, nil
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			default:
				return m, tea.Batch(progressIncrementCmd())
			}
		}
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case progressIncrementMsg:
		// hardcoded increment
		cmd := m.progress.IncrPercent(0.25)
		return m, tea.Batch(cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m ProgressModel) View() string {
	return docStyle.Render(func() string {
		s := "Welcome to Addetz' Bubble Tea Tutorial!\n\n"
		s = s + "I'm the progress model.\n\n"
		s = s + m.progress.View() + "\n\n"
		if m.progress.Percent() != 1.0 {
			s = s + "Press any key to increment progress!\n\n"
		}

		return s + helpStyle("\nPress q to quit.\n")
	}())
}

func progressIncrementCmd() tea.Cmd {
	return func() tea.Msg {
		return progressIncrementMsg{}
	}
}
