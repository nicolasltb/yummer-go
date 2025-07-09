import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

interface Reserva {
  id: number;
  cliente_id: number;
  mesa_id: number;
  data_hora: string; // Assuming ISO 8601 string from backend
  numero_pessoas: number;
}

const ReservaList: React.FC = () => {
  const [reservas, setReservas] = useState<Reserva[]>([]);

  useEffect(() => {
    fetchReservas();
  }, []);

  const fetchReservas = async () => {
    try {
      const response = await fetch('/reservas');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setReservas(data);
    } catch (error) {
      console.error("Error fetching reservas:", error);
    }
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Tem certeza que deseja deletar esta reserva?')) {
      try {
        const response = await fetch(`/reservas/${id}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        fetchReservas(); // Refresh the list
      } catch (error) {
        console.error("Error deleting reserva:", error);
      }
    }
  };

  return (
    <div>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <h2>Reservas</h2>
        <Link to="/reservas/new" className="btn btn-primary">Adicionar Reserva</Link>
      </div>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Cliente ID</th>
            <th>Mesa ID</th>
            <th>Data e Hora</th>
            <th>Número de Pessoas</th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          {reservas.length > 0 ? (
            reservas.map((reserva) => (
              <tr key={reserva.id}>
                <td>{reserva.id}</td>
                <td>{reserva.cliente_id}</td>
                <td>{reserva.mesa_id}</td>
                <td>{new Date(reserva.data_hora).toLocaleString()}</td>
                <td>{reserva.numero_pessoas}</td>
                <td>
                  <Link to={`/reservas/edit/${reserva.id}`} className="btn btn-sm btn-warning me-2">Editar</Link>
                  <button onClick={() => handleDelete(reserva.id)} className="btn btn-sm btn-danger">Deletar</button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={6}>Nenhuma reserva encontrada.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default ReservaList;
