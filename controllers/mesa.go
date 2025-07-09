package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func GetMesasHTML(c *gin.Context) {
	var mesas []models.Mesa
	rows, err := database.DB.Query("SELECT id, restaurante_id, numero, capacidade, disponivel FROM mesas")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var mesa models.Mesa
		if err := rows.Scan(&mesa.ID, &mesa.RestauranteID, &mesa.Numero, &mesa.Capacidade, &mesa.Disponivel); err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		mesas = append(mesas, mesa)
	}
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Mesas",
		"bodyData": mesas,
	})
}

func CreateMesaHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Adicionar Mesa",
		"bodyData": models.Mesa{},
	})
}

func EditMesaHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid mesa ID"})
		return
	}

	var mesa models.Mesa
	err = database.DB.QueryRow("SELECT id, restaurante_id, numero, capacidade, disponivel FROM mesas WHERE id = ?", id).Scan(&mesa.ID, &mesa.RestauranteID, &mesa.Numero, &mesa.Capacidade, &mesa.Disponivel)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": "Mesa not found"})
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Editar Mesa",
		"bodyData": mesa,
	})
}

func SaveMesaHTML(c *gin.Context) {
	var mesa models.Mesa
	mesa.ID, _ = strconv.Atoi(c.PostForm("id"))
	mesa.RestauranteID, _ = strconv.Atoi(c.PostForm("restaurante_id"))
	mesa.Numero, _ = strconv.Atoi(c.PostForm("numero"))
	mesa.Capacidade, _ = strconv.Atoi(c.PostForm("capacidade"))
	mesa.Disponivel = c.PostForm("disponivel") == "true"

	if mesa.ID == 0 { // Create new
		query, err := database.DB.Prepare("INSERT INTO mesas (restaurante_id, numero, capacidade, disponivel) VALUES (?, ?, ?, ?)")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(mesa.RestauranteID, mesa.Numero, mesa.Capacidade, mesa.Disponivel)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return		}
	} else { // Update existing
		query, err := database.DB.Prepare("UPDATE mesas SET restaurante_id = ?, numero = ?, capacidade = ?, disponivel = ? WHERE id = ?")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(mesa.RestauranteID, mesa.Numero, mesa.Capacidade, mesa.Disponivel, mesa.ID)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	}
	c.Redirect(http.StatusFound, "/mesas")
}

func DeleteMesaHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid mesa ID"})
		return
	}

	query, err := database.DB.Prepare("DELETE FROM mesas WHERE id = ?")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/mesas")
}