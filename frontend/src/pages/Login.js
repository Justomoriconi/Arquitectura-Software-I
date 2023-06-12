import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Grid, Paper, Typography, TextField, Button } from '@mui/material';

const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (event) => {
    setEmail(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          "Email":email,
          "Pwd":password,
        }),
      });

      const data = await response.json();

      console.log(data); // Imprimir la respuesta del backend
      console.log(password, email);
      if (data.success) {
        navigate("/")
        console.log('Inicio de sesión ');
      } else {
        // Inicio de sesión fallido
        console.log('Inicio de sesión fallido');
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Container maxWidth="xs">
      <Grid container justifyContent="center" alignItems="center" style={{ height: '100vh' }}>
        <Grid item xs={12}>
          <Paper elevation={3} sx={{ padding: '2rem' }}>
            <Typography variant="h4" align="center" gutterBottom>
              Login
            </Typography>
            <form onSubmit={handleSubmit}>
              <TextField
                label="Email"
                type="email"
                fullWidth
                margin="normal"
                value={email}
                onChange={handleEmailChange}
                required
              />
              <TextField
                label="Password"
                type="password"
                fullWidth
                margin="normal"
                value={password}
                onChange={handlePasswordChange}
                required
              />
              <Button type="submit" variant="contained" color="primary" fullWidth>
                Login
              </Button>
            </form>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default Login;
