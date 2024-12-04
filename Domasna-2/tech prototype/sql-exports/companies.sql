-- Table: public.companies

-- DROP TABLE IF EXISTS public.companies;

CREATE TABLE IF NOT EXISTS public.companies
(
	created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    name character varying(100) COLLATE pg_catalog."default",
    address character varying(128) COLLATE pg_catalog."default",
    city character varying(128) COLLATE pg_catalog."default",
    country character varying(128) COLLATE pg_catalog."default",
    email character varying(128) COLLATE pg_catalog."default",
    website character varying(128) COLLATE pg_catalog."default",
    contact_name character varying(128) COLLATE pg_catalog."default",
    contact_phone character varying(128) COLLATE pg_catalog."default",
    contact_email character varying(128) COLLATE pg_catalog."default",
    phone character varying(128) COLLATE pg_catalog."default",
    fax character varying(128) COLLATE pg_catalog."default",
    prospect character varying(128) COLLATE pg_catalog."default",
    ticker character varying(10) COLLATE pg_catalog."default",
    url character varying(128) COLLATE pg_catalog."default",
    CONSTRAINT companies_pkey PRIMARY KEY (id),
    CONSTRAINT companies_ticker_key UNIQUE (ticker)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.companies
    OWNER to root;