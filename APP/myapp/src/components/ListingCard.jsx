// src/components/ListingCard.js
import React, {useEffect, useState} from 'react';
import { Card, CardMedia, CardContent, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import ImageService from "../services/ImageService";
import placeholderImage from '../assets/placeholder.png';


const ListingCard = ({ listing }) => {

    const navigate = useNavigate();
    const [images, setImages] = useState([]);

    useEffect(() => {
        const fetchImages = async () => {
            const imagesFromServer = await ImageService.getImagesByListingId(listing.listing_id)
            //console.log(imagesFromServer[0].image_data)
            setImages(imagesFromServer)
        }
        fetchImages()
    }, []);


    // Complete onClick function to navigate to the detailed listing page
    const handleCardClick = () => {
        navigate(`/listing/${listing.listing_id}`);
    };

    return (
        <Card sx={{ minWidth: 300, maxWidth: 300, flexShrink: 0, backgroundColor : '#FEFEFE'}}>
            {/* Display the image with a placeholder for click-to-open modal */}
            <CardMedia
                component="img"
                height="140"
                image={images.length > 0 ? `data:image/png;base64,${images[0].image_data}` : placeholderImage}
                // Make sure base64 data is properly formatted
                alt={listing.title}
                sx={{ cursor: 'pointer' }}
                onClick={handleCardClick}
            />
            <CardContent>
                <Typography variant="h6">{listing.title}</Typography>
                <Typography variant="body2" color="text.secondary">
                    {listing.description}
                </Typography>
                <Typography variant="subtitle2" color="primary">
                    {/* Link to user profile */}
                    <Link component={RouterLink} to={`/profile/${listing.user_id}`}>
                        {listing.username}
                    </Link>
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    {listing.country + "    "}
                    {listing.city}
                </Typography>
            </CardContent>
        </Card>
    );
};

export default ListingCard;
