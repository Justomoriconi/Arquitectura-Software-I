import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import  Login  from  "./pages/Login"
import Signup from "./pages/Signup"
import ButtonAppBar from './componets/ButtonAppBar';
import Home from './pages/Home';
import { Box } from '@mui/material';



function App() {
  return (
   
    <Box > 
      <ButtonAppBar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="login" element={<Login />} />
        <Route path="signup" element={<Signup />} />
      </Routes>
    </Box>
 
 
 
  );
}

export default App;
