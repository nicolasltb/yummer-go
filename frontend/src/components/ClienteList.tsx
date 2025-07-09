import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

interface Cliente {
  id: number;
  nome: string;
  email: string;
  telefone: string;
}

const ClienteList: React.FC = () => {
  const [clientes, setClientes] = useState<Cliente[]>([]);

  useEffect(() => {
    fetchClientes();
  }, []);

  const fetchClientes = async () => {
    try {
      const response = await fetch('/clientes');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setClientes(data);
    } catch (error) {
      console.error("Error fetching clientes:", error);
    }
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Tem certeza que deseja deletar este cliente?')) {
      try {
        const response = await fetch(`/clientes/${id}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        fetchClientes(); // Refresh the list
      } catch (error) {
        console.error("Error deleting cliente:", error);
      }
    }
  };

  return (
    <div>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <h2>Clientes</h2>
        <Link to="/clientes/new" className="btn btn-primary">Adicionar Cliente</Link>
      </div>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nome</th>
            <th>Email</th>
            <th>Telefone</th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          {clientes.length > 0 ? (
            clientes.map((cliente) => (
              <tr key={cliente.id}>
                <td>{cliente.id}</td>
                <td>{cliente.nome}</td>
                <td>{cliente.email}</td>
                <td>{cliente.telefone}</td>
                <td>
                  <Link to={`/clientes/edit/${cliente.id}`} className="btn btn-sm btn-warning me-2">Editar</Link>
                  <button onClick={() => handleDelete(cliente.id)} className="btn btn-sm btn-danger">Deletar</button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={5}>Nenhum cliente encontrado.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};

export default ClienteList;
