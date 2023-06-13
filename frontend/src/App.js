import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import ButtonAppBar from './componets/ButtonAppBar';
import { Box } from '@mui/material';
import BookNow from './pages/BookNow';
import Reserves from './pages/Reserves';
import Search from './pages/Search';
import  Login  from  "./pages/Login"
import Signup from "./pages/Signup"
import Home from './pages/Home';



function App() {
  return (
   
      <Box>
        <ButtonAppBar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/hotel/id/:hotelId" element={<BookNow />} />
          <Route path="/reserves" element={<Reserves />} />
          <Route path="/search" element={<Search />} />
        </Routes>
      </Box>
   
  );
}

export default App;
