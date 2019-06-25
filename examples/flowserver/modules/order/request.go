
package order

import "github.com/micro-plat/hydra/component"

type IRequest interface {
}

type Request struct {
	c component.IContainer
}

func NewRequest(c component.IContainer) *Request {
	return &Request{
		c: c,
	}
}
