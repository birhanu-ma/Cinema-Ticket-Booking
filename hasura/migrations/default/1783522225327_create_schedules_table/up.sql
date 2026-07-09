CREATE TABLE schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    movies_id UUID REFERENCES movies(id) ON DELETE SET NULL,
    cinema_id UUID REFERENCES cinemas(id) ON DELETE SET NULL,
    price NUMERIC(10, 2) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);