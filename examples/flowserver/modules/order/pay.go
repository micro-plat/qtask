
package order

import "github.com/micro-plat/hydra/component"

type IPay interface {
}

type Pay struct {
	c component.IContainer
}

func NewPay(c component.IContainer) *Pay {
	return &Pay{
		c: c,
	}
}
