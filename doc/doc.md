### LoanTrackerAPI - User Management Endpoints
This folder contains various API endpoints related to user management within the LoanTrackerAPI. The endpoints are handled by the `UserController` and include functionalities for user registration, login, email verification, profile management, and password reset. Below is a brief overview of each endpoint:

1. **User Registration**  
    **Endpoint:** `POST /users/register`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).RegisterUser-fm`  
    **Description:** Registers a new user with the system.

2. **Email Verification**  
    **Endpoint:** `GET /users/verify-email`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).VerifyEmail-fm`  
    **Description:** Verifies a user's email address.

3. **User Login**  
    **Endpoint:** `POST /users/login`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).Login-fm`  
    **Description:** Authenticates a user and provides access and refresh tokens.

4. **Token Refresh**  
    **Endpoint:** `POST /users/token/refresh`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).RefreshToken-fm`  
    **Description:** Refreshes the user's access token using a valid refresh token.

5. **Get User Profile**  
    **Endpoint:** `GET /users/profile`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).GetUserProfile-fm`  
    **Description:** Retrieves the profile information of the authenticated user.

6. **Request Password Reset**  
    **Endpoint:** `POST /users/password-reset`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).RequestPasswordReset-fm`  
    **Description:** Initiates a password reset process for a user.

7. **Update Password After Reset**  
    **Endpoint:** `POST /users/password-update`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).UpdatePasswordAfterReset-fm`  
    **Description:** Updates the user's password after a successful password reset.

8. **Admin: Get All Users**  
    **Endpoint:** `GET /admin/users/`  
    **Handler:** `LoanTrackerAPI/Delivery/controllers.(*UserController).GetAllUsers-fm`  
    **Description:** Retrieves a list of all users in the system. (Admin access required)

9. **Admin: Delete User by ID**  
    **Endpoint:** `DELETE /admin/users/:id`  
    **Description:** Deletes a user by their ID. (Admin access required)

For more details, refer to the [Postman Documentation](https://documenter.getpostman.com/view/37364622/2sAXjJ6De8).