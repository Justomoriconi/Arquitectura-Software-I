import React, { useState } from "react";
import MediaCard from "../componets/mediaCard";
import { Container, Grid} from '@mui/material';




const Home = () =>{

    const [hotel] = useState([
        {
            id: "1", 
            name: "Gardi Hotel & Suites",
            description:"El Gardi Hotel & Suites se encuentra en Buenos Aires, a 400 metros del teatro Colón, y ofrece alojamiento con bar, estacionamiento privado y jardín.",
        },
        
        {
            id: "2", 
            name: "Ayres Recoleta Uriburu",
            description:"El Ayres Recoleta Uriburu se encuentra en Buenos Aires, a 400 metros del cementerio de la Recoleta, y ofrece alojamiento con salón compartido, wifi gratis, centro de negocios y servicio de conserjería. Se encuentra a menos de 1 km del Museo Nacional de Bellas Artes y ofrece un mostrador de información turística.",
        },
        
        {
            id: "3", 
            name: "Clayton Buenos Aires",
            description:"El Clayton Buenos Aires goza de una buena ubicación en el barrio de Palermo de Buenos Aires, a menos de 1 km de Bosques de Palermo, a 8 minutos a pie de los lagos de Palermo y a menos de 1 km del parque El Rosedal. El alojamiento se encuentra a 1,5 km del jardín japonés de Buenos Aires y a 2,7 km del Museo de Arte Latinoamericano de Buenos Aires MALBA y de la plaza Serrano. Hay recepción 24 horas, salón compartido y servicio de cambio de divisa.",
        },
        
        {
            id: "4", 
            name: "Trendy & Luxury",
            description:"El Trendy & Luxury se encuentra en Buenos Aires, en el barrio de Palermo, a 500 metros de la plaza Serrano y a 2,3 km del parque ecológico de Buenos Aires, y ofrece jardín. El departamento tiene terraza.",
        },
        
        {
            id: "5", 
            name: "Hotel Regis",
            description:"El Hotel Regis está bien situado en el centro de Buenos Aires, a 800 metros del Obelisco de Buenos Aires, a menos de 1 km del teatro Colón y a 13 minutos a pie de la basílica del Santísimo Sacramento. El alojamiento cuenta con recepción 24 horas, servicio de conserjería y wifi gratis en todas las instalaciones. El hotel dispone de habitaciones familiares.",
        }
]);




    return (
        <Container maxWidth="x2">
            <Grid 
                container
                direction="column"
                alignItems="center"
                justifyContent="center"
                sx={{minHeight: "75vh"}}
            >
            {hotel.map( (hotel) => {
                console.log(hotel);
                return( <MediaCard 
                            key={hotel.id}
                            title="Foto de hotel"
                            name={hotel.name} 
                            description={hotel.description}  
                        />
                );
            })}
         </Grid>
        </Container>
    );
}

export default Home;