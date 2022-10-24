CREATE TABLE users (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "full_name" STRING NULL,
    "phone" varchar UNIQUE NOT NULL,
    "password" STRING NULL,
    "created_at" timestamp default now(),
    "updated_at" timestamp NULL
);