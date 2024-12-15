import React, {useEffect, useState} from 'react'
import {useNavigate, useParams} from "react-router-dom";
import UserService from "../../services/UserService";
import ListingService from "../../services/ListingService";
import TransactionService from "../../services/TransactionService";


const TransactionDetails = () => {

    const { transactionId } = useParams();

    const [listing, setListing] = useState(null);
    const [offeringUser, setOfferingUser] = useState(null);
    const [transaction, setTransaction] = useState(null);
    const navigate = useNavigate();


    useEffect(() => {

        async function loadListing() {

            var transaction
            try {
                transaction = TransactionService.getTransactionByID(transactionId);
                setTransaction(transaction);
            }catch(err) {
                return;
            }

            if (transaction === null ||
                transaction.listing_id === null ||
                transaction.user_offering_id == null)
                return


            try {
                const userData = await UserService.getUserById(transaction.user_offering_id);
                setOfferingUser(userData);
            } catch (err) {
                return;
            }

            try {
                const listingData = await ListingService.getListingById(transaction.listing_id);
                setListing(listingData.data);
            } catch (err) {
                return;
            }
        }
        loadListing()
    }, []);



    return (
        <>
            <h1>{transaction.transaction_id}</h1>
        </>
    )
}

export default TransactionDetails;