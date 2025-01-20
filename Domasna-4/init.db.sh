#!/bin/bash
service postgresql start
sleep 10
su - postgres
psql -c "CREATE ROLE root WITH SUPERUSER LOGIN PASSWORD 'root';"
psql -c "CREATE TABLE IF NOT EXISTS public.companies (
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    id character varying(20) COLLATE pg_catalog.\"default\" NOT NULL,
    name character varying(100) COLLATE pg_catalog.\"default\",
    address character varying(128) COLLATE pg_catalog.\"default\",
    city character varying(128) COLLATE pg_catalog.\"default\",
    country character varying(128) COLLATE pg_catalog.\"default\",
    email character varying(128) COLLATE pg_catalog.\"default\",
    website character varying(128) COLLATE pg_catalog.\"default\",
    contact_name character varying(128) COLLATE pg_catalog.\"default\",
    contact_phone character varying(128) COLLATE pg_catalog.\"default\",
    contact_email character varying(128) COLLATE pg_catalog.\"default\",
    phone character varying(128) COLLATE pg_catalog.\"default\",
    fax character varying(128) COLLATE pg_catalog.\"default\",
    prospect character varying(128) COLLATE pg_catalog.\"default\",
    ticker character varying(10) COLLATE pg_catalog.\"default\",
    url character varying(128) COLLATE pg_catalog.\"default\",
    CONSTRAINT companies_pkey PRIMARY KEY (id),
    CONSTRAINT companies_ticker_key UNIQUE (ticker)
) TABLESPACE pg_default;"

psql -c "CREATE TABLE IF NOT EXISTS public.history_records (
    id character varying(20) COLLATE pg_catalog.\"default\" NOT NULL,
    date date NOT NULL,
    ticker character varying(10) COLLATE pg_catalog.\"default\",
    price_last_transaction numeric(15, 2),
    max numeric(15, 2),
    min numeric(15, 2),
    average_price numeric(15, 2),
    revenue_percent numeric(10, 2),
    amount numeric(15, 2),
    revenue_best numeric(15, 2),
    revenue_total numeric(15, 2),
    currency character varying(5) COLLATE pg_catalog.\"default\",
    CONSTRAINT history_records_pkey PRIMARY KEY (id),
    CONSTRAINT fk_ticker FOREIGN KEY (ticker) REFERENCES public.companies (ticker) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
) TABLESPACE pg_default;"

psql -c "ALTER TABLE IF EXISTS public.history_records OWNER TO root;"
