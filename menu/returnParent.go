package menu

import (
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

const (
	BACK_TO_MENU_FLAG = 99999
)

// ReturnParentMenu 返回上一级
type ReturnParentMenu struct {
	MenuSubItem
	infoStrings []map[int]string
}

// InitReturnParentMenu 初始化
func (ths *ReturnParentMenu) InitReturnParentMenu(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "返回上一级",
		publicdefine.L_English: "Go back",
	}
}

func (ths *ReturnParentMenu) execute(isSync bool) {
	if !isSync {
		ths.ASyncChan <- BACK_TO_MENU_FLAG
	}
}
