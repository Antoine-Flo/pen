package ui

import (
	"log"
	"pen/style"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type Editor struct {
	textearea     textarea.Model
	err           error
	Focused       bool
	Width, Height int
}

func InitEditor() *Editor {
	editor := textarea.New()
	editor.Focus()

	return &Editor{
		textearea: editor,
		err:       nil,
	}
}

func (e *Editor) Init() tea.Cmd {
	return textarea.Blink
}

func (e *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case errMsg:
		log.Print(msg)
		e.err = msg
		return e, nil
	}

	e.textearea, cmd = e.textearea.Update(msg)

	return e, cmd
}

func (e *Editor) View() string {
	var v lipgloss.Style

	if e.Focused {
		v = style.Focused
	} else {
		v = style.Base
	}

	// Set textarea dimensions accounting for borders
	e.textearea.SetWidth(e.Width - 2)
	e.textearea.SetHeight(e.Height - 2)

	return v.Render(e.textearea.View())
}
