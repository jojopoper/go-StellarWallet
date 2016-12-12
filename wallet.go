package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Ledgercn/ConsoleColor"
	"github.com/jojopoper/go-StellarWallet/menu"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

func main() {
	// publicdefine.CurrProxyInfo = &publicdefine.UseProxyInfo{}
	// publicdefine.CurrProxyInfo.Enabled = false
	// _, filename, _, _ := runtime.Caller(1)
	// publicdefine.CURRENT_DIR = path.Dir(filename)

	setCurrentArgs(os.Args)

	file, _ := exec.LookPath(os.Args[0])
	fpath, _ := filepath.Abs(file)
	publicdefine.CURRENT_DIR = path.Dir(fpath)

	currLanguage, errMsg := selectLanguage(3)
	if errMsg == nil {
		menu.MainMenuInstace.SetLanguageType(currLanguage)
		networkTitle := []string{
			publicdefine.L_Chinese: " 当前连接网络为 : ",
			publicdefine.L_English: " Current connecting network is : ",
		}
		ConsoleColor.Println(ConsoleColor.C_YELLOW,
			"\r\n", networkTitle[currLanguage], publicdefine.GetDefaultNWString(), "\r\n")
		menu.MainMenuInstace.ExecuteFunc(true)
	}
}

func setCurrentArgs(args []string) {
	lenArgs := len(args)
	for i := 1; i < lenArgs; i++ {
		switch strings.ToLower(args[i]) {
		case "--live":
			publicdefine.STELLAR_DEFAULT_NETWORK = publicdefine.STELLAR_LIVE_NETWORK
		case "--test":
			publicdefine.STELLAR_DEFAULT_NETWORK = publicdefine.STELLAR_TEST_NETWORK
		case "--proxy":
			i += setCurrentProxy(args[i:])
		}
	}
}

// format is --proxy "IP;PORT;UserName;Password"
func setCurrentProxy(args []string) int {
	// fmt.Println(args)
	// fmt.Println(len(args))
	if args == nil || len(args) < 2 {
		return 0
	}
	tmps := strings.Split(args[1], ";")
	if len(tmps) != 4 {
		return 0
	}
	publicdefine.CurrProxyInfo.IP = tmps[0]
	publicdefine.CurrProxyInfo.Port = tmps[1]
	publicdefine.CurrProxyInfo.UserName = tmps[2]
	publicdefine.CurrProxyInfo.Password = tmps[3]
	publicdefine.CurrProxyInfo.Enabled = true
	// fmt.Println(*publicdefine.CurrProxyInfo)
	return 1
}

func selectLanguage(maxRetry int) (int, error) {
	for i := 0; i < maxRetry; i++ {
		fmt.Printf("选择语言(select language):\r\n中文输入 %d 回车，For English press %d + Enter: ",
			publicdefine.L_Chinese+1, publicdefine.L_English+1)
		var input string
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Println(err)
			fmt.Scanf("%s\n", &input)
			return -1, err
		}

		switch input {
		case strconv.Itoa(publicdefine.L_Chinese + 1):
			return publicdefine.L_Chinese, nil
		case strconv.Itoa(publicdefine.L_English + 1):
			return publicdefine.L_English, nil
		default:
			ConsoleColor.Println(ConsoleColor.C_RED, "语言选择错误(Language selection error)\r\n")
			// fmt.Println("语言选择错误(Language selection error)\r\n")
		}
	}
	ConsoleColor.Println(ConsoleColor.C_YELLOW, "ByeBye!\r\n")
	return -1, errors.New("Language selection error")
}
