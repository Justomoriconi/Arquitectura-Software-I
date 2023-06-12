import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { Link } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../AuthContext';
import axios from 'axios';



export default function ButtonAppBar() {
  const { isLoggedIn, setIsLoggedIn } = useContext(AuthContext);

  const handleLogout = () => {
    axios.post('/logout')
      .then(response => {
        setIsLoggedIn(false);
      })
      .catch(error => {
        console.error(error);
      });
  };

  return (
    <AppBar position="static">
      <Toolbar sx={{ display: "flex", flexDirection: "row", justifyContent: "space-between" }}>
        <Link to="/">
          <Typography style={{ cursor: 'pointer' }} >
            Home
          </Typography>
        </Link>
        {isLoggedIn ? (
          <>
            <Box>
            <Link to="/reserves" >
              <Typography  >
                My Reserves
              </Typography>
            </Link>
            </Box>
            <Box>
              <Typography onClick={handleLogout} style={{ cursor: 'pointer' }}>
              Logout
              </Typography>
            </Box>
          </>
        ) : (
          <Box>
            <Link to="/login">
              <Typography style={{ cursor: 'pointer' }}>
                Login
              </Typography>
            </Link>
            <Link to="/signup">
              <Typography style={{ cursor: 'pointer' }}>
                Signup
              </Typography>
            </Link>
          </Box>
        )}
      </Toolbar>
    </AppBar>
  );
}
