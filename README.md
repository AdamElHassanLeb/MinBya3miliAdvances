# MinBya3mili – Local Tradesman Marketplace

## Beta Stage

### Overview

MinBya3mili is a platform designed to connect users with local tradesmen, providing a service discovery and communication experience based on proximity. By leveraging modern technologies such as React, Go, Chi Router, JWT, MySQL, AWS, and Blockchain, we aim to revolutionize how tradesmen and clients connect. The platform allows for secure, transparent legal agreements, while ensuring scalability, high availability, and a seamless user experience.

This project is currently in its Beta stage and focuses on solving the common struggle of finding trustworthy tradesmen or contractors for various tasks.

### Key Features

- **Local Proximity-Based Service Discovery**: Users can find tradesmen in their area and connect with them based on proximity, helping to streamline the process of hiring services locally.
- **Blockchain Smart Contracts**: To ensure security and transparency in agreements, MinBya3mili leverages blockchain-based smart contracts, transferring legal responsibility to users. This eliminates the need for traditional intermediaries and offers a modern approach to legal contracts.
- **User Profiles & Social Media Integration**: Tradesmen can build their portfolios and clients can view examples of their work. This social media-like experience allows for user engagement and fosters trust in the tradesmen's expertise.
- **React-based Single Page Application (SPA)**: The frontend is built using React, providing a responsive, almost desktop-like experience, enabling users to navigate through listings, profiles, and other sections seamlessly.
- **AWS Deployment**: Both the backend and database are deployed on AWS to ensure high availability, fault tolerance, and scalability, making sure the platform can handle increasing traffic and data in the future.
- **Dependency Injection for Database**: The database implementation uses dependency injection, which ensures that changes to the database system can be handled without disruptions, providing scalability and ease of maintenance in production environments.
- **Secure Authentication with JWT**: The platform uses JSON Web Tokens (JWT) to secure authentication and authorization, ensuring that only authorized users can perform certain actions, such as creating or updating listings, transactions, or images.

### Technological Stack

- **Frontend**: React, Material UI (MUI)
- **Backend**: Go, Chi Router, JWT Authentication
- **Database**: MySQL (Deployed on AWS)
- **Cloud Infrastructure**: AWS
- **Blockchain**: Smart Contracts
- **API Integration**: Nominatim API (for location services)

### API Endpoints

#### User Routes
- `GET /api/v1/user/users`: Get all users
- `GET /api/v1/user/userId/{id}`: Get user by ID
- `GET /api/v1/user/userName/{name}`: Get user by username
- `POST /api/v1/user/create`: Create a new user
- `DELETE /api/v1/user/delete/{id}`: Delete a user
- `PUT /api/v1/user/update/{id}`: Update user details
- `POST /api/v1/user/auth`: User authentication

#### Listing Routes
- `GET /api/v1/listing/listings/{type}`: Get all listings by type (Offer/Request)
- `GET /api/v1/listing/listingId/{id}`: Get listing by ID
- `POST /api/v1/listing/create`: Create a new listing
- `PUT /api/v1/listing/update/{id}`: Update a listing
- `DELETE /api/v1/listing/delete/{id}`: Delete a listing

#### Image Routes
- `POST /api/v1/image/uploadForListing/{listing_id}`: Upload an image for a listing
- `POST /api/v1/image/uploadProfilePicture/{user_id}`: Upload profile picture for a user
- `GET /api/v1/image/imageId/{image_id}`: Get image by ID
- `GET /api/v1/image/image/{image_id}`: Get image by UUID
- `GET /api/v1/image/listing/{listing_id}`: Get images by listing ID
- `GET /api/v1/image/user/{user_id}`: Get images by user ID
- `DELETE /api/v1/image/{image_id}`: Delete an image
- `PUT /api/v1/image/{image_id}/{show_on_profile}`: Update image visibility

#### Transaction Routes
- `POST /api/v1/transaction/create`: Create a new transaction
- `GET /api/v1/transaction/transactionId/{id}`: Get transaction by ID
- `GET /api/v1/transaction/offered/{user_id}/{status}`: Get transactions by offered user and status
- `GET /api/v1/transaction/offering/{user_id}/{status}`: Get transactions by offering user and status
- `GET /api/v1/transaction/listing/{listing_id}/{status}`: Get transactions by listing and status
- `PUT /api/v1/transaction/{id}`: Update a transaction
- `DELETE /api/v1/transaction/{id}`: Delete a transaction

### Technological Decisions

**Excluding WebRTC and Messaging Integration**

After researching competitors such as Facebook Marketplace, we found that implementing WebRTC and messaging services would have been resource-intensive and unnecessary for the core functionality of MinBya3mili. Instead, we decided that integrating a cost-effective solution by linking directly to WhatsApp for communication between users and tradesmen is the most efficient choice. This approach leverages the already widespread use of WhatsApp, ensuring a smooth communication process without requiring the additional overhead of maintaining a custom messaging platform.

### Future Considerations

- **Scaling**: We are focused on scaling the platform to support more users and transactions as we move out of the beta stage, ensuring that the infrastructure can handle the increasing demand.
- **Feature Expansion**: Future features may include enhanced filters for listing search, advanced transaction management, and even AI-powered recommendations for tradesmen based on user preferences and past interactions.

### Conclusion

MinBya3mili aims to create a more connected and efficient environment for users seeking local tradesmen and services. By focusing on proximity, blockchain contracts, and a user-friendly experience, we are confident in the platform's potential to meet the needs of both tradesmen and customers in Lebanon. We are continually improving the platform and welcome any feedback from users as we move forward.
