import React, { useState, useEffect } from "react";
import { useParams } from 'react-router-dom';
import { Container, Grid, Card, CardContent, CardMedia, Typography, TextField, Button, Box } from '@mui/material';

const BookNow = () => {
  const { hotelId } = useParams();
  const [hotel, setHotel] = useState(null);
  const [checkIn, setCheckIn] = useState("");
  const [checkOut, setCheckOut] = useState("");

  const fetchHotelById = async (id) => {
    try {
      const response = await fetch(`http://127.0.0.1:8080/hotel/id/${id}`);
      const data = await response.json();
      setHotel(data);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    if (hotelId) {
      fetchHotelById(hotelId);
    }
  }, [hotelId]);

  const handleCheckInChange = (event) => {
    const value = event.target.value;
    console.log("Check-in:", value);
    setCheckIn(value);
  };

  const handleCheckOutChange = (event) => {
    const value = event.target.value;
    console.log("Check-out:", value);
    setCheckOut(value);
  };

  return (
    <Container maxWidth="md">
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          minHeight: '100vh',
        }}
      >
        <Grid container spacing={2}>
          <Grid item xs={12} sm={6}>
            <Card>
              <CardMedia
                component="img"
                height="200"
                image="../images/hotel.jpg"
                alt="Hotel Image"
              />
              <CardContent>
                <Typography variant="h6" component="div">
                  {hotel && hotel.name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {hotel && hotel.description}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Card>
              <CardContent>
                <Box sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                  <Typography variant="h6" component="div">
                    Select Dates
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
                  <Typography variant="body2" color="text.secondary">
                    Check Out
                  </Typography>
                  <TextField
                    id="check-out-date"
                    type="date"
                    fullWidth
                    value={checkOut}
                    onChange={handleCheckOutChange}
                  />
                  <Button variant="contained" color="primary" fullWidth>
                    Book Now
                  </Button>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
}

export default BookNow;
