package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func GetClientesHTML(c *gin.Context) {
	var clientes []models.Cliente
	rows, err := database.DB.Query("SELECT id, nome, email, telefone FROM clientes")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone); err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		clientes = append(clientes, cliente)
	}
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Clientes",
		"bodyData": clientes,
	})
}

func CreateClienteHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Adicionar Cliente",
		"bodyData": models.Cliente{},
	})
}

func EditClienteHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid client ID"})
		return
	}

	var cliente models.Cliente
	err = database.DB.QueryRow("SELECT id, nome, email, telefone FROM clientes WHERE id = ?", id).Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": "Cliente not found"})
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title":    "Editar Cliente",
		"bodyData": cliente,
	})
}

func SaveClienteHTML(c *gin.Context) {
	var cliente models.Cliente
	cliente.ID, _ = strconv.Atoi(c.PostForm("id"))
	cliente.Nome = c.PostForm("nome")
	cliente.Email = c.PostForm("email")
	cliente.Telefone = c.PostForm("telefone")

	if cliente.ID == 0 { // Create new
		query, err := database.DB.Prepare("INSERT INTO clientes (nome, email, telefone) VALUES (?, ?, ?)")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(cliente.Nome, cliente.Email, cliente.Telefone)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	} else { // Update existing
		query, err := database.DB.Prepare("UPDATE clientes SET nome = ?, email = ?, telefone = ? WHERE id = ?")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
		_, err = query.Exec(cliente.Nome, cliente.Email, cliente.Telefone, cliente.ID)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
			return
		}
	}
	c.Redirect(http.StatusFound, "/clientes")
}

func DeleteClienteHTML(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "layout.html", gin.H{"error": "Invalid client ID"})
		return
	}

	query, err := database.DB.Prepare("DELETE FROM clientes WHERE id = ?")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "layout.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/clientes")
}