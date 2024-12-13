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
    Grid, IconButton
} from '@mui/material';
import { UserContext } from '../../utils/UserContext';
import ListingCard from './ListingCard';
import ListingService from '../../services/ListingService';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import UserService from "../../services/UserService";
import ImageService from "../../services/ImageService";
import {Link} from "react-router-dom";
import {ArrowBack, ArrowForward} from "@mui/icons-material";
import MapIcon from "../../utils/Icons";

const ScrollableListings = (listingType) => {
    const { user } = useContext(UserContext); // Access user data from context
    const [listings, setListings] = useState([]);
    const [loading, setLoading] = useState(true);
    const [searchQuery, setSearchQuery] = useState(''); // For search query input
    const [filterType, setFilterType] = useState('date'); // For dropdown filter type (distance or date)
    const [selectedLocation, setSelectedLocation] = useState([0, 0]); // For map's selected location (longitude, latitude)
    const [maxDistance, setMaxDistance] = useState(60);
    const [zoomLevel, setZoomLevel] = useState(1); // Default zoom level
    const [leftPanelVisible, setLeftPanelVisible] = useState(true);

    //Zoom
    const handleZoom = () => {
        const zoom = window.devicePixelRatio;
        setZoomLevel(zoom);
        if(zoom < 2) {
            setLeftPanelVisible(true);
            return
        }
        setLeftPanelVisible(false);
    };


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

    useEffect(() => {
        window.addEventListener('resize', handleZoom);
        handleZoom(); // Check zoom on initial load
        return () => window.removeEventListener('resize', handleZoom);
    }, []);

    // Function for handling search button click
    const handleSearch = async () => {
        setListings([])
        var currListings;
        if (filterType === "date"){
            if(searchQuery === ""){
                try {
                    currListings = await ListingService.getListingsByDate(listingType.listingType)
                    setListings(currListings.data)
                }catch(e){
                    setListings([])
                }
                return;
            }

            try {
                currListings = await ListingService.getListingsByDateAndSearch(searchQuery, listingType.listingType)
                setListings(currListings.data)
            }catch(e){
                setListings([])
            }
            return
        }
        //distance
        if(searchQuery === ""){

            try {
                currListings = await  ListingService.getListingsByDistance(selectedLocation[1], selectedLocation[0], maxDistance, listingType.listingType)
                if(currListings && currListings.data)
                    setListings(currListings.data)
            }catch(e){
                setListings([])
            }
            return;
        }

        try{
            currListings = await ListingService.getListingsByDistanceAndSearch(selectedLocation[1], selectedLocation[0], maxDistance, listingType.listingType, searchQuery)
            setListings(currListings.data)
        }catch(e){
            setListings([])
        }
    };

    return (<>

        <Box sx={{ display: 'flex', gap: 2, overflow : 'hidden' }}>
            {/* Left Panel with Search Bar, Button, and Filter Dropdown */}
            <Box
                sx={{
                    display: !leftPanelVisible ? 'none' : 'flex',
                    flexDirection: 'column',
                    gap: 2,
                    padding: 2,
                    width: '15vw',
                    height: `90vh`,
                    border: '2px solid #ccc',
                    borderRadius: '8px',
                    boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                    position: 'sticky',
                    top: '10px',
                    marginTop: '1vh',
                    overflow: 'auto',

                }}
            className="MUIBox">
                {/* Search Bar */}
                <TextField
                    sx={{
                        marginTop: '10vh',
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
                    <Box sx={{  width: '100%',
                        minHeight: '300px',
                        minWidth: '200px',
                        marginTop: '1vh',
                        marginBottom: '2vh',
                        overflow: 'hidden',}}>
                        <MapContainer
                            center={selectedLocation}
                            zoom={13}
                            style={{ width: '100%',
                                height: '100%',
                                position: 'relative',}}
                        >
                            <TileLayer
                                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                            />
                            <Marker position={selectedLocation} icon={MapIcon}>
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
                        marginTop : '2vh'
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
            {/* Toggle Button with Arrows */}
            <IconButton
                sx={{
                    display: zoomLevel < 2 ? '' : 'none',
                    position: 'absolute',
                    left: leftPanelVisible ? '15.3vw' : '0px', // Dynamically position button
                    top: '45vh',
                    zIndex: 10,
                    transform: 'translateY(-50%)',
                    backgroundColor: '',
                    borderRadius: '50%',
                }}
                className={"LeftTabOpenClose"}
                onClick={() => setLeftPanelVisible(!leftPanelVisible)} // Toggle visibility
            >
                {leftPanelVisible ? <ArrowBack /> : <ArrowForward />} {/* Toggle arrows */}
            </IconButton>

            {/* Right Panel with Listings */}
            <Grid
                container
                spacing={2.5}
                sx={{
                    padding: 3,
                    overflowY: 'auto',
                    maxHeight: '100vh',
                    margin: '0 auto',
                    width: '100%',
                    maxWidth: '1200px',
                    paddingTop: '100px',
                    alignContent: 'flex-start',
                }}
            >
                {loading ? (
                    <Grid item xs={12}>
                        <div>Loading...</div>
                    </Grid>
                ) : listings && listings.length > 0 ? (
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
        </>
    );
};

export default ScrollableListings;
