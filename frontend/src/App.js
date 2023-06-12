import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import  Login  from  "./pages/Login"
import Signup from "./pages/Signup"
import ButtonAppBar from './componets/ButtonAppBar';
import Home from './pages/Home';
import { Box } from '@mui/material';
import Search from './pages/Search';
import { AuthProvider } from './AuthContext';
import Reserves from './pages/Reserves';



function App() {
  return (
    <AuthProvider>
      <Box>
        <ButtonAppBar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/search" element={<Search />} />
          <Route path="/reserves" element={<Reserves />} />
        </Routes>
      </Box>
    </AuthProvider>
  );
}

export default App;
