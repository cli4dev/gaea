package conf

const Login = `package member

import (
	"time"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/modules/member"
	"{{.projectName}}/modules/util"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//LoginHandler 用户登录对象
type LoginHandler struct {
	c component.IContainer
	m member.IUser
}

//NewLoginHandler 创建登录对象
func NewLoginHandler(container component.IContainer) (u *LoginHandler) {
	return &LoginHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 处理用户登录，登录成功后转跳到指定的系统
func (u *LoginHandler) Handle(ctx *context.Context) (r interface{}) {

	//检查输入参数
	if err := ctx.Request.Check("username", "password"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	//签名
	var raw string
	secret := app.GetConf(u.c).GetSecret()
	ctx.Log.Info("secret:", secret)
	m := make(db.QueryRow)
	m["username"] = ctx.Request.GetString("username")
	m["password"] = ctx.Request.GetString("password")
	m["ident"] = "{{getAppconf .login 3}}"
	m["timestamp"] = time.Now().Unix()
	raw, m["sign"] = util.MakeSign(m, secret)
	ctx.Log.Infof("请求登录数据：%v[%s]", m, raw)
	//处理用户登录
	member,code, err := u.m.Login(m)
	if err != nil {
		ctx.Response.SetPlain()
		return context.NewError(code, err)
	}
	//设置jwt数据
	ctx.Response.SetJWT(member)
	//返回用户名和角色信息
	return map[string]interface{}{
		"name": member.UserName,
		"role": member.RoleName,
	}

}`
const Info = `package member

import (
	"time"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/modules/member"
	"{{.projectName}}/modules/util"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//InfoHandler is
type InfoHandler struct {
	c component.IContainer
	m member.IUser
}

//NewInfoHandler is
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 获取菜单
func (u *InfoHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取系统信息----------")
	//签名
	var raw string
	secret := app.GetConf(u.c).GetSecret()
	m := make(db.QueryRow)
	m["ident"] = "{{getAppconf .login 3}}"
	m["timestamp"] = time.Now().Unix()
	raw, m["sign"] = util.MakeSign(m, secret)
	ctx.Log.Infof("请求系统信息数据：%v[%s]", m, raw)
	ctx.Log.Info("2. 执行操作")
	info, err := u.m.GetSysInfo(m)
	if err != nil {
		ctx.Response.SetStatus(context.ERR_FORBIDDEN)
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return info
}`

const Menu = `package member

import (
	"time"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/modules/member"
	"{{.projectName}}/modules/util"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//MenuHandler is
type MenuHandler struct {
	c component.IContainer
	m member.IUser
}

//NewMenuHandler is
func NewMenuHandler(container component.IContainer) (u *MenuHandler) {
	return &MenuHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 获取菜单
func (u *MenuHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("--------获取菜单----------")
	ctx.Log.Info("1. 获取参数")
	mem := member.Get(ctx)
	//签名
	var raw string
	secret := app.GetConf(u.c).GetSecret()
	m := make(db.QueryRow)
	m["user_id"] = mem.UserID
	m["system_id"] = mem.SystemID
	m["ident"] = mem.SysIdent
	m["timestamp"] = time.Now().Unix()
	raw, m["sign"] = util.MakeSign(m, secret)
	ctx.Log.Infof("请求菜单数据：%v[%s]", m, raw)
	ctx.Log.Info("2. 执行操作")
	Menus, err := u.m.GetMenu(m)
	if err != nil {
		ctx.Response.SetStatus(context.ERR_FORBIDDEN)
		return err
	}
	//设置jwt数据
	ctx.Log.Info("3. 返回数据")
	return Menus
}`

const ChangePwd = `package member

import (
	"fmt"
	"time"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/modules/member"
	"{{.projectName}}/modules/util"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//ChangePwdHandler 用户登录对象
type ChangePwdHandler struct {
	c component.IContainer
	m member.IUser
}

//NewChangePwdHandler 创建登录对象
func NewChangePwdHandler(container component.IContainer) (u *ChangePwdHandler) {
	return &ChangePwdHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle 处理用户修改密码
func (u *ChangePwdHandler) Handle(ctx *context.Context) (r interface{}) {

	//检查输入参数
	if err := ctx.Request.Check("password", "password_old"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	fmt.Println("password", ctx.Request.GetString("password"))
	fmt.Println("password_old", ctx.Request.GetString("password_old"))
	// if ctx.Request.GetString("password") != ctx.Request.GetString("password_old") {
	// 	return context.NewError(context.ERR_NOT_ACCEPTABLE, "两次密码不相同")
	// }
	//签名
	mem := member.Get(ctx)
	fmt.Println("mem", mem)
	var raw string
	secret := app.GetConf(u.c).GetSecret()
	ctx.Log.Info("secret:", secret)
	m := make(db.QueryRow)
	m["user_id"] = mem.UserID
	m["password"] = ctx.Request.GetString("password")
	m["password_old"] = ctx.Request.GetString("password_old")
	m["ident"] = mem.SysIdent
	m["timestamp"] = time.Now().Unix()
	raw, m["sign"] = util.MakeSign(m, secret)
	ctx.Log.Infof("请求修改密码数据：%v[%s]", m, raw)
	//处理用户登录
	err := u.m.ChangePwd(m)
	if err != nil {

		return err
	}
	return "success"
}`

const User = `package member

import (
	"time"
	"{{.projectName}}/modules/app"
	"{{.projectName}}/modules/member"
	"{{.projectName}}/modules/util"
	"github.com/micro-plat/lib4go/db"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

//UserHandler 用户登录对象
type UserHandler struct {
	c component.IContainer
	m member.IUser
}

//NewUserHandler 创建登录对象
func NewUserHandler(container component.IContainer) (u *UserHandler) {
	return &UserHandler{
		c: container,
		m: member.NewUser(container),
	}
}

//Handle .
func (u *UserHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取用户----------")
	ctx.Log.Info("1. 获取参数")
	mem := member.Get(ctx)
	//签名
	var raw string
	secret := app.GetConf(u.c).GetSecret()
	ctx.Log.Info("secret:", secret)
	m := make(db.QueryRow)
	m["ident"] = mem.SysIdent
	m["timestamp"] = time.Now().Unix()
	raw, m["sign"] = util.MakeSign(m, secret)
	ctx.Log.Infof("请求登录数据：%v[%s]", m, raw)
	//处理用户登录
	list, err := u.m.GetUser(m)
	if err != nil {
		return context.NewError(context.ERR_SERVER_ERROR, err)
	}
	return map[string]interface{}{"list": list["alluser"]}
}`

const AppModuleConf = `package app

import (
	'fmt'

	'github.com/asaskevich/govalidator'
	'github.com/micro-plat/hydra/component'
)

//Conf 应用程序配置
type Conf struct {
	SSOHost               string "json:'sso-host'"
	Secret                string "json:'secret'"
	Ident                 string "json:'ident'"
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//GetLoginURL 获取sso服务器登录地址
func (c *Conf) GetLoginURL() string {
	return c.SSOHost + '/subsys/login'
}

//GetLoginURL 获取sso服务器修改密码地址
func (c *Conf) GetChangePwdURL() string {
	return c.SSOHost + '/subsys/pwd'
}

//GetMenuURL 获取菜单请求地址
func (c *Conf) GetMenuURL() string {
	return c.SSOHost + '/subsys/menu'
}

//GetAreacompanyURL 获取地市公司请求地址
func (c *Conf) PostAreacompanyURL() string {
	return c.SSOHost + '/subsys/user'
}

//GetSysInfoURL 或取系统信息链接
func (c *Conf) GetSysInfoURL() string {
	return c.SSOHost + '/subsys/info'
}

//GetSecret 获取签名的secret
func (c *Conf) GetSecret() string {
	return c.Secret
}

//GetSystemName 获取系统名称
func (c *Conf) GetIdent() string {
	return c.Ident
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set('__AppConf__', m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get('__AppConf__').(*Conf)
}`
