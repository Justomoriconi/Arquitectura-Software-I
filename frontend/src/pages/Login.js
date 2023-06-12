import React, { useState } from 'react';
import Box from '@mui/material/Box';
import { Container, Grid, Paper, Button, Typography, TextField } from '@mui/material';
import { Link } from 'react-router-dom';





const Login=() =>{
    const [loginData, setLoginData] = useState({
        email: '',
        password: '',
    });

    

    const handleOnChange = (e) =>{
        setLoginData({...loginData, [e.target.name] : e.target.value })
    }

   /* const handleSubmit = (email, password) => {
        alert('datos formulariol')
    }*/


    const onSubmit = () =>{
        console.log(loginData)
    }

    
  



    return  (
        <Container maxWidth="xl">
            <Grid 
                container
                direction="column"
                alignItems="center"
                justifyContent="center"
                sx={{minHeight: "75vh"}}
            >
                <Grid item>
                    <Paper sx={{padding:"1.2em", borderRadius:"0.5em"}}>
                        <Typography variant="h4">Login</Typography>
                        <Box component="form">
                            
                                <TextField  variant='outlined'
                                            margin='normal'
                                            required
                                            fullWidth
                                            id="email"
                                            label="email"
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
                                            label="password"
                                            type="password"
                                            id="password"
                                            autoComplete='current-password'
                                            sx={{mt: 1.5, mb: 1.5}}
                                            value={loginData.password}
                                            onChange={handleOnChange}
                                />
                            
                            <Button fullWidth onClick={()=> onSubmit()} >Login</Button>
                            <Typography variant="h7" alignItems='center'>Create an account </Typography>
                            <Link  to="/signup">
                                <Typography variant="h7" alignItems='center'>Signup</Typography>
                            </Link>
                        </Box>
                    </Paper>
                </Grid>
            </Grid>
        </Container>
    )

}
export default Login;