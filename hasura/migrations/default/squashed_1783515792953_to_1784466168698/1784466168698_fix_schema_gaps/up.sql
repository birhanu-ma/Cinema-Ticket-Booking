
ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_schedule_seat_id_fkey
    FOREIGN KEY (schedule_seat_id)
    REFERENCES public.schedule_seats(id)
    ON DELETE SET NULL;

ALTER TABLE public.payments
    ALTER COLUMN paid_at TYPE timestamp with time zone
    USING (CURRENT_DATE + paid_at);


ALTER TABLE public.schedule_seats
    ALTER COLUMN schedule_id DROP DEFAULT;

ALTER TABLE public.schedule_seats
    ALTER COLUMN seat_id DROP DEFAULT;

ALTER TABLE public.seats
    ALTER COLUMN hall_id DROP DEFAULT;

ALTER TABLE public.tickets
    ALTER COLUMN schedule_seat_id DROP DEFAULT;