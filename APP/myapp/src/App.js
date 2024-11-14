import React from 'react';
import { Link, Route, Routes, useLocation } from 'react-router-dom';
import Login from './components/Login';
import Header from './static/Header';
import Home from "./components/Home";

function App() {
    const location = useLocation(); // Get the current location

    return (
        <div className="App">
            {/* Conditionally render the Header only if not on the /login route */}
            {location.pathname !== '/login' && <Header />}

            <h1>Hello World</h1>
            <Link to={'/home'}>Login</Link>
            <Routes>
                <Route path="/home" component={<Home/>} />
                <Route path="/login" element={<Login />} />
            </Routes>
        </div>
    );
}

export default App;
