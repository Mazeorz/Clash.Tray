package main

import (
	"C"
	"Clash.Tray/controller"
	. "Clash.Tray/icon"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"os"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(Date0)
	systray.SetTitle("Clash.Tray")
	systray.SetTooltip("A Tray tool for Clash")
	mTitle := systray.AddMenuItem("Clash.Tray", "")
	systray.AddSeparator()
	updateConfig := systray.AddMenuItem("更新配置", "")
	mConfig := systray.AddMenuItem("配置管理", "")
	systray.AddSeparator()
	mWork := systray.AddMenuItem("运作模式", "")
	mGeneral := mWork.AddSubMenuItem("普通模式", "")
	mAdvanced := mWork.AddSubMenuItem("Tun模式", "")
	mUwp := systray.AddMenuItem("UWP回环", "")
	systray.AddSeparator()
	RuleSwitch := systray.AddMenuItem("代理模式", "")
	mGlobal := RuleSwitch.AddSubMenuItem("全局", "")
	mRule := RuleSwitch.AddSubMenuItem("规则", "")
	mDirect := RuleSwitch.AddSubMenuItem("直连", "")
	mSysProxy := systray.AddMenuItem("系统代理", "")
	mSysOn := mSysProxy.AddSubMenuItem("开启系统代理", "")
	mSysOff := mSysProxy.AddSubMenuItem("关闭系统代理", "")
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("退出", "Quit")

	go func() {
		for {
			select {
			case <-mTitle.ClickedCh:
				err := open.Start("https://github.com/Mazeorz/Clash.Tray")
				if err != nil {
					return
				}
			case <-mConfig.ClickedCh:
				controller.MenuConfig()
			case <-updateConfig.ClickedCh:

			case <-mGeneral.ClickedCh:

			case <-mAdvanced.ClickedCh:

			case <-mUwp.ClickedCh:

			case <-mSysOn.ClickedCh:

			case <-mSysOff.ClickedCh:

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
