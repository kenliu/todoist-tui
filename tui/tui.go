package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct{}

func (m MainModel) Init() tea.Cmd {
	//TODO implement me
	panic("implement me")
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//TODO implement me
	panic("implement me")
}

func (m MainModel) View() string {
	//TODO implement me
	panic("implement me")
}

func InitialModel() MainModel {
	return MainModel{}
}
