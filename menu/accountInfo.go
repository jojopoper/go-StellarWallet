package menu

import (
	"fmt"

	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

// AccountInfo 账户信息
type AccountInfo struct {
	MenuSubItem
	infoStrings []map[int]string
}

// InitAccInfo 初始化
func (ths *AccountInfo) InitAccInfo(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户信息",
		publicdefine.L_English: "Account Informations",
	}

	ths.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{},
		publicdefine.L_English: map[int]string{},
	}

	baseinfo := AccountInfoBase{}
	baseinfo.InitAccInfoBase(ths, "1")
	ths.AddSubItem(&baseinfo)

	payinfo := AccountInfoPayment{}
	payinfo.InitAccInfoPayment(ths, "1")
	ths.AddSubItem(&payinfo)

	operquary := AccountInfoOperationQuary{}
	operquary.InitAccInfoOperQuary(ths, "1")
	ths.AddSubItem(&operquary)

	returnParent := ReturnParentMenu{}
	returnParent.InitReturnParentMenu(ths, "1")
	ths.AddSubItem(&returnParent)

	exitapp := ExitApp{}
	exitapp.InitExitApp(ths, "1")
	ths.AddSubItem(&exitapp)
}

func (ths *AccountInfo) execute(isSync bool) {
	for {
		fmt.Printf("\n\n%s\r\n\n", ths.GetTitlePath(ths.languageIndex))
		ths.PrintSubmenu()
		fmt.Printf("\n %s", ths.GetInputMemo(ths.languageIndex))

		var input string

		_, err := fmt.Scanf("%s\n", &input)
		if err == nil {
			selectIndex, b := publicdefine.IsNumber(input)
			if b {
				if selectIndex <= len(ths.subItems) && selectIndex >= 0 {
					ths.subItems[selectIndex-1].ExecuteFunc(false)
					ret := ths.subItems[selectIndex-1].ExecFlag()
					if ret == BACK_TO_MENU_FLAG {
						break
					}
				}
			}
		} else {
			fmt.Println(err)
		}
	}
	if !isSync {
		ths.ASyncChan <- 1
	}
}
