import React, { useContext } from 'react';
import { Route, Routes, useLocation, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Header from './static/Header';
import Home from './components/Home';
import { UserContext, UserProvider } from './utils/UserContext';

// PrivateRoute component to protect routes
const PrivateRoute = ({ element }) => {
    const { user } = useContext(UserContext);

    // If the user is not logged in, redirect to /Login
    if (!user) {
        return <Navigate to="/Login" />;
    }

    // If the user is logged in, return the route's element
    return element;
};

function App() {
    const location = useLocation();
    //const { user } = useContext(UserContext);

    return (
        <div className="App">
            {/* Conditionally render the Header only if not on the /login route */}
            {location.pathname !== '/Login' && <Header />}

            <Routes>
                {/* All routes below are protected by the PrivateRoute wrapper */}
                <Route path="/home" element={<PrivateRoute element={<Home />} />} />
                <Route path="/Login" element={<Login />} />
            </Routes>
        </div>
    );
}

export default function AppWrapper() {
    return (
        <UserProvider>
            <App />
        </UserProvider>
    );
}
