package entity

import (
	"time"

	"github.com/google/uuid"
)

type Mail struct {
	Id          string
	Insured     Customer
	Destination string
	Description string
	Message     string
	Date        time.Time
}

func (m *Mail) GetTemplateMail(insured Customer) string {
	m.Id = uuid.NewString()
	m.Description = "Apolice contratada."
	m.Destination = insured.Email
	m.Message = "Olá " + insured.Name + ", você acabou de adquirir o seguro de vida, segue os dados da apolice..."

	return "mensagem formatada"
}
