package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

type ShortAnswerField struct {
	textinput textinput.Model
}

func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "Your answer here"
	return &ShortAnswerField{ti}
}

func (sa *ShortAnswerField) Value() string {
	return sa.textinput.Value()
}

func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textinput.Blur
}

func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textinput, cmd = sa.textinput.Update(msg)
	return sa, cmd
}

func (sa *ShortAnswerField) View() string {
	return sa.textinput.View()
}

type LongAnswerField struct {
	textinput textarea.Model
}

func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Focus()
	ta.Placeholder = "Your answer here"
	return &LongAnswerField{ta}
}

func (la *LongAnswerField) Value() string {
	return la.textinput.Value()
}

func (la *LongAnswerField) Blur() tea.Msg {
	return la.textinput.Blur
}

func (sl *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sl.textinput, cmd = sl.textinput.Update(msg)
	return sl, cmd
}

func (sl *LongAnswerField) View() string {
	return sl.textinput.View()
}
