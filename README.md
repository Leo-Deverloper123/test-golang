# Hospital Middleware System

## Project Structure

```
.
├── config/
│   └── database.go
├── handlers/
│   ├── patient_handler.go
│   └── staff_handler.go
├── middleware/
│   └── auth.go
├── models/
│   ├── hospital.go
│   ├── patient.go
│   └── staff.go
├── routes/
│   └── routes.go
├── tests/
│   ├── patient_test.go
│   └── staff_test.go
├── utils/
│   └── jwt.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── main.go
├── nginx.conf
└── README.md
```

## API Specification

### Staff APIs

#### Create Staff
- **Endpoint**: POST /staff/create
- **Input**:
  ```json
  {
    "username": "string",
    "password": "string",
    "hospital_id": "number"
  }
  ```

#### Staff Login
- **Endpoint**: POST /staff/login
- **Input**:
  ```json
  {
    "username": "string",
    "password": "string",
    "hospital": "number"
  }
  ```

### Patient APIs

#### Search Patient
- **Endpoint**: POST /patient/search
- **Authentication**: Required (JWT Token)
- **Input**:
  ```json
  {
    "national_id": "string",
    "passport_id": "string",
    "first_name": "string",
    "middle_name": "string",
    "last_name": "string",
    "date_of_birth": "string",
    "phone_number": "string",
    "email": "string"
  }
  ```

#### External Patient Search
- **Endpoint**: GET /patient/search/:id
- **Authentication**: Required (JWT Token)
- **Parameters**: id (national_id or passport_id)

## Database Schema

### Hospital Table
```sql
CREATE TABLE hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    api_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Staff Table
```sql
CREATE TABLE staff (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    hospital_id INTEGER REFERENCES hospitals(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Patient Table
```sql
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    first_name_th VARCHAR(255),
    middle_name_th VARCHAR(255),
    last_name_th VARCHAR(255),
    first_name_en VARCHAR(255),
    middle_name_en VARCHAR(255),
    last_name_en VARCHAR(255),
    date_of_birth DATE,
    patient_hn VARCHAR(255) UNIQUE,
    national_id VARCHAR(13) UNIQUE,
    passport_id VARCHAR(255) UNIQUE,
    phone_number VARCHAR(20),
    email VARCHAR(255),
    gender CHAR(1),
    hospital_id INTEGER REFERENCES hospitals(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
```

## Setup Instructions

1. Clone the repository
2. Run `docker-compose up --build`
3. The application will be available at http://localhost:80

## Testing

Run tests using:
```bash
go test ./tests/...
```