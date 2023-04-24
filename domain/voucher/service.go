package voucher

import (
	"errors"
	"time"

	"github.com/brunno98/voucher-manager/domain/code"
)

type Service interface {
	Recover(subscriptionId string, activationDate time.Time, voucherKey string) (code.Code, error)
	AvailableVouchers(subscriotionId string, activationDate time.Time) (AvailableVouchersDTO, error)
}

type voucherService struct {
	Repository
	code.RecoverService
}

func NewService(r Repository, rs code.RecoverService) *voucherService {
	return &voucherService{r, rs}
}

func (service *voucherService) Recover(subscriptionId string, activationDate, now time.Time, voucherKey string) (code.Code, error) {
	recoverParameter := RecoverParameter{DaysToRenew: 30, DaysToExpiry: 60}

	// encontra a data de recuperação mais antiga disponivel
	voucher, err := service.FindByKey(voucherKey)
	if err != nil {
		return code.Code{}, err
	}

	availableDates := recoverParameter.GetDateOfAvailableVouchers(activationDate, now)
	if len(availableDates) == 0 {
		// não há data de resgate disponível
		return code.Code{}, errors.New("unavailable date to recover")
	}

	// Verificar se já houve resgate para as data disponives, as data que já tem
	// resgate devem ser desconsidaras e removidas da lista. Caso a lista esteja vazia, retornar erro
	earliestRetrievalDate, found := service.RecoverService.GetEarliestRecoveryDateNotUsed(subscriptionId, availableDates)
	if !found {
		return code.Code{}, errors.New("unavailable date to recover")
	}

	//busca o codigo referente a voucherKey e "queima" ele

	recovered, err := service.Repository.Recover(subscriptionId, earliestRetrievalDate, &voucher)
	if err != nil {
		return code.Code{}, errors.New("error while recovering code")
	}

	// (TODO) persiste em historico quem resgatou, qual codigo e quando

	// retorna o codigo resgatado
	return recovered, nil
}

func (service *voucherService) AvailableVouchers(subscriptionId string, activationDate time.Time) (AvailableVouchersDTO, error) {
	now := time.Now()
	recoverParameter := RecoverParameter{DaysToRenew: 30, DaysToExpiry: 60}

	availableDTO := AvailableVouchersDTO{}

	// busca os resgates disponiveis baseado na data de ativação que não estão expirados

	availableDates := recoverParameter.GetDateOfAvailableVouchers(activationDate, now)
	if len(availableDates) == 0 {
		// não há data de resgate disponível
		return availableDTO, errors.New("unavailable date to recover")
	}

	// remove os resgates já usados

	availableDates = service.RecoverService.FilterRecoveryAlreadyUsed(availableDates)

	availableDTO.Quantity = len(availableDates)
	availableDTO.NextRecoveryDate = recoverParameter.GetNextRecoverDate(activationDate, now)

	var availableRecoveryInfo []AvailableRecoveryInfoDTO
	for _, d := range availableDates {
		expiryDate := recoverParameter.GetExpireDate(d)
		info := AvailableRecoveryInfoDTO{
			ReferenceDate:  d,
			ExpirationDate: expiryDate,
		}
		availableRecoveryInfo = append(availableRecoveryInfo, info)
	}

	availableDTO.AvailableRecoveryInfo = availableRecoveryInfo

	// busca os vouchers ofertados

	vouchers := service.Repository.FindAll()

	availableDTO.AvailableVouchers = vouchers

	// retorna as informações do DTO
	return availableDTO, nil
}
