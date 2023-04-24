package voucher

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type VoucherRouters struct {
	voucherService Service
}

func NewVoucherRouter(s Service) *VoucherRouters {
	return &VoucherRouters{s}
}

func (v *VoucherRouters) InitRouters(router *gin.RouterGroup) {
	router.GET("available", v.AvailblesVouchers)
	// router.POST("recover")
}

func (v *VoucherRouters) AvailblesVouchers(c *gin.Context) {
	var requestBody struct {
		SubscriptionId string `json:"subscriptionId" form:"subscriptionId"`
		ActivationDate string `json:"activationDate" form:"activationDate"`
	}

	fmt.Printf("%#v\n", requestBody)

	if err := c.Bind(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	fmt.Printf("%#v\n", requestBody)

	activationDate, err := time.Parse("2006-01-02", requestBody.ActivationDate)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	availableVouchers, err := v.voucherService.AvailableVouchers(requestBody.SubscriptionId, activationDate)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, availableVouchers)
}
