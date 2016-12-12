package menu

import (
	"fmt"

	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

const (
	SA_INFO_MEMO = iota
)

// SoftwareAbout 关于
type SoftwareAbout struct {
	MenuSubItem
	infoStrings []map[int]string
}

// InitAbout 初始化
func (ths *SoftwareAbout) InitAbout(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "关于",
		publicdefine.L_English: "About",
	}

	ths.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			SA_INFO_MEMO: "   软件版本 : 1.0.0.20161202\r\n" +
				"   钱包源码 : https://www.github.com/jojopoper/go-StellarWallet\r\n" +
				" 我们的QQ群 : 452779719\r\n" +
				" 支持和打赏 : GBP23RS4PCE73ZBUBXZAKZIRUFQNJIAYKB27KQC6BVRMVMJJERLNHUAZ\r\n",
		},
		publicdefine.L_English: map[int]string{
			SA_INFO_MEMO: "     Wallet Version : 1.0.0.20151118\r\n" +
				"        Source code : https://www.github.com/jojopoper/go-StellarWallet\r\n" +
				"       Our QQ group : 452779719\r\n" +
				" Support and reward : GBP23RS4PCE73ZBUBXZAKZIRUFQNJIAYKB27KQC6BVRMVMJJERLNHUAZ\r\n",
		},
	}

}

func (ths *SoftwareAbout) execute(isSync bool) {
	fmt.Println("")
	fmt.Println(ths.infoStrings[ths.languageIndex][SA_INFO_MEMO])
	fmt.Println("")

	var input string

	fmt.Scanf("%s\n", &input)

	if !isSync {
		ths.ASyncChan <- 0
	}
}
