//import { Button } from "@mui/material";
//import { Link } from 'react-router-dom';
import React, { useState, useEffect } from "react";
import MediaCard from "../componets/mediaCard";
import { Container, Grid} from '@mui/material';









const Home = () =>{
   
    
        
        const url = "http://127.0.0.1:8080/hotels/";
        
        const [hotel, setHotel] = useState();
      
        const fetchApi = async () => {
          const response = await fetch(url);
          const responseJSON = await response.json();
          console.log(responseJSON);
          setHotel(responseJSON);
        }
      
        useEffect(() => {
          fetchApi();
        }, []);



    return (
        <Container maxWidth="x2">
        <Grid 
            container
            direction="column"
            alignItems="center"
            justifyContent="center"
            sx={{minHeight: "75vh"}}
        >
        {hotel && hotel.map( (hotel) => {
          console.log(hotel);
          return( <MediaCard 
                      key={hotel.id}
                      image={require('../images/hotel.jpg').default}
                      title="Foto de hotel"
                      id={hotel.id}
                      name={hotel.Name} 
                      description={hotel.Description}  
                  />
            );
        })}
     </Grid>
    </Container>
    );
}

export default Home;