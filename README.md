# GO tickets CRUD API
A CRUD API that shows simple tasks such as create, retrieve, update, delete in context to GO tickets. Here we can create a new GO ticket for the passenger that includes the attributes such as ticket id, name, source, destination and ticket price etc. We can perform tasks like retrieving details using ticket id, update name or price and so on...

## Server running
![image](https://github.com/user-attachments/assets/3102622f-42fe-4855-8a5e-01e7ff5fb060)

### Following tasks can be perfomed
1. Create new ticket
2. Get all tickets
3. Get tickets by ID
4. Update an existing ticket
5. Delete a ticket

### 1. Create a new ticket
This API allows you to create a GO ticket by passing the following parameters:
- ID: Describes the ticket ID
- passenger: The name of the ticket holder
- origin: The source point of journey 
- destination: The destination city
- price: the toatl price of the ticket (in $)
- date: Date of the ticket printed (in YYYY/MM/DD)

![image](https://github.com/user-attachments/assets/0caa9a1e-df32-4267-9fa1-a909a82ae84c)

### 2. Get All tickets
The API provides a list of all GO tickets in the database.
![image](https://github.com/user-attachments/assets/52c8ca85-2a66-4265-b9c1-2c611db86000)

### 3. Get ticket by ID
This API allows you to retrieve a ticket by its ID.
![image](https://github.com/user-attachments/assets/2b189f66-9ca9-4b2d-aee2-a731b455a585)


### 4. Update the ticket by ID
This API allows you to update an existing ticket by its ID. You can modify the passenger, travel date, and ticket price.
Existing details of a ticket
![image](https://github.com/user-attachments/assets/ae8fe393-a187-4366-9105-7856056daafb)

Updated the price of the ticket
![image](https://github.com/user-attachments/assets/3bb85ba2-f32c-47f0-9bd7-335364d90334)

Verified the update on ticket price
![image](https://github.com/user-attachments/assets/6d89d478-f9ac-481f-b547-35afeb180246)

### 5. Delete a ticket by ID
This API allows you to delete a ticket based on its ID.
![image](https://github.com/user-attachments/assets/08182717-ed04-4287-918f-c30da5e75881)

