package style

import (
	"github.com/charmbracelet/lipgloss"
)

var Base = lipgloss.NewStyle().
	Width(40).
	Height(20).
	BorderStyle(lipgloss.HiddenBorder())

var Focused = lipgloss.NewStyle().
	Width(40).
	Height(20).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("69"))

var ItemStyle = lipgloss.NewStyle().PaddingLeft(4)
var SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))