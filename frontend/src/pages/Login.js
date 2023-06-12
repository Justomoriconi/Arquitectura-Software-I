import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, Link} from 'react-router-dom';
import { Container, Grid, Paper, Button, Typography, TextField, Box } from '@mui/material';

const Login = () => {
  const [loginData, setLoginData] = useState({
    email: '',
    password: '',
  });
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate(); 

  const handleOnChange = (e) => {
    setLoginData({ ...loginData, [e.target.name]: e.target.value });
  };

  const onSubmit = async () => {
    try {
      const response = await axios.post('http://127.0.0.1:8080/login', loginData);
      console.log(response.data);

      if (response.data.success) {
        setIsLoggedIn(true);
        navigate('/home'); 
      } else {
        console.log('Inicio de sesión fallido');
      }
    } catch (error) {
      console.error(error);
    }
  };
 

  return (
    <Container maxWidth="xl">
      <Grid container direction="column" alignItems="center" justifyContent="center" sx={{ minHeight: "75vh" }}>
        <Grid item>
          <Paper sx={{ padding: "1.2em", borderRadius: "0.5em" }}>
            <Typography variant="h4">Login</Typography>
            <Box component="form">
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="email"
                name="email"
                autoFocus
                sx={{ mt: 2, mb: 1.5 }}
                value={loginData.email}
                onChange={handleOnChange}
              />
              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="password"
                type="password"
                id="password"
                autoComplete="current-password"
                sx={{ mt: 1.5, mb: 1.5 }}
                value={loginData.password}
                onChange={handleOnChange}
              />
              <Button fullWidth onClick={onSubmit}>
                Login
              </Button>
              {isLoggedIn ? (
                <Typography variant="h7" alignItems="center">
                  Logged in successfully!
                </Typography>
              ) : (
                <Typography variant="h7" alignItems="center">
                  Create an account
                </Typography>
              )}
              <Link to="/signup">
                <Typography variant="h7" alignItems="center">
                  Signup
                </Typography>
              </Link>
            </Box>
          </Paper>
        </Grid>
      </Grid>
      {isLoggedIn && <p>Usuario autenticado</p>}
    </Container>
  );
};

export default Login;
