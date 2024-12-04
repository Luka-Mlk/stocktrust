-- Table: public.history_records

-- DROP TABLE IF EXISTS public.history_records;

CREATE TABLE IF NOT EXISTS public.history_records
(
    id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    date date NOT NULL,
    ticker character varying(10) COLLATE pg_catalog."default",
    price_last_transaction numeric(15,2),
    max numeric(15,2),
    min numeric(15,2),
    average_price numeric(15,2),
    revenue_percent numeric(10,2),
    amount numeric(15,2),
    revenue_best numeric(15,2),
    revenue_total numeric(15,2),
    currency character varying(5) COLLATE pg_catalog."default",
    CONSTRAINT history_records_pkey PRIMARY KEY (id),
    CONSTRAINT fk_ticker FOREIGN KEY (ticker)
        REFERENCES public.companies (ticker) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.history_records
    OWNER to root;