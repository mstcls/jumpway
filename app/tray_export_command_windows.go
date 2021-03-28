package app

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/wzshiming/logger"
)

func (a *App) itemExportCommand(menu *systray.MenuItem) {
	mShell := menu.AddSubMenuItem("Shell", "")
	go a.itemExportCommandShell(mShell)

	mCmd := menu.AddSubMenuItem("Cmd", "")
	go a.itemExportCommandCmd(mCmd)

	mPowerShell := menu.AddSubMenuItem("PowerShell", "")
	go a.itemExportCommandPowerShell(mPowerShell)
}

func (a *App) itemExportCommandCmd(menu *systray.MenuItem) {
	for range menu.ClickedCh {
		command := fmt.Sprintf("set http_proxy=http://127.0.0.1:%d && set https_proxy=http://127.0.0.1:%d", a.Port, a.Port)
		err := clipboard.WriteAll(command)
		if err != nil {
			logger.Log.Error(err, "write clipboard")
		}
	}
}

func (a *App) itemExportCommandPowerShell(menu *systray.MenuItem) {
	for range menu.ClickedCh {
		command := fmt.Sprintf("$env:http_proxy='http://127.0.0.1:%d'; $env:https_proxy='http://127.0.0.1:%d'; ", a.Port, a.Port)
		err := clipboard.WriteAll(command)
		if err != nil {
			logger.Log.Error(err, "write clipboard")
		}
	}
}