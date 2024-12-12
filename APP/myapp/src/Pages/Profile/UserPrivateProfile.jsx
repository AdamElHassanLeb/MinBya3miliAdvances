import React, { useState, useEffect, useContext } from 'react';
import {
    Box,
    Typography,
    Avatar,
    Divider,
    Grid,
    Button,
    TextField,
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    FormControl, InputLabel, Select, MenuItem
} from '@mui/material';
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css'; // Import carousel styles
import UserService from '../../services/UserService';
import ListingService from '../../services/ListingService';
import { UserContext } from '../../utils/UserContext';
import ListingCard from '../../components/Listings/ListingCard'; // Assuming there's a ListingCard component
import { MapContainer, TileLayer, Marker, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import ImageService from "../../services/ImageService";
import {useNavigate} from "react-router-dom";
import serverAddress from "../../utils/ServerAddress";

const UserProfile = () => {
    const { user } = useContext(UserContext);
    const [userData, setUserData] = useState(null);
    const [listings, setListings] = useState([]);
    const [isEditing, setIsEditing] = useState(false);
    const [editableData, setEditableData] = useState({});
    const [deleteModalOpen, setDeleteModalOpen] = useState(false);
    const [selectedLocation, setSelectedLocation] = useState([0, 0]); // Default coordinates
    const [selectedImages, setSelectedImages] = useState([]);
    const [listingType, setListingType] = useState("Any");
    const navigate = useNavigate()

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const userDetails = await UserService.getUserById(user.user_id);
                setUserData(userDetails);
                setEditableData(userDetails);
                setSelectedLocation(userDetails.location || [0, 0]); // Set initial map location

                // Fetch user listings
                const userListings = await ListingService.getListingsByUserId(user.user_id, listingType);
                if(userListings && userListings.data)
                    setListings(userListings.data);
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData();
    }, [listingType]);


    function handleChangeType(event) {
        setListingType(event.target.value);
    }

    const handleImageChange = async (event) => {
        setSelectedImages(Array.from(event.target.files));
        //console.log(Array.from(event.target.files))
        try {
            const res = await ImageService.uploadProfileImage(userData.user_id, Array.from(event.target.files))
        }catch (error){

        }
       const user1 = await UserService.getUserById(userData.user_id)
        setEditableData(user1)
        setUserData(user1)
        navigate('/UserPrvateProfile')
    };

    const handleUpdateClick = () => {
        setIsEditing(true);
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setEditableData({ ...editableData, [name]: value });
    };

    const handleMapClick = (e) => {
        const { lat, lng } = e.latlng;
        console.log(lng, lat)
        setSelectedLocation([lat, lng]);
        setEditableData({ ...editableData, location: [lng, lat] });
    };

    const LocationPicker = () => {
        useMapEvents({
            click: handleMapClick
        });
        return null;
    };

    const handleSubmit = async () => {
        try {
            // Update user details
            const res = await UserService.updateUser(user.user_id, {
                ...editableData,
                location: selectedLocation
            });
            setUserData(editableData);
            setIsEditing(false);
        } catch (error) {
            console.error('Error updating user data:', error);
        }
    };

    const handleDeleteClick = () => {
        setDeleteModalOpen(true);
    };

    const handleDeleteConfirm = () => {
        console.log('Delete confirmed');
        setDeleteModalOpen(false);
    };

    const handleDeleteCancel = () => {
        setDeleteModalOpen(false);
    };

    if (!userData) return <div>Loading...</div>;

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
            className="MUIBox">
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
                {isEditing ? (
                    <>
                        <Grid item xs={6}>
                            <TextField
                                label="First Name"
                                name="first_name"
                                value={editableData.first_name}
                                onChange={handleInputChange}
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                label="Last Name"
                                name="last_name"
                                value={editableData.last_name}
                                onChange={handleInputChange}
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                label="Phone"
                                name="phone_number"
                                value={editableData.phone_number}
                                onChange={handleInputChange}
                                fullWidth
                            />
                        </Grid>

                        <Grid item xs={6}>
                            <TextField
                                label="Profession"
                                name="profession"
                                value={editableData.profession}
                                onChange={handleInputChange}
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                label="Date of Birth"
                                name="date_of_birth"
                                type="date"
                                value={editableData.date_of_birth}
                                onChange={handleInputChange}
                                fullWidth
                                InputLabelProps={{ shrink: true }}
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                label="Password"
                                name="password"
                                type="password"
                                value={editableData.password}
                                onChange={handleInputChange}
                                fullWidth
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <Typography variant="body2" gutterBottom>
                                <strong>Select Location on the Map:</strong>
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
                                <LocationPicker />
                            </MapContainer>
                        </Grid>
                    </>
                ) : (
                    <>
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
                                <strong>City:</strong> {userData.loc_details.country || 'N/A'}
                            </Typography>

                        <Grid item xs={4}>
                            <Typography variant="body2">
                                <strong>Country:</strong> {userData.loc_details.country || 'N/A'}
                            </Typography>
                        </Grid>
                        </Grid>
                    </>
                )}
            </Grid>

            {/* Buttons */}
            <Box sx={{ display: 'flex', justifyContent: 'space-between', marginTop: '20px' }}>
                <Button variant="contained" color="error" onClick={handleDeleteClick}>
                    Delete
                </Button>
                <Box sx={{ display: 'flex', justifyContent: 'center', marginBottom: '20px' }}>
                    <Button
                        variant="contained"
                        component="label"
                        color="info"
                        fullWidth
                    >
                        Select Image
                        <input
                            type="file"
                            hidden
                            onChange={handleImageChange}
                        />
                    </Button>
                </Box>
                {isEditing ? (
                    <Button variant="contained" color="primary" onClick={handleSubmit}>
                        Save
                    </Button>
                ) : (
                    <Button variant="contained" color="primary" onClick={handleUpdateClick}>
                        Update
                    </Button>
                )}
            </Box>

            {/* Delete Confirmation Modal */}
            <Dialog open={deleteModalOpen} onClose={handleDeleteCancel}>
                <DialogTitle>Confirm Delete</DialogTitle>
                <DialogContent>
                    <Typography>Are you sure you want to delete your profile?</Typography>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleDeleteCancel} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={handleDeleteConfirm} color="error" variant="contained">
                        Confirm
                    </Button>
                </DialogActions>
            </Dialog>

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
                        <MenuItem value={"Any"}>All</MenuItem>
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

export default UserProfile;
