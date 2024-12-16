# MinBya3mili – Local Tradesman Marketplace

## Table of Contents
1. [Overview](#overview)
2. [Key Business Aspects](#key-business-aspects)
    - [Value Proposition](#value-proposition)
    - [Monetization Strategy](#monetization-strategy)
3. [Key Features](#key-features)
    - [Business-Oriented Features](#business-oriented-features)
    - [Technical Features](#technical-features)
4. [Technological Stack](#technological-stack)
5. [API Endpoints](#api-endpoints)
    - [User Management](#user-management)
    - [Listings Management](#listings-management)
    - [Image Management](#image-management)
    - [Transaction Management](#transaction-management)
6. [Technical and Business Decisions](#technical-and-business-decisions)
    - [Simplicity and Scalability](#simplicity-and-scalability)
    - [Blockchain-Powered Agreements](#blockchain-powered-agreements)
    - [Dependency Injection for Flexibility](#dependency-injection-for-flexibility)
7. [Future Roadmap](#future-roadmap)
    - [Feature Enhancements](#feature-enhancements)
    - [Scaling and Optimization](#scaling-and-optimization)
    - [Business Growth](#business-growth)
8. [Conclusion](#conclusion)

---

## Overview

MinBya3mili is an innovative platform designed to revolutionize how users connect with local tradesmen. By combining cutting-edge technology and user-centric design, MinBya3mili offers a seamless experience for discovering and hiring trustworthy tradesmen for various tasks. The platform leverages a modern tech stack—including React, Go, MySQL, AWS, and Blockchain—to ensure security, scalability, and simplicity.

As the platform enters its Beta stage, MinBya3mili focuses on solving the widespread challenge of finding reliable local contractors while fostering transparency and trust through technology.

## Key Business Aspects

### Value Proposition

- **Convenience**: Simplifies the process of finding local tradesmen by providing an easy-to-use, proximity-based service discovery system.
- **Trust and Transparency**: Blockchain-powered smart contracts eliminate the need for intermediaries, offering secure, legally binding agreements.
- **Engagement and Growth**: Social media-like features allow tradesmen to showcase their portfolios, enhancing their visibility and helping them grow their businesses.
- **Cost Efficiency**: Avoids the overhead of custom messaging systems by integrating with widely-used platforms like WhatsApp.
- **Market Differentiation**: Focuses on combining modern technology and localized needs, targeting underserved markets such as Lebanon.

### Monetization Strategy

- **Subscription Plans**: Premium plans for tradesmen to unlock advanced features, such as enhanced profile visibility and analytics.
- **Transaction Fees**: A small fee for transactions facilitated through the platform.
- **Advertising**: Allow tradesmen to promote their listings to targeted users.
- **API Licenses**: Provide location and service APIs to third-party businesses.

## Key Features

### Business-Oriented Features

1. **Proximity-Based Service Discovery**: Users can locate tradesmen in their vicinity, enabling faster and more personalized service connections.
2. **Tradesmen Portfolios**: Allows professionals to display their skills, past projects, and customer reviews.
3. **Blockchain Smart Contracts**: Creates secure, transparent, and legally binding agreements, shifting legal responsibility to users and eliminating intermediaries.
4. **Social Proof**: Customer reviews and ratings build trust, helping clients choose the best tradesmen.
5. **WhatsApp Integration**: Facilitates communication without requiring additional messaging infrastructure.

### Technical Features

1. **Scalable Frontend**: React-based SPA ensures a fast, responsive, and user-friendly interface.
2. **Robust Backend**: Built with Go and Chi Router for high-performance API handling.
3. **Secure Authentication**: JWT tokens ensure secure login and user-specific access.
4. **Cloud Deployment**: AWS ensures high availability, fault tolerance, and the ability to scale dynamically.
5. **Database Flexibility**: Dependency injection allows smooth transitions and maintenance of the database system.
6. **Smart Contracts**: Blockchain technology ensures tamper-proof agreements between tradesmen and clients.

## Technological Stack

- **Frontend**: React, Material UI (MUI)
- **Backend**: Go, Chi Router, JWT Authentication
- **Database**: MySQL (Deployed on AWS)
- **Cloud Infrastructure**: AWS EC2, S3, and RDS
- **Blockchain**: Smart Contracts for transaction agreements
- **APIs**: Nominatim API for location-based services

## API Endpoints

### User Management
- **GET /api/v1/user/users**: Retrieve all users.
- **GET /api/v1/user/userId/{id}**: Get details of a specific user by ID.
- **POST /api/v1/user/create**: Register a new user.
- **PUT /api/v1/user/update/{id}**: Update user information.
- **DELETE /api/v1/user/delete/{id}**: Remove a user from the system.

### Listings Management
- **GET /api/v1/listing/listings/{type}**: View listings filtered by type (Offer/Request).
- **POST /api/v1/listing/create**: Create a new service listing.
- **PUT /api/v1/listing/update/{id}**: Edit an existing listing.
- **DELETE /api/v1/listing/delete/{id}**: Remove a listing.

### Image Management
- **POST /api/v1/image/uploadForListing/{listing_id}**: Upload an image for a listing.
- **GET /api/v1/image/listing/{listing_id}**: Retrieve images associated with a specific listing.
- **DELETE /api/v1/image/{image_id}**: Delete an image.

### Transaction Management
- **POST /api/v1/transaction/create**: Initiate a new transaction.
- **GET /api/v1/transaction/transactionId/{id}**: Retrieve transaction details by ID.
- **PUT /api/v1/transaction/{id}**: Update transaction details.
- **DELETE /api/v1/transaction/{id}**: Cancel a transaction.

## Technical and Business Decisions

### Simplicity and Scalability
- **Decision to Use WhatsApp Integration**: Instead of investing resources into developing an in-house messaging system, MinBya3mili integrates with WhatsApp. This decision reduces development complexity and operational costs while leveraging an already widely-used platform.
- **AWS Deployment**: Hosting on AWS ensures that the platform can scale effortlessly with user growth, maintaining performance during peak usage.

### Blockchain-Powered Agreements
Using blockchain technology aligns with the platform’s commitment to trust and transparency. Smart contracts automate agreement enforcement, reducing disputes and eliminating intermediary fees.

### Dependency Injection for Flexibility
The use of dependency injection in the database design ensures that updates or migrations to new database systems can occur with minimal disruption, supporting long-term scalability and maintenance.

## Future Roadmap

### Feature Enhancements
- **Mobile Application**: Develop native iOS and Android apps to expand accessibility.
- **Advanced Search Filters**: Allow users to filter listings by ratings, price ranges, and availability.
- **Enhanced Listings**: Add multimedia support like videos and richer descriptions for tradesmen profiles.

### Scaling and Optimization
- **Serverless Architecture**: Transition to serverless computing for cost-effective scaling.
- **Database Optimization**: Implement caching solutions to reduce load times and improve database performance.
- **Global Expansion**: Adapt the platform to support international markets, considering multilingual support and region-specific features.

### Business Growth
- **Partnerships**: Collaborate with local businesses and tradesmen’s associations to onboard more professionals.
- **Marketing Campaigns**: Targeted digital advertising to increase user adoption.
- **Customer Support**: Introduce AI-powered chatbots to provide instant support and guidance.

## Conclusion

MinBya3mili combines technical innovation with practical business solutions to redefine local service discovery. By focusing on user convenience, trust, and scalability, the platform is well-positioned to become a leading marketplace for tradesmen in Lebanon and beyond. MinBya3mili is committed to continuous improvement and welcomes feedback as it progresses beyond the Beta stage to a full-scale launch.
