-- auto-generated definition
create table product
(
    sku          varchar(12)    not null
        constraint product_pkey primary key,
    name         varchar(255),
    price_amount numeric(12, 2) not null,
    gallery      varchar(255),
    thumbnail    varchar(255),
    colors       varchar(255)
)

create index product_sku_partial
    on product (sku);


create table product_size
(
    full_sku    varchar(50) not null
        constraint product_size_pkey
            primary key,

    product_sku varchar(12)
        constraint product_size_product_sku_fkey
            references product
            on delete cascade
);
