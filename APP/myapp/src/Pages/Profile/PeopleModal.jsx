import React, {useContext, useState} from 'react';
import { Modal, Box, Typography, Button, TextField } from '@mui/material';
import UserSearchRow from "../../components/User/UserSearchRow";
import UserService from "../../services/UserService";
import {Link} from "react-router-dom";
import {UserContext} from "../../utils/UserContext";

const PeopleModal = ({ isOpen, onClose }) => {
    const [users, setUsers] = useState([]);
    const { user } = useContext(UserContext);

    const removeUserById = (idToRemove) => {
        setUsers((prevUsers) =>
            prevUsers.filter((user) => user.user_id !== idToRemove)
        );
    };

    const handleSearch = async (event) => {
        const term = event.target.value;

        let users;
        try {
            users = await UserService.getUsersByUsername(term);
        } catch (error) {
            users = [];
        }

        setUsers(users);

        if(user)
            removeUserById(user.user_id);

    };

    return (
        <Modal
            open={isOpen}
            onClose={onClose}
            aria-labelledby="people-modal-title"
            aria-describedby="people-modal-description"
        >
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
                <Typography id="people-modal-title" variant="h5" gutterBottom>
                    People
                </Typography>

                {/* Search Bar */}
                <TextField
                    label="Search Users"
                    variant="outlined"
                    fullWidth
                    onChange={handleSearch}
                    sx={{ marginBottom: 2 }}
                />

                {/* Scrollable List */}
                <Box
                    sx={{
                        flex: 1,
                        overflowY: 'auto',
                        padding: 1,
                    }}
                >
                    {users && users.length > 0 ? (
                        <ul style={{ padding: 0, listStyle: "none", margin: 0, background: "transparent" }}>
                            {users.map((user, index) => (
                                <li key={index} style={{ marginBottom: "16px", background: "transparent" }}>
                                    <UserSearchRow user={user} onClose={onClose} />
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <Typography>No users found.</Typography>
                    )}
                </Box>

                {/* Close Button */}
                <Button onClick={onClose} variant="contained" color="primary" sx={{ width : '10vw', maxWidth: '100px', minHeight: '3vh', minWidth : '10vw' }}>
                    Close
                </Button>
            </Box>
        </Modal>
    );
};

export default PeopleModal;
