package main

import "github.com/charmbracelet/lipgloss"

//style bar
var (
	w             = lipgloss.Width
	h             = lipgloss.Height
	barStyleColor = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#353533",
		Dark:  "#F2F2D9",
	}).Background(lipgloss.AdaptiveColor{
		Light: "#CACAC6",
		Dark:  "#DCDCD5",
	})
	textFg        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FEFFFE")).Padding(0, 1)
	statusStyle   = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#EA52BC"))
	msgStyle      = textFg.Copy().Inherit(barStyleColor)
	encodingStyle = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#F5BAF7"))
	versionStyle  = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#9152EA"))
)
