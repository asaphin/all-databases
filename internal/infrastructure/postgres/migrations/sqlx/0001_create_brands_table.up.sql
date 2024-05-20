CREATE TABLE brands
(
    id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name   VARCHAR(100) UNIQUE NOT NULL,
    slogan TEXT
);

CREATE INDEX idx_brand_name ON brands (name);
