import React, { useContext, useEffect, useState } from 'react';
import {Box, TextField, Button, FormControl, InputLabel, Select, MenuItem, Typography, Slider} from '@mui/material';
import { UserContext } from '../utils/UserContext';
import ListingCard from './ListingCard';
import ListingService from '../services/ListingService';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import UserService from "../services/UserService";
import ImageService from "../services/ImageService";

const ScrollableListings = (listingType) => {
    const { user } = useContext(UserContext); // Access user data from context
    const [listings, setListings] = useState([]);
    const [loading, setLoading] = useState(true);
    const [searchQuery, setSearchQuery] = useState(''); // For search query input
    const [filterType, setFilterType] = useState(''); // For dropdown filter type (distance or date)
    const [selectedLocation, setSelectedLocation] = useState([0, 0]); // For map's selected location (longitude, latitude)
    const [maxDistance, setMaxDistance] = useState(60);

    // Function to handle location click on the map
    const MapEvents = () => {
        useMapEvents({
            click(e) {
                const { lat, lng } = e.latlng;
                setSelectedLocation([lat, lng]); // Set new location on map click
            },
        });
        return null;
    };

    // Fetch listings only on first render
    useEffect(() => {
        if (user) {
            const fetchListings = async () => {
                try {
                    const longitude = user.location[0];
                    const latitude = user.location[1];
                    const maxDistance = 60; // Set your desired max distance in km

                    // Fetch listings based on the user's location and distance
                    const response = await ListingService.getListingsByDistance(longitude, latitude, maxDistance, listingType.listingType);
                    console.log(longitude, latitude, maxDistance, listingType.listingType)
                    setListings(response.data); // Assuming response contains listings data
                    setLoading(false);
                } catch (error) {
                    console.error("Error fetching listings:", error);
                    setLoading(false);
                }
            };

            fetchListings();
        }
    }, []);

    // Function for handling search button click
    const handleSearch = async () => {
        console.log('Search button clicked');
        console.log('Search Query:', searchQuery);
        console.log('Filter Type:', filterType);
        console.log('Selected Location:', selectedLocation);
        setListings([])
        var currListings;
        if (filterType === "date"){
            if(searchQuery === ""){
                currListings = await ListingService.getListingsByDate(listingType)
                setListings(currListings.data)
                return;
            }
            currListings = await ListingService.getListingsByDateAndSearch(searchQuery, listingType)
            setListings(currListings.data)
            return
        }
        //distance
        if(searchQuery === ""){

            currListings = await  ListingService.getListingsByDistance(selectedLocation[0], selectedLocation[1], maxDistance, listingType)
            setListings(currListings.data)
            return;
        }

        currListings = await ListingService.getListingsByDistanceAndSearch(selectedLocation[0], selectedLocation[1], maxDistance, listingType, searchQuery)
        setListings(currListings.data)
    };

    return (
        <Box sx={{ display: 'flex', gap: 2 }}>
            {/* Left Panel with Search Bar, Button, and Filter Dropdown */}
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: 2,
                    padding: 2,
                    width: '300px',
                    backgroundColor: 'white',
                    borderRadius: '8px',
                    boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                    position: 'sticky',
                    top: '10px',
                }}
            >
                {/* Search Bar */}
                <TextField
                    label="Search"
                    variant="outlined"
                    fullWidth
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                />

                {/* Search Button */}
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleSearch}
                    fullWidth
                >
                    Search
                </Button>

                {/* Dropdown Menu for Distance or Date */}
                <FormControl fullWidth>
                    <InputLabel>Sort By</InputLabel>
                    <Select
                        value={filterType}
                        onChange={(e) => setFilterType(e.target.value)}
                        label="Filter By"
                    >
                        <MenuItem value="distance">Distance</MenuItem>
                        <MenuItem value="date">Date</MenuItem>
                    </Select>
                </FormControl>


                {filterType === 'distance' && (
                    <>
                        <Typography gutterBottom>Max Distance (0 - 250 km)</Typography>
                        <Slider
                            value={maxDistance}
                            onChange={(e, newValue) => setMaxDistance(newValue)}
                            aria-labelledby="max-distance-slider"
                            min={0}
                            max={250}
                            valueLabelDisplay="auto"
                        />
                    </>
                )}


                {/* Show Map for Distance Option */}
                {filterType === 'distance' && (
                    <Box sx={{ width: '100%', height: '300px' }}>
                        <MapContainer
                            center={selectedLocation}
                            zoom={13}
                            style={{ width: '100%', height: '100%' }}
                        >
                            <TileLayer
                                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                            />
                            <Marker position={selectedLocation}>
                                <Popup>Your selected location</Popup>
                            </Marker>
                            <MapEvents />
                        </MapContainer>
                    </Box>
                )}
            </Box>

            {/* Right Panel with Listings */}
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: 2,
                    padding: 2,
                    overflowY: 'auto',
                    maxHeight: '80vh',
                    alignItems: 'center',
                    justifyContent: 'center',
                    margin: '0 auto',
                    width: '100%',
                    maxWidth: '900px',
                    paddingTop: '150px',  // Use a relative unit for scalability
                }}
            >
                {loading ? (
                    <div>Loading...</div> // Show a loading message
                ) : listings.length > 0 ? (
                    listings.map((listing) => (
                        <ListingCard key={listing.listing_id} listing={listing} />
                    ))
                ) : (
                    <div>No listings found within the specified distance.</div>
                )}
            </Box>
        </Box>
    );
};

export default ScrollableListings;
