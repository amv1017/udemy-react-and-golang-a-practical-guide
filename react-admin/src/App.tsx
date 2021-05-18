import React from 'react';
import logo from './logo.svg';
import './App.css';
import Users from './pages/Users';
import Dashboard from './pages/Dashboard';
import { BrowserRouter, Route } from 'react-router-dom';
import Register from './pages/Register';

export default function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path={'/'} exact component={Dashboard} />
        <Route path={'/users'} component={Users} />
        <Route path={'/register'} component={Register} />
      </BrowserRouter>
    </div>
  )
}

