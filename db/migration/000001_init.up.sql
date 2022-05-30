CREATE TABLE manager
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(127) NOT NULL UNIQUE,
    password   TEXT         NOT NULL,
    first_name VARCHAR(127) NOT NULL,
    last_name  VARCHAR(127) NOT NULL,
    phone      VARCHAR(20)  NOT NULL,
    is_deleted BOOL         NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE store
(
    id                 SERIAL PRIMARY KEY,
    title              VARCHAR(127) NOT NULL,
    info               TEXT,
    address            VARCHAR(127) NOT NULL,
    image              VARCHAR(127) NOT NULL,
    avg_delivery_time  BIGINT       NOT NULL,
    working_hour_begin TIME         NOT NULL,
    working_hour_end   TIME         NOT NULL,
    min_order_price    INTEGER      NOT NULL,
    delivery_price     INTEGER      NOT NULL,
    contact_phone      VARCHAR(20)  NOT NULL,
    manager_id         INTEGER REFERENCES manager (id),
    is_deleted         BOOL         NOT NULL DEFAULT false,
    created_at         TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ  NOT NULL DEFAULT now(),
    CONSTRAINT min_order_price_check CHECK ( min_order_price > 0 ),
    CONSTRAINT delivery_price_check CHECK ( delivery_price > 0 )
);

CREATE TABLE category
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(127) NOT NULL,
    store_id   INTEGER REFERENCES store (id),
    is_deleted BOOL         NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE product
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(127) NOT NULL,
    description TEXT,
    price       INTEGER      NOT NULL,
    image       VARCHAR(127) NOT NULL,
    is_active   BOOL         NOT NULL DEFAULT true,
    category_id INTEGER REFERENCES category (id),
    is_deleted  BOOL         NOT NULL DEFAULT false,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
    CONSTRAINT price_check CHECK ( price > 0 )
);

CREATE TABLE customer
(
    id          SERIAL PRIMARY KEY,
    external_id BIGINT      NOT NULL,
    username    VARCHAR(127),
    phone       VARCHAR(127),
    first_name  VARCHAR(127),
    last_name   VARCHAR(127),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TYPE order_payment_method AS ENUM ('CASH', 'CARD');
CREATE TYPE order_state AS ENUM ('PENDING', 'CANCELED', 'ACCEPTED', 'DELIVERY_IN_PROGRESS', 'DELIVERED');

CREATE TABLE orders
(
    id                  SERIAL PRIMARY KEY,
    customer_id         INTEGER REFERENCES customer (id),
    store_id            INTEGER REFERENCES store (id),
    address             VARCHAR(127) NOT NULL,
    state               order_state,
    payment_method      order_payment_method,
    cancellation_reason TEXT,
    is_deleted          BOOL         NOT NULL DEFAULT false,
    created_at          TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE order_item
(
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER REFERENCES orders (id),
    product_id INTEGER REFERENCES product (id),
    amount     INTEGER NOT NULL DEFAULT 1,
    CONSTRAINT amount_check CHECK ( amount > 0 )
);