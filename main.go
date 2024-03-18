package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var types []string = []string{"File", "Database"}

type model struct {
	data    map[int]textinput.Model
	step    int
	stepMax int
	err     error
}

type errMsg error

type File struct {
	Delimiter string
	Pattern   string
}

type Database struct {
	Catalogue string
	Schema    string
	Table     string
}

func NewModel() model {
	data := make(map[int]textinput.Model)
	tmpFile := File{}
	fields := reflect.TypeOf(tmpFile)

	total := 0
	for i := 0; i < fields.NumField(); i++ {
		ti := textinput.New()
		ti.Prompt = fields.Field(i).Name + ": "
		ti.Placeholder = "placeholder"
		ti.Focus()
		ti.Width = 20
		ti.CharLimit = 150
		data[total] = ti
		total++
	}

	tmpDatabase := Database{}
	fields = reflect.TypeOf(tmpDatabase)

	for i := 0; i < fields.NumField(); i++ {
		ti := textinput.New()
		ti.Prompt = fields.Field(i).Name + ": "
		ti.Placeholder = "placeholder"
		ti.Focus()
		ti.Width = 20
		ti.CharLimit = 150
		data[total] = ti
		total++
	}

	m := model{
		step:    0,
		data:    data,
		stepMax: total,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.step < m.stepMax {
				m.step++
			}
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	ti, ok := m.data[m.step]
	if ok {
		ti, cmd = ti.Update(msg)
	}

	m.data[m.step] = ti
	return m, cmd
}

func (m model) View() string {
	var s string

	switch m.step {
	case m.stepMax:
		for i := 0; i < m.stepMax; i++ {
			s += fmt.Sprintf("%s\n", m.data[i].Value())
		}
	default:
		s = fmt.Sprintf(
			"%s\n\n",
			m.data[m.step].View(),
		)
	}

	s += "(esc or ctrl+c to quit)"
	return s
}

func main() {
	_, err := tea.NewProgram(NewModel()).Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
