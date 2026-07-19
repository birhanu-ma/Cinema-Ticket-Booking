
ALTER TABLE public.tickets
    ALTER COLUMN schedule_seat_id SET DEFAULT gen_random_uuid();

ALTER TABLE public.seats
    ALTER COLUMN hall_id SET DEFAULT gen_random_uuid();

ALTER TABLE public.schedule_seats
    ALTER COLUMN seat_id SET DEFAULT gen_random_uuid();

ALTER TABLE public.schedule_seats
    ALTER COLUMN schedule_id SET DEFAULT gen_random_uuid();

ALTER TABLE public.payments
    ALTER COLUMN paid_at TYPE time with time zone
    USING (paid_at::time with time zone);

ALTER TABLE ONLY public.tickets
    DROP CONSTRAINT tickets_schedule_seat_id_fkey;