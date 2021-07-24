package gochoice

import (
	"fmt"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestPickFirstChoice(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "A" {
		t.Error()
	}
	if index != 0 {
		t.Error("Choice 'A' should have returned index 0")
	}
}

func TestPickSecondChoice(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "B" {
		t.Error()
	}
	if index != 1 {
		t.Error("Choice 'B' should have returned index 1")
	}
}

func TestPickThirdChoice(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "C" {
		t.Error()
	}
	if index != 2 {
		t.Error("Choice 'C' should have returned index 2")
	}
}

func TestPickLastChoiceWithEndKey(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyEnd, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "C" {
		t.Error()
	}
	if index != 2 {
		t.Error("Choice 'C' should have returned index 2")
	}
}

func TestPickFirstChoiceWithHomeKey(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyHome, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "A" {
		t.Error()
	}
	if index != 0 {
		t.Error("Choice 'A' should have returned index 0")
	}
}

func TestPickThirdChoiceButWithExtraKeyDownInputs(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "C" {
		t.Error()
	}
	if index != 2 {
		t.Error("Choice 'C' should have returned index 2")
	}
}

func TestPickSecondChoiceButWithMoreComplexKeyCombo(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyDown, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyUp, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, _ := pick("question", []string{"A", "B", "C"}, screen, &config)
	if choice != "B" {
		t.Error()
	}
	if index != 1 {
		t.Error("Choice 'B' should have returned index 1")
	}
}

func TestPickQuit(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyEscape, 0, tcell.ModNone)
	_, _, err = pick("question", []string{"A", "B", "C"}, screen, &config)
	if err == nil {
		t.Error()
	}
}

func TestPickSearch(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyRune, 'd', tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, err := pick("question", []string{"john", "doe", "jane"}, screen, &config)
	if err != nil {
		t.Fatal(err.Error())
	}
	if choice != "doe" {
		t.Error("expected doe, got", choice)
	}
	if index != 1 {
		t.Error("expected 1, got", index)
	}
}

func TestPickSearchNoResult(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyRune, 'z', tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	_, _, err = pick("question", []string{"john", "doe", "jane"}, screen, &config)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestPickSearchComplex(t *testing.T) {
	config := defaultConfig
	screen, err := createSimulationScreen()
	if err != nil {
		t.Errorf("encountered error while creating simulation screen: %v", err)
	}
	defer screen.Fini()
	screen.SetStyle(tcell.StyleDefault.Background(config.BackgroundColor))
	screen.Show()
	screen.InjectKey(tcell.KeyRune, 'j', tcell.ModNone)
	screen.InjectKey(tcell.KeyRune, 'a', tcell.ModNone)
	screen.InjectKey(tcell.KeyRune, 'n', tcell.ModNone)
	screen.InjectKey(tcell.KeyRune, 'e', tcell.ModNone)
	screen.InjectKey(tcell.KeyRune, 'z', tcell.ModNone)
	screen.InjectKey(tcell.KeyBackspace, 0, tcell.ModNone)
	screen.InjectKey(tcell.KeyEnter, 0, tcell.ModNone)
	choice, index, err := pick("question", []string{"john", "doe", "jane"}, screen, &config)
	if err != nil {
		t.Fatal(err.Error())
	}
	if choice != "jane" {
		t.Error("expected jane, got", choice)
	}
	if index != 2 {
		t.Error("expected 2, got", index)
	}
}

func createSimulationScreen() (tcell.SimulationScreen, error) {
	screen := tcell.NewSimulationScreen("UTF-8")
	if err := screen.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize screen: %v", err)
	}
	return screen, nil
}
