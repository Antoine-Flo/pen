package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Layout constants
const (
	FileExplorerWidth = 30
	EditorMargin      = 12
	ViewMargin        = 2
	VerticalPadding   = 4
)

var ItemStyle = lipgloss.NewStyle().PaddingLeft(2)
var SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

var Base = lipgloss.NewStyle().
	BorderStyle(lipgloss.HiddenBorder())

var Focused = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("69"))

