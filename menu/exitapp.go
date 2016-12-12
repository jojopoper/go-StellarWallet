package menu

import (
	"os"

	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

// ExitApp 退出钱包程序
type ExitApp struct {
	MenuSubItem
}

// InitExitApp 初始化
func (ths *ExitApp) InitExitApp(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "退出钱包程序",
		publicdefine.L_English: "Exit",
	}
}

func (ths *ExitApp) execute(isSync bool) {
	os.Exit(0)
}
