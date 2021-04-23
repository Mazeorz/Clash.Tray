package main

import (
	"C"
	. "Clash.Tray/icon"
	"github.com/getlantern/systray"
	"os"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(Date0)
	systray.SetTitle("Clash.Tray")
	systray.SetTooltip("A Tray tool for Clash")
	systray.AddMenuItem("Clash.Tray", "")
	systray.AddSeparator()

	RuleSwitch := systray.AddMenuItem("Rule", "")
	mGlobal := RuleSwitch.AddSubMenuItem("Global", "")
	mRule := RuleSwitch.AddSubMenuItem("Rule", "")
	mDirect := RuleSwitch.AddSubMenuItem("Direct", "")
	mQuit := systray.AddMenuItem("Exit", "Quit")

	go func() {
		for {
			select {
			case <-mDirect.ClickedCh:
				systray.SetIcon(Date2)
				mDirect.Check()
				mRule.Uncheck()
				mGlobal.Uncheck()
			case <-mRule.ClickedCh:
				systray.SetIcon(Date3)
				mDirect.Uncheck()
				mRule.Check()
				mGlobal.Uncheck()
			case <-mGlobal.ClickedCh:
				systray.SetIcon(Date4)
				mDirect.Uncheck()
				mRule.Uncheck()
				mGlobal.Check()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	os.Exit(1)
}
