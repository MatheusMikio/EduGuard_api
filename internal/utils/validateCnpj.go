package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type brasilAPICnpjResponse struct {
	CNPJRaiz                   string `json:"cnpj"`
	DescricaoSituacaoCadastral string `json:"descricao_situacao_cadastral"`
}

func ValidateCnpjExists(cnpj string) error {
	clean := strings.NewReplacer(".", "",
		"/", "",
		"-", "",
	).Replace(cnpj)

	url := fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%s", clean)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return errors.New("Não foi possível verificar o CNPJ.")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("Não encontrado.")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Erro ao consultar CNPJ.")
	}

	var result brasilAPICnpjResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return errors.New("Erro ao processar resposta da Receita Federal")
	}

	if !strings.EqualFold(result.DescricaoSituacaoCadastral, "ATIVA") {
		return errors.New("CNPJ com situação cadastral inativa")
	}

	return nil
}
