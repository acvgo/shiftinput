package tui

import "github.com/charmbracelet/lipgloss"

const PINK_HEX = "#f699cd"
var docStyle = lipgloss.NewStyle().Padding(2).Margin(2).Border(lipgloss.NormalBorder()).Foreground(lipgloss.Color(PINK_HEX))
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type itemSelectionMsg struct{}

type progressIncrementMsg struct{}

type progressSelectedMsg struct{}
