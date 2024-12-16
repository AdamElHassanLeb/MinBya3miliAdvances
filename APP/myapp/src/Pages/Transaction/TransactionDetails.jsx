import React, { useContext, useEffect, useState } from 'react';
import { useNavigate, useParams } from "react-router-dom";
import { format } from 'date-fns';
import { Box, Button, Divider, Typography, Stack, Alert, CircularProgress, Snackbar, Chip } from "@mui/material";
import ListingCard from "../../components/Listings/ListingCard";
import { UserContext } from "../../utils/UserContext";
import DeleteIcon from "@mui/icons-material/Delete";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import HourglassEmptyIcon from "@mui/icons-material/HourglassEmpty";
import UserService from "../../services/UserService";
import ListingService from "../../services/ListingService";
import TransactionService from "../../services/TransactionService";
import ReceiptIcon from '@mui/icons-material/Receipt';

const TransactionDetails = () => {
    const { transactionId } = useParams();
    const { user } = useContext(UserContext);
    const [listing, setListing] = useState(null);
    const [offeringUser, setOfferingUser] = useState(null);
    const [transaction, setTransaction] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [snackbarMessage, setSnackbarMessage] = useState('');
    const [snackbarOpen, setSnackbarOpen] = useState(false);
    const navigate = useNavigate();


    //Contract

    const getContract = async () => {

        try{

            const data = await TransactionService.getContract(transactionId);
            console.log("data", data);

            // Create the English contract file
            const englishContractBlob = new Blob([data.english_contract], { type: 'text/plain' });
            const englishContractLink = document.createElement('a');
            englishContractLink.href = URL.createObjectURL(englishContractBlob);
            englishContractLink.download = 'english_contract.txt';  // Filename for the English contract
            englishContractLink.click();  // Trigger the download

            // Create the Arabic contract file
            const arabicContractBlob = new Blob([data.arabic_contract], { type: 'text/plain' });
            const arabicContractLink = document.createElement('a');
            arabicContractLink.href = URL.createObjectURL(arabicContractBlob);
            arabicContractLink.download = 'arabic_contract.txt';  // Filename for the Arabic contract
            arabicContractLink.click();  // Trigger the download

        }catch(e){

        }
    }


    // Snackbar close handler
    const handleSnackbarClose = () => setSnackbarOpen(false);

    // Load transaction details
    async function loadTransactionDetails() {
        try {
            setLoading(true);
            setError(null);

            const transactionData = await TransactionService.getTransactionByID(transactionId);
            setTransaction(transactionData);

            if (!transactionData?.listing_id || !transactionData?.user_offering_id) {
                throw new Error("Invalid transaction data.");
            }

            const userData = await UserService.getUserById(transactionData.user_offering_id);
            setOfferingUser(userData);

            const listingData = await ListingService.getListingById(transactionData.listing_id);
            setListing(listingData.data);
        } catch (err) {
            setError(err.response?.data?.message || "Failed to load transaction details.");
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        loadTransactionDetails();
    }, []);

    // Action handlers
    const acceptOffer = async () => {
        try {
            await TransactionService.updateTransaction(transactionId, { ...transaction, status: "Accepted" });
            setSnackbarMessage("Offer Accepted");
            setSnackbarOpen(true);
            await loadTransactionDetails();
        } catch (err) {
            console.error(err);
        }
    };

    const completeOffer = async () => {
        try {
            await TransactionService.updateTransaction(transactionId, { ...transaction, status: "Completed" });
            setSnackbarMessage("Offer Completed");
            setSnackbarOpen(true);
            await loadTransactionDetails();
        } catch (err) {
            console.error(err);
        }
    };

    const deleteOffer = async () => {
        try {
            await TransactionService.deleteTransaction(transactionId);
            setSnackbarMessage("Transaction Deleted");
            setSnackbarOpen(true);
            navigate('/Home');
        } catch (err) {
            setError("Failed to delete transaction.");
        }
    };

    // Render transaction actions
    const renderTransactionAction = () => {
        if (transaction.user_offered_id === user.user_id) {
            if (transaction.status === "Completed") {
                return <Button variant="contained" color="success" disabled>Completed</Button>;
            } else if (transaction.status === "Accepted") {
                return <Button variant="contained" color="primary" onClick={completeOffer}>Mark as Completed</Button>;
            }
            return <Button variant="contained" color="primary" onClick={acceptOffer}>Accept Offer</Button>;
        }
        return (
            <Typography variant="body1" color="textSecondary">
                {transaction.status === "Completed" ? "Transaction Completed" : transaction.status === "Accepted" ? "Offer Accepted" : "Pending"}
            </Typography>
        );
    };

    if (loading) {
        return (
            <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
                <CircularProgress />
            </Box>
        );
    }

    if (error) {
        return <Alert severity="error">{error}</Alert>;
    }

    return (
        <Box
            sx={{
                maxWidth: '900px',
                margin: 'auto',
                padding: '20px',
                border: '1px solid #ccc',
                borderRadius: '10px',
                boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
                marginTop: '10vh',
                minHeight: '80vh',
            }}
            className="MUIBox"
        >
            <Typography variant="h4" gutterBottom>
                Transaction Details
            </Typography>

            <Divider sx={{ marginBottom: '20px' }} />

            {/* Listing Information */}
            <Typography variant="h6">Listing Details:</Typography>
            <ListingCard listing={listing}/>

            <Divider sx={{ marginY: '20px' }} />

            {/* Transaction Information */}
            <Stack spacing={2}>
                <Typography variant="h6">Transaction Information:</Typography>
                <Chip
                    label={transaction.status}
                    color={transaction.status === "Completed" ? "success" : transaction.status === "Accepted" ? "primary" : "error"}
                    icon={transaction.status === "Completed" ? <CheckCircleIcon /> : <HourglassEmptyIcon />}
                    variant="outlined"
                />
                <Typography><strong>Price:</strong> {transaction.price_with_currency} {transaction.currency_code}</Typography>
                <Typography><strong>Job Start Date:</strong> {format(new Date(transaction.job_start_date), 'PPP')}</Typography>
                <Typography><strong>Job End Date:</strong> {format(new Date(transaction.job_end_date), 'PPP')}</Typography>
                <Typography><strong>Details from Offering User:</strong> {transaction.details_from_offering}</Typography>
                <Typography><strong>Details from Offered User:</strong> {transaction.details_from_offered}</Typography>
            </Stack>

            <Divider sx={{ marginY: '20px' }} />

            {/* Offering User Information */}
            <Typography variant="h6">User Offering:</Typography>
            <Typography><strong>Name:</strong> {offeringUser.first_name} {offeringUser.last_name}</Typography>
            <Typography><strong>Phone:</strong> {offeringUser.phone_number}</Typography>
            <Typography><strong>Profession:</strong> {offeringUser.profession}</Typography>

            <Divider sx={{ marginY: '20px' }} />

            {/* Actions */}
            <Stack direction="row" spacing={2}>
                {renderTransactionAction()}
                <Button variant="contained" color="error" onClick={deleteOffer} startIcon={<DeleteIcon />} aria-label="Delete Transaction">
                    Delete Transaction
                </Button>
                <Button variant="contained" color = "primary" onClick={getContract} startIcon={<ReceiptIcon/>}> Request Contract </Button>
            </Stack>

            {/* Snackbar */}
            <Snackbar
                open={snackbarOpen}
                autoHideDuration={3000}
                onClose={handleSnackbarClose}
                message={snackbarMessage}
            />
        </Box>
    );
};

export default TransactionDetails;
