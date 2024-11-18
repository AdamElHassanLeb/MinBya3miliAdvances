import React, { useContext } from 'react';
import { Route, Routes, Navigate, useLocation } from 'react-router-dom';
import Login from './components/Login';
import Header from './static/Header';
import Home from './components/Home';
import Offers from './components/Offers';
import Requests from './components/Requests';
import CreateListing from './components/CreateListing';
import { UserContext, UserProvider } from './utils/UserContext';
import ListingDetail from "./components/ListingDetails";

// PrivateRoutes wrapper to protect multiple routes at once
const PrivateRoutes = ({ children }) => {
    const { user } = useContext(UserContext);

    // Redirect to /Login if not logged in
    return user ? children : <Navigate to="/Login" />;
};

function App() {
    const location = useLocation();

    return (
        <>
            <div className="App">
                {/* Conditionally render the Header only if not on the /Login route */}
                {location.pathname !== '/Login' && <Header />}
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
                                <Route path="/listing/:listingId" element={<ListingDetail/>} />
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
