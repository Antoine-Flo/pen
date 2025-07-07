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

var base = lipgloss.NewStyle().
	BorderStyle(lipgloss.HiddenBorder())

var focused = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("69"))

var ItemStyle = lipgloss.NewStyle().PaddingLeft(2)
var SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

// File Explorer styles with fixed width
var FileExplorerBase = base.
	Width(FileExplorerWidth)

var FileExplorerFocused = focused.
	Width(FileExplorerWidth)

// Editor styles - will be sized dynamically
var EditorBase = base.
	PaddingLeft(ViewMargin)

var EditorFocused = focused.
	PaddingLeft(ViewMargin)
