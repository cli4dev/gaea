package tmpls

//BaseTpl .
const BaseTpl = `
package %s
import "github.com/micro-plat/hydra/component"
type I%s interface {
}
type %s struct {
	c component.IContainer
}
func New%s(c component.IContainer) *%s{
	return &%s{
		c: c,
	}
}
`
