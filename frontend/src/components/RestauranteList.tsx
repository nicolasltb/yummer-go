import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

interface Restaurante {
  id: number;
  nome: string;
  endereco: string;
  tipo_cozinha: string;
  horario_funcionamento: string;
}

const RestauranteList: React.FC = () => {
  const [restaurantes, setRestaurantes] = useState<Restaurante[]>([]);

  useEffect(() => {
    fetchRestaurantes();
  }, []);

  const fetchRestaurantes = async () => {
    try {
      const response = await fetch('/restaurantes');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setRestaurantes(data);
    } catch (error) {
      console.error("Error fetching restaurantes:", error);
    }
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Tem certeza que deseja deletar este restaurante?')) {
      try {
        const response = await fetch(`/restaurantes/${id}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        fetchRestaurantes(); // Refresh the list
      } catch (error) {
        console.error("Error deleting restaurante:", error);
      }
    }
  };

  return (
    <div>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <h2>Restaurantes</h2>
        <Link to="/restaurantes/new" className="btn btn-primary">Adicionar Restaurante</Link>
      </div>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nome</th>
            <th>Endereço</th>
            <th>Tipo de Cozinha</th>
            <th>Horário de Funcionamento</th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          {restaurantes.length > 0 ? (
            restaurantes.map((restaurante) => (
              <tr key={restaurante.id}>
                <td>{restaurante.id}</td>
                <td>{restaurante.nome}</td>
                <td>{restaurante.endereco}</td>
                <td>{restaurante.tipo_cozinha}</td>
                <td>{restaurante.horario_funcionamento}</td>
                <td>
                  <Link to={`/restaurantes/edit/${restaurante.id}`} className="btn btn-sm btn-warning me-2">Editar</Link>
                  <button onClick={() => handleDelete(restaurante.id)} className="btn btn-sm btn-danger">Deletar</button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={6}>Nenhum restaurante encontrado.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default RestauranteList;
