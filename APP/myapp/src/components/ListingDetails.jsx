import React, { useState, useEffect, useContext } from 'react';
import {useNavigate, useParams} from 'react-router-dom';
import ListingService from '../services/ListingService'; // Service to fetch listing details
import ImageService from '../services/ImageService'; // Service to fetch images
import UserService from '../services/UserService';  // Import the new UserService
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css'; // Import carousel styles
import { Box, Typography, Divider, Button, Dialog, DialogActions, DialogContent, DialogTitle } from '@mui/material';
import { UserContext } from "../utils/UserContext"; // Assuming you have a context for user data
import placeholderImage from '../assets/placeholder.png';

const ListingDetail = () => {
    const { listingId } = useParams(); // Get the listingId from URL params
    const [listing, setListing] = useState(null);
    const [images, setImages] = useState([]);
    const [listingUser, setListingUser] = useState(null);  // State to store user data
    const [openDeleteModal, setOpenDeleteModal] = useState(false); // State to control modal visibility
    const navigate = useNavigate();
    // Get the logged-in user from context
    const { user } = useContext(UserContext);  // Assuming currentUser contains the logged-in user's data

    useEffect(() => {
        const fetchListing = async () => {
            try {
                const listingData = await ListingService.getListingById(listingId);
                setListing(listingData.data);


                const imagesData = await ImageService.getImagesByListingId(listingId);
                setImages(imagesData);

                console.log(listingData.data)

                // Fetch user details using the userId from the listing
                if (listingData.data) {
                    const userData = await UserService.getUserById(listingData.data.user_id);
                    setListingUser(userData);
                }
            } catch (error) {
                console.error('Error fetching listing data:', error);
            }
        };

        fetchListing();
    }, [listingId]);

    if (!listing || !listingUser) return <div>Loading...</div>;  // Wait until both are fetched

    // Format the date_created string into a readable format
    const formattedDate = new Date(listing.date_created).toLocaleDateString();

    // Handle delete listing
    const handleDelete = async () => {
        try {
            await ListingService.deleteListing(listingId);
            //alert('Listing deleted successfully');
            // Optionally, navigate to another page after deletion
        } catch (error) {
            console.error('Error deleting listing:', error);
            alert('Failed to delete the listing');
        } finally {
            // Close the modal after the deletion attempt
            setOpenDeleteModal(false);
            navigate("/Home")
        }
    };

    const handleCloseModal = () => {
        setOpenDeleteModal(false); // Close the modal if the user cancels
    };

    const handleOpenModal = () => {
        setOpenDeleteModal(true); // Open the modal when the delete button is clicked
    };

    return (
        <Box
            sx={{
                maxWidth: '900px',
                margin: 'auto',
                padding: '20px',
                border: '1px solid #ccc',
                borderRadius: '10px',
                boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
            }}
        >

            {/* Image Carousel */}
            <Box sx={{ width: '100%', marginBottom: '20px' }}>
                <Carousel showThumbs={false} infiniteLoop autoPlay>
                    {images.length > 0 ? (
                        images.map((image, index) => (
                            <div key={index}>
                                <img
                                    src={`data:image/png;base64,${image.image_data}`}
                                    alt={`Listing Image ${index + 1}`}
                                    style={{ width: '100%', height: 'auto', objectFit: 'cover' }}
                                />
                            </div>
                        ))
                    ) : (
                        <div>
                            <img
                                src={placeholderImage} // Use the imported placeholder image
                                alt="Placeholder Image"
                                style={{ width: '100%', height: 'auto', objectFit: 'cover' }}
                            />
                        </div>
                    )}
                </Carousel>
            </Box>


            {/* Listing Title */}
            <Typography variant="h3" sx={{
                marginBottom: '15px',
                fontWeight: 'bold',
                color: '#333',
                textAlign: 'center',
            }}>
                {listing.title}
            </Typography>

            {/* Listing Description */}
            <Typography variant="body1" sx={{
                marginBottom: '20px',
                fontSize: '1.1rem',
                lineHeight: 1.5,
                color: '#555',
            }}>
                {listing.description}
            </Typography>

            {/* User Information */}
            <Divider sx={{ marginBottom: '15px' }} />
            <Typography variant="h6" sx={{ fontWeight: 'bold', marginBottom: '10px' }}>
                Seller Info:
            </Typography>
            <Typography variant="body1">Name: {listingUser.first_name + ' ' + listingUser.last_name}</Typography>
            <Typography variant="body1">City: {listingUser.loc_details.city}</Typography>
            <Typography variant="body1">Country: {listingUser.loc_details.country}</Typography>

            {/* Additional Details */}
            <Divider sx={{ marginBottom: '15px', marginTop: '20px' }} />
            <Typography variant="h6" sx={{ fontWeight: 'bold', marginBottom: '10px' }}>
                Listing Details:
            </Typography>
            <Typography variant="body1">Created: {formattedDate}</Typography>
            <Typography variant="body1">Listing Type: {listing.type === 'offer' ? 'Offer' : 'Request'}</Typography>
            <Typography variant="body1">City: {listing.city}</Typography>
            <Typography variant="body1">Country: {listing.country}</Typography>

            {/* Delete Button if the user is the owner of the listing */}
            {user && user.user_id === listing.user_id && (
                <Button
                    variant="contained"
                    color="error"
                    sx={{ marginTop: '20px' }}
                    onClick={handleOpenModal}
                >
                    Delete Listing
                </Button>
            )}

            {/* Delete Confirmation Modal */}
            <Dialog
                open={openDeleteModal}
                onClose={handleCloseModal}
            >
                <DialogTitle>Confirm Deletion</DialogTitle>
                <DialogContent>
                    <Typography>Are you sure you want to delete this listing?</Typography>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleCloseModal} color="primary">Cancel</Button>
                    <Button onClick={handleDelete} color="error">Delete</Button>
                </DialogActions>
            </Dialog>
        </Box>
    );
};

export default ListingDetail;
