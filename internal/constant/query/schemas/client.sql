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