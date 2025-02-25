#                                               `Smart Home Microservices Project`

##  This project consists of three microservices: User, Device, and API Gateway. It is designed to manage a smart home system with devices such as smart speakers, vacuum cleaners, smart alarms, and smart lock doors. The project can be run using Docker Compose.

## Table of Contents
# `Overview`
#    `Microservices`
#    `User Service`
#   ` Device Service`
#   `API Gateway`
## Overview:

# *This project is a backend system for a smart home application. It allows users to manage various smart devices in their homes through the following microservices*:

# User Service: Manages user information and authentication.
- Device Service: Handles device registration, management, and operations.
- API Gateway: Routes requests to the appropriate microservice and ensures secure communication.
- Microservices
- User Service
- The User Service handles all user-related operations including registration, login, and user profile management. Authentication is managed using JWT.

##    `Device Service`
# *The Device Service allows users to register and manage their smart home devices. Supported devices include:*

#   `Smart Speaker`
#   `Vacuum Cleaner`
#   `Smart Alarm`
#   `Smart Lock Door`


#   `API Gateway`
The API Gateway serves as the single entry point for all client requests. It routes requests to the appropriate microservice and handles cross-cutting concerns such as authentication and rate limiting.

###     `Devices`
# The system supports the following smart devices:
#    `Smart Speaker`
#   `Vacuum Cleaner`
#   `Smart Alarm`
#   `Smart Lock Door`
Users can add these devices to their home by providing the device name through the appropriate endpoint in the Device Service.


#   Getting Started
        To run the project, ensure you have Docker and Docker Compose installed. Follow the steps below:

# Copy code
# `docker-compose up`

### *This will start all three microservices along with their dependencies.*

#   Issues
        Currently, the project works perfectly on Postman but has issues with Swagger. The Swagger documentation is not functioning as expected.