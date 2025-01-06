package main


import "github.com/charmbracelet/lipgloss"

func styledBox(content string, color string) string {
  style := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color(color)).
    BorderStyle(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color(color)).
    Width(100).
    Align(lipgloss.Left)
  return style.Render(content)
}
