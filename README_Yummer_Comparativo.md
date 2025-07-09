
# 📊 Comparativo Técnico — Projeto Yummer: Go (Gin) + React vs Python (Django) + React

Este repositório tem como objetivo apresentar uma comparação técnica entre duas abordagens distintas para o desenvolvimento da aplicação **Yummer**, um sistema de reservas para restaurantes.

## 🎯 Objetivo

Comparar duas stacks diferentes aplicadas à mesma aplicação:

- **Yummer-Go**: Backend em Go com o framework Gin + Frontend React
- **Yummer-Web**: Backend em Python com Django + Django REST Framework + Frontend React

---

## 🛠️ Abordagens Utilizadas

### 🟦 Yummer-Go (Gin + React)

- Backend escrito em **Go** com o framework **Gin**
- ORM utilizado: **GORM** com banco SQLite
- Endpoints REST definidos diretamente nas rotas do Gin
- Frontend desacoplado em React, comunicação via `fetch`
- Container Docker multi-stage para otimização da imagem final

### 🐍 Yummer-Web (Django + React)

- Backend em **Python** utilizando **Django** e **Django REST Framework**
- ORM nativo do Django e banco SQLite
- Estrutura modular de apps, views e serializers
- Frontend separado em React, consumo da API via `fetch`
- Docker configurado com ambiente completo

---

## 📊 Tabela Comparativa

| Item                        | Go + Gin                                | Django                                     |
|-----------------------------|------------------------------------------|--------------------------------------------|
| Linguagem                   | Go                                       | Python                                     |
| Framework Backend           | Gin                                      | Django + Django REST Framework             |
| ORM                         | GORM                                     | ORM nativo do Django                       |
| Banco de Dados              | SQLite                                   | SQLite                                     |
| Frontend                    | React (fetch)                            | React (fetch)                              |
| Integração com Frontend     | Desacoplado via REST API                 | Desacoplado via REST API                   |
| Validações                  | Feitas manualmente                       | Automatizadas com serializers              |
| Estrutura MVC               | Parcial (handlers, models, services)     | Completa (views, models, serializers)      |
| Admin Interface             | Não implementado                         | Interface administrativa pronta            |
| Docker                      | Multi-stage otimizado                    | Estrutura padrão com backend + frontend    |

---

## ✅ Vantagens Observadas

### Go + Gin
- Performance superior e baixo consumo de recursos
- Estrutura simples, ideal para microserviços
- Imagem Docker leve e rápida de subir

### Django
- Alta produtividade com recursos prontos
- Admin automático facilita gestão dos dados
- Estrutura robusta e modular por natureza

---

## ⚠️ Limitações

### Go + Gin
- Validações e manipulação de dados mais verbosas
- Falta de interface administrativa nativa
- Curva de aprendizado para GORM

### Django
- Mais pesado para sistemas pequenos
- Build e inicialização mais lentos
- Acoplamento inicial elevado, se não modularizado

---

## 🚀 Oportunidades Futuras

### Para Go
- Adição de Swagger (OpenAPI) para documentação automática
- Autenticação via JWT
- Testes unitários e integração com `httptest` e `testify`

### Para Django
- Estrutura CI/CD com cobertura de testes
- Deploy separado por ambiente (dev, prod)
- Autenticação avançada e permissões com DRF

---

## 🔎 Exemplos de Código

### Criando uma reserva — Go + Gin

```go
func CreateReserva(c *gin.Context) {
    var reserva models.Reserva
    if err := c.ShouldBindJSON(&reserva); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&reserva)
    c.JSON(http.StatusOK, &reserva)
}
```

### Criando uma reserva — Django

```python
class ReservaViewSet(viewsets.ModelViewSet):
    queryset = Reserva.objects.all()
    serializer_class = ReservaSerializer
```

---

## 🖥️ Diferenças Visuais

Ambas as aplicações usam React como frontend, com listagem de reservas, criação de clientes, mesas e restaurantes. As rotas REST fornecem as mesmas operações CRUD, com diferença apenas na estruturação e no estilo de código backend.

---

## 📌 Conclusão

Ambas as stacks são viáveis para a construção de APIs modernas e funcionais. A escolha ideal depende do perfil do time, requisitos de performance, manutenibilidade e curva de aprendizado desejada:

- **Go + Gin** é recomendável para APIs leves, performáticas e altamente escaláveis.
- **Django + DRF** é ideal para aplicações completas, com autenticação, painel admin e produtividade alta.

---

## 📂 Repositórios

- 🔵 Go + Gin: [nicolasltb/yummer-go](https://github.com/nicolasltb/yummer-go)
- 🟠 Django + React: [PdRabelo/yummer-web](https://github.com/PdRabelo/yummer-web)
