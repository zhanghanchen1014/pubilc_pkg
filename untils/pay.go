package untils

import "github.com/zhanghanchen1014/pubilc_pkg/pkg"

type Pay interface {
	AlipayOrder(orderSn, total string) string
}

type ALiPay struct {
}

func (p *ALiPay) AlipayOrder(orderSn, total string) string {
	return pkg.PayOrder(orderSn, total)
}
