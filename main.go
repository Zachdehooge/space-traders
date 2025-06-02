package main

import (
	"fmt"
	"strconv"

	"github.com/Zachdehooge/space-traders/api"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#f20023")).
			PaddingTop(1).
			PaddingLeft(2)

	agentStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#f25c00")).
			PaddingTop(1).
			PaddingLeft(4)
)

type item struct {
	title string
	desc  string
}

func (i item) FilterValue() string { return i.title }
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }

type model struct {
	list       list.Model
	state      string // "menu" or "agentInfo"
	agentData  string
	serverData string
}

func (m model) Init() tea.Cmd {
	return nil
}

// Helper function to get main menu items
func getMainMenuItems() []list.Item {
	return []list.Item{
		item{title: "Server Information", desc: "View server details"},
		item{title: "Agent Information", desc: "View your agent details and credits"},
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.state == "menu" {
				selectedItem := m.list.SelectedItem()
				if selectedItem != nil {
					switch selectedItem.(item).title {
					case "Agent Information":
						m.state = "agentInfo"
						m.agentData = m.getAgentInfo()
						// Create a new list for the agent info view with back option
						items := []list.Item{
							item{title: "← Back to Main Menu", desc: "Return to the main menu"},
						}
						m.list.SetItems(items)
						return m, nil
					case "Server Information":
						m.state = "serverInfo"
						m.serverData = m.getServerInfo()
						items := []list.Item{
							item{title: "← Back to Main Menu", desc: "Return to the main menu"},
						}
						m.list.SetItems(items)
						return m, nil
					}
				}
			} else if m.state == "agentInfo" {
				selectedItem := m.list.SelectedItem()
				if selectedItem != nil && selectedItem.(item).title == "← Back to Main Menu" {
					m.state = "menu"
					// Restore main menu items
					m.list.SetItems(getMainMenuItems())
					return m, nil
				}
			} else if m.state == "serverInfo" {
				selectedItem := m.list.SelectedItem()
				if selectedItem != nil && selectedItem.(item).title == "← Back to Main Menu" {
					m.state = "menu"
					// Restore main menu items
					m.list.SetItems(getMainMenuItems())
					return m, nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.state == "agentInfo" {
		return titleStyle.Render("Space Traders - Agent Information") + "\n\n" +
			m.agentData + "\n\n" +
			m.list.View()
	}
	if m.state == "serverInfo" {
		return titleStyle.Render("Space Traders - Server Information") + "\n\n" +
			m.serverData + "\n\n" +
			m.list.View()
	}

	return titleStyle.Render("Space Traders - Main Menu") + "\n\n" + m.list.View()
}

// API calls being implemented into the TUI

func (m model) getAgentInfo() string {
	agentName := api.PlayerAgent()
	agentCredits := strconv.Itoa(api.PlayerCredits())

	return agentStyle.Render("Agent Name: "+agentName) + "\n" +
		agentStyle.Render("Agent Credits: "+agentCredits)
}

func (m model) getServerInfo() string {
	serverStatus := api.ServerStatus()
	serverReset := api.ServerResetDate()

	return agentStyle.Render("Server Status: "+serverStatus) + "\n" +
		agentStyle.Render("Last Server Reset Date: "+serverReset)
}

func main() {
	const defaultWidth = 80
	const listHeight = 14

	l := list.New(getMainMenuItems(), list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = ""
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle

	m := model{
		list:  l,
		state: "menu",
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
