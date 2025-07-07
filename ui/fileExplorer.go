package ui

import (
	"fmt"
	"io"
	"strings"
	"pen/style"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item string

func (i item) FilterValue() string { return "" }

type FileExplorer struct {
	list   list.Model
	choice string
}

func InitFileExplorer() FileExplorer {
	items := []list.Item{
		item("Ramen"),
		item("Tomato Soup"),
		item("Hamburgers"),
		item("Cheeseburgers"),
		item("Currywurst"),
		item("Okonomiyaki"),
		item("Pasta"),
		item("Fillet Mignon"),
		item("Caviar"),
		item("Just Wine"),
		item("Hamburgers"),
		item("Cheeseburgers"),
		item("Currywurst"),
		item("Okonomiyaki"),
		item("Pasta"),
		item("Fillet Mignon"),
		item("Caviar"),
		item("Currywurst"),
		item("Okonomiyaki"),
		item("Pasta"),
		item("Fillet Mignon"),
		item("Caviar"),
		item("Just Wine"),
		item("Hamburgers"),
		item("Cheeseburgers"),
		item("Currywurst"),
	}

	l := list.New(items, itemDelegate{}, 20, 20)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.Title = "Notes"

	return FileExplorer{
		list:   l,
		choice: "list",
	}
}

func (f FileExplorer) Init() tea.Cmd {
	return nil
}

func (f FileExplorer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	f.list, cmd = f.list.Update(msg)
	return f, cmd
}

func (f FileExplorer) View() string {
	return "\n" + f.list.View()
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprint(i)

	fn := style.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return style.SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}