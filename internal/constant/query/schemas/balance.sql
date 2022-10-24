CREATE TABLE balance (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "client_id" STRING UNIQUE NOT NULL ,
                         "amount" DECIMAL NOT NULL,
                         "status" STRING NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);

