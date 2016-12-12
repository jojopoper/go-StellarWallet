package menu

import (
	"fmt"
	"os"
	"time"

	"github.com/Ledgercn/ConsoleColor"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
	"github.com/stellar/go/keypair"
)

const (
	CA_INFO_SECRET_SEED = iota
	CA_INFO_PUBLIC_ADDR
	CA_INFO_MEMO_TEXT
	CA_INFO_MEMO_SAVEFILE
	CA_INFO_MEMO_SAVEFILE_ERR
)

// CreateAccount 创建账户
type CreateAccount struct {
	MenuSubItem
	infoStrings []map[int]string
}

// InitCreator 初始化
func (ths *CreateAccount) InitCreator(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "创建账户",
		publicdefine.L_English: "Create new account",
	}
	ths.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			CA_INFO_SECRET_SEED:       " Secret seed:",
			CA_INFO_PUBLIC_ADDR:       " Public:",
			CA_INFO_MEMO_TEXT:         " 需要保存账户信息到文件请输入s，否则输入任意键返回菜单: ",
			CA_INFO_MEMO_SAVEFILE:     " 保存账户信息完成，请妥善保存好文件 %s",
			CA_INFO_MEMO_SAVEFILE_ERR: " 保存账户信息失败，错误信息: ",
		},
		publicdefine.L_English: map[int]string{
			CA_INFO_SECRET_SEED:       " Secret seed:",
			CA_INFO_PUBLIC_ADDR:       " Public:",
			CA_INFO_MEMO_TEXT:         " If you need to save account informations to a file then press s, or press any key to return menu: ",
			CA_INFO_MEMO_SAVEFILE:     " Save account information is complete，please keep safe the file : %s",
			CA_INFO_MEMO_SAVEFILE_ERR: " Save account information is failure, the error message: ",
		},
	}
}

func (ths *CreateAccount) execute(isSync bool) {
	var input string

	keyp, err := keypair.Random()

	if err == nil {
		fmt.Printf("\r\n"+ths.infoStrings[ths.languageIndex][CA_INFO_SECRET_SEED]+" %s\r\n", keyp.Seed())
		fmt.Printf(ths.infoStrings[ths.languageIndex][CA_INFO_PUBLIC_ADDR]+" %s\r\n", keyp.Address())
		fmt.Printf("\r\n" + ths.infoStrings[ths.languageIndex][CA_INFO_MEMO_TEXT])
		fmt.Scanf("%s\n", &input)

		if input == "s" {
			err = ths.savefile("account_info.txt", keyp.Seed(), keyp.Address())
			if err == nil {
				ConsoleColor.Printf(ConsoleColor.C_YELLOW,
					"\r\n"+ths.infoStrings[ths.languageIndex][CA_INFO_MEMO_SAVEFILE]+"\r\n\r\n", "account_info.txt")
			} else {
				ConsoleColor.Printf(ConsoleColor.C_RED,
					"\r\n%s\r\n%v\r\n\r\n", ths.infoStrings[ths.languageIndex][CA_INFO_MEMO_SAVEFILE_ERR], err)
			}
		}
	} else {
		fmt.Println(err.Error())
		fmt.Scanf("%s\n", &input)
	}

	if !isSync {
		ths.ASyncChan <- 0
	}
}

func (ths *CreateAccount) savefile(filepath, seed, addr string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeType)
	if err != nil {
		return err
	}
	defer f.Close()
	sav := fmt.Sprintf("\r\n================= %s =================\r\n", time.Now().Format("2006-01-02 15:04:05"))
	sav += fmt.Sprintf(ths.infoStrings[ths.languageIndex][CA_INFO_SECRET_SEED]+" %s\r\n", seed)
	sav += fmt.Sprintf(ths.infoStrings[ths.languageIndex][CA_INFO_PUBLIC_ADDR]+" %s\r\n\r\n", addr)
	_, err = f.WriteString(sav)
	return err
}
