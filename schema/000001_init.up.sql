CREATE TABLE users
(
    id            bytea                not null unique primary key,
    login         varchar(255)          not null unique,
    email         varchar(255)          not null unique,
    password_hash text                                not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    updated_at    timestamp default CURRENT_TIMESTAMP not null
);

-- alter sequence users_id_seq RESTART WITH 1;

CREATE TABLE users_data
(
    user_id          bytea references users (id)  on delete cascade,
    full_name        varchar(255)        not null,
    date_of_birth    date        not null,
    number           varchar(50) not null,
    address          varchar(255)        not null,
    discount_percent integer          default 0,
    discount_amount  double precision default 0,
    created_at       timestamp        default CURRENT_TIMESTAMP,
    updated_at       timestamp        default CURRENT_TIMESTAMP
);
CREATE TABLE users_tokens
(
    user_id      bytea  references users (id)  on delete cascade,
    access_hash  varchar(255)                        not null,
    refresh_hash varchar(255)                        not null,
    created_at   timestamp default CURRENT_TIMESTAMP not null
);
CREATE TABLE suppliers
(
    id                   bytea                                unique,
    external_supplier_id int,
    supplier_name        text                                not null,
    image                text,
    description          text,
    created_at           timestamp default CURRENT_TIMESTAMP not null
);
CREATE TABLE products_categories
(
    id            serial        unique,
    category_name text not null,
    created_at    timestamp default CURRENT_TIMESTAMP
);
CREATE TABLE products
(
    id                  bytea        unique,
    product_name        text                                                                 not null,
    category_id         int    references products_categories (id)  on delete cascade     not null,
    external_product_id int                                                              not null,
    created_at          timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE products_suppliers
(
    product_id           bytea references products (id)  on delete cascade                     not null,
    supplier_id          bytea references suppliers (id)  on delete cascade not null,
    external_product_id  int,
    external_supplier_id int,
    price                double precision                    not null,
    image                text                                not null,
    description          text                                not null,
    quantity             integer   default 100,
    created_at           timestamp default CURRENT_TIMESTAMP not null,
    updated_at           timestamp default CURRENT_TIMESTAMP not null
);
CREATE TABLE orders
(
    id               bytea                             not null unique,
    user_id          bytea references users (id)  on delete cascade not null,
    total_price      double precision default 0                 not null,
    status           varchar(20)                                   not null,
    payment_method   varchar(20),
    discount_amount  double precision default 0,
    discount_percent integer          default 0,
    created_at       timestamp        default CURRENT_TIMESTAMP not null,
    updated_at       timestamp        default CURRENT_TIMESTAMP not null


);
CREATE TABLE orders_products
(
--     id                  serial unique,
    product_id          bytea references products (id)  on delete cascade,
    order_id            bytea references orders (id)  on delete cascade,
    numbers_of_products integer,
    purchase_price      double precision                    not null,
    created_at          timestamp default CURRENT_TIMESTAMP not null,
    updated_at          timestamp default CURRENT_TIMESTAMP not null
);


create function trigger_set_timestamp() returns trigger
    language plpgsql
as
$$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$;


