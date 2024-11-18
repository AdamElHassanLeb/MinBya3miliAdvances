// src/services/ListingService.js
import axiosInstance from "../utils/Axios";

const URL = '/listing';

const getAllListings = async (type) => {
    try {
        return await axiosInstance.get(`${URL}/listings/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const getListingById = async (id) => {
    try {
        return await axiosInstance.get(`${URL}/listingId/${id}`);
    } catch (error) {
        console.log(error);
    }
};

const getListingsByUserId = async (user_id, type) => {
    try {
        return await axiosInstance.get(`${URL}/listings/user/${user_id}/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const searchListings = async (query, type) => {
    try {
        return await axiosInstance.get(`${URL}/search/${query}/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const getListingsByDate = async (type) => {
    try {
        return await axiosInstance.get(`${URL}/date/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const getListingsByDateAndSearch = async (query, type) => {
    try {
        return await axiosInstance.get(`${URL}/date/search/${query}/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const getListingsByDistance = async (longitude, latitude, max_distance, type) => {
    try {
        return await axiosInstance.get(`${URL}/distance/${longitude}/${latitude}/${max_distance}/${type}`);
    } catch (error) {
        console.log(error);
    }
};

const createListing = async (listingData) => {
    try {
        return await axiosInstance.post(`${URL}/create`, listingData);
    } catch (error) {
        console.log(error);
    }
};

const updateListing = async (id, updatedData) => {
    try {
        return await axiosInstance.put(`${URL}/update/${id}`, updatedData);
    } catch (error) {
        console.log(error);
    }
};

const deleteListing = async (id) => {
    try {
        return await axiosInstance.delete(`${URL}/delete/${id}`);
    } catch (error) {
        console.log(error);
    }
};

// Function to fetch listings by distance and search query
const getListingsByDistanceAndSearch = async (longitude, latitude, maxDistance, listingType, searchQuery) => {
    try {
        const response = await axiosInstance.get(`/api/v1/listing/distance/${longitude}/${latitude}/${maxDistance}/${listingType}/${searchQuery}`);
        return response.data; // Assuming your API returns the data in a structure like { listings: [...], ... }
    } catch (error) {
        console.error("Error fetching listings by distance and search:", error);
        throw error;
    }
};

export default {
    getAllListings,
    getListingById,
    getListingsByUserId,
    searchListings,
    getListingsByDate,
    getListingsByDateAndSearch,
    getListingsByDistance,
    createListing,
    updateListing,
    deleteListing,
    getListingsByDistanceAndSearch
};
