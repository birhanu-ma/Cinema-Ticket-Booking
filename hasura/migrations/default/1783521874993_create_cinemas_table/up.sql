CREATE TABLE cinemas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL ,
    location TEXT NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    update_at TIMESTAMPTZ NOT NULL DEFAULT now()
)