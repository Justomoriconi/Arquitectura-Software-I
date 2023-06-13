import React, { useEffect, useState } from "react";
import { Typography, Table, TableHead, TableBody, TableRow, TableCell, Box } from "@mui/material";

const Reserves = () => {
  const [reserves, setReserves] = useState([]);
 // const [hotels, setHotels] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/mybookings/", {
      credentials: 'include',
    })
      .then(response => response.json())
      .then(data => {
        setReserves(data);
        console.log(data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);
/*
  const getHotelsByIds = (ids) => {
    Promise.all(ids.map(id => fetch(`http://localhost:8080/hotel/id/${id}`, { credentials: 'include' })
      .then(response => response.json())))
      .then(data => {
        setHotels(data);
        console.log(data);
      })
      .catch(error => {
        console.error(error);
      });
  };*/

  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', minHeight: '100vh' }}>
      <div>
        <Typography variant="h4" sx={{ textAlign: 'center', marginBottom: '20px' }}>Tus Reservas</Typography>
        {reserves.length > 0 ? (
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ID Reserve</TableCell>
                <TableCell>ID Hotel</TableCell>
                <TableCell>Check In</TableCell>
                <TableCell>Check Out</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {reserves.map((reserve) => {
                return (
                  <TableRow key={reserve.id}>
                    <TableCell>{reserve.id}</TableCell>
                    <TableCell>{reserve.HotelID}</TableCell>
                    <TableCell>{reserve.checkin}</TableCell>
                    <TableCell>{reserve.checkout}</TableCell>
                  </TableRow>
                );
              })}
            </TableBody>
          </Table>
        ) : (
          <Typography variant="body1" align="center">No tienes reservas.</Typography>
        )}
      </div>
    </Box>
  );
};

export default Reserves;
