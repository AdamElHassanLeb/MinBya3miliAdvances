import React, {useContext, useState} from 'react';
import { Box, Button, TextField, FormControl, InputLabel, Select, MenuItem, Typography } from '@mui/material';
import { MapContainer, TileLayer, Marker, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import axios from 'axios';
import {UserContext} from "../utils/UserContext";
import ListingService from "../services/ListingService";
import ImageService from "../services/ImageService";

const CreateListing = () => {
    const [listingType, setListingType] = useState('');
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [selectedLocation, setSelectedLocation] = useState([0, 0]);
    const [selectedImages, setSelectedImages] = useState([]);

    const {user} = useContext(UserContext)

    // Function to handle map click
    const MapEvents = () => {
        useMapEvents({
            click(e) {
                const { lat, lng } = e.latlng;
                setSelectedLocation([lat, lng]);
            },
        });
        return null;
    };

    const handleImageChange = (event) => {
        setSelectedImages(Array.from(event.target.files));
    };

    const handleImageUpload = async () => {
        if (selectedImages.length > 0) {
            const formData = new FormData();
            selectedImages.forEach((file, index) => {
                formData.append(`image${index + 1}`, file);
            });

            try {
                const response = await axios.post('/your-upload-endpoint', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                });
                console.log('Upload successful:', response.data);
            } catch (error) {
                console.error('Error uploading images:', error);
            }
        }
    };

    const handleFormSubmit = async (e) => {
        e.preventDefault();

        // Log form data for debugging
        console.log({
            listingType,
            title,
            description,
            location: selectedLocation,
        });

        // Check if location is not set, default to user's location
        if (selectedLocation[0] === 0 && selectedLocation[1] === 0) {
            setSelectedLocation([user.location[0], user.location[1]]);
        }

        try {
            // Create the listing
            const res = await ListingService.createListing({
                type: listingType,
                title,
                description,
                location: selectedLocation,
                user_id: user.user_id,
            });

            console.log("Listing creation response:", res);

            // If listing creation is successful, upload images
            if (res && res.data && res.data.listing_id) {
                const imageUploadResponse = await ImageService.uploadListingImage(res.data.listing_id, selectedImages);
                console.log("Image upload response:", imageUploadResponse);
            } else {
                console.error("Listing creation failed or no listing ID returned.");
            }
        } catch (error) {
            console.error("Error during form submission:", error);
        }
    };


    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, maxWidth: '600px', margin: 'auto', padding: 2 }}>
            <Typography variant="h4">Create Listing</Typography>

            <FormControl fullWidth>
                <InputLabel>Type</InputLabel>
                <Select
                    value={listingType}
                    onChange={(e) => setListingType(e.target.value)}
                    label="Type"
                >
                    <MenuItem value="offer">Offer</MenuItem>
                    <MenuItem value="request">Request</MenuItem>
                </Select>
            </FormControl>

            <TextField
                label="Title"
                variant="outlined"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                fullWidth
            />

            <TextField
                label="Description"
                variant="outlined"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                multiline
                rows={4}
                fullWidth
            />

            <Box sx={{ width: '100%', height: '300px', marginBottom: 2 }}>
                <MapContainer center={[51.505, -0.09]} zoom={13} style={{ width: '100%', height: '100%' }}>
                    <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
                    <Marker position={selectedLocation}>
                    </Marker>
                    <MapEvents />
                </MapContainer>
            </Box>

            {/* Button to select multiple images */}
            <Button
                variant="contained"
                component="label"
                color="secondary"
                fullWidth
            >
                Select Images
                <input
                    type="file"
                    multiple
                    hidden
                    onChange={handleImageChange}
                />
            </Button>


            <Button variant="contained" color="primary" onClick={handleFormSubmit}>
                Create Listing
            </Button>
        </Box>
    );
};

export default CreateListing;