package key

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/micro-plat/gaea/cmds"
	"github.com/urfave/cli"
)

type keyCmd struct {
}

//NewKeyCmd .
func NewKeyCmd() cli.Command {
	p := &keyCmd{}
	return cli.Command{
		Name:   "key",
		Usage:  "生成 jwt 加密所需的 key",
		Flags:  p.geStartFlags(),
		Action: p.action,
	}
}

func (p *keyCmd) geStartFlags() []cli.Flag {

	flags := make([]cli.Flag, 0, 4)
	flags = append(flags, cli.StringFlag{
		Name:  "s",
		Usage: "生成 jwt 加密 key 所需原字符串,不输入此参数则每次的key都一样",
	})
	return flags
}

func (p *keyCmd) action(c *cli.Context) (err error) {
	cmds.Log.Infof("jwt Key: %s", MakeSign(c.String("s")))
	return nil
}

//MakeSign .
func MakeSign(s string) (sign string) {
	if s != "" {
		s = s + "key=" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
