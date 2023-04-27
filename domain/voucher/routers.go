package voucher

import (
	"fmt"
	"log"
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
	router.POST("recover", v.Recover)
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

func (v *VoucherRouters) Recover(c *gin.Context) {
	var requestBody struct {
		SubscriptionId string `json:"subscriptionId" form:"subscriptionId"`
		ActivationDate string `json:"activationDate" form:"activationDate"`
		OfferKey       string `json:"offerKey" form:"offerKey"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("Falha no Bind de requestBody na ação de recover")
		c.AbortWithError(400, err)
	}

	activationDate, err := time.Parse("2006-01-02", requestBody.ActivationDate)
	if err != nil {
		log.Println("Falha no parse da date de ativação na ação de recover")
		c.AbortWithError(400, err)
	}
	code, err := v.voucherService.Recover(requestBody.SubscriptionId, activationDate, time.Now(), requestBody.OfferKey)
	if err != nil {
		log.Println("Falha na recuperação de voucher na ação de recover")
		c.AbortWithError(500, err)
	}

	c.JSON(200, code)
}
