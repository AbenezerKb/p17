CREATE TABLE system_config (
                                 "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                 "setting_name" STRING NOT NULL,
                                 "setting_value" STRING NOT NULL,
                                 "created_at" timestamp default now(),
                                 "updated_at" timestamp
)