package validators

import (
	"errors"
	"myapi/internal/models"
)

func ValidateItem(item *models.Item) error {
	if item.Preco <= 0 {
		return errors.New("preço deve ser maior que zero")
	}

	if len(item.Codigo) != 6 {
		return errors.New("o código deve ter 6 caracteres")
	}

	if item.Quantidade < 0 {
		return errors.New("quantidade não pode ser negativa")
	}

	return nil
}
