CREATE TABLE public.movie_images (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    movie_id uuid NOT NULL,
    image_url text NOT NULL,
    is_featured boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY public.movie_images
    ADD CONSTRAINT movie_images_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.movie_images
    ADD CONSTRAINT movie_images_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;
