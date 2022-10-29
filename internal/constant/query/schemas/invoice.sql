
CREATE TYPE payment_type AS ENUM (
    'Prepaid',
    'Postpaid'
    );


CREATE TABLE invoice (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "invoice_number" UUID DEFAULT gen_random_uuid() ,
                         "client_id" STRING NOT NULL,
                         "payment_type" payment_type NOT NULL,
                         "current_balance" DECIMAL NOT NULL,
                         "balance_at_beginning" DECIMAL NOT NULL,
                         "discount" DECIMAL NOT NULL,
                         "message_count" jsonb NOT NULL ,
                         "client_transaction" jsonb NOT NULL ,
                         "tax" DECIMAL NOT NULL,
                         "tax_rate" DECIMAL NOT NULL,
                         "created_at" timestamp default now()
);
