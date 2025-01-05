package models

import "time"

type GtmMessage struct {
	ID              int       `gorm:"primaryKey;column:id"`
	Status          int8      `gorm:"column:status;not null"`
	PhoneNumber     string    `gorm:"column:phone_number;type:varchar(255);not null"`
	Commentary      string    `gorm:"column:commentary;type:text"`
	UserID          int       `gorm:"column:user_id;not null"`
	CreatedAt       time.Time `gorm:"column:created_at;not null"`
	Acao            string    `gorm:"column:acao;type:varchar(255)"`
	ResponsavelAcao string    `gorm:"column:responsavel_acao;type:varchar(255)"`
	Descricao       string    `gorm:"column:descricao;type:text"`
	ValorOrcado     string    `gorm:"column:valor_orcado;type:decimal(15,2)"`
	ValorAjustado   string    `gorm:"column:valor_ajustado;type:decimal(15,2)"`
	Cliente         string    `gorm:"column:cliente;type:varchar(255)"`
	IDGtm           int       `gorm:"column:id_gtm;type:varchar(100)"`
}
