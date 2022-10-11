package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
)

//调用方法主结构体
type UserInfo struct {
	ServerId string `json:"ServerId"`
}

//用户全部配置信息结构体
type UserSettingStruct struct {
	Webhook    map[string]string `json:"webhook"`
	LogList    map[string]string `json:"log_list"`
	FooterText string            `json:"FOOTER_TEXT"`
	FooterIcon string            `json:"FOOTER_ICON"`
	ColorCode  string            `json:"COLOR_CODE"`
	GroupName  string            `json:"GROUP_NAME"`
	UserEmail  string            `json:"User_Email"`
	UserStatus string            `json:"User_Status"`
}

//用户单其他设置结构体
type UserSpecialSettingStruct struct {
	FooterText string `json:"FOOTER_TEXT"`
	FooterIcon string `json:"FOOTER_ICON"`
	ColorCode  string `json:"COLOR_CODE"`
	GroupName  string `json:"GROUP_NAME"`
	UserEmail  string `json:"User_Email"`
	UserStatus string `json:"User_Status"`
}

type CustomerAllSettings map[string]UserSettingStruct

//系统Chekcout Channel List
type SystemCheckoutList struct {
	Checkout map[string]string `json:"Checkout"`
	Update   map[string]string `json:"Update"`
	Announce map[string]string `json:"Announce"`
	Release  map[string]string `json:"Release"`
	Special  map[string]string `json:"Special"`
}

//系统CheckoutBotName文件
type SystemCheckoutBotList struct {
	Checkout []string `json:"Checkout"`
	Update   []string `json:"Update"`
	Announce []string `json:"Announce"`
	Release  []string `json:"Release"`
	Special  []string `json:"Special"`
}

type ResWebhookBotList struct {
	Webhooks map[string]string `json:"Webhooks"`
	BotList  []string          `json:"BotList"`
}

//创建一个用户实例
func NewUserInfo(serverId string) UserInfo {
	return UserInfo{ServerId: serverId}
}

//获取用户所有的 Checkout webhook 和Bot List
func (u *UserInfo) GetUserAllWebhooks() (CustomerAllSettings, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return CustomerAllSettings{}, err
	}
	//返回数据
	return UserWebhookJson, nil
}

//获取用户所有的 Checkout webhook 和Bot List
func (u *UserInfo) GetUserCheckoutWebhooks() (ResWebhookBotList, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var ResWebhookAndBotNameList = ResWebhookBotList{}
	ResWebhookAndBotNameList.Webhooks = map[string]string{}

	//读取系统BotName比对文件 强制排序
	readSystemBotNameFile, _ := ioutil.ReadFile("./settings/systembotname.json")
	botNameMap := SystemCheckoutBotList{}
	errUnmarshal1 := jsonDo.Unmarshal(readSystemBotNameFile, &botNameMap)

	if errUnmarshal1 != nil {
		return ResWebhookBotList{}, errUnmarshal1
	}
	//读取系统BotName文件 强制排序
	for _, botName := range botNameMap.Checkout {
		ResWebhookAndBotNameList.BotList = append(ResWebhookAndBotNameList.BotList, botName)
	}
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return ResWebhookBotList{}, err
	}
	//读取系统比对webhook名字文件
	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	fmt.Println("ecaz = ", userWebhooksMap)
	if errUnmarshal != nil {
		return ResWebhookBotList{}, errUnmarshal
	}
	fmt.Println("dwdwdwdwdwd")
	for channelName, webhook := range UserWebhookJson[u.ServerId].Webhook {
		fmt.Println("webhook = ", webhook)
		ResWebhookAndBotNameList.Webhooks[userWebhooksMap.Checkout[channelName]] = webhook
	}
	//返回数据
	return ResWebhookAndBotNameList, nil
}
func (u *UserInfo) GetUserUpdateWebhooks() (ResWebhookBotList, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var ResWebhookAndBotNameList = ResWebhookBotList{}
	ResWebhookAndBotNameList.Webhooks = map[string]string{}

	//读取系统BotName比对文件 强制排序
	readSystemBotNameFile, _ := ioutil.ReadFile("./settings/systembotname.json")
	botNameMap := SystemCheckoutBotList{}
	errUnmarshal1 := jsonDo.Unmarshal(readSystemBotNameFile, &botNameMap)
	if errUnmarshal1 != nil {
		return ResWebhookBotList{}, errUnmarshal1
	}
	//读取系统BotName文件 强制排序
	for _, botName := range botNameMap.Update {
		ResWebhookAndBotNameList.BotList = append(ResWebhookAndBotNameList.BotList, botName)
	}
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return ResWebhookBotList{}, err
	}
	//读取系统比对webhook名字文件
	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return ResWebhookBotList{}, errUnmarshal
	}
	for channelName, webhook := range UserWebhookJson[u.ServerId].Webhook {

		ResWebhookAndBotNameList.Webhooks[userWebhooksMap.Update[channelName]] = webhook
	}
	//返回数据
	return ResWebhookAndBotNameList, nil
}
func (u *UserInfo) GetUserAnnounceWebhooks() (ResWebhookBotList, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var ResWebhookAndBotNameList = ResWebhookBotList{}
	ResWebhookAndBotNameList.Webhooks = make(map[string]string)

	//读取系统BotName比对文件 强制排序
	readSystemBotNameFile, _ := ioutil.ReadFile("./settings/systembotname.json")
	botNameMap := SystemCheckoutBotList{}
	errUnmarshal1 := jsonDo.Unmarshal(readSystemBotNameFile, &botNameMap)
	if errUnmarshal1 != nil {
		return ResWebhookBotList{}, errUnmarshal1
	}
	//读取系统BotName文件 强制排序
	for _, botName := range botNameMap.Announce {
		ResWebhookAndBotNameList.BotList = append(ResWebhookAndBotNameList.BotList, botName)
	}
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return ResWebhookBotList{}, err
	}
	//读取系统比对webhook名字文件
	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return ResWebhookBotList{}, errUnmarshal
	}
	for channelName, webhook := range UserWebhookJson[u.ServerId].Webhook {
		if channelName != "" && channelName != " " {
			ResWebhookAndBotNameList.Webhooks[userWebhooksMap.Announce[channelName]] = webhook
		}
	}
	//返回数据
	return ResWebhookAndBotNameList, nil
}
func (u *UserInfo) GetUserReleaseWebhooks() (ResWebhookBotList, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var ResWebhookAndBotNameList = ResWebhookBotList{}
	ResWebhookAndBotNameList.Webhooks = make(map[string]string)

	//读取系统BotName比对文件 强制排序
	readSystemBotNameFile, _ := ioutil.ReadFile("./settings/systembotname.json")
	botNameMap := SystemCheckoutBotList{}
	errUnmarshal1 := jsonDo.Unmarshal(readSystemBotNameFile, &botNameMap)
	if errUnmarshal1 != nil {
		return ResWebhookBotList{}, errUnmarshal1
	}
	//读取系统BotName文件 强制排序
	for _, botName := range botNameMap.Release {
		ResWebhookAndBotNameList.BotList = append(ResWebhookAndBotNameList.BotList, botName)
	}
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return ResWebhookBotList{}, err
	}
	//读取系统比对webhook名字文件
	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return ResWebhookBotList{}, errUnmarshal
	}
	for channelName, webhook := range UserWebhookJson[u.ServerId].Webhook {
		if channelName != "" && channelName != " " {
			ResWebhookAndBotNameList.Webhooks[userWebhooksMap.Release[channelName]] = webhook
		}
	}
	//返回数据
	return ResWebhookAndBotNameList, nil
}
func (u *UserInfo) GetUserSpecialWebhooks() (ResWebhookBotList, error) {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var ResWebhookAndBotNameList = ResWebhookBotList{}
	ResWebhookAndBotNameList.Webhooks = make(map[string]string)

	//读取系统BotName比对文件 强制排序
	readSystemBotNameFile, _ := ioutil.ReadFile("./settings/systembotname.json")
	botNameMap := SystemCheckoutBotList{}
	errUnmarshal1 := jsonDo.Unmarshal(readSystemBotNameFile, &botNameMap)
	if errUnmarshal1 != nil {
		return ResWebhookBotList{}, errUnmarshal1
	}
	//读取系统BotName文件 强制排序
	for _, botName := range botNameMap.Special {
		ResWebhookAndBotNameList.BotList = append(ResWebhookAndBotNameList.BotList, botName)
	}
	//读取用户Json文件
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	err := jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	if err != nil {
		return ResWebhookBotList{}, err
	}
	//读取系统比对webhook名字文件
	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return ResWebhookBotList{}, errUnmarshal
	}
	for channelName, webhook := range UserWebhookJson[u.ServerId].Webhook {
		if channelName != "" && channelName != " " {
			ResWebhookAndBotNameList.Webhooks[userWebhooksMap.Special[channelName]] = webhook
		}
	}
	//返回数据
	return ResWebhookAndBotNameList, nil
}

func (u *UserInfo) UpdateCustomerJsonFileCheckoutWebhook(webhooks map[string]string) bool {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var order = map[string]string{}

	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return false
	}
	for channel, channelInput := range userWebhooksMap.Checkout {
		order[channelInput] = channel
	}
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	for botName, webhook := range webhooks {
		UserWebhookJson[u.ServerId].Webhook[order[botName]] = webhook
	}
	d, _ := jsonDo.Marshal(UserWebhookJson)
	var writeWebhooks bytes.Buffer
	json.Indent(&writeWebhooks, d, "", "\t")
	err := ioutil.WriteFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId), writeWebhooks.Bytes(), 0777)
	if err != nil {
		return false
	}
	return true
}
func (u *UserInfo) UpdateCustomerJsonFileAnnounceWebhook(webhooks map[string]string) bool {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var order = map[string]string{}

	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return false
	}
	for channel, channelInput := range userWebhooksMap.Announce {
		order[channelInput] = channel
	}
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	for botName, webhook := range webhooks {
		UserWebhookJson[u.ServerId].Webhook[order[botName]] = webhook
	}
	d, _ := jsonDo.Marshal(UserWebhookJson)
	var writeWebhooks bytes.Buffer
	json.Indent(&writeWebhooks, d, "", "\t")
	err := ioutil.WriteFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId), writeWebhooks.Bytes(), 0777)
	if err != nil {
		return false
	}
	return true
}
func (u *UserInfo) UpdateCustomerJsonFileUpdateWebhook(webhooks map[string]string) bool {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var order = map[string]string{}

	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return false
	}
	for channel, channelInput := range userWebhooksMap.Update {
		order[channelInput] = channel
	}
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	for botName, webhook := range webhooks {
		UserWebhookJson[u.ServerId].Webhook[order[botName]] = webhook
	}
	d, _ := jsonDo.Marshal(UserWebhookJson)
	var writeWebhooks bytes.Buffer
	json.Indent(&writeWebhooks, d, "", "\t")
	err := ioutil.WriteFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId), writeWebhooks.Bytes(), 0777)
	if err != nil {
		return false
	}
	return true
}
func (u *UserInfo) UpdateCustomerJsonFileReleaseWebhook(webhooks map[string]string) bool {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var order = map[string]string{}

	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return false
	}
	for channel, channelInput := range userWebhooksMap.Release {
		order[channelInput] = channel
	}
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	for botName, webhook := range webhooks {
		UserWebhookJson[u.ServerId].Webhook[order[botName]] = webhook
	}
	d, _ := jsonDo.Marshal(UserWebhookJson)
	var writeWebhooks bytes.Buffer
	json.Indent(&writeWebhooks, d, "", "\t")
	err := ioutil.WriteFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId), writeWebhooks.Bytes(), 0777)
	if err != nil {
		return false
	}
	return true
}
func (u *UserInfo) UpdateCustomerJsonFilesSpecialWebhook(webhooks map[string]string) bool {
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	var order = map[string]string{}

	readSystemWebhookNameFile, _ := ioutil.ReadFile("./settings/systemsettings.json")
	userWebhooksMap := SystemCheckoutList{}
	errUnmarshal := jsonDo.Unmarshal(readSystemWebhookNameFile, &userWebhooksMap)
	if errUnmarshal != nil {
		return false
	}
	for channel, channelInput := range userWebhooksMap.Special {
		order[channelInput] = channel
	}
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserWebhookJson := CustomerAllSettings{}
	jsonDo.Unmarshal(readUserFile, &UserWebhookJson)
	for botName, webhook := range webhooks {
		UserWebhookJson[u.ServerId].Webhook[order[botName]] = webhook
	}
	d, _ := jsonDo.Marshal(UserWebhookJson)
	var writeWebhooks bytes.Buffer
	json.Indent(&writeWebhooks, d, "", "\t")
	err := ioutil.WriteFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId), writeWebhooks.Bytes(), 0777)
	if err != nil {
		return false
	}
	return true
}

//获取用户所有其他设置
func (u *UserInfo) GetUserSpecialSettings() (map[string]UserSpecialSettingStruct, error) {
	readUserFile, _ := ioutil.ReadFile(fmt.Sprintf("/Customer/webhook/%s.json", u.ServerId))
	UserSettingJson := map[string]UserSpecialSettingStruct{}
	var jsonDo = jsoniter.ConfigCompatibleWithStandardLibrary
	err := jsonDo.Unmarshal(readUserFile, &UserSettingJson)
	if err != nil {
		return map[string]UserSpecialSettingStruct{}, err
	}
	return UserSettingJson, nil
}

//写入Checkout Webhook到文件中

//func main() {
//	a := NewUserInfo("775822627133456424")
//	file, err := a.GetUserCheckoutWebhooks()
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		for n, vlue := range file {
//			fmt.Println(n," : ",vlue)
//		}
//
//	}
//}