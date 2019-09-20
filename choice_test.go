package gochoice

import (
	"fmt"
	"github.com/gdamore/tcell"
	"testing"
)

func TestPickFirstChoice(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor.toTcellColor()))
	screen.Show()
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "A" {
		t.Error()
	}
}

func TestPickSecondChoice(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor.toTcellColor()))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "B" {
		t.Error()
	}
}

func TestPickQuit(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor.toTcellColor()))
	screen.Show()
	screen.InjectKey(tcell.KeyRune, 'q', tcell.ModNone)
	_, err = pick("question", []string{"A", "B", "C"}, screen, &config)
	if err == nil {
		t.Error()
	}
}

func createSimulationScreen() (tcell.SimulationScreen, error) {
	screen := tcell.NewSimulationScreen("UTF-8")
	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize screen: %v", err)
	}
	return screen, nil
}
