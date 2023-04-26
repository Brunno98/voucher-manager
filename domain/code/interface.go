package code

import "time"

type Repository interface {
	// Retorna a lista de datas de resgate já usadas por um usuario
	// ordenada a partir da mais recente
	GetLastRecoveredDates(subscriptionId string, limit int) ([]time.Time, error)
	GetRecoveredByReferenceDates(subscriptionId string, dates []time.Time) []Recovered
}
