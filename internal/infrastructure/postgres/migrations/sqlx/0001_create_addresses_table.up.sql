CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS addresses
(
    id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    type            VARCHAR(50),  -- Type of address (e.g., Home, Work, etc.)
    in_care_of_name VARCHAR(100), -- Name of the person or organization the address is care of
    street          VARCHAR(255), -- Street name
    street_number   VARCHAR(20),  -- Street number (could be alphanumeric)
    apartment       VARCHAR(20),  -- Apartment number or unit
    suite           VARCHAR(20),  -- Suite number
    floor           VARCHAR(20),  -- Floor number
    city            VARCHAR(100), -- City name
    state           VARCHAR(100), -- State or region name
    province        VARCHAR(100), -- Province name (if applicable)
    zip             VARCHAR(20),  -- ZIP code or postal code
    postal_code     VARCHAR(20),  -- Postal code (if applicable)
    country         VARCHAR(100), -- Country name
    latitude        DOUBLE PRECISION,
    longitude       DOUBLE PRECISION
);
