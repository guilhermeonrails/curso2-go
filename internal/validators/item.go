package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateItem(item *models.Item) error {
	if len(item.Nome) < 3 {
		return errors.New("nome deve ter no mínimo 3 caracteres")
	}

	if len(item.Codigo) == 6 {
		return errors.New("código deve ter 6 caracteres")
	}

	if item.Preco <= 0 {
		return errors.New("o preço não pode ser negativo")
	}

	if item.Quantidade < 0 {
		return errors.New("quantidade não pode ser negativa")
	}
	return nil
}
