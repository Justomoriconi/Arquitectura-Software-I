import React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import { Link } from 'react-router-dom';



export default function MediaCard(props) {
  return (
    <Card sx={{ maxWidth: 1000 }}>
      <CardMedia
        component="img"
        image={require('../images/hotel.jpg').default}
        alt="Hotel"
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {props.name}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {props.description}
        </Typography>
      </CardContent>
      <CardActions>
        <Link to={`/hotel/id/${props.id}`}>
          <Button size="small" >Book now</Button>
        </Link>
      </CardActions>
    </Card>
  );
}
