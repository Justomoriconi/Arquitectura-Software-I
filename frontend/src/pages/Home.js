import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Container, Grid, Card, CardContent, Typography, TextField, Button, Box } from '@mui/material';

const Home = () => {
  const navigate = useNavigate();
  const [checkIn, setCheckIn] = useState("");
  const [checkOut, setCheckOut] = useState("");
  const [availableHotels, setAvailableHotels] = useState([]);

  const [expandedHotelId, setExpandedHotelId] = useState(null);

  const handleCheckInChange = (event) => {
    setCheckIn(event.target.value);
  };

  const handleCheckOutChange = (event) => {
    setCheckOut(event.target.value);
  };

  const handleSearch = () => {
    fetch('http://localhost:8080/availablehotels/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "Checkin": checkIn,
        "Checkout": checkOut,
      })
    })
      .then(response => response.json())
      .then(data => {
        // Procesar la respuesta del servidor
        setAvailableHotels(data);
      })
      .catch(error => {
        // Manejar el error de la solicitud
        console.error(error);
      });
  };

  const handleBookNow = (hotelId) => {
    const reservationData = {
      "Hotelid": hotelId.toString(),
      "Checkin": checkIn,
      "CheckOut": checkOut
    };

    fetch('http://localhost:8080/reserve/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include',
      body: JSON.stringify(reservationData)
    })
      .then(response => {
        if (response.status === 200) {
          console.log('Hotel reserved successfully!');
          setCheckIn("");
          setCheckOut("");
          navigate("/reserves");

        } else {

          console.log('Reservation failed. Please try again.');
        }
      })
      .then(data => {
        // Procesar la respuesta del servidor
        console.log(data);
      })
      .catch(error => {
        // Manejar el error de la solicitud
        console.error(error);
        console.log('An error occurred. Please try again.');
      });
  };


  return (
    <Container maxWidth="md" sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '100vh' }}>
      <Grid item xs={12} sm={6}>
        <Card>
          <CardContent>
            <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '500px' }}>
              <Typography variant="h6" component="div">
                Search Available Hotels
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Check In
              </Typography>
              <TextField
                id="check-in-date"
                type="date"
                fullWidth
                value={checkIn}
                onChange={handleCheckInChange}
              />
              <Typography variant="body2" color="text.secondary" sx={{ mt: 2 }}>
                Check Out
              </Typography>
              <TextField
                id="check-out-date"
                type="date"
                fullWidth
                value={checkOut}
                onChange={handleCheckOutChange}
              />
              <Box sx={{ mt: 2 }}>
                <Button onClick={handleSearch} variant="contained" color="primary" fullWidth>
                  Search
                </Button>
              </Box>
            </Box>
          </CardContent>
        </Card>
      </Grid>
      {availableHotels.length > 0 && (
        <Grid item xs={12} sm={6}>
          <Card>
            <CardContent>
              <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '500px' }}>
                <Typography variant="h6" component="div">
                  Available Hotels
                </Typography>
                {availableHotels.map(hotel => (
                  <Box
                    key={hotel.id}
                    sx={{
                      marginTop: '10px',
                      backgroundColor: '#f5f5f5',
                      padding: '10px',
                      borderRadius: '5px',
                    }}
                  >
                    <Typography variant="h6" component="div">
                      {hotel.Name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Rooms: {hotel.rooms}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                      Description: {expandedHotelId === hotel.id ? hotel.Description : `${hotel.Description.slice(0, 100)}...`}
                    </Typography>
                    {hotel.Description.length > 100 && (
                      <Button
                        variant="text"
                        color="primary"
                        sx={{ marginTop: '10px' }}
                        onClick={() => setExpandedHotelId(hotel.id)}
                      >
                        {expandedHotelId === hotel.id ? 'Show Less' : 'Read More'}
                      </Button>
                    )}
                    <Button
                      variant="contained"
                      color="primary"
                      sx={{ marginTop: '10px' }}
                      onClick={() => handleBookNow(hotel.id)}
                    >
                      Book Now
                    </Button>
                  </Box>
                ))}
              </Box>
            </CardContent>
          </Card>
        </Grid>
      )}
    </Container>
  );
}

export default Home;
