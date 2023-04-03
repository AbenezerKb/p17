CREATE TABLE users (
                       "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       "full_name" STRING NULL,
                       "phone" varchar UNIQUE NOT NULL,
                       "password" STRING NULL,
                       "created_at" timestamp default now(),
                       "updated_at" timestamp NULL
);
CREATE INDEX ON "users" ("phone");

ALTER TABLE
    users
ADD
    CONSTRAINT unique_phone UNIQUE (phone);

CREATE TABLE clients (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "title" STRING NOT NULL ,
                         "phone" varchar UNIQUE NOT NULL,
                         "email" STRING NOT NULL ,
                         "password" STRING NOT NULL ,
                         "status" STRING NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);
ALTER TABLE clients ADD CONSTRAINT unique_phone UNIQUE (phone);

CREATE TABLE templates (
                           "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           "template_id" STRING NOT NULL,
                           "client" STRING NOT NULL,
                           "template" STRING NOT NULL,
                           "category" STRING NOT NULL,
                           "created_at" timestamp default now(),
                           "updated_at" timestamp NULL
);


CREATE TYPE message_type AS ENUM (
    'Incoming',
    'OutGoing'
    );

CREATE TABLE messages (
                          "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          "sender_phone" STRING NOT NULL,
                          "content" STRING NOT NULL,
                          "price" DECIMAL NOT NULL,
                          "receiver_phone" STRING NOT NULL,
                          "type" message_type NOT NULL,
                          "status" STRING NULL,
                          "delivery_status" STRING NULL,
                          "message_id" STRING NULL,
                          "created_at" timestamp default now()
);


CREATE TABLE balance (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "client_id" STRING UNIQUE NOT NULL ,
                         "amount" DECIMAL NOT NULL,
                         "status" STRING NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);

CREATE TABLE system_config (
                                 "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                 "setting_name" STRING NOT NULL,
                                 "setting_value" STRING NOT NULL,
                                 "created_at" timestamp default now(),
                                 "updated_at" timestamp
);


ALTER TABLE balance ADD CONSTRAINT unique_client UNIQUE (client_id);


ALTER TABLE balance ADD CONSTRAINT check_amount check ( amount>=0 );


CREATE TYPE payment AS ENUM (
    'Prepaid',
    'Postpaid'
    );


CREATE TABLE invoice (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "invoice_number" bigint NOT NULL ,
                         "client_id" STRING NOT NULL ,
                         "payment_type" payment NOT NULL,
                         "current_balance" DECIMAL NOT NULL,
                         "balance_at_beginning" DECIMAL NOT NULL,
                         "discount" DECIMAL NOT NULL,
                         "MessageCount" jsonb NOT NULL ,
                         "ClientTransaction" jsonb NOT NULL ,
                         "tax" DECIMAL NOT NULL,
                         "tax_rate" DECIMAL NOT NULL,
                         "created_at" timestamp default now()
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
