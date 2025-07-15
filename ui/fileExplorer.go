package ui

import (
	"fmt"
	"io"

	"pen/note"
	"pen/style"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NoteItem struct {
	note *note.Note
}

const fileExplorerWidth = 30

func (i NoteItem) FilterValue() string { return "" }

func (i NoteItem) Title() string {
	return i.note.Name
}

func (i NoteItem) Description() string {
	return ""
}

type FileExplorer struct {
	list          list.Model
	focused       bool
	Width, Height int
	err           error
}

func InitFileExplorer() *FileExplorer {
	notes, err := note.ListNotes("./testFiles")

	items := make([]list.Item, len(notes))
	for i, note := range notes {
		items[i] = NoteItem{note: note}
	}

	l := list.New(items, itemDelegate{}, 50, 20)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.Title = "📁 Notes"

	return &FileExplorer{
		list: l,
		err:  err,
	}
}

func (f *FileExplorer) Init() tea.Cmd {
	return nil
}

func (f *FileExplorer) SetFocused(focused bool) {
	f.focused = focused
}

func (f *FileExplorer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	f.list, cmd = f.list.Update(msg)
	return f, cmd
}

func (f *FileExplorer) View() string {

	v := lipgloss.NewStyle().Width(fileExplorerWidth)

	if f.focused {
		v = v.Inherit(style.Focused)
	} else {
		v = v.Inherit(style.Base)
	}

	// Set list dimensions accounting for borders
	f.list.SetWidth(fileExplorerWidth - 2)
	f.list.SetHeight(f.Height - 2)

	return v.Render(f.list.View())
}

func (f *FileExplorer) SetHeight(height int) {
	f.list.SetHeight(height)
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(NoteItem)
	if !ok {
		return
	}

	str := i.Title()

	fn := style.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return style.SelectedItemStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
