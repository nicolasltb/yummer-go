
package models

type Restaurante struct {
	ID                   int    `json:"id"`
	Nome                 string `json:"nome"`
	Endereco             string `json:"endereco"`
	TipoCozinha          string `json:"tipo_cozinha"`
	HorarioFuncionamento string `json:"horario_funcionamento"`
}
