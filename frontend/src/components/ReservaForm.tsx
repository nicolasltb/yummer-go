import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

interface Reserva {
  id: number;
  cliente_id: number;
  mesa_id: number;
  data_hora: string; // ISO 8601 string
  numero_pessoas: number;
}

const ReservaForm: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [reserva, setReserva] = useState<Reserva>({
    id: 0,
    cliente_id: 0,
    mesa_id: 0,
    data_hora: '',
    numero_pessoas: 0,
  });

  useEffect(() => {
    if (id) {
      fetchReserva(parseInt(id));
    }
  }, [id]);

  const fetchReserva = async (reservaId: number) => {
    try {
      const response = await fetch(`/reservas/${reservaId}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setReserva({
        ...data,
        data_hora: data.data_hora.substring(0, 16), // Format for datetime-local input
      });
    } catch (error) {
      console.error("Error fetching reserva:", error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setReserva((prevReserva) => ({
      ...prevReserva,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const method = reserva.id ? 'PUT' : 'POST';
      const url = reserva.id ? `/reservas/${reserva.id}` : '/reservas';
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(reserva),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      navigate('/reservas');
    } catch (error) {
      console.error("Error saving reserva:", error);
    }
  };

  return (
    <div>
      <h2>{reserva.id ? 'Editar Reserva' : 'Adicionar Reserva'}</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="cliente_id" className="form-label">Cliente ID</label>
          <input
            type="number"
            className="form-control"
            id="cliente_id"
            name="cliente_id"
            value={reserva.cliente_id}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="mesa_id" className="form-label">Mesa ID</label>
          <input
            type="number"
            className="form-control"
            id="mesa_id"
            name="mesa_id"
            value={reserva.mesa_id}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="data_hora" className="form-label">Data e Hora</label>
          <input
            type="datetime-local"
            className="form-control"
            id="data_hora"
            name="data_hora"
            value={reserva.data_hora}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="numero_pessoas" className="form-label">Número de Pessoas</label>
          <input
            type="number"
            className="form-control"
            id="numero_pessoas"
            name="numero_pessoas"
            value={reserva.numero_pessoas}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-success">Salvar</button>
        <button type="button" className="btn btn-secondary ms-2" onClick={() => navigate('/reservas')}>Cancelar</button>
      </form>
    </div>
  );
};

export default ReservaForm;
