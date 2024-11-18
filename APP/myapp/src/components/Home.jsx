import React, { useState } from 'react';
import ScrollableListings from "./ListingPage";

const Home = () => {

    const exampleListings = [
        {
            listing_id: 1,
            type: 'Offer',
            location: { lon: -73.935242, lat: 40.73061 },
            user_id: 101,
            username: 'JohnDoe',
            title: 'Professional Lawn Mowing Service',
            description: 'Offering reliable lawn mowing services with 5 years of experience.',
            date_created: '2024-11-01 10:30:00',
            active: true,
            city: 'New York',
            country: 'USA',
            images: ['https://littlesunnykitchen.com/wp-content/uploads/2020/10/Batata-Harra-13.jpg', 'https://example.com/image2.jpg'],
        },
        {
            listing_id: 2,
            type: 'Request',
            location: { lon: -0.127758, lat: 51.507351 },
            user_id: 102,
            username: 'JaneSmith',
            title: 'Looking for a Personal Trainer',
            description: 'Seeking a certified personal trainer for sessions 3 times a week.',
            date_created: '2024-10-25 14:15:00',
            active: true,
            city: 'London',
            country: 'UK',
            images: ['https://www.saborbrasil.it/wp-content/uploads/2021/06/Batata-doce-frita-1024x768.jpg', 'https://example.com/image4.jpg'],
        },
        {
            listing_id: 3,
            type: 'Offer',
            location: { lon: 139.691706, lat: 35.689487 },
            user_id: 103,
            username: 'TaroYamada',
            title: 'Guitar Lessons for Beginners',
            description: 'Learn to play the guitar with comprehensive lessons and practice sessions.',
            date_created: '2024-11-10 09:00:00',
            active: true,
            city: 'Tokyo',
            country: 'Japan',
            images: ['https://s2-receitas.glbimg.com/_GkYC8FQvE7JNZoILsLV7BvmD-I=/0x0:1000x750/984x0/smart/filters:strip_icc()/i.s3.glbimg.com/v1/AUTH_1f540e0b94d8437dbbc39d567a1dee68/internal_photos/bs/2024/t/d/W4TOLSQxAXiLig8w82IA/batata-frita-sequinha.jpg'],
        }
    ];



    return (
        <ScrollableListings listingType={"Offer"} />
    )
}

export default Home