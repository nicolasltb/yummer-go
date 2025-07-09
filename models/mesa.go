
package models

type Mesa struct {
	ID            int  `json:"id"`
	RestauranteID int  `json:"restaurante_id"`
	Numero        int  `json:"numero"`
	Capacidade    int  `json:"capacidade"`
	Disponivel    bool `json:"disponivel"`
}
