CREATE TABLE templates (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "template_id" STRING NOT NULL,
                         "client" STRING NOT NULL,
                         "template" STRING NOT NULL,
                         "category" STRING NOT NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);

