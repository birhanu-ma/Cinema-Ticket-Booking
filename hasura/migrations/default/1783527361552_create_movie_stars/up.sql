CREATE TABLE movie_stars (
    movie_id UUID REFERENCES movies(id) ON DELETE CASCADE,
    star_id UUID REFERENCES stars(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (movie_id, star_id)
);