// src/components/HorizontalHeader.js
import React from 'react';
import { Link } from 'react-router-dom';
import '../App.css'; // Import the CSS for styling

const Header = () => {
    return (
        <div className="horizontal-header">
            <nav>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/login">Login</Link></li>
                    <li><Link to="/dashboard">Dashboard</Link></li>
                    {/* Add more links as needed */}
                </ul>
            </nav>
        </div>
    );
};

export default Header;
