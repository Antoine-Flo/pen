package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"pen/ui"
)

type focusedPane uint

const (
	rightPane focusedPane = iota
	leftPane
)

type Pen struct {
	focused       focusedPane
	fileExplorer  *ui.FileExplorer
	editor        *ui.Editor
	width, height int
}

func main() {
	p := tea.NewProgram(Pen{
		focused:      rightPane, // Editor focused by default
		fileExplorer: ui.InitFileExplorer(),
		editor:       ui.InitEditor(),
	}, tea.WithAltScreen())

	f, _ := tea.LogToFile("logs/debug.log", "debug")
	defer f.Close()

	if _, err := p.Run(); err != nil {
		log.Print(err)
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (p Pen) Init() tea.Cmd {
	// Set initial focus state using single source of truth
	p.updateFocusState()

	return tea.Batch(
		p.fileExplorer.Init(),
		p.editor.Init(),
	)
}

// updateFocusState syncs component focus based on main focus state
func (p *Pen) updateFocusState() {
	p.editor.SetFocused(p.focused == rightPane)
	p.fileExplorer.SetFocused(p.focused == leftPane)
}

func (p Pen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		p.width, p.height = msg.Width, msg.Height
		p.fileExplorer.Height, p.editor.Height = p.height, p.height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return p, tea.Quit
		case "ctrl+w":
			p.focused = focusedPane(1 - p.focused)
			p.updateFocusState()
			return p, nil
		}
	}

	var cmd tea.Cmd
	var model tea.Model

	if p.focused == rightPane {
		model, cmd = p.editor.Update(msg)
		p.editor = model.(*ui.Editor)
	}
	if p.focused == leftPane {
		model, cmd = p.fileExplorer.Update(msg)
		p.fileExplorer = model.(*ui.FileExplorer)
	}

	cmds = append(cmds, cmd)
	return p, tea.Batch(cmds...)
}

func (p Pen) View() string {

	fileExplorer := p.fileExplorer.View()
	p.editor.Width = p.width - lipgloss.Width(fileExplorer)
	editor := p.editor.View()

	v := lipgloss.JoinHorizontal(
		lipgloss.Top,
		fileExplorer,
		editor,
	)

	return v
}
