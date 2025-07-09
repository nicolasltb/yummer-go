import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

interface Mesa {
  id: number;
  restaurante_id: number;
  numero: number;
  capacidade: number;
  disponivel: boolean;
}

const MesaList: React.FC = () => {
  const [mesas, setMesas] = useState<Mesa[]>([]);

  useEffect(() => {
    fetchMesas();
  }, []);

  const fetchMesas = async () => {
    try {
      const response = await fetch('/mesas');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setMesas(data);
    } catch (error) {
      console.error("Error fetching mesas:", error);
    }
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Tem certeza que deseja deletar esta mesa?')) {
      try {
        const response = await fetch(`/mesas/${id}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        fetchMesas(); // Refresh the list
      } catch (error) {
        console.error("Error deleting mesa:", error);
      }
    }
  };

  return (
    <div>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <h2>Mesas</h2>
        <Link to="/mesas/new" className="btn btn-primary">Adicionar Mesa</Link>
      </div>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Restaurante ID</th>
            <th>Número</th>
            <th>Capacidade</th>
            <th>Disponível</th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          {mesas.length > 0 ? (
            mesas.map((mesa) => (
              <tr key={mesa.id}>
                <td>{mesa.id}</td>
                <td>{mesa.restaurante_id}</td>
                <td>{mesa.numero}</td>
                <td>{mesa.capacidade}</td>
                <td>{mesa.disponivel ? 'Sim' : 'Não'}</td>
                <td>
                  <Link to={`/mesas/edit/${mesa.id}`} className="btn btn-sm btn-warning me-2">Editar</Link>
                  <button onClick={() => handleDelete(mesa.id)} className="btn btn-sm btn-danger">Deletar</button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={6}>Nenhuma mesa encontrada.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default MesaList;
