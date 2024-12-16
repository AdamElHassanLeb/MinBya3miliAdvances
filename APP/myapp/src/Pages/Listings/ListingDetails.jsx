import React, { useState, useEffect, useContext } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import ListingService from '../../services/ListingService';
import ImageService from '../../services/ImageService';
import UserService from '../../services/UserService';
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import {
    Box,
    Typography,
    Divider,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    TextField,
    Avatar
} from '@mui/material';
import { UserContext } from '../../utils/UserContext';
import placeholderImage from '../../assets/placeholder.png';
import serverAddress from '../../utils/ServerAddress';
import DeleteIcon from '@mui/icons-material/Delete';
import UpdateIcon from '@mui/icons-material/Update';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import SendOfferModal from '../Transaction/SendOfferModal';

const ListingDetail = () => {
    const { listingId } = useParams();
    const navigate = useNavigate();
    const { user } = useContext(UserContext);

    const [listing, setListing] = useState(null);
    const [editableListing, setEditableListing] = useState(null);
    const [images, setImages] = useState([]);
    const [listingUser, setListingUser] = useState(null);
    const [isEditing, setIsEditing] = useState(false);
    const [openDeleteModal, setOpenDeleteModal] = useState(false);
    const [isOfferModalOpen, setIsOfferModalOpen] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const listingData = await ListingService.getListingById(listingId);
                setListing(listingData.data);
                setEditableListing(listingData.data);

                const imagesData = await ImageService.getImagesByListingId(listingId);
                setImages(imagesData);

                if (listingData.data) {
                    const userData = await UserService.getUserById(listingData.data.user_id);
                    setListingUser(userData);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, [listingId]);

    if (!listing || !listingUser) return <div>Loading...</div>;

    const formattedDate = new Date(listing.date_created).toLocaleDateString();

    const handleEditClick = () => setIsEditing(true);

    const handleSaveClick = async () => {
        try {
            await ListingService.updateListing(listingId, editableListing);
            setListing(editableListing);
            setIsEditing(false);
        } catch (error) {
            console.error('Error updating listing:', error);
        }
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setEditableListing({ ...editableListing, [name]: value });
    };

    const handleDelete = async () => {
        try {
            await ListingService.deleteListing(listingId);
            navigate('/Home');
        } catch (error) {
            console.error('Error deleting listing:', error);
        } finally {
            setOpenDeleteModal(false);
        }
    };

    return (
        <>
            <Box
                sx={{
                    maxWidth: '900px',
                    margin: 'auto',
                    padding: '20px',
                    border: '1px solid #ccc',
                    borderRadius: '10px',
                    boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                    marginTop: '10vh',
                }}
                className={"MUIBox"}
            >
                {/* Image Carousel */}
                <Box sx={{ width: '100%', marginBottom: '20px' }}>
                    <Carousel showThumbs={false} infiniteLoop autoPlay>
                        {images.length > 0 ? (
                            images.map((image, index) => (
                                <div key={index}>
                                    <img
                                        src={serverAddress() + `/api/v1/image/image/${image.url}`}
                                        alt={`Listing Image ${index + 1}`}
                                        style={{ width: '100%', height: 'auto', objectFit: 'cover' }}
                                    />
                                </div>
                            ))
                        ) : (
                            <img
                                src={placeholderImage}
                                alt="Placeholder"
                                style={{ width: '100%', height: 'auto', objectFit: 'cover' }}
                            />
                        )}
                    </Carousel>
                </Box>

                {/* Listing Title */}
                {isEditing ? (
                    <TextField
                        name="title"
                        label="Title"
                        value={editableListing.title}
                        onChange={handleInputChange}
                        fullWidth
                        sx={{ marginBottom: '15px' }}
                    />
                ) : (
                    <Typography variant="h3" sx={{
                        marginBottom: '15px',
                        fontWeight: 'bold',
                        color: '#333',
                        textAlign: 'center',
                    }}>
                        {listing.title}
                    </Typography>
                )}

                <Divider sx={{ marginBottom: '15px' }} />

                {/* Listing Description */}
                {isEditing ? (
                    <TextField
                        name="description"
                        label="Description"
                        value={editableListing.description}
                        onChange={handleInputChange}
                        fullWidth
                        multiline
                        rows={4}
                        sx={{ marginBottom: '20px' }}
                    />
                ) : (
                    <Typography variant="body1" sx={{
                        marginBottom: '20px',
                        fontSize: '1.1rem',
                        lineHeight: 1.5,
                        color: '#555',
                    }}>
                        {listing.description}
                    </Typography>
                )}

                {/* User Information */}
                <Divider sx={{ marginBottom: '15px' }} />
                <Typography variant="h6" sx={{ fontWeight: 'bold', marginBottom: '10px' }}>
                    Seller Info:
                </Typography>
                <Avatar
                    src={serverAddress() + `/api/v1/image/imageId/${listingUser.image_id}` || placeholderImage}
                    alt="Profile Picture"
                    sx={{ width: 70, height: 70 }}
                    component={Link}
                    to={`/User/${listingUser.user_id}`}
                />
                <Typography variant="body1">Name: {listingUser.first_name} {listingUser.last_name}</Typography>
                <Typography variant="body1">City: {listingUser.loc_details.city}</Typography>
                <Typography variant="body1">Country: {listingUser.loc_details.country}</Typography>

                {/* Listing Details */}
                <Divider sx={{ marginBottom: '15px', marginTop: '20px' }} />
                <Typography variant="h6" sx={{ fontWeight: 'bold', marginBottom: '10px' }}>
                    Listing Details:
                </Typography>
                <Typography variant="body1">Created: {formattedDate}</Typography>
                <Typography variant="body1">Listing Type: {listing.type === 'offer' ? 'Offer' : 'Request'}</Typography>
                <Typography variant="body1">City: {listing.city}</Typography>
                <Typography variant="body1">Country: {listing.country}</Typography>

                {/* Actions */}
                <Box sx={{ display: 'flex', gap: '10px', marginTop: '20px' }}>
                    {user?.user_id === listing.user_id ? (
                        isEditing ? (
                            <Button variant="contained" color="primary" onClick={handleSaveClick}>
                                Save
                            </Button>
                        ) : (
                            <Button
                                variant="contained"
                                color="primary"
                                onClick={handleEditClick}
                                startIcon={<UpdateIcon />}
                            >
                                Update Listing
                            </Button>
                        )
                    ) : (
                        <Button
                            variant="contained"
                            color="primary"
                            startIcon={<LocalOfferIcon />}
                            onClick={() => setIsOfferModalOpen(true)}
                        >
                            Send Offer
                        </Button>
                    )}

                    {user?.user_id === listing.user_id && (
                        <Button
                            variant="contained"
                            color="error"
                            onClick={() => setOpenDeleteModal(true)}
                            startIcon={<DeleteIcon />}
                        >
                            Delete Listing
                        </Button>
                    )}
                </Box>

                {/* Delete Confirmation Modal */}
                <Dialog open={openDeleteModal} onClose={() => setOpenDeleteModal(false)}>
                    <DialogTitle>Confirm Deletion</DialogTitle>
                    <DialogContent>
                        <Typography>Are you sure you want to delete this listing?</Typography>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={() => setOpenDeleteModal(false)} color="primary">Cancel</Button>
                        <Button onClick={handleDelete} color="error">Delete</Button>
                    </DialogActions>
                </Dialog>
            </Box>

            {/* Offer Modal */}
            <SendOfferModal
                isOpen={isOfferModalOpen}
                onClose={() => setIsOfferModalOpen(false)}
                User={user}
                OfferingUser={listingUser}
                Listing={listing}
            />
        </>
    );
};

export default ListingDetail;
