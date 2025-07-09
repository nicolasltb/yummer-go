package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func GetReservasHTML(c *gin.Context) {
	var reservas []models.Reserva
	rows, err := database.DB.Query("SELECT id, cliente_id, mesa_id, data_hora, numero_pessoas FROM reservas")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var reserva models.Reserva
		if err := rows.Scan(&reserva.ID, &reserva.ClienteID, &reserva.MesaID, &reserva.DataHora, &reserva.NumeroPessoas); err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		reservas = append(reservas, reserva)
	}
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Reservas",
		"bodyData": reservas,
	})
}

func CreateReservaHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Adicionar Reserva",
		"bodyData": models.Reserva{},
	})
}

func EditReservaHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid reserva ID"})
		return
	}

	var reserva models.Reserva
	err = database.DB.QueryRow("SELECT id, cliente_id, mesa_id, data_hora, numero_pessoas FROM reservas WHERE id = ?", id).Scan(&reserva.ID, &reserva.ClienteID, &reserva.MesaID, &reserva.DataHora, &reserva.NumeroPessoas)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": "Reserva not found"})
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Editar Reserva",
		"bodyData": reserva,
	})
}

func SaveReservaHTML(c *gin.Context) {
	var reserva models.Reserva
	reserva.ID, _ = strconv.Atoi(c.PostForm("id"))
	reserva.ClienteID, _ = strconv.Atoi(c.PostForm("cliente_id"))
	reserva.MesaID, _ = strconv.Atoi(c.PostForm("mesa_id"))
	reserva.NumeroPessoas, _ = strconv.Atoi(c.PostForm("numero_pessoas"))

	dataHoraStr := c.PostForm("data_hora")
	parsedTime, err := time.Parse("2006-01-02T15:04", dataHoraStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid date/time format"})
		return
	}
	reserva.DataHora = parsedTime

	if reserva.ID == 0 { // Create new
		query, err := database.DB.Prepare("INSERT INTO reservas (cliente_id, mesa_id, data_hora, numero_pessoas) VALUES (?, ?, ?, ?)")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(reserva.ClienteID, reserva.MesaID, reserva.DataHora, reserva.NumeroPessoas)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	} else { // Update existing
		query, err := database.DB.Prepare("UPDATE reservas SET cliente_id = ?, mesa_id = ?, data_hora = ?, numero_pessoas = ? WHERE id = ?")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(reserva.ClienteID, reserva.MesaID, reserva.DataHora, reserva.NumeroPessoas, reserva.ID)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	}
	c.Redirect(http.StatusFound, "/reservas")
}

func DeleteReservaHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid reserva ID"})
		return
	}

	query, err := database.DB.Prepare("DELETE FROM reservas WHERE id = ?")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/reservas")
}