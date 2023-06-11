import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { Link } from 'react-router-dom';


export default function ButtonAppBar() {
  return (
  
      <AppBar position="static">
        <Toolbar sx={{display: "flex", flexDirection: "row", justifyContent: "space-between"}}>
          <Link to="/">
            <Typography gutterBottom variant='h6' >
              Home
            </Typography>
          </Link>
         <Box> 
            <Link  to="/login">
              <Typography>
                Login
              </Typography> 
            </Link>
            <Link  to="/signup">
              <Typography>
                Signup
              </Typography> 
            </Link>
          </Box>
         
        </Toolbar>
      </AppBar>
  );
}