package menu

import (
	"fmt"

	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

// MenuInfo 菜单
type MenuInfo struct {
	MenuSubItem
	currentLevel  int
	WelcomeString []string
}

// MainMenuInstace 菜单唯一实例
var MainMenuInstace *MenuInfo

func init() {
	MainMenuInstace = new(MenuInfo)
	MainMenuInstace.MenuSubItem.InitMenu("0")
	MainMenuInstace.currentLevel = 0
	MainMenuInstace.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "菜单",
		publicdefine.L_English: "Menu",
	}
	MainMenuInstace.WelcomeString = []string{
		publicdefine.L_Chinese: " ##  欢迎使用 恒星币钱包，请选择您需要的功能  ##\n",
		publicdefine.L_English: " ##   Welcome to use the stellar wallet   ##\n" +
			" ##      please choose the function you need       ##\n",
	}
	MainMenuInstace.MenuSubItem.Exec = MainMenuInstace.Execute

	// locAcc := LocalAccount{}
	// locAcc.InitLocalAccount(MainMenuInstace, "0")
	// MainMenuInstace.MenuSubItem.AddSubItem(&locAcc)

	creatAcc := CreateAccount{}
	creatAcc.InitCreator(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&creatAcc)

	accInfo := AccountInfo{}
	accInfo.InitAccInfo(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&accInfo)

	mergeAcc := MergeAccount{}
	mergeAcc.InitMerge(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&mergeAcc)

	about := SoftwareAbout{}
	about.InitAbout(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&about)

	exitapp := ExitApp{}
	exitapp.InitExitApp(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&exitapp)
}

// Execute 执行函数
func (ths *MenuInfo) Execute(isSync bool) {
	for {
		fmt.Println("\r\n******************************************************")
		fmt.Println(ths.getWelcomeString(ths.languageIndex))
		fmt.Println(" " + ths.GetTitle(ths.languageIndex))
		ths.PrintSubmenu()
		fmt.Printf("\n %s", ths.GetInputMemo(ths.languageIndex))

		var input string

		_, err := fmt.Scanf("%s\n", &input)
		if err == nil {
			selectIndex, b := publicdefine.IsNumber(input)
			if b {
				if selectIndex <= len(ths.subItems) && selectIndex > 0 {
					ths.subItems[selectIndex-1].ExecuteFunc(false)
					ths.subItems[selectIndex-1].ExecFlag()
				}
			}
		}
	}
}

func (ths *MenuInfo) getWelcomeString(langType int) string {
	return ths.WelcomeString[langType]
}
