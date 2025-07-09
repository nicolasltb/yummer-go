package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func GetRestaurantesHTML(c *gin.Context) {
	var restaurantes []models.Restaurante
	rows, err := database.DB.Query("SELECT id, nome, endereco, tipo_cozinha, horario_funcionamento FROM restaurantes")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var restaurante models.Restaurante
		if err := rows.Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Endereco, &restaurante.TipoCozinha, &restaurante.HorarioFuncionamento); err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		restaurantes = append(restaurantes, restaurante)
	}
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Restaurantes",
		"bodyData": restaurantes,
	})
}

func CreateRestauranteHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Adicionar Restaurante",
		"bodyData": models.Restaurante{},
	})
}

func EditRestauranteHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid restaurant ID"})
		return
	}

	var restaurante models.Restaurante
	err = database.DB.QueryRow("SELECT id, nome, endereco, tipo_cozinha, horario_funcionamento FROM restaurantes WHERE id = ?", id).Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Endereco, &restaurante.TipoCozinha, &restaurante.HorarioFuncionamento)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": "Restaurante not found"})
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Editar Restaurante",
		"bodyData": restaurante,
	})
}

func SaveRestauranteHTML(c *gin.Context) {
	var restaurante models.Restaurante
	restaurante.ID, _ = strconv.Atoi(c.PostForm("id"))
	restaurante.Nome = c.PostForm("nome")
	restaurante.Endereco = c.PostForm("endereco")
	restaurante.TipoCozinha = c.PostForm("tipo_cozinha")
	restaurante.HorarioFuncionamento = c.PostForm("horario_funcionamento")

	if restaurante.ID == 0 { // Create new
		query, err := database.DB.Prepare("INSERT INTO restaurantes (nome, endereco, tipo_cozinha, horario_funcionamento) VALUES (?, ?, ?, ?)")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(restaurante.Nome, restaurante.Endereco, restaurante.TipoCozinha, restaurante.HorarioFuncionamento)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	} else { // Update existing
		query, err := database.DB.Prepare("UPDATE restaurantes SET nome = ?, endereco = ?, tipo_cozinha = ?, horario_funcionamento = ? WHERE id = ?")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(restaurante.Nome, restaurante.Endereco, restaurante.TipoCozinha, restaurante.HorarioFuncionamento, restaurante.ID)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	}
	c.Redirect(http.StatusFound, "/restaurantes")
}

func DeleteRestauranteHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid restaurant ID"})
		return
	}

	query, err := database.DB.Prepare("DELETE FROM restaurantes WHERE id = ?")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/restaurantes")
}