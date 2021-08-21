CREATE TABLE customer (
    id SERIAL NOT NULL primary key,
    name VARCHAR(100),
    phone VARCHAR(11),
    cpf VARCHAR(14),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product (
    id SERIAL NOT NULL primary key,
    name VARCHAR(100),
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE costumer_address (
    id SERIAL NOT NULL PRIMARY KEY,
    street VARCHAR(100) NOT NULL,
    district VARCHAR(100) NOT NULL,
    number INTEGER NOT NULL,
    complement VARCHAR(100),
    id_costumer INTEGER NOT NULL CONSTRAINT fk_costumer_address REFERENCES customer (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL NOT NULL primary key,
    id_costumer INTEGER NOT NULL CONSTRAINT fk_orders_customer REFERENCES customer (id),
    id_product INTEGER NOT NULL CONSTRAINT fk_orders_product REFERENCES product (id),
    total_amount INTEGER NOT NULL,
    type_payment VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_status (
    id SERIAL NOT NULL primary key,
    status TEXT NOT NULL,
    id_order INTEGER NOT NULL CONSTRAINT fk_id_order_status REFERENCES orders (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_stock (
    id SERIAL NOT NULL primary key,
    id_product INTEGER NOT NULL CONSTRAINT fk_id_product_stock REFERENCES product (id),
    stock_out INTEGER NOT NULL,
    stock_input INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);