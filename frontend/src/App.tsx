import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import './App.css';
import ClienteList from './components/ClienteList';
import ClienteForm from './components/ClienteForm';
import RestauranteList from './components/RestauranteList';
import RestauranteForm from './components/RestauranteForm';
import MesaList from './components/MesaList';
import MesaForm from './components/MesaForm';
import ReservaList from './components/ReservaList';
import ReservaForm from './components/ReservaForm';

function App() {
  return (
    <Router>
      <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">Yummer Go</Link>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item">
                <Link className="nav-link" to="/clientes">Clientes</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/restaurantes">Restaurantes</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/mesas">Mesas</Link>
              </li>
              <li className="nav-item">
                <Link className="nav-link" to="/reservas">Reservas</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>

      <div className="container mt-4">
        <Routes>
          <Route path="/" element={<h2>Welcome to Yummer Go!</h2>} />
          <Route path="/clientes" element={<ClienteList />} />
          <Route path="/clientes/new" element={<ClienteForm />} />
          <Route path="/clientes/edit/:id" element={<ClienteForm />} />
          <Route path="/restaurantes" element={<RestauranteList />} />
          <Route path="/restaurantes/new" element={<RestauranteForm />} />
          <Route path="/restaurantes/edit/:id" element={<RestauranteForm />} />
          <Route path="/mesas" element={<MesaList />} />
          <Route path="/mesas/new" element={<MesaForm />} />
          <Route path="/mesas/edit/:id" element={<MesaForm />} />
          <Route path="/reservas" element={<ReservaList />} />
          <Route path="/reservas/new" element={<ReservaForm />} />
          <Route path="/reservas/edit/:id" element={<ReservaForm />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;