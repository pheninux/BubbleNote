package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type PartialModel struct {
}

func (sm *StateManager) titleView() string {
	BoxTitle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
	return lipgloss.Place(with-2, 2, lipgloss.Center, lipgloss.Center, BoxTitle.Foreground(lipgloss.Color("#D7E00F")).Padding(0, 2, 0, 2).Bold(true).Render("Block Note"), lipgloss.WithWhitespaceChars("#"), lipgloss.WithWhitespaceForeground(lipgloss.Color("66")))
	//return docStyle.Align(lipgloss.Center).Foreground(lipgloss.Color("#D7E00F")).Render("Block Note\n")
}
