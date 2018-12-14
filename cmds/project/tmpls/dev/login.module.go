package dev

const AppModuleMemberMenu = `package member

type Menu struct {
	ID       string "json:'id'"
	Name     string "json:'name'"
	Level    string "json:'level_id'"
	IsOpen   string "json:'is_open'"
	Icon     string "json:'icon'"
	SystemID string "json:'sys_id'"
	Parent   string "json:'parent'"
	Path     string "json:'path'"
	Sortrank string "json:'sortrank'"
	Children []Menu "json:'children,omitempty'"
}
`

const AppModuleMemberState = `package member

import 'github.com/micro-plat/hydra/context'

//LoginState 用户信息
type LoginState struct {
	UserID         int    "json:'user_id' m2s:'user_id'"
	UserName       string "json:'user_name' m2s:'user_name'"
	RoleName       string "json:'role_name' m2s:'role_name'"
	SystemID       int    "json:'sys_id' "
	SysIdent       string "json:'ident' "
	RoleID         int    "json:'role_id'"
	Status         int    "json:'status' m2s:'status'"
	IndexURL       string "json:'index_url'"
	Code           string "json:'code'"
	ProfilePercent int    "json:'profile_percent'"
	LoginTimeout   int    "json:'login_timeout' m2s:'login_timeout'"
	Password       string "json:'password.omitempty' "
}

//Save 保存member信息
func Save(ctx *context.Context, m *LoginState) error {
	ctx.Meta.Set('login-state', m)
	return nil
}

//Get 获取member信息
func Get(ctx *context.Context) *LoginState {
	v, _ := ctx.Meta.Get('login-state')
	if v == nil {
		return nil
	}
	return v.(*LoginState)
}`

const AppModuleMemberUser = `package member

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"{{.projectName}}/modules/app"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
)

type IUser interface {
	Login(m db.QueryRow) (*LoginState, int, error)
	ChangePwd(m db.QueryRow) error
	GetMenu(m db.QueryRow) ([]*Menu, error)
	GetUser(m db.QueryRow) (map[string]interface{}, error)
	GetSysInfo(m db.QueryRow) (map[string]interface{}, error)
}

type User struct {
	c     component.IContainer
	http  *http.Client
}

func NewUser(c component.IContainer) *User {
	return &User{
		c:     c,
		http:  &http.Client{},
	}
}

//Login 远程登录
func (u *User) Login(m db.QueryRow) (*LoginState, int, error) {
	state, code, err := u.remoteLogin(m)
	return state, code, err
}

func (u *User) remoteLogin(m db.QueryRow) (*LoginState, int, error) {

	url := fmt.Sprintf("%s?username=%s&password=%s&ident=%s&sign=%s&timestamp=%s",
		app.GetConf(u.c).GetLoginURL(), m.GetString("username"),
		m.GetString("password"), m.GetString("ident"), m.GetString("sign"), m.GetString("timestamp"))
	fmt.Println("请求url:", url)
	resp, err := u.http.Get(url)
	if err != nil {
		return nil, 400, fmt.Errorf("请求:%v失败(%v)", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 400, fmt.Errorf("%s", string(body))
	}
	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, fmt.Errorf("%s", string(body))
	}
	var state LoginState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, 400, fmt.Errorf("解析返回结果失败 %s：%v(%s)", url, err, string(body))
	}

	return &state, resp.StatusCode, nil
}

//GetMenu is
func (u *User) GetMenu(m db.QueryRow) ([]*Menu, error) {
	return u.remoteMenuQuery(m)
}

func (u *User) remoteMenuQuery(m db.QueryRow) ([]*Menu, error) {
	// "user_id", "system_id", "ident", "timestamp", "sign"
	url := fmt.Sprintf("%s?user_id=%s&system_id=%s&ident=%s&timestamp=%s&sign=%s",
		app.GetConf(u.c).GetMenuURL(), m.GetString("user_id"),
		m.GetString("system_id"), m.GetString("ident"), m.GetString("timestamp"), m.GetString("sign"))
	fmt.Println("请求url:", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}
	var menus []*Menu
	err = json.Unmarshal(body, &menus)
	if err != nil {
		return nil, fmt.Errorf("解析返回结果失败 %s：%v(%s)", url, err, string(body))
	}
	return menus, nil
}

//changePwd
func (u *User) ChangePwd(m db.QueryRow) error {
	_, _, err := u.remoteChangePwd(m)
	return err
}

func (u *User) remoteChangePwd(m db.QueryRow) (http.Header, []*http.Cookie, error) {

	url := fmt.Sprintf("%s?user_id=%s&password=%s&password_old=%s&sign=%s&timestamp=%s&ident=%s",
		app.GetConf(u.c).GetChangePwdURL(), m.GetString("user_id"),
		m.GetString("password"), m.GetString("password_old"), m.GetString("sign"), m.GetString("timestamp"), m.GetString("ident"))
	resp, err := u.http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("登录失败,HttpStatus:%d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}
	d := map[string]string{}
	err = json.Unmarshal(body, &d)
	if !strings.EqualFold(d["data"], "success") {
		return nil, nil, fmt.Errorf("修改密码失败 %s：%v(%s)", url, err, string(body))
	}

	return resp.Header, resp.Cookies(), nil
}

func (u *User) GetUser(m db.QueryRow) (map[string]interface{}, error) {
	data, err := u.remoteGetUser(m)
	return data, err
}

func (u *User) remoteGetUser(m db.QueryRow) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s?&sign=%s&timestamp=%s&ident=%s",
		app.GetConf(u.c).PostAreacompanyURL(), m.GetString("sign"), m.GetString("timestamp"), m.GetString("ident"))
	// resp, err := u.http.NewRequest("Post", url, nil)
	// if err != nil {
	// 	return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	// }

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("登录失败,HttpStatus:%d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}
	// d := map[string]string{}
	// err = json.Unmarshal(body, &d)
	// if !strings.EqualFold(d["data"], "success") {
	// 	return nil, fmt.Errorf("获取用户失败 %s：%v(%s)", url, err, string(body))
	// }

	listByte := []byte(body)
	var returnList map[string]interface{}
	err = json.Unmarshal(listByte, &returnList)
	if err != nil {
		return nil, fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	return returnList, nil
}

//GetSysInfo .
func (u *User) GetSysInfo(m db.QueryRow) (map[string]interface{}, error) {
	data, err := u.remoteGetSysInfo(m)
	return data, err
}

func (u *User) remoteGetSysInfo(m db.QueryRow) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s?&sign=%s&timestamp=%s&ident=%s",
		app.GetConf(u.c).GetSysInfoURL(), m.GetString("sign"), m.GetString("timestamp"), m.GetString("ident"))

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.http.Do(request)
	if err != nil {
		return nil, fmt.Errorf("请求:%v失败(%v)", url, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("登录失败,HttpStatus:%d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取远程数据失败 %s %v", url, err)
	}

	listByte := []byte(body)
	var returnList map[string]interface{}
	err = json.Unmarshal(listByte, &returnList)
	if err != nil {
		return nil, fmt.Errorf("字符串转json发生错误，err：%v", err)
	}

	return returnList, nil
}`

const Util = `package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

//MakeSign 远程登录，生成签名
func MakeSign(mReq map[string]interface{}, secret string) (raw string, sign string) {
	fmt.Println("签名计算, API KEY:", secret)
	//1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	//2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//3, 在键值对的最后加上key=secret
	if secret != "" {
		signStrings = signStrings + "key=" + secret
	}

	//4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return signStrings, upperSign
}`
