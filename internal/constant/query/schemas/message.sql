
CREATE TYPE message_type AS ENUM (
'Incoming',
'OutGoing'
);

CREATE TABLE messages (
                           "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           "client_id" STRING NOT NULL,
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

