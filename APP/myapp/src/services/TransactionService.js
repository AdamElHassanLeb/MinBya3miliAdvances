import axiosInstance from "../utils/Axios";
import { getTokenBearer } from "../utils/Token";

const URL = '/transaction';

// Create a transaction
const createTransaction = async (price, currency, startDate, endDate, details, listingId, userId, listingUserId) => {
    try {
        const token = getTokenBearer(); // Retrieve the bearer token
        const payload = {
            user_offered_id: userId,
            user_offering_id: listingUserId,
            listing_id: listingId,
            price_with_currency: parseFloat(price),
            currency_code: currency,
            job_start_date: startDate,
            job_end_date: endDate,
            details_from_offered: details,
        };

        const response = await axiosInstance.post(`${URL}/create`, payload, {
            headers: { Authorization: token },
        });
        return response.data; // Return the API response data
    } catch (err) {
        console.error("Error creating transaction:", err);
        throw new Error(`Failed to create transaction: ${err.response?.data?.message || err.message}`);
    }
};

// Get transaction by ID
const getTransactionByID = async (transactionId) => {
    try {
        const token = getTokenBearer();
        const response = await axiosInstance.get(`${URL}/transactionId/${transactionId}`, {
            headers: { Authorization: token },
        });
        return response.data;
    } catch (err) {
        console.error("Error fetching transaction by ID:", err);
        throw new Error(`Failed to fetch transaction by ID: ${err.response?.data?.message || err.message}`);
    }
};

// Get transactions by offered user and status
const getTransactionsByOfferedUserAndStatus = async (userId, status) => {
    try {
        if (!status || status == "") {status="Any"}
        const token = getTokenBearer();
        const response = await axiosInstance.get(`${URL}/offered/${userId}/${status}`, {
            headers: { Authorization: token },
        });
        return response.data;
    } catch (err) {
        console.error("Error fetching transactions by offered user and status:", err);
        throw new Error(`Failed to fetch transactions: ${err.response?.data?.message || err.message}`);
    }
};

// Get transactions by offering user and status
const getTransactionsByOfferingUserAndStatus = async (userId, status) => {
    try {
        const token = getTokenBearer();
        const response = await axiosInstance.get(`${URL}/offering/${userId}/${status}`, {
            headers: { Authorization: token },
        });
        return response.data;
    } catch (err) {
        console.error("Error fetching transactions by offering user and status:", err);
        throw new Error(`Failed to fetch transactions: ${err.response?.data?.message || err.message}`);
    }
};

// Get transactions by listing and status
const getTransactionsByListingAndStatus = async (listingId, status) => {
    try {
        const token = getTokenBearer();
        const response = await axiosInstance.get(`${URL}/listing/${listingId}/${status}`, {
            headers: { Authorization: token },
        });
        return response.data;
    } catch (err) {
        console.error("Error fetching transactions by listing and status:", err);
        throw new Error(`Failed to fetch transactions: ${err.response?.data?.message || err.message}`);
    }
};

export default {
    createTransaction,
    getTransactionByID,
    getTransactionsByOfferedUserAndStatus,
    getTransactionsByOfferingUserAndStatus,
    getTransactionsByListingAndStatus,
};
