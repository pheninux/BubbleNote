package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type PartialModel struct {
}

func (b *BaseModel) titleView() string {
	return lipgloss.NewStyle().Align(lipgloss.Center).Width(with - 2).Foreground(lipgloss.Color("#D7E00F")).Render("Block Note\n")
}
