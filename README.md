# THE USER

[![Go](https://img.shields.io/badge/Go-blue)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-green)](https://gin-gonic.com/)
[![Docker](https://img.shields.io/badge/Docker-blue)](https://www.docker.com/)
[![MongoDB](https://img.shields.io/badge/MongoDB-green)](https://www.mongodb.com/)
[![AWS](https://img.shields.io/badge/AWS-yellow)](https://aws.amazon.com/)

This is a RESTful API built with Golang and Gin web framework, using MongoDB as the database. The API provides user-related services such as user authentication, profile picture upload to AWS S3 using presigned URL, and Google OAuth authentication. The project is also dockerized and includes a docker-compose.yml file to easily spin up the application.

### Prerequisites

- Docker and docker-compose installed
- AWS S3 bucket for file storage
- Google OAuth credentials

## Endpoints

| Method | Endpoint                   | Description                                     |
| ------ | -------------------------- | ----------------------------------------------- |
| POST   | `/auth/signup`             | Create a new user                               |
| POST   | `/auth/login`              | a user and receive a JWT                        |
| POST   | `/auth/google`             | Google OAuth login/signup                       |
| GET    | `/auth/token`              | Get new access token using refresh token        |
| GET    | `/user`                    | Admin route to get all users information        |
| PUT    | `/user/profile`            | Get a user's information                        |
| GET    | `/user/profile/upload-url` | Get a presigned URL to upload a profile picture |

## Authentication

The API uses a JWT-based authentication system for user login and signup. When a user logs in or signs up, the API generates a JWT access token and a refresh token and returns them to the client in the response body. The access token contains the user's information and has a short expiration time, while the refresh token has a longer expiration time and can be used to generate new access tokens.

## Profile Picture Upload

The API allows users to upload profile pictures to AWS S3 using presigned URLs. When a user requests an upload URL for their picture, the API generates a presigned URL and returns it in the response body. The client can then use this URL to upload the picture directly to S3.

## Google OAuth2

The API also supports Google OAuth authentication for user login and signup. When a user logs in or signs up with their Google account, the API retrieves their profile information from Google and creates a new user if they have not already created or logged in to the database with that information.
