import React, { useContext, useState, useEffect } from 'react';
import { Route, Routes, Navigate, useLocation } from 'react-router-dom';
import Login from './Pages/Login';
import Header from './static/Header';
import Home from './Pages/Home';
import Offers from './Pages/Listings/Offers';
import Requests from './Pages/Listings/Requests';
import CreateListing from './Pages/Listings/CreateListing';
import { UserContext, UserProvider } from './utils/UserContext';
import ListingDetail from "./Pages/Listings/ListingDetails";
import UserPrivateProfile from "./Pages/Profile/UserPrivateProfile";
 // Light Mode CSS
// Dark Mode CSS (if needed for specific overrides)
import './StyleSheets/App.css'

const PrivateRoutes = ({ children }) => {
    const { user } = useContext(UserContext);

    // Redirect to /Login if not logged in
    return user ? children : <Navigate to="/Login" />;
};

function App() {

    const location = useLocation();
    const [isDarkMode, setIsDarkMode] = useState(false);

    // Toggle between Dark and Light mode
    const toggleTheme = () => {
        setIsDarkMode(!isDarkMode);
    };

    // Set the theme class on the body element
    useEffect(() => {
        if (isDarkMode) {
            document.body.classList.add('dark-mode');
        } else {
            document.body.classList.remove('dark-mode');

        }
    }, [isDarkMode]);

    return (
        <>
            <div className="App">
                {/* Conditionally render the Header only if not on the /Login route */}
                {location.pathname !== '/Login' && <Header toggleTheme={toggleTheme} />}
            </div>
            <Routes>
                {/* Public route */}
                <Route path="/Login" element={<Login />} />

                {/* Protected routes */}
                <Route
                    path="/*"
                    element={
                        <PrivateRoutes>
                            <Routes>
                                <Route path="/home" element={<Home />} />
                                <Route path="/Offers" element={<Offers />} />
                                <Route path="/Requests" element={<Requests />} />
                                <Route path="/CreateListing" element={<CreateListing />} />
                                <Route path="/listing/:listingId" element={<ListingDetail />} />
                                <Route path="/UserPrvateProfile" element={<UserPrivateProfile />} />
                            </Routes>
                        </PrivateRoutes>
                    }
                />
            </Routes>
        </>
    );
}

export default function AppWrapper() {
    return (
        <UserProvider>
            <App />
        </UserProvider>
    );
}
