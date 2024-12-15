import React, { useState } from 'react';
import { Box, Modal, Typography, TextField, MenuItem, Button } from "@mui/material";
import TransactionService from "../../services/TransactionService";

const SendOfferModal = ({ isOpen, onClose, User, OfferingUser, Listing }) => {
    // State to manage form inputs
    const [price, setPrice] = useState('');
    const [currency, setCurrency] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');
    const [details, setDetails] = useState('');
    const [isConfirmOpen, setIsConfirmOpen] = useState(false); // State for confirmation modal

    // Available currencies
    const currencies = ['USD', 'EUR', 'GBP', 'INR']; // Add more as needed

    // Handler to open confirmation modal
    const handleConfirm = () => {
        setIsConfirmOpen(true);
    };

    // Handler for final submission
    const handleSubmit = async () => {
        const offerData = {
            price,
            currency,
            startDate,
            endDate,
            details,
            listingId: Listing.listing_id,
            userId: User.user_id,
            offeringUserId: OfferingUser.user_id,
        };
        console.log('Offer Data:', offerData);


        try {
            const resData = await TransactionService.createTransaction(price, currency, startDate, endDate, details, Listing.listing_id, User.user_id, Listing.user_id)
            console.log(resData);
        }
        catch (error) {
            console.error(error);
        }
        // Add API call or further processing here
        setIsConfirmOpen(false); // Close confirmation modal
        onClose(); // Close the main modal after submission
    };

    return (
        <>
            {/* Main Modal */}
            <Modal
                open={isOpen}
                onClose={onClose}
                aria-labelledby="Send-Offer-modal-title"
                aria-describedby="Send-Offer-modal-description">
                <Box
                    className="modal-people"
                    sx={{
                        position: 'absolute',
                        top: '5%',
                        left: '50%',
                        transform: 'translate(-50%, 0)',
                        width: '80%',
                        height: '80%',
                        borderRadius: 2,
                        boxShadow: 24,
                        p: 4,
                        display: 'flex',
                        flexDirection: 'column',
                    }}
                >
                    {/* Title */}
                    <Typography id="offer-modal-title" variant="h5" gutterBottom>
                        Send Offer
                    </Typography>

                    {/* Form Inputs */}
                    <Box
                        sx={{
                            flex: 1,
                            overflowY: 'auto',
                            padding: 1,
                            display: 'flex',
                            flexDirection: 'column',
                            gap: 2,
                        }}
                    >
                        {/* Price Input */}
                        <TextField
                            label="Price"
                            variant="outlined"
                            type="number"
                            value={price}
                            onChange={(e) => setPrice(e.target.value)}
                            fullWidth
                        />

                        {/* Currency Selector */}
                        <TextField
                            label="Currency"
                            variant="outlined"
                            select
                            value={currency}
                            onChange={(e) => setCurrency(e.target.value)}
                            fullWidth
                        >
                            {currencies.map((curr) => (
                                <MenuItem key={curr} value={curr}>
                                    {curr}
                                </MenuItem>
                            ))}
                        </TextField>

                        {/* Start Date */}
                        <TextField
                            label="Start Date"
                            variant="outlined"
                            type="date"
                            InputLabelProps={{ shrink: true }}
                            value={startDate}
                            onChange={(e) => setStartDate(e.target.value)}
                            fullWidth
                        />

                        {/* End Date */}
                        <TextField
                            label="End Date"
                            variant="outlined"
                            type="date"
                            InputLabelProps={{ shrink: true }}
                            value={endDate}
                            onChange={(e) => setEndDate(e.target.value)}
                            fullWidth
                        />

                        {/* Details Input */}
                        <TextField
                            label="Details"
                            variant="outlined"
                            multiline
                            rows={4}
                            value={details}
                            onChange={(e) => setDetails(e.target.value)}
                            fullWidth
                        />
                    </Box>

                    {/* Submit and Cancel Buttons */}
                    <Box sx={{ display: 'flex', justifyContent: 'flex-end', gap: 2, mt: 2 }}>
                        <Button variant="outlined" onClick={onClose}>
                            Cancel
                        </Button>
                        <Button
                            variant="contained"
                            onClick={handleConfirm} // Open confirmation modal
                            disabled={!price || !currency || !startDate || !endDate || !details}
                        >
                            Submit
                        </Button>
                    </Box>
                </Box>
            </Modal>

            {/* Confirmation Modal */}
            <Modal
                open={isConfirmOpen}
                onClose={() => setIsConfirmOpen(false)}
                aria-labelledby="confirm-offer-modal-title"
                aria-describedby="confirm-offer-modal-description">
                <Box
                    sx={{
                        position: 'absolute',
                        top: '50%',
                        left: '50%',
                        transform: 'translate(-50%, -50%)',
                        width: 400,
                        borderRadius: 2,
                        boxShadow: 24,
                        p: 4,
                        textAlign: 'center',
                    }}
                    className={"MUIBox"}
                >
                    <Typography id="confirm-offer-modal-title" variant="h6" gutterBottom>
                        Confirm Offer Details
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                        <strong>Price:</strong> {price} {currency}
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                        <strong>Start Date:</strong> {startDate}
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                        <strong>End Date:</strong> {endDate}
                    </Typography>
                    <Typography variant="body1" gutterBottom>
                        <strong>Details:</strong> {details}
                    </Typography>

                    {/* Confirmation Buttons */}
                    <Box sx={{ display: 'flex', justifyContent: 'center', gap: 2, mt: 2 }}>
                        <Button variant="outlined" onClick={() => setIsConfirmOpen(false)}>
                            Edit
                        </Button>
                        <Button variant="contained" onClick={handleSubmit}>
                            Confirm
                        </Button>
                    </Box>
                </Box>
            </Modal>
        </>
    );
};

export default SendOfferModal;
