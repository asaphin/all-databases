CREATE TABLE IF NOT EXISTS addresses
(
    id              VARCHAR(255) PRIMARY KEY,
    type            VARCHAR(255),
    in_care_of_name VARCHAR(255),
    street          VARCHAR(255),
    street_number   VARCHAR(255),
    apartment       VARCHAR(255),
    suite           VARCHAR(255),
    floor           VARCHAR(255),
    city            VARCHAR(255),
    state           VARCHAR(255),
    province        VARCHAR(255),
    zip             VARCHAR(255),
    postal_code     VARCHAR(255),
    country         VARCHAR(255),
    latitude        DOUBLE PRECISION,
    longitude       DOUBLE PRECISION
);
