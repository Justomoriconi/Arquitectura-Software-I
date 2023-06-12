import React, { useState, useEffect } from "react";
import MediaCard from "../componets/mediaCard";
import { Container, Grid} from '@mui/material';









const Home = () => {
  const url = "http://127.0.0.1:8080/hotels/";
  const [hotels, setHotels] = useState([]);

  const fetchApi = async () => {
      try {
          const response = await fetch(url);
          const responseJSON = await response.json();
          console.log(responseJSON);
          setHotels(responseJSON);
      } catch (error) {
          console.error(error);
      }
  }

  useEffect(() => {
      fetchApi();
  }, []);

  return (
    <Container maxWidth="x2" sx={{ marginBottom: 4 }}>
      <Grid
        container
        direction="column"
        alignItems="center"
        justifyContent="center"
        sx={{ minHeight: "75vh" }}
      >
        {hotels && hotels.map((hotel, index) => (
          <Grid item key={hotel.id} sx={{ marginTop: index !== 0 ? 4 : 2 }}>
            <MediaCard
              image={require('../images/hotel.jpg').default}
              title="Foto de hotel"
              id={hotel.id}
              name={hotel.Name}
              description={hotel.Description}
              sx={{
                marginBottom: "2rem",
                marginTop: index === 0 ? "2rem" : 0, // Margen superior solo para el primer elemento
              }}
            />
          </Grid>
        ))}
      </Grid>
    </Container>
  );
  
}

export default Home;