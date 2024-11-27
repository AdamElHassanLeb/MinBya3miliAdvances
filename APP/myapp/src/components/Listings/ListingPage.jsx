import React, { useContext, useEffect, useState } from 'react';
import {
    Box,
    TextField,
    Button,
    FormControl,
    InputLabel,
    Select,
    MenuItem,
    Typography,
    Slider,
    Grid
} from '@mui/material';
import { UserContext } from '../../utils/UserContext';
import ListingCard from './ListingCard';
import ListingService from '../../services/ListingService';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import UserService from "../../services/UserService";
import ImageService from "../../services/ImageService";
import {Link} from "react-router-dom";

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

                    setSelectedLocation([longitude, latitude])
                    // Fetch listings based on the user's location and distance
                    const response = await ListingService.getListingsByDate(listingType.listingType);
                    //console.log(longitude, latitude, maxDistance, listingType.listingType)
                    //console.log(response.data)
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
        //console.log('Search button clicked');
        //console.log('Search Query:', searchQuery);
        //console.log('Filter Type:', filterType);
        //console.log('Selected Location:', selectedLocation);
        setListings([])
        var currListings;
        if (filterType === "date"){
            if(searchQuery === ""){
                //console.log(listingType)
                currListings = await ListingService.getListingsByDate(listingType.listingType)
                setListings(currListings.data)
                return;
            }
            currListings = await ListingService.getListingsByDateAndSearch(searchQuery, listingType.listingType)
            setListings(currListings.data)
            return
        }
        //distance
        if(searchQuery === ""){

            currListings = await  ListingService.getListingsByDistance(selectedLocation[0], selectedLocation[1], maxDistance, listingType.listingType)
            console.log(currListings)
            if(currListings && currListings.data)
                setListings(currListings.data)
            return;
        }

        currListings = await ListingService.getListingsByDistanceAndSearch(selectedLocation[0], selectedLocation[1], maxDistance, listingType.listingType, searchQuery)
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
                    height: `100vh`,
                    backgroundColor: 'white',
                    border: '2px solid #ccc',
                    borderRadius: '8px',
                    boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                    position: 'sticky',
                    top: '10px',

                }}
            >
                {/* Search Bar */}
                <TextField
                    sx={{
                        marginTop: '7vh',
                    }}
                    label="Search"
                    variant="outlined"
                    fullWidth
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                />



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
                {/* Search Button */}
                <Button
                    variant="contained"
                    color="primary"
                    onClick={handleSearch}
                    fullWidth
                >
                    Search
                </Button>

                <Button
                    sx={{
                        padding: 2,
                        position: 'sticky',
                        marginTop : '10vh'
                    }}
                    component={Link}
                    variant="contained"
                    color="warning"
                    to="/CreateListing"
                    fullWidth
                >
                    Create Listing
                </Button>

            </Box>

            {/* Right Panel with Listings */}
            <Grid
                container
                spacing={2}
                sx={{
                    padding: 2,
                    overflowY: 'auto',
                    maxHeight: '80vh',
                    margin: '0 auto',
                    width: '100%',
                    maxWidth: '1200px',
                    paddingTop: '100px',
                }}
            >
                {loading ? (
                    <Grid item xs={12}>
                        <div>Loading...</div>
                    </Grid>
                ) : listings.length > 0 ? (
                    listings.map((listing) => (
                        <Grid item xs={12} sm={6} md={4} key={listing.listing_id}>
                            <ListingCard listing={listing} />
                        </Grid>
                    ))
                ) : (
                    <Grid item xs={12}>
                        <div>No listings found within the specified distance.</div>
                    </Grid>
                )}
            </Grid>
        </Box>
    );
};

export default ScrollableListings;
