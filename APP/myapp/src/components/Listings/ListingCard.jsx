// src/components/ListingCard.js
import React, {useEffect, useState} from 'react';
import { Card, CardMedia, CardContent, Typography, Link } from '@mui/material';
import { Link as RouterLink, useNavigate } from 'react-router-dom';
import ImageService from "../../services/ImageService";
import placeholderImage from '../../assets/placeholder.png';
import serverAddress from "../../utils/ServerAddress";


const ListingCard = ({ listing }) => {

    const navigate = useNavigate();
    const [images, setImages] = useState([]);

    useEffect(() => {
        const fetchImages = async () => {
            try {
                const imagesFromServer = await ImageService.getImagesByListingId(listing.listing_id)
                setImages(imagesFromServer)
            }catch (error) {
                console.error(error)
            }
        }
        fetchImages()
    }, []);


    // Complete onClick function to navigate to the detailed listing page
    const handleCardClick = () => {
        navigate(`/listing/${listing.listing_id}`);
    };

    return (<>
        {listing ? (
    <Card sx={{
        minWidth: 300, maxWidth: 300, flexShrink: 0, cursor: 'pointer',
        background: 'linear-gradient(to top, #d9d9d9, #6e8b9d);',
        borderRadius: 3,
    }}
          onClick={handleCardClick} className="zoom" elevation={20}>
        {/* Display the image with a placeholder for click-to-open modal */}
        <CardMedia
            component="img"
            height="140"
            image={images.length > 0 ? serverAddress() + `/api/v1/image/image/${images[0].url}` : placeholderImage}
            // Make sure base64 data is properly formatted
            alt={listing.title}

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
            ): null
        }
        </>
    );
};

export default ListingCard;
