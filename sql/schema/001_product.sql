CREATE TABLE item
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(255)   NOT NULL,
    brand      VARCHAR(255)   NOT NULL,
    price      NUMERIC(10, 2) NOT NULL,
    available  BOOLEAN        NOT NULL,
    url        VARCHAR(255)   NOT NULL,
    last_check TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT product_data_name_brand_unique UNIQUE (name, brand)
);

CREATE TABLE size
(
    id         UUID PRIMARY KEY,
    product_id UUID        NOT NULL REFERENCES item (id) ON DELETE CASCADE ,
    size       VARCHAR(50) NOT NULL,
    quantity   INTEGER     NOT NULL
);
