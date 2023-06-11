import React, { useState } from 'react';
import Box from '@mui/material/Box';
import { Container, Grid, Paper, Button, Typography, TextField } from '@mui/material';
import { Link } from 'react-router-dom';




const Signup = () =>{
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
                                    type="name"
                                    id="name"
                                    sx={{mt: 1.5, mb: 1.5}}
                        />
                          <TextField  variant='outlined'
                                    margin='normal'
                                    required
                                    fullWidth
                                    name="lastName"
                                    label="Last Name"
                                    type="lastName"
                                    id="lastName"
                                    autoComplete='current-password'
                                    sx={{mt: 1.5, mb: 1.5}}
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
                        />
                    
                    <Button fullWidth type="submit" >Login</Button>
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