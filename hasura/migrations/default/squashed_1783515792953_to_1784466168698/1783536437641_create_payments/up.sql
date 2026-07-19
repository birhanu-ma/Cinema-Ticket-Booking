CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    transaction_ref TEXT, 
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(), 
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now() 
);