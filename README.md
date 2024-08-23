# API Testing with Thunder Client

This project includes a set of API endpoint tests using Thunder Client.

## How to Run the Tests

1. Install the Thunder Client extension in VS Code if you haven't already.
2. Open the Thunder Client tab in VS Code.
3. Import the test collection by clicking on the "Collections" tab, then click on the import button.
4. Navigate to the root folder of this project and select the `thunder-collection.json` file.
5. After importing, you can run the tests by selecting the collection and clicking the "Run All" button.

## Test Documentation

The `thunder-collection.json` file in the root directory contains all the endpoint tests for this project. Below is a summary of the included tests:

- **Create User**: Tests the user creation endpoint `/api/v1/users`.
- **Get User**: Tests retrieving a user by ID `/api/v1/users/`.
- **Get All Users**: Tests retrieving all users `/api/v1/users`.
