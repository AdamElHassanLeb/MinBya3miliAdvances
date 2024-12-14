// src/components/Login.js
import React, {useState, useEffect, useContext} from 'react';
import {
    TextField,
    Button,
    Typography,
    Container,
    Box,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Grid,
    IconButton
} from '@mui/material';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents} from 'react-leaflet';
import L from 'leaflet';
import UserService from '../services/UserService';
import 'leaflet/dist/leaflet.css';
import { toast } from 'react-toastify';
import {UserContext} from "../utils/UserContext";
import {useNavigate} from "react-router-dom";
import MapIcon from "../utils/Icons"
import LightModeIcon from "@mui/icons-material/LightMode";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import {Brightness3, Brightness5} from "@mui/icons-material";

const Login = ({ toggleTheme, isDarkMode }) => {
    const [openModal, setOpenModal] = useState(false);
    const [formData, setFormData] = useState({
        first_name: '',
        last_name: '',
        phone_number: '',
        date_of_birth: '',
        profession: '',
        location: '', // latitude, longitude
        password: '',
    });
    const navigate = useNavigate();
    const { setUser } = useContext(UserContext);

    const [userLocation, setUserLocation] = useState([0, 0]); // Default to [0, 0] if no location

    // Get user's current location using the Geolocation API
    useEffect(() => {
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition((position) => {
                const { latitude, longitude } = position.coords;
                setUserLocation([latitude, longitude]);
            });
        }
    }, []);

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    // Function to toggle modal
    const handleModalOpen = () => setOpenModal(true);
    const handleModalClose = () => setOpenModal(false);

    // Handle form submission
    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            // Make the login request
            const result = await UserService.login(formData.phone_number, formData.password);
/*
            switch (result.status) {
                case 200:
                    toast.success('Login Success');
                    break
                case 500:

                    toast.error('Login Failed');
                    break

                case 401:
                    toast.error('Incorrect Credentials');
                    break

                case 400:
                    toast.error('Bad Credentials');
            }

            console.log(result.status);
*/
            // Assuming the response contains the user object in data.user
            if (result.data && result.data.user) {
                // Set the user in the context (from UserContext)
                setUser(result.data.user);

                // Optionally, store the token in localStorage
                if (result.data.token) {
                    localStorage.setItem('token', result.data.token);
                }

                // Redirect or navigate to the home page
                navigate('/home');
            } else {
                toast.error('Login failed');
            }
        } catch (error) {
            toast.error('Login failed: ' + error.message);
        }
    };


    // Handle form submission
    const handleSignUp = async (e) => {
        e.preventDefault();

        try {

            const result = await UserService.signUp(formData.first_name,
                formData.last_name, formData.phone_number, formData.date_of_birth,
                formData.profession, [formData.location[1], formData.location[0]], formData.password);

            if(result.status == 201){
                toast.success('Signup successfully');
                navigate('/Login');
                handleModalClose();
            }

        } catch (error) {
            toast.error(`Signup Failed ` + error.message);
        }
    };




    // Custom hook to handle map interactions
    const MapEvents = () => {
        useMapEvents({
            click(e) {
                const { lat, lng } = e.latlng;
                setFormData((prevData) => ({
                    ...prevData,
                    location: [lat, lng],
                }));
                setUserLocation([lat, lng]);
            },
        });
        return null;
    };

    return (
        <Container maxWidth="xs">
            <Box
                sx={{
                    padding: 3,
                    borderRadius: 2,
                    boxShadow: 3,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    marginTop: 8,
                }}
                className="MUIContainer-root"
            >
                <IconButton color="inherit" onClick={toggleTheme} sx = {{marginRight: -40, marginTop : -1}} >
                    {!isDarkMode ? <LightModeIcon /> : <Brightness3 />} {/* Change based on current theme */}
                </IconButton>
                <Typography variant="h4" gutterBottom>
                    Login
                </Typography>

                {/* Login Form */}
                <TextField
                    placeholder="Phone Number"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    name="phone_number"
                    value={formData.phone_number}
                    onChange={handleChange}
                />
                <TextField
                    placeholder="Password"
                    type="password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    name="password"
                    value={formData.password}
                    onChange={handleChange}
                />
                <Button
                    variant="contained"
                    color="primary"
                    fullWidth
                    sx={{ marginTop: 2 }}
                    onClick={handleLogin}
                >
                    Login
                </Button>

                {/* Sign Up Button */}
                <Button
                    variant="outlined"
                    color="secondary"
                    fullWidth
                    sx={{ marginTop: 2 }}
                    onClick={handleModalOpen} // Open the sign-up modal
                >
                    Sign Up
                </Button>
            </Box>

            {/* Sign Up Modal */}
            <Dialog open={openModal} onClose={handleModalClose}>
                <DialogTitle
                    className="MUIContainer-root">Sign Up</DialogTitle>
                <DialogContent
                    className="MUIContainer-root">
                    <Grid container spacing={2}>
                        {/* First Name Input */}
                        <Grid item xs={6}>
                            <TextField
                                label="First Name"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="first_name"
                                value={formData.first_name}
                                onChange={handleChange}
                            />
                        </Grid>
                        {/* Last Name Input */}
                        <Grid item xs={6}>
                            <TextField
                                label="Last Name"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="last_name"
                                value={formData.last_name}
                                onChange={handleChange}
                            />
                        </Grid>

                        {/* Phone Number Input */}
                        <Grid item xs={12}>
                            <TextField
                                label="Phone Number"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="phone_number"
                                value={formData.phone_number}
                                onChange={handleChange}
                            />
                        </Grid>

                        {/* Date of Birth Input */}
                        <Grid item xs={12}>
                            <TextField
                                label="Date of Birth"
                                type="date"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="date_of_birth"
                                value={formData.date_of_birth}
                                onChange={handleChange}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                            />
                        </Grid>

                        {/* Profession Input */}
                        <Grid item xs={12}>
                            <TextField
                                label="Profession"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="profession"
                                value={formData.profession}
                                onChange={handleChange}
                            />
                        </Grid>

                        {/* Location Input (Map) */}
                        <Grid item xs={12}>
                            <Typography variant="h6">Select Your Location</Typography>
                            <MapContainer
                                center={userLocation}
                                zoom={13}
                                style={{ width: '100%', height: '300px', marginBottom: '16px' }}
                            >
                                <TileLayer
                                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                                />
                                <Marker position={userLocation} icon={MapIcon}>
                                    <Popup>Your current location</Popup>
                                </Marker>
                                <MapEvents />
                            </MapContainer>
                            <TextField
                                sx = {{color : "#ffffff"}}
                                label="Location (Latitude, Longitude)"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="location"
                                value={formData.location}
                                onChange={handleChange}
                                placeholder="Location will be set on map click"
                                disabled
                            />
                        </Grid>

                        {/* Password Input */}
                        <Grid item xs={12}>
                            <TextField
                                label="Password"
                                type="password"
                                variant="outlined"
                                fullWidth
                                margin="normal"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </Grid>
                    </Grid>
                </DialogContent>
                <DialogActions
                    className="MUIContainer-root">
                    <Button onClick={handleModalClose} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleSignUp} color="primary">
                        Sign Up
                    </Button>
                </DialogActions>
            </Dialog>
        </Container>
    );
};

export default Login;
