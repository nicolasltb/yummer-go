
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func CreateReserva(c *gin.Context) {
	var reserva models.Reserva
	if err := c.ShouldBindJSON(&reserva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("INSERT INTO reservas (cliente_id, mesa_id, data_hora, numero_pessoas) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := query.Exec(reserva.ClienteID, reserva.MesaID, reserva.DataHora, reserva.NumeroPessoas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	reserva.ID = int(id)

	c.JSON(http.StatusCreated, reserva)
}

func GetReservas(c *gin.Context) {
	var reservas []models.Reserva
	rows, err := database.DB.Query("SELECT id, cliente_id, mesa_id, data_hora, numero_pessoas FROM reservas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var reserva models.Reserva
		if err := rows.Scan(&reserva.ID, &reserva.ClienteID, &reserva.MesaID, &reserva.DataHora, &reserva.NumeroPessoas); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		reservas = append(reservas, reserva)
	}

	c.JSON(http.StatusOK, reservas)
}

func GetReserva(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var reserva models.Reserva
	err := database.DB.QueryRow("SELECT id, cliente_id, mesa_id, data_hora, numero_pessoas FROM reservas WHERE id = ?", id).Scan(&reserva.ID, &reserva.ClienteID, &reserva.MesaID, &reserva.DataHora, &reserva.NumeroPessoas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Reserva not found"})
		return
	}

	c.JSON(http.StatusOK, reserva)
}

func UpdateReserva(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var reserva models.Reserva
	if err := c.ShouldBindJSON(&reserva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("UPDATE reservas SET cliente_id = ?, mesa_id = ?, data_hora = ?, numero_pessoas = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(reserva.ClienteID, reserva.MesaID, reserva.DataHora, reserva.NumeroPessoas, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reserva.ID = id
	c.JSON(http.StatusOK, reserva)
}

func DeleteReserva(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	query, err := database.DB.Prepare("DELETE FROM reservas WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva deleted successfully"})
}
