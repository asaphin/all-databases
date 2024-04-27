CREATE TABLE addresses
(
    id              SERIAL PRIMARY KEY,
    type            VARCHAR(50) NOT NULL,
    in_care_of_name VARCHAR(255),
    street          VARCHAR(255),
    street_number   VARCHAR(50),
    apartment       VARCHAR(50),
    locality        VARCHAR(100),
    region          VARCHAR(100),
    postal_code     VARCHAR(50),
    country         VARCHAR(100),
    additional_info JSONB,
    latitude        DOUBLE PRECISION,
    longitude       DOUBLE PRECISION
);
