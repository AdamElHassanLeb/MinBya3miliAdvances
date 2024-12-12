import React, {useState, useEffect, useContext} from 'react';
import {
    Box,
    Typography,
    Avatar,
    Grid,
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';
import {useNavigate, useParams} from 'react-router-dom';
import UserService from '../../services/UserService';
import ListingService from '../../services/ListingService';
import ListingCard from '../../components/Listings/ListingCard'; // Assuming there's a ListingCard component
import { MapContainer, TileLayer, Marker } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import {UserContext} from "../../utils/UserContext";
import serverAddress from "../../utils/ServerAddress";

const UserPublicProfile = () => {
    const { userId } = useParams(); // Get userId from URL
    const [userData, setUserData] = useState(null);
    const [listings, setListings] = useState([]);
    const [selectedLocation, setSelectedLocation] = useState([0, 0]); // Default coordinates
    const [listingType, setListingType] = useState("Any");
    const navigate = useNavigate()
    const { user } = useContext(UserContext);
    useEffect(() => {
        const fetchUserData = async () => {
            try {
                if(user && user.user_id == userId) {
                    navigate("/UserPrvateProfile")
                }

                const userDetails = await UserService.getUserById(userId);
                setUserData(userDetails);
                setSelectedLocation(userDetails.location ? [userDetails.location.latitude, userDetails.location.longitude] : [0, 0]);

                // Fetch user listings
                const userListings = await ListingService.getListingsByUserId(userId, listingType);
                setListings(userListings.data);
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData();
    }, [listingType]);

    if (!userData) return <div>Loading...</div>;

    function handleChangeType(event) {
        setListingType(event.target.value);
    }

    return (
        <Box
            sx={{
                maxWidth: '1000px',
                margin: 'auto',
                padding: '20px',
                border: '1px solid #ccc',
                borderRadius: '10px',
                boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                marginTop: '10vh'
            }}
            className="MUIBox"
        >
            {/* Profile Picture */}
            <Box sx={{ display: 'flex', justifyContent: 'center', marginBottom: '20px' }}>
                <Avatar
                    src={serverAddress() + `/api/v1/image/imageId/${userData.image_id}` || `../assets/default-avatar.png`}
                    alt="Profile Picture"
                    sx={{ width: 150, height: 150 }}
                />
            </Box>

            {/* User Details */}
            <Grid container spacing={2}>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>First Name:</strong> {userData.first_name || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>Last Name:</strong> {userData.last_name || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>Phone:</strong> {userData.phone_number || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>Profession:</strong> {userData.profession || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>Date of Birth:</strong> {userData.date_of_birth || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>City:</strong> {userData.loc_details.city || 'N/A'}
                    </Typography>
                </Grid>
                <Grid item xs={4}>
                    <Typography variant="body2">
                        <strong>Country:</strong> {userData.loc_details.country || 'N/A'}
                    </Typography>
                </Grid>
            </Grid>
            {/*

            {selectedLocation && (
                <Grid item xs={12} sx={{ marginTop: '20px' }}>
                    <Typography variant="body2" gutterBottom>
                        <strong>Location:</strong>
                    </Typography>
                    <MapContainer
                        center={selectedLocation}
                        zoom={13}
                        style={{ height: '300px', width: '100%' }}
                    >
                        <TileLayer
                            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                        />
                        <Marker position={selectedLocation} />
                    </MapContainer>
                </Grid>
            )}
            */}
            {/* Scrollable Collection of User Listings */}
            <Box sx={{ marginTop: '20px' }}>
                <Typography variant="h6" gutterBottom>User Listings</Typography>

                <FormControl sx = {{width : "10vw", maxWidth:"200px"}}>
                    <InputLabel id="demo-simple-select-label" sx={{ marginBottom: '5vh' }}>Type</InputLabel>
                    <Select
                        labelId="demo-simple-select-label"
                        id="demo-simple-select"
                        label="Type"
                        onChange={handleChangeType}
                        defaultValue={"Any"}
                    >
                        <MenuItem value={"Any"}>Any</MenuItem>
                        <MenuItem value={"Offer"}>Offers</MenuItem>
                        <MenuItem value={"Request"}>Requests</MenuItem>
                    </Select>
                </FormControl>

                <Box
                    sx={{
                        display: 'flex',
                        flexWrap: 'wrap',
                        gap: '50px', // Space between items
                        maxHeight: '400px',
                        overflowY: 'auto',
                        marginLeft: '50px',
                    }}
                >
                    {listings.length > 0 ? (
                        listings.map((listing, index) => (
                            <ListingCard key={index} listing={listing} />
                        ))
                    ) : (
                        <Typography variant="body2">No listings available.</Typography>
                    )}
                </Box>
            </Box>
        </Box>
    );
};

export default UserPublicProfile;
