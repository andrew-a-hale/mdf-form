package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColour lipgloss.Color
	InputField   lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColour = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColour).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(80)

	return s
}

type model struct {
	styles    Styles
	index     int
	questions []Question
	width     int
	height    int
	done      bool
}

type Question struct {
	question string
	answer   string
	input    Input
}

func NewQuestion(question string) Question {
	return Question{question: question}
}

func newShortQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewShortAnswerField()
	q.input = field

	return q
}

func newLongQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewLongAnswerField()
	q.input = field

	return q
}

func New(questions []Question) *model {
	styles := DefaultStyles()
	return &model{questions: questions, styles: *styles}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			current.answer = current.input.Value()
			log.Printf(
				"question: %s, answer: %s\n",
				current.question,
				current.answer,
			)
			m.Next()
			return m, current.input.Blur
		}
	}

	// update model
	current.input, cmd = current.input.Update(msg)

	return m, cmd
}

func (m model) View() string {
	current := m.questions[m.index]
	if m.done {
		var output string
		for _, q := range m.questions {
			output += fmt.Sprintf("%s: %s\n", q.question, q.answer)
		}
		return output
	}

	if m.width == 0 {
		return "loading..."
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index].question,
			m.styles.InputField.Render(current.input.View()),
		),
	)
}

func (m *model) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
		m.done = true
	}
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	questions := []Question{
		newShortQuestion("first question"),
		newLongQuestion("second question"),
	}
	p := tea.NewProgram(New(questions), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
