import React, { useState } from 'react';
import Box from '@mui/material/Box';
import { Container, Grid, Paper, Button, Typography, TextField } from '@mui/material';
import { Link } from 'react-router-dom';
import axios from 'axios';






const Signup = () =>{




    
    const initialState = {
        name: '',
        lastName: '',
        email: '',
        password: '',
        username: '',
      };

    const [loginData, setLoginData] = useState(initialState);


    

    const handleOnChange = (e) =>{
        setLoginData({...loginData, [e.target.name] : e.target.value })
    }

    const onSubmit = async () => {
        
        if (!loginData.name || !loginData.lastName || !loginData.email || !loginData.password) {
            alert('Por favor, completa todos los campos');
            return;
          }
        
        try {
          const response = await axios.post('http://127.0.0.1:8080/Singup', loginData);
          console.log(response.data); 
          setLoginData(initialState);
        } catch (error) {
          console.error(error);
        }
      };




    return  <Container maxWidth="xl">
    <Grid 
        container
        direction="column"
        alignItems="center"
        justifyContent="center"
        sx={{minHeight: "75vh"}}
    >
        <Grid item>
            <Paper sx={{padding:"1.2em", borderRadius:"0.5em"}}>
                <Typography variant="h4">Signup</Typography>
                <Box component="form">
                    
                      
                          <TextField  variant='outlined'
                                    margin='normal'
                                    required
                                    fullWidth
                                    name="name"
                                    label="Name"
                                    id="name"
                                    sx={{mt: 1.5, mb: 1.5}}
                                    value={loginData.name}
                                    onChange={handleOnChange}
                        />
                          <TextField  variant='outlined'
                                    margin='normal'
                                    required
                                    fullWidth
                                    name="lastName"
                                    label="Last Name"
                                    id="lastName"
                                    autoComplete='current-password'
                                    sx={{mt: 1.5, mb: 1.5}}
                                    value={loginData.lastName}
                                    onChange={handleOnChange}
                        />
                          <TextField  variant='outlined'
                                    margin='normal'
                                    required
                                    fullWidth
                                    id="email"
                                    label="Email"
                                    name="email"
                                    autoFocus
                                    sx={{mt: 2, mb: 1.5}}
                                    value={loginData.email}
                                    onChange={handleOnChange}
                        />
                        <TextField  variant='outlined'
                                    margin='normal'
                                    required
                                    fullWidth
                                    name="password"
                                    label="Password"
                                    type="password"
                                    id="password"
                                    autoComplete='current-password'
                                    sx={{mt: 1.5, mb: 1.5}}
                                    value={loginData.password}
                                    onChange={handleOnChange}
                        />
                    
                    <Button fullWidth onClick={()=> onSubmit()} >Login</Button>
                    <Typography variant="h7" alignItems='center'>Already have an account ? </Typography>
                    <Link  to="/login">
                        <Typography variant="h7" alignItems='center'>Login</Typography>
                    </Link>
                </Box>
            </Paper>
        </Grid>
    </Grid>
</Container>
}

export default Signup;