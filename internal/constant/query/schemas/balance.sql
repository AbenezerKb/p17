CREATE TABLE balance (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "client_id" STRING UNIQUE NOT NULL ,
                         "amount" DECIMAL NOT NULL,
                         "client_email" STRING UNIQUE NOT NULL ,
                         "status" STRING NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);

CREATE TYPE transfer AS ENUM (
    'CREDITING',
    'DEBITING'
    );

CREATE TABLE client_transaction(
                                   "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                   "client_id" STRING UNIQUE NOT NULL ,
                                   "amount" DECIMAL NOT NULL,
                                   "type" transfer NOT NULL,
                                   "created_at" timestamp default now()
);