package main

import (
	"C"
	"Clash.Tray/icon"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"os"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data2)
	systray.SetTitle("Clash.Tray")
	systray.SetTooltip("A Tray tool for Clash")
	systray.AddMenuItem("Clash.Tray", "")
	systray.AddSeparator()
	RuleSwitch := systray.AddMenuItem("Rule", "")
	RuleSwitch.AddSubMenuItem("Global", "")
	RuleSwitch.AddSubMenuItem("Rule", "")
	RuleSwitch.AddSubMenuItem("Direct", "")
	mQuit := systray.AddMenuItem("Exit", "Quit")
	go func() {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		}
	}()
}

func onExit() {
	os.Exit(1)
}
