CREATE DATABASE vaccine ENCODING = 'UTF8' LOCALE = 'en_US.utf8';

ALTER DATABASE vaccine OWNER TO postgres;

\connect vaccine

CREATE TABLE public.flag (
    id integer NOT NULL,
    text text
);


ALTER TABLE public.flag OWNER TO postgres;

CREATE SEQUENCE public.flag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.flag_id_seq OWNER TO postgres;

ALTER SEQUENCE public.flag_id_seq OWNED BY public.flag.id;

CREATE TABLE public.vaccines (
    id integer NOT NULL,
    name text,
    description text,
    manufacturer text
);

ALTER TABLE public.vaccines OWNER TO postgres;

CREATE SEQUENCE public.vaccines_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vaccines_id_seq OWNER TO postgres;

ALTER SEQUENCE public.vaccines_id_seq OWNED BY public.vaccines.id;


ALTER TABLE ONLY public.flag ALTER COLUMN id SET DEFAULT nextval('public.flag_id_seq'::regclass);

ALTER TABLE ONLY public.vaccines ALTER COLUMN id SET DEFAULT nextval('public.vaccines_id_seq'::regclass);

INSERT INTO public.flag (id, text) VALUES (1, 'HZU18{sQL-!nj3ct3on---nice}');

INSERT INTO public.vaccines (id, name, description, manufacturer) VALUES (2, 'AZD1222 Vaxzevria', 'Recombinant ChAdOx1 adenoviral vector encoding the Spike protein antigen of the SARS-CoV-2.', 'AstraZeneca, AB');
INSERT INTO public.vaccines (id, name, description, manufacturer) VALUES (3, 'mRNA-1273', 'mRNA-based vaccine encapsulated in lipid nanoparticle (LNP)', 'Moderna Biotech');
INSERT INTO public.vaccines (id, name, description, manufacturer) VALUES (4, 'COVID-19 Vaccine (Vero Cell), Inactivated/CoronavacTM', 'Inactivated, produced in Vero cells', 'Sinovac Life Sciences Co., Ltd');
INSERT INTO public.vaccines (id, name, description, manufacturer) VALUES (5, 'Sputnik V', 'Human Adenovirus Vector-based Covid-19 vaccine', 'Russian Direct Investment Fund');
INSERT INTO public.vaccines (id, name, description, manufacturer) VALUES (1, 'BNT162b2/COMIRNATY Tozinameran (INN)', 'Nucleoside modified mRNA', 'BioNTech Manufacturing
GmbH');

SELECT pg_catalog.setval('public.flag_id_seq', 1, true);
SELECT pg_catalog.setval('public.vaccines_id_seq', 5, true);

ALTER TABLE ONLY public.flag
    ADD CONSTRAINT flag_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.vaccines
    ADD CONSTRAINT vaccines_pkey PRIMARY KEY (id);
