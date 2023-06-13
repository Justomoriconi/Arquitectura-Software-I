import React, { createContext, useState, useEffect } from 'react';
import axios from 'axios';

const AuthContext = createContext();

const AuthProvider = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    // Recupera el estado de inicio de sesión almacenado en el localStorage
    const storedLoggedInStatus = localStorage.getItem('isLoggedIn');

    if (storedLoggedInStatus) {
      setIsLoggedIn(JSON.parse(storedLoggedInStatus));
    }
  }, []);

  const login = () => {
    // Realiza la lógica de inicio de sesión
    setIsLoggedIn(true);

    // Almacena el estado de inicio de sesión en el localStorage
    localStorage.setItem('isLoggedIn', JSON.stringify(true));
  };

  const logout = async () => {
    try {
      const response = await axios.post('http://localhost:8080/logout', null, {
        withCredentials: true,
      });

  
      if (response.status === 200) {
        // El cierre de sesión fue exitoso, realiza las acciones correspondientes
        console.log('Cierre de sesión exitoso');
      } else {
        // El cierre de sesión falló, muestra un mensaje de error o realiza las acciones correspondientes
        console.log('Cierre de sesión fallido');
      }
    } catch (error) {
      // Ocurrió un error durante la solicitud, muestra un mensaje de error o realiza las acciones correspondientes
      console.error(error);
    }

    // Realiza la lógica de cierre de sesión
    setIsLoggedIn(false);

    // Elimina el estado de inicio de sesión del localStorage
    localStorage.removeItem('isLoggedIn');
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
