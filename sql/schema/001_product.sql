CREATE TABLE item
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255)   NOT NULL,
    brand      VARCHAR(255)   NOT NULL,
    price      NUMERIC(10, 2) NOT NULL,
    available  BOOLEAN        NOT NULL,
    url        VARCHAR(255)   NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT product_data_name_brand_unique UNIQUE (name, brand)
);

CREATE TABLE size
(
    id         BIGSERIAL PRIMARY KEY,
    product_id BIGINT      NOT NULL REFERENCES item (id),
    size       VARCHAR(50) NOT NULL,
    quantity   INTEGER     NOT NULL
);
