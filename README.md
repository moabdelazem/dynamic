# Dynamic DevOps API

A RESTful API service built with Go that provides product management capabilities.

## Project Structure

```
.
├── cmd/                    # Application entry points
│   └── main.go             # Main application
├── internal/               # Private application code
│   ├── api/                # API-related code
│   │   ├── handlers/       # HTTP handlers
│   │   ├── middleware/     # HTTP middleware
│   │   └── router.go       # Router setup
│   ├── models/             # Data models
│   └── store/              # Data storage
├── pkg/                    # Public libraries
│   └── config/             # Configuration
├── k8s/                    # Kubernetes configuration
├── Dockerfile              # Docker configuration
├── Makefile                # Build automation
└── README.md               # Documentation
```

## API Endpoints

### Base Endpoints

- `GET /` - Welcome message
- `GET /health` - Health check endpoint
- `GET /info` - Information about the developer

### Product Endpoints

- `GET /products` - List all products (with pagination)
- `POST /products` - Create a new product
- `GET /products/{id}` - Get a specific product
- `PUT /products/{id}` - Update a specific product
- `DELETE /products/{id}` - Delete a specific product

## Pagination

The list products endpoint supports pagination with the following query parameters:

- `page` - Page number (default: 1)
- `limit` - Number of items per page (default: 10)

Example: `GET /products?page=2&limit=5`

## Environment Variables

The application can be configured using the following environment variables:

- `API_PORT` - Port to run the server on (default: 8080)

## Running the Application

### Local Development

```bash
# Run the application
make run

# Build the application
make build

# Run the built binary
./bin/dynamic
```

### Using Docker

```bash
# Build the Docker image
docker build -t dynamic-devops-api .

# Run the Docker container
docker run -p 8080:8080 dynamic-devops-api
```

### Using Kubernetes

```bash
# Apply the Kubernetes configuration
kubectl apply -f k8s/
```

## API Examples

### Create a Product

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "id": "4",
    "name": "New Product",
    "description": "A new product description",
    "price": 150
  }'
```

### List Products

```bash
curl -X GET http://localhost:8080/products
```

### Get a Product

```bash
curl -X GET http://localhost:8080/products/1
```

### Update a Product

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Product Name",
    "price": 120
  }'
```

### Delete a Product

```bash
curl -X DELETE http://localhost:8080/products/1
```
