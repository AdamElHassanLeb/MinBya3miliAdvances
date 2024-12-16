import React, {useContext, useEffect, useState} from 'react'
import {Avatar, Grid, Typography} from "@mui/material";
import ListingService from "../../services/ListingService";
import {useNavigate} from "react-router-dom";
import UserService from "../../services/UserService";
import serverAddress from "../../utils/ServerAddress";
import {UserContext} from "../../utils/UserContext";


const TransactionListItem = ({transaction}) => {


    const [listing, setListing] = useState(null);
    const [offeringUser, setOfferingUser] = useState(null);

    const navigate = useNavigate();
    const { user } = useContext(UserContext);


    useEffect(() => {

        async function loadListing() {

            if (transaction === null ||
                transaction.listing_id === null ||
                transaction.user_offering_id == null)
                return


            try {
                var userData
                if (transaction.user_offering_id == user.user_id)
                    userData = await UserService.getUserById(transaction.user_offered_id);
                else
                    userData = await UserService.getUserById(transaction.user_offering_id);
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
            {listing != null && offeringUser != null ? (
            <Grid  className = {"User-Row"}
                   container
                   alignItems="center"
                   spacing={0.5}
                   wrap="nowrap"
                   sx={{ padding: 1, width: "80%", marginLeft : 'auto', marginRight : "auto", background: "transparent", borderRadius: "5px" }}
                    onClick={() => {navigate(`/Transaction/${transaction.transaction_id}`)}}
            >


                <Grid item>
                    <Avatar
                        src={serverAddress() + `/api/v1/image/imageId/${offeringUser.image_id}` || `../assets/default-avatar.png`}
                        alt="Profile Picture"
                        sx={{ width: 60, height: 60 }}
                    />
                </Grid>
                {/* User Details */}
                <Grid item xs={1}>
                    <Typography variant="body1" fontWeight="bold">
                        Name:
                    </Typography>
                    <Typography variant="body2">{offeringUser.first_name + ' ' + offeringUser.last_name}</Typography>
                </Grid>


                <Grid item xs={3}>
                    <Typography variant="body1" fontWeight="bold">
                        Listing Title:
                    </Typography>
                    <Typography variant="body2">{listing.title}</Typography>
                </Grid>

                <Grid item xs={1}>
                    <Typography variant="body1" fontWeight="bold">
                        Price:
                    </Typography>
                    <Typography variant="body2">{transaction.price_with_currency + ' ' + transaction.currency_code}</Typography>
                </Grid>
                <Grid item xs={1}>
                    <Typography variant="body1" fontWeight="bold">
                        Start Date:
                    </Typography>
                    <Typography variant="body2">{transaction.job_start_date}</Typography>
                </Grid>
                <Grid item xs={1}>
                    <Typography variant="body1" fontWeight="bold">
                        End Date:
                    </Typography>
                    <Typography variant="body2">{transaction.job_end_date}</Typography>
                </Grid>
                { transaction.status !== "Pending" ?
                <Grid item xs={3}>
                    <Typography variant="body1" fontWeight="bold">
                        Description:
                    </Typography>
                    <Typography variant="body2">{listing.details_from_offered}</Typography>
                </Grid>
                : null}
                <Grid item xs={1}>
                    <Typography variant="body1" fontWeight="bold">
                        Status:
                    </Typography>
                    <Typography variant="body2">{transaction.status}</Typography>
                </Grid>
            </Grid>
            ) : null}
        </>
    )
}

export default TransactionListItem