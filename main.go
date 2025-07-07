package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"pen/style"
	"pen/ui"
)

type focusedView uint

const (
	editorView focusedView = iota
	fileExplorerView
)

type Pen struct {
	focused      focusedView
	fileExplorer ui.FileExplorer
	editor       ui.Editor
}

func main() {
	p := tea.NewProgram(Pen{
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
	return tea.Batch(
		p.fileExplorer.Init(),
		p.editor.Init(),
	)
}

func (p Pen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		p.handleResize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return p, tea.Quit
		case "ctrl+w":
			if p.focused == editorView {
				p.focused = fileExplorerView
			} else {
				p.focused = editorView
			}
		}
	}

	var cmd tea.Cmd
	var model tea.Model

	if p.focused == editorView {
		model, cmd = p.editor.Update(msg)
		p.editor = model.(ui.Editor)
		cmds = append(cmds, cmd)
	} else {
		model, cmd = p.fileExplorer.Update(msg)
		p.fileExplorer = model.(ui.FileExplorer)
		cmds = append(cmds, cmd)
	}

	return p, tea.Batch(cmds...)
}

func (p Pen) View() string {
	var v string

	if p.focused == editorView {
		v += lipgloss.JoinHorizontal(
			lipgloss.Top,
			style.FileExplorerBase.Render(p.fileExplorer.View()),
			style.EditorFocused.Render(p.editor.View()),
		)
	} else {
		v += lipgloss.JoinHorizontal(
			lipgloss.Top,
			style.FileExplorerFocused.Render(p.fileExplorer.View()),
			style.EditorBase.Render(p.editor.View()),
		)
	}

	return v
}

// Helper function to handle window resize
func (p *Pen) handleResize(width, height int) {

	editorWidth := width - style.FileExplorerWidth - style.EditorMargin - 6
	p.fileExplorer.SetHeight(height - style.VerticalPadding)
	p.editor.SetSize(editorWidth, height-style.VerticalPadding)
}
