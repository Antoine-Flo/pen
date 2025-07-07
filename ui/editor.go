package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Editor struct {
	textearea textarea.Model
	text      string
	err       error
}

func InitEditor() Editor {
	editor := textarea.New()
	editor.Focus()

	return Editor{
		textearea: editor,
		text:      "Hello there",
		err:       nil,
	}
}

func (e Editor) Init() tea.Cmd {
	return textarea.Blink
}

func (e Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (e Editor) View() string {
	return e.textearea.View()
}
