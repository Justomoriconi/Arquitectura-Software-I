import React, { useContext } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { Link, useNavigate } from 'react-router-dom';
import { AuthContext } from '../AuthProvider';

export default function ButtonAppBar() {
  const navigate = useNavigate();
  const { isLoggedIn, logout } = useContext(AuthContext);

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  return (
    <AppBar position="static">
      <Toolbar sx={{ display: 'flex', flexDirection: 'row', justifyContent: 'space-between' }}>
        <Link to="/">
          <Typography gutterBottom variant="h6">
            Home
          </Typography>
        </Link>
        <Box >
          {isLoggedIn ? (
            
           <Box>
                <Link to="/reserves">
                  <Typography >
                    My Reserves
                  </Typography>
                </Link>

                <Typography onClick={handleLogout} variant="body1">
                  Cerrar sesión
                </Typography>
              </Box>
              
          ) : (
            <>
              <Link to="/login">
                <Typography variant="body1">Login</Typography>
              </Link>
              <Link to="/signup">
                <Typography variant="body1">Signup</Typography>
              </Link>
            </>
          )}
        </Box>
      </Toolbar>
    </AppBar>
  );
  
  

}