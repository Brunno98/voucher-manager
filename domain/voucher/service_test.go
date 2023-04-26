package voucher

import (
	"testing"
	"time"

	code "github.com/brunno98/voucher-manager/domain/code"
	"github.com/golang/mock/gomock"
)

func TestRecover(t *testing.T) {
	activationDate := time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC)
	now := time.Date(2023, 4, 22, 0, 0, 0, 0, time.UTC)
	subscriptionId := "SOME ID"
	voucherKey := "SOME KEY"

	ctrl := gomock.NewController(t)

	mockRepository := NewMockRepository(ctrl)
	mockRecoverService := code.NewMockRecoverService(ctrl)

	mockRepository.
		EXPECT().
		FindByKey(gomock.Any()).
		Return(Voucher{Key: voucherKey}, nil)

	mockRecoverService.
		EXPECT().
		GetEarliestRecoveryDateNotUsed(gomock.Any(), gomock.Any()).
		Return(activationDate, nil)

	mockRepository.
		EXPECT().
		Recover(subscriptionId, activationDate, gomock.Any()).
		Return(code.Code{Code: "SOME CODE"}, nil)

	service := NewService(mockRepository, mockRecoverService)

	code, _ := service.Recover(subscriptionId, activationDate, now, "some key")

	if code.Code != "SOME CODE" {
		t.Errorf("Expected code: %s, but received: %s", "SOME CODE", code.Code)
	}

}
