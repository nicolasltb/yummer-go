import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

interface Cliente {
  id: number;
  nome: string;
  email: string;
  telefone: string;
}

const ClienteForm: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [cliente, setCliente] = useState<Cliente>({
    id: 0,
    nome: '',
    email: '',
    telefone: '',
  });

  useEffect(() => {
    if (id) {
      fetchCliente(parseInt(id));
    }
  }, [id]);

  const fetchCliente = async (clienteId: number) => {
    try {
      const response = await fetch(`/clientes/${clienteId}`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      setCliente(data);
    } catch (error) {
      console.error("Error fetching cliente:", error);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setCliente((prevCliente) => ({
      ...prevCliente,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const method = cliente.id ? 'PUT' : 'POST';
      const url = cliente.id ? `/clientes/${cliente.id}` : '/clientes';
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(cliente),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      navigate('/clientes');
    } catch (error) {
      console.error("Error saving cliente:", error);
    }
  };

  return (
    <div>
      <h2>{cliente.id ? 'Editar Cliente' : 'Adicionar Cliente'}</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">Nome</label>
          <input
            type="text"
            className="form-control"
            id="nome"
            name="nome"
            value={cliente.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="email" className="form-label">Email</label>
          <input
            type="email"
            className="form-control"
            id="email"
            name="email"
            value={cliente.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="telefone" className="form-label">Telefone</label>
          <input
            type="text"
            className="form-control"
            id="telefone"
            name="telefone"
            value={cliente.telefone}
            onChange={handleChange}
          />
        </div>
        <button type="submit" className="btn btn-success">Salvar</button>
        <button type="button" className="btn btn-secondary ms-2" onClick={() => navigate('/clientes')}>Cancelar</button>
      </form>
    </div>
  );
};

export default ClienteForm;
