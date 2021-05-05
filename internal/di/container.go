package di

import (
	"ob-product/internal/controller"
	ir "ob-product/internal/infra/repository"
	is "ob-product/internal/infra/service"
	"sync"
)

var once sync.Once

type container struct {
}

var containerInstance *container

func GetInstance() *container {
	if containerInstance == nil {
		once.Do(
			func() {
				containerInstance = &container{}
			},
		)
	}

	return containerInstance
}

func (c *container) Injec() *controller.ProductController {
	repo := &ir.ProductInfraRepository{}
	svc := &is.ProductInfraService{repo}
	ctl := &controller.ProductController{svc}

	return ctl
}
