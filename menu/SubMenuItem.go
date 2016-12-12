package menu

import (
	"fmt"

	"github.com/jojopoper/go-StellarWallet/publicdefine"
	"golang.org/x/net/proxy"
)

// MenuSubItem 子菜单定义
type MenuSubItem struct {
	title         []string
	Exec          func(isSync bool)
	subItems      []MenuSubItemInterface
	parentItem    MenuSubItemInterface
	keyPath       string
	ASyncChan     chan int
	languageIndex int
	inputMemo     []string
}

// MenuSubItemInterface 子菜单接口定义
type MenuSubItemInterface interface {
	InitMenu(key string)
	SetLanguageType(langType int)
	GetTitle(langType int) string
	HasTitle() bool
	SetTitle(langType int, t string)
	GetSubItems() []MenuSubItemInterface
	AddSubItem(itm MenuSubItemInterface) int
	GetParentItem() MenuSubItemInterface
	SetParentItem(p MenuSubItemInterface)
	GetKeyPath() string
	SetKeyPath(kp string)
	GetTitlePath(langType int) string
	ExecuteFunc(isSync bool)
	ExecFlag() int
	PrintSubmenu()
	GetInputMemo(langType int) string
	httpGet(geturl string) (map[string]interface{}, error)
	httppostForm(address, data string) (map[string]interface{}, error)
}

// InitMenu 初始化
func (ths *MenuSubItem) InitMenu(key string) {
	ths.ASyncChan = make(chan int)
	ths.title = make([]string, 2)
	ths.subItems = make([]MenuSubItemInterface, 0)
	ths.keyPath = key
	ths.inputMemo = []string{
		publicdefine.L_Chinese: "请选择菜单列表项目对应的数字并回车: ",
		publicdefine.L_English: "Select the number of items on the menu list item and press enter: ",
	}
}

// SetLanguageType 设置语言
func (ths *MenuSubItem) SetLanguageType(langType int) {
	ths.languageIndex = langType
	for _, sub := range ths.subItems {
		sub.SetLanguageType(langType)
	}
}

// GetTitle 获取标题
func (ths *MenuSubItem) GetTitle(langType int) string {
	if ths.HasTitle() {
		return ths.title[langType]
	}
	return ""
}

// HasTitle 获取是否含有标题
func (ths *MenuSubItem) HasTitle() bool {
	return len(ths.title) > 0
}

// SetTitle 设置标题
func (ths *MenuSubItem) SetTitle(langType int, t string) {
	ths.title[langType] = t
}

// GetSubItems 获取子菜单
func (ths *MenuSubItem) GetSubItems() []MenuSubItemInterface {
	return ths.subItems
}

// AddSubItem 添加子菜单
func (ths *MenuSubItem) AddSubItem(itm MenuSubItemInterface) int {
	length := len(ths.subItems)
	itm.SetParentItem(ths)
	itm.SetKeyPath(fmt.Sprintf("%s.%d", ths.keyPath, length))
	ths.subItems = append(ths.subItems, itm)
	return length
}

// GetParentItem 得到父菜单
func (ths *MenuSubItem) GetParentItem() MenuSubItemInterface {
	return ths.parentItem
}

// SetParentItem 设置父菜单
func (ths *MenuSubItem) SetParentItem(p MenuSubItemInterface) {
	ths.parentItem = p
}

// GetKeyPath 获取路径
func (ths *MenuSubItem) GetKeyPath() string {
	return ths.keyPath
}

// SetKeyPath 设置路径
func (ths *MenuSubItem) SetKeyPath(kp string) {
	ths.keyPath = kp
}

// GetTitlePath 获取标题路径
func (ths *MenuSubItem) GetTitlePath(langType int) (ret string) {
	if ths.parentItem == nil {
		ret = ths.title[langType]
	} else {
		ret = ths.parentItem.GetTitlePath(langType) + " > " + ths.title[langType]
	}
	return ret
}

// ExecuteFunc 执行
func (ths *MenuSubItem) ExecuteFunc(isSync bool) {
	if ths.Exec != nil {
		if isSync {
			ths.Exec(isSync)
		} else {
			go ths.Exec(isSync)
		}
	}
}

// ExecFlag 执行标识
func (ths *MenuSubItem) ExecFlag() int {
	return <-ths.ASyncChan
}

// PrintSubmenu 打印子菜单
func (ths *MenuSubItem) PrintSubmenu() {
	length := len(ths.subItems)
	for i := 0; i < length; i++ {
		fmt.Printf(" %d.\t%s\r\n", i+1, ths.subItems[i].GetTitle(ths.languageIndex))
	}
}

// GetInputMemo 得到说明
func (ths *MenuSubItem) GetInputMemo(langType int) string {
	return ths.inputMemo[langType]
}

func (ths *MenuSubItem) httpGet(geturl string) (map[string]interface{}, error) {
	if publicdefine.CurrProxyInfo == nil || publicdefine.CurrProxyInfo.Enabled == false {
		return publicdefine.HttpGet(geturl)
	} else if len(publicdefine.CurrProxyInfo.UserName) == 0 {
		return publicdefine.HttpProxyGet(geturl, publicdefine.CurrProxyInfo.IP,
			publicdefine.CurrProxyInfo.Port)
	}

	return publicdefine.HttpProxyGet(geturl, publicdefine.CurrProxyInfo.IP,
		publicdefine.CurrProxyInfo.Port,
		&proxy.Auth{
			User:     publicdefine.CurrProxyInfo.UserName,
			Password: publicdefine.CurrProxyInfo.Password,
		})
}

func (ths *MenuSubItem) httppostForm(address, data string) (map[string]interface{}, error) {
	if publicdefine.CurrProxyInfo == nil || publicdefine.CurrProxyInfo.Enabled == false {
		return publicdefine.HttpPostForm(address, data)
	} else if len(publicdefine.CurrProxyInfo.UserName) == 0 {
		return publicdefine.HttpProxyPostForm(address, data, publicdefine.CurrProxyInfo.IP,
			publicdefine.CurrProxyInfo.Port)
	}

	return publicdefine.HttpProxyPostForm(address, data, publicdefine.CurrProxyInfo.IP,
		publicdefine.CurrProxyInfo.Port,
		&proxy.Auth{
			User:     publicdefine.CurrProxyInfo.UserName,
			Password: publicdefine.CurrProxyInfo.Password,
		})
}
