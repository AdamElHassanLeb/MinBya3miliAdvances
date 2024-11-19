// src/components/HorizontalHeader.js
import React from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton, Box, Container } from '@mui/material';
import { Link } from 'react-router-dom';
import HomeIcon from '@mui/icons-material/Home';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import NotificationsIcon from '@mui/icons-material/Notifications';
import LogoutIcon from '@mui/icons-material/Logout';


const Header = () => {
    return (
        <AppBar position="static" color="primary">
            <Container maxWidth="lg">
                <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
                    {/* "Min Bya3mili" text as a link to home */}
                    <Typography variant="h6" component={Link} to="/Home" color="inherit" sx={{ textDecoration: 'none' }}>
                        Min Bya3mili
                    </Typography>

                    {/* Center Text Options */}
                    <Box sx={{ display: 'flex', gap: 2 }}>
                        <Button color="inherit" component={Link} to="/Offers">
                            Offers
                        </Button>
                        <Button color="inherit" component={Link} to="/Requests">
                            Requests
                        </Button>
                        <Button color="inherit" component={Link} to="/option3">
                            People
                        </Button>
                        <Button color="inherit" component={Link} to="/option3">
                            Messages
                        </Button>
                    </Box>

                    {/* Right-side Icons */}
                    <Box sx={{ display: 'flex', gap: 1 }}>
                    <IconButton color="inherit" component={Link} to="/Home">
                        <HomeIcon />
                    </IconButton>
                    <IconButton color="inherit" component={Link} to="/notifications">
                        <NotificationsIcon />
                    </IconButton>
                    <IconButton color="inherit" component={Link} to="/UserPrvateProfile">
                        <AccountCircleIcon />
                    </IconButton>
                        <IconButton color="inherit" component={Link} to="/Login">
                            <LogoutIcon />
                        </IconButton>
                </Box>

                </Toolbar>
            </Container>
        </AppBar>
    );
};

export default Header;
