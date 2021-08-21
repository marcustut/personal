-- migrate:up
ALTER TABLE public.todos ADD updated_at TIMESTAMP DEFAULT NOW();

-- migrate:down
ALTER TABLE public.todos DROP COLUMN updated_at;
