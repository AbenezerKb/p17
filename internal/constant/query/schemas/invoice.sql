
CREATE TYPE payment AS ENUM (
    'Prepaid',
    'Postpaid'
    );


CREATE TABLE invoice (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "invoice_number" bigint NOT NULL ,
                         "client_id" STRING NOT NULL ,
                         "payment" payment NOT NULL,
                         "amount" DECIMAL NOT NULL,
                         "total_monthly" DECIMAL NOT NULL,
                         "discount" DECIMAL NULL,
                         "tax" DECIMAL NULL,
                         "tax_rate" DECIMAL NULL,
                         "created_at" timestamp default now(),
                         "updated_at" timestamp NULL
);
