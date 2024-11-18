// src/components/ListingCard.js
import React from 'react';
import { Card, CardMedia, CardContent, Typography, Link } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';

const ListingCard = ({ listing }) => {
    return (
        <Card sx={{ minWidth: 300, maxWidth: 300, flexShrink: 0 }}>
            {/* Display the image with a placeholder for click-to-open modal */}
            <CardMedia
                component="img"
                height="140"
                image={ '../assets/placeholder.png'} // Placeholder if no image
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
