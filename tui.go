package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	inputs      []textinput.Model // 0:input,1:output,2:fromW,3:fromH,4:toW,5:toH
	recursive   bool
	focusIndex  int
	ready       bool
	viewport    viewport.Model
	logs        textarea.Model
	processing  bool
	cfg         Config
	processor   *Processor
	logCh       chan string
	width       int
	height      int
}

func NewModel(cfg Config) model {
	var inputs []textinput.Model
	labels := []string{"Input Path", "Output Path", "From Width", "From Height", "To Width", "To Height"}
	values := []string{
		cfg.InputPath,
		cfg.OutputPath,
		intOrEmpty(cfg.FromWidth),
		intOrEmpty(cfg.FromHeight),
		intOrEmpty(cfg.ToWidth),
		intOrEmpty(cfg.ToHeight),
	}
	for i := 0; i < 6; i++ {
		ti := textinput.New()
		ti.Placeholder = labels[i]
		ti.CharLimit = 256
		ti.SetValue(values[i])
		inputs = append(inputs, ti)
	}
	return model{
		inputs:     inputs,
		recursive:  cfg.Recursive,
		focusIndex: 0,
		logs:       textarea.New(),
		logCh:      make(chan string, 100),
		cfg:        cfg,
	}
}

func intOrEmpty(v int) string {
	if v == 0 {
		return ""
	}
	return fmt.Sprintf("%d", v)
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if !m.ready {
			m.viewport = viewport.New(msg.Width-4, msg.Height-12)
			m.viewport.Style = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
			m.logs.SetWidth(msg.Width - 4)
			m.logs.SetHeight(8)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width - 4
			m.viewport.Height = msg.Height - 12
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			if m.processing {
				break
			}
			s := msg.String()
			if s == "enter" && m.focusIndex == len(m.inputs) {
				// Запуск обработки
				m.cfg.InputPath = m.inputs[0].Value()
				m.cfg.OutputPath = m.inputs[1].Value()
				m.cfg.FromWidth, _ = parseInt(m.inputs[2].Value())
				m.cfg.FromHeight, _ = parseInt(m.inputs[3].Value())
				m.cfg.ToWidth, _ = parseInt(m.inputs[4].Value())
				m.cfg.ToHeight, _ = parseInt(m.inputs[5].Value())
				m.cfg.Recursive = m.recursive

				if m.cfg.InputPath == "" || m.cfg.OutputPath == "" {
					m.logs.InsertString("ERROR: Input and Output paths are required\n")
					break
				}
				m.processing = true
				m.processor = NewProcessor(m.cfg)
				go m.processor.Process(m.logCh)
				return m, m.listenForLogs()
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}
			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			}

			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds = append(cmds, m.inputs[i].Focus())
				} else {
					m.inputs[i].Blur()
				}
			}
			return m, tea.Batch(cmds...)
		}
	}

	if !m.processing {
		var cmd tea.Cmd
		if m.focusIndex >= 0 && m.focusIndex < len(m.inputs) {
			m.inputs[m.focusIndex], cmd = m.inputs[m.focusIndex].Update(msg)
			cmds = append(cmds, cmd)
		}
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "r" {
			m.recursive = !m.recursive
		}
	}

	if logMsg, ok := msg.(logMessage); ok {
		m.logs.InsertString(string(logMsg) + "\n")
		m.viewport.SetContent(m.logs.Value())
		m.viewport.GotoBottom()
	}

	select {
	case _, ok := <-m.logCh:
		if !ok {
			m.processing = false
		}
	default:
	}

	return m, tea.Batch(cmds...)
}

type logMessage string

func (m model) listenForLogs() tea.Cmd {
	return func() tea.Msg {
		for msg := range m.logCh {
			return logMessage(msg)
		}
		return nil
	}
}

func parseInt(s string) (int, error) {
	var val int
	_, err := fmt.Sscan(s, &val)
	return val, err
}

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	var b strings.Builder
	b.WriteString("Batch Image Resizer (TUI)\n\n")

	for i, input := range m.inputs {
		prefix := "  "
		if i == m.focusIndex {
			prefix = "> "
		}
		b.WriteString(prefix + input.View() + "\n")
	}

	recText := "[ ]"
	if m.recursive {
		recText = "[x]"
	}
	if m.focusIndex == len(m.inputs) {
		b.WriteString("> " + recText + " Recursive (press r to toggle)\n")
	} else {
		b.WriteString("  " + recText + " Recursive (press r to toggle)\n")
	}

	button := "[ Start Resize ]"
	if m.focusIndex == len(m.inputs) {
		button = "> " + button
	} else {
		button = "  " + button
	}
	b.WriteString("\n" + button + "\n\n")

	b.WriteString("Logs:\n")
	b.WriteString(m.viewport.View())

	return lipgloss.NewStyle().
		Width(m.width-2).
		Height(m.height-2).
		BorderStyle(lipgloss.RoundedBorder()).
		Render(b.String())
}