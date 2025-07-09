
package models

import "time"

type Reserva struct {
	ID            int       `json:"id"`
	ClienteID     int       `json:"cliente_id"`
	MesaID        int       `json:"mesa_id"`
	DataHora      time.Time `json:"data_hora"`
	NumeroPessoas int       `json:"numero_pessoas"`
}
