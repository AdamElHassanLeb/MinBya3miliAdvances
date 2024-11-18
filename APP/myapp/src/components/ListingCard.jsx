// src/components/ListingCard.js
import React, {useEffect, useState} from 'react';
import { Card, CardMedia, CardContent, Typography, Link } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import ImageService from "../services/ImageService";

const ListingCard = ({ listing }) => {

    const [images, setImages] = useState([]);

    useEffect(() => {
        const fetchImages = async () => {
            const imagesFromServer = await ImageService.getImagesByListingId(listing.listing_id)
            //console.log(imagesFromServer[0].image_data)
            setImages(imagesFromServer)
        }
        fetchImages()
    }, []);



    return (
        <Card sx={{ minWidth: 300, maxWidth: 300, flexShrink: 0 }}>
            {/* Display the image with a placeholder for click-to-open modal */}
            <CardMedia
                component="img"
                height="140"
                image={`data:image/png;base64,${images[0].image_data}`} // Make sure base64 data is properly formatted
                alt={listing.title}
                sx={{ cursor: 'pointer' }}
                onClick={() => console.log(`Open modal for listing ID: ${listing.listing_id}`)}
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
