package menu

import (
	"fmt"
	"strings"

	"github.com/Ledgercn/ConsoleColor"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

const (
	AIOQ_INFO_INPUT_ADDR = iota
	AIOQ_INFO_QUARY_WAITING
	AIOQ_INFO_ADDR_FORMAT_ERR
	AIOQ_INFO_NEXT_RECORDS
)

// AccountInfoOperationQuary 账户操作查询
type AccountInfoOperationQuary struct {
	MenuSubItem
	infoStrings []map[int]string
}

// InitAccInfoOperQuary 初始化
func (ths *AccountInfoOperationQuary) InitAccInfoOperQuary(parent MenuSubItemInterface, key string) {
	ths.MenuSubItem.InitMenu(key)
	ths.parentItem = parent
	ths.MenuSubItem.Exec = ths.execute

	ths.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户操作查询",
		publicdefine.L_English: "Account Operations Quary",
	}

	ths.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			AIOQ_INFO_INPUT_ADDR:      " 请输入要查询的账户地址\r\n > ",
			AIOQ_INFO_QUARY_WAITING:   " 正在查询请稍后...",
			AIOQ_INFO_ADDR_FORMAT_ERR: "\r\n ** 输入的地址无效 [%s]\r\n",
			AIOQ_INFO_NEXT_RECORDS:    " > 查看下10条操作，请输入n回车，结束请输入回车: ",
		},
		publicdefine.L_English: map[int]string{
			AIOQ_INFO_INPUT_ADDR:      " Please input the account address you want to query\r\n > ",
			AIOQ_INFO_QUARY_WAITING:   " Searching for a query...",
			AIOQ_INFO_ADDR_FORMAT_ERR: "\r\n ** Stellar address is invalid [%s]\r\n",
			AIOQ_INFO_NEXT_RECORDS:    " > Quary next 10 operations, input n + enter, or press enter to end quary: ",
		},
	}

}

func (ths *AccountInfoOperationQuary) execute(isSync bool) {
	addr := ths.inputAddr()
	if len(addr) > 0 && publicdefine.VerifyGAddress(addr) == nil {
		ths.quary(addr)
	} else {
		ConsoleColor.Printf(ConsoleColor.C_RED,
			"\r\n"+ths.infoStrings[ths.languageIndex][AIOQ_INFO_ADDR_FORMAT_ERR]+"\r\n\r\n", addr)
	}
	if !isSync {
		ths.ASyncChan <- 0
	}
}

func (ths *AccountInfoOperationQuary) inputAddr() string {
	fmt.Printf(ths.infoStrings[ths.languageIndex][AIOQ_INFO_INPUT_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		return strings.Trim(input, " ")
	}
	return ""
}

func (ths *AccountInfoOperationQuary) quary(addr string) {
	result := &publicdefine.StellarAccOperationQuary{
		QuaryCursor: "",
		IsEnd:       false,
	}
	for {
		ConsoleColor.Print(ConsoleColor.C_BLUE,
			ths.infoStrings[ths.languageIndex][AIOQ_INFO_QUARY_WAITING])

		reqUrl := fmt.Sprintf("%s%s/%s/%s?order=desc&limit=%d&cursor=%s", publicdefine.STELLAR_DEFAULT_NETWORK,
			publicdefine.STELLAR_NETWORK_ACCOUNTS, addr,
			publicdefine.STELLAR_NETWORK_OPERATIONS, 10, result.QuaryCursor)
		resMap, err := ths.httpGet(reqUrl)

		if err == nil {
			result.PutMapBody(resMap)
			ths.PrintResult(result)
		} else {
			ConsoleColor.Println(ConsoleColor.C_RED, err)
			break
		}

		if result.IsEnd {
			ConsoleColor.Println(ConsoleColor.C_BLUE, "\r\n End\r\n")
			break
		}
		if ths.inputNext() != "n" {
			break
		}
	}
}

// PrintResult 打印结果
func (ths *AccountInfoOperationQuary) PrintResult(r *publicdefine.StellarAccOperationQuary) {
	fmt.Print("\r")
	for i := 0; i < len(r.Records); i++ {
		ConsoleColor.Printf(ConsoleColor.C_GREEN,
			" %02d %s\r\n", i+1, strings.Repeat("-", 80))

		ConsoleColor.Printf(ConsoleColor.C_BLUE, "%s\r\n\r\n", r.Records[i].ToString())
	}
}

func (ths *AccountInfoOperationQuary) inputNext() string {
	fmt.Printf(ths.infoStrings[ths.languageIndex][AIOQ_INFO_NEXT_RECORDS])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		return strings.Trim(input, " ")
	}
	return ""
}
