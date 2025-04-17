package services

import (
	"encoding/json"
	"errors"
	"myapi/internal/models"
	"myapi/internal/validators"
	"net/http"
)

func DecodeAndValidateItem(r *http.Request) (*models.Item, error) {
	var item models.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		return nil, errors.New("Erro ao decodificar o item: " + err.Error())
	}

	err := validators.ValidateItem(&item)
	if err != nil {
		return nil, errors.New("Erro de validação: " + err.Error())
	}

	return &item, nil
}
