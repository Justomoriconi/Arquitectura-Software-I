import { Button } from "@mui/material";
import React, { useState } from "react";
import { Link } from 'react-router-dom';







const Home = () =>{


    return (

        <Link to="/search">
            <Button >Search</Button>
        </Link>
    );
}

export default Home;