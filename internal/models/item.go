package models

import "errors"

type Item struct {
	Id         uint    `gorm:"primaryKey" json:"id"`
	Nome       string  `json:"nome"`
	Codigo     string  `gorm:"unique" json:"codigo"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func (i *Item) Validate() error {
	if i.Preco < 0 {
		return errors.New("o preço não pode ser negativo")
	}
	return nil
}
