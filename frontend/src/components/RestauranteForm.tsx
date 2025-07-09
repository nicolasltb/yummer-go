import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

interface Restaurante {
  id: number;
  nome: string;
  endereco: string;
  tipo_cozinha: string;
  horario_funcionamento: string;
}

const RestauranteForm: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [restaurante, setRestaurante] = useState<Restaurante>({
    id: 0,
    nome: '',
    endereco: '',
    tipo_cozinha: '',
    horario_funcionamento: '',
  });

  useEffect(() => {
    if (id) {
      fetchRestaurante(parseInt(id));
    }
  }, [id]);

  const fetchRestaurante = async (restauranteId: number) => {
    try {
      const response = await fetch(`/restaurantes/${restauranteId}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setRestaurante(data);
    } catch (error) {
      console.error("Error fetching restaurante:", error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setRestaurante((prevRestaurante) => ({
      ...prevRestaurante,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const method = restaurante.id ? 'PUT' : 'POST';
      const url = restaurante.id ? `/restaurantes/${restaurante.id}` : '/restaurantes';
      let bodyData: Omit<Restaurante, 'id'> | Restaurante;
      if (method === 'POST') {
        const { id, ...rest } = restaurante;
        bodyData = rest;
      } else {
        bodyData = restaurante;
      }

      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(bodyData),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      navigate('/restaurantes');
    } catch (error) {
      console.error("Error saving restaurante:", error);
    }
  };

  return (
    <div>
      <h2>{restaurante.id ? 'Editar Restaurante' : 'Adicionar Restaurante'}</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">Nome</label>
          <input
            type="text"
            className="form-control"
            id="nome"
            name="nome"
            value={restaurante.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="endereco" className="form-label">Endereço</label>
          <input
            type="text"
            className="form-control"
            id="endereco"
            name="endereco"
            value={restaurante.endereco}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="tipo_cozinha" className="form-label">Tipo de Cozinha</label>
          <input
            type="text"
            className="form-control"
            id="tipo_cozinha"
            name="tipo_cozinha"
            value={restaurante.tipo_cozinha}
            onChange={handleChange}
          />
        </div>
        <div className="mb-3">
          <label htmlFor="horario_funcionamento" className="form-label">Horário de Funcionamento</label>
          <input
            type="text"
            className="form-control"
            id="horario_funcionamento"
            name="horario_funcionamento"
            value={restaurante.horario_funcionamento}
            onChange={handleChange}
          />
        </div>
        <button type="submit" className="btn btn-success">Salvar</button>
        <button type="button" className="btn btn-secondary ms-2" onClick={() => navigate('/restaurantes')}>Cancelar</button>
      </form>
    </div>
  );
};

export default RestauranteForm;
