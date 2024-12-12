import React from "react";
import serverAddress from "../../utils/ServerAddress";
import { Avatar, Grid, Typography } from "@mui/material";
import {useNavigate} from "react-router-dom";

const UserSearchRow = ({ user, onClose }) => {
    const navigate = useNavigate();

    return (
        <Grid
            className = {"User-Row"}
            container
            alignItems="center"
            spacing={2}
            wrap="nowrap"
            sx={{ padding: 1, width: "75%", marginLeft : 'auto', marginRight : "auto" }}
            onClick={() => {
                navigate(`/User/${user.user_id}`);
                onClose();
            }}
        >
            {/* Profile Picture */}
            <Grid item>
                <Avatar
                    src={serverAddress() + `/api/v1/image/imageId/${user.image_id}` || `../assets/default-avatar.png`}
                    alt="Profile Picture"
                    sx={{ width: 60, height: 60 }}
                />
            </Grid>

            {/* User Details */}
            <Grid item xs={3}>
                <Typography variant="body1" fontWeight="bold">
                    First Name:
                </Typography>
                <Typography variant="body2">{user.first_name}</Typography>
            </Grid>
            <Grid item xs={3}>
                <Typography variant="body1" fontWeight="bold">
                    Last Name:
                </Typography>
                <Typography variant="body2">{user.last_name}</Typography>
            </Grid>
            <Grid item xs={3}>
                <Typography variant="body1" fontWeight="bold">
                    Phone:
                </Typography>
                <Typography variant="body2">{user.phone_number}</Typography>
            </Grid>
            <Grid item xs={3}>
                <Typography variant="body1" fontWeight="bold">
                    Profession:
                </Typography>
                <Typography variant="body2">{user.profession}</Typography>
            </Grid>
        </Grid>
    );
};

export default UserSearchRow;
