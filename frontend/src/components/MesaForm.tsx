import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

interface Mesa {
  id: number;
  restaurante_id: number;
  numero: number;
  capacidade: number;
  disponivel: boolean;
}

const MesaForm: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [mesa, setMesa] = useState<Mesa>({
    id: 0,
    restaurante_id: 0,
    numero: 0,
    capacidade: 0,
    disponivel: false,
  });

  useEffect(() => {
    if (id) {
      fetchMesa(parseInt(id));
    }
  }, [id]);

  const fetchMesa = async (mesaId: number) => {
    try {
      const response = await fetch(`/mesas/${mesaId}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setMesa(data);
    } catch (error) {
      console.error("Error fetching mesa:", error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setMesa((prevMesa) => ({
      ...prevMesa,
      [name]: type === 'checkbox' ? checked : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const method = mesa.id ? 'PUT' : 'POST';
      const url = mesa.id ? `/mesas/${mesa.id}` : '/mesas';
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(mesa),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      navigate('/mesas');
    } catch (error) {
      console.error("Error saving mesa:", error);
    }
  };

  return (
    <div>
      <h2>{mesa.id ? 'Editar Mesa' : 'Adicionar Mesa'}</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="restaurante_id" className="form-label">Restaurante ID</label>
          <input
            type="number"
            className="form-control"
            id="restaurante_id"
            name="restaurante_id"
            value={mesa.restaurante_id}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="numero" className="form-label">Número da Mesa</label>
          <input
            type="number"
            className="form-control"
            id="numero"
            name="numero"
            value={mesa.numero}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="capacidade" className="form-label">Capacidade</label>
          <input
            type="number"
            className="form-control"
            id="capacidade"
            name="capacidade"
            value={mesa.capacidade}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3 form-check">
          <input
            type="checkbox"
            className="form-check-input"
            id="disponivel"
            name="disponivel"
            checked={mesa.disponivel}
            onChange={handleChange}
          />
          <label className="form-check-label" htmlFor="disponivel">Disponível</label>
        </div>
        <button type="submit" className="btn btn-success">Salvar</button>
        <button type="button" className="btn btn-secondary ms-2" onClick={() => navigate('/mesas')}>Cancelar</button>
      </form>
    </div>
  );
};

export default MesaForm;
