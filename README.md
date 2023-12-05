# Scooter Microservices Gateway INE5670

This is a Go project for mobile and emmbeded systems class at Federal College of Santa Catarina.
Consists in a API gateway with multiple microservices behind, used to control scooters, as in a shared scooter app.

## API Gateway

It's a simple Go server, acting just as a proxy for the microservices routing the requests based on path.

## Microservices:

1. User:  User CRUD related operations
2. Scooter: Scooter related operations:
    - Getting near me scooters
    - Scooter CRUD
    - Get scooter location
    - Update scooter location
    - Listing
3. Rent: Scooter rent related operations:
    -  Rent a scooter
    -  Finish scooter rent
4. Payment: Payment transactions for billing scooter usage.

