--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4
-- Dumped by pg_dump version 14.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: operations; Type: TABLE; Schema: public; Owner: bulkadmin
--

CREATE TABLE public.operations (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    created_by character varying(50) NOT NULL,
    type character varying(20) NOT NULL,
    port character varying(50) NOT NULL,
    startop timestamp(0) with time zone NOT NULL,
    endop timestamp(0) with time zone NOT NULL,
    vessel bigint NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.operations OWNER TO bulkadmin;

--
-- Name: operations_id_seq; Type: SEQUENCE; Schema: public; Owner: bulkadmin
--

CREATE SEQUENCE public.operations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.operations_id_seq OWNER TO bulkadmin;

--
-- Name: operations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bulkadmin
--

ALTER SEQUENCE public.operations_id_seq OWNED BY public.operations.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: bulkadmin
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    created_by character varying(50) NOT NULL,
    sales_number character varying(20) NOT NULL,
    purchasing_number character varying(20),
    customer character varying(100) NOT NULL,
    loading_berth character varying(50) NOT NULL,
    destination_port character varying(50) NOT NULL,
    destination_berth character varying(50) NOT NULL,
    product character varying(50) NOT NULL,
    volume numeric(7,2) NOT NULL,
    sales_rep character varying(50) NOT NULL,
    crp character varying(50) NOT NULL,
    vessel bigint NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.orders OWNER TO bulkadmin;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: bulkadmin
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_id_seq OWNER TO bulkadmin;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bulkadmin
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: vessels; Type: TABLE; Schema: public; Owner: bulkadmin
--

CREATE TABLE public.vessels (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    created_by character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    voyage character varying(50) NOT NULL,
    service character varying(50) NOT NULL,
    status character varying(50),
    tolerance character varying(50) NOT NULL,
    booking character varying(10),
    internal_note text,
    external_note text,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.vessels OWNER TO bulkadmin;

--
-- Name: vessel_list_view; Type: VIEW; Schema: public; Owner: bulkadmin
--

CREATE VIEW public.vessel_list_view AS
 SELECT DISTINCT ON (c.sales_number) a.name,
    c.sales_number,
    c.purchasing_number,
    c.customer,
    c.volume,
    c.product,
    b.startop
   FROM ((public.vessels a
     JOIN public.operations b ON ((b.vessel = a.id)))
     JOIN public.orders c ON ((c.vessel = a.id)))
  WHERE ((b.type)::text = 'load'::text)
  ORDER BY c.sales_number, b.startop DESC;


ALTER TABLE public.vessel_list_view OWNER TO bulkadmin;

--
-- Name: vessels_id_seq; Type: SEQUENCE; Schema: public; Owner: bulkadmin
--

CREATE SEQUENCE public.vessels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vessels_id_seq OWNER TO bulkadmin;

--
-- Name: vessels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bulkadmin
--

ALTER SEQUENCE public.vessels_id_seq OWNED BY public.vessels.id;


--
-- Name: operations id; Type: DEFAULT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.operations ALTER COLUMN id SET DEFAULT nextval('public.operations_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: vessels id; Type: DEFAULT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.vessels ALTER COLUMN id SET DEFAULT nextval('public.vessels_id_seq'::regclass);


--
-- Data for Name: operations; Type: TABLE DATA; Schema: public; Owner: bulkadmin
--

COPY public.operations (id, created_at, created_by, type, port, startop, endop, vessel, version) FROM stdin;
1	2022-06-25 17:11:41-03	Heather Valiant	load	Houston	2022-07-02 00:00:00-03	2022-07-06 00:00:00-03	1	1
2	2022-06-25 17:11:41-03	Heather Valiant	discharge	Santos	2022-08-02 00:00:00-03	2022-08-06 00:00:00-03	1	1
3	2022-06-25 17:11:41-03	Heather Valiant	discharge	Campana	2022-08-10 00:00:00-03	2022-08-13 00:00:00-03	1	1
4	2022-06-25 17:11:41-03	Heather Valiant	load	Houston	2022-07-15 00:00:00-03	2022-07-19 00:00:00-03	2	1
5	2022-06-25 17:11:41-03	Heather Valiant	discharge	Buenaventura	2022-08-15 00:00:00-03	2022-08-19 00:00:00-03	2	1
6	2022-06-25 17:11:41-03	Heather Valiant	discharge	Callao	2022-08-21 00:00:00-03	2022-08-23 00:00:00-03	2	1
7	2022-06-28 16:26:10-03	Heather Valiant	load	Houston	2022-07-12 00:00:00-03	2022-07-15 00:00:00-03	1	1
8	2022-06-28 16:32:22-03	Heather Valiant	discharge	Santos	2022-08-12 00:00:00-03	2022-08-15 00:00:00-03	1	1
9	2022-06-28 16:33:29-03	Heather Valiant	discharge	Campana	2022-08-17 00:00:00-03	2022-08-20 00:00:00-03	1	1
10	2022-06-29 13:20:55-03	Heather Valiant	load	Houston	2022-07-17 00:00:00-03	2022-07-20 00:00:00-03	1	1
11	2022-06-29 13:21:42-03	Heather Valiant	load	Houston	2022-07-19 00:00:00-03	2022-07-22 00:00:00-03	1	1
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: bulkadmin
--

COPY public.orders (id, created_at, created_by, sales_number, purchasing_number, customer, loading_berth, destination_port, destination_berth, product, volume, sales_rep, crp, vessel, version) FROM stdin;
7	2022-06-28 17:57:16-03	Cindy Pecina	123456	654321	Quimica Callegari	LBC#2	Campana	Tacsa	n-Propanol	512.35	Rodrigo Valente	Cindy Pecina	2	3
8	2022-06-29 09:04:30-03	Cindy Pecina	123457	654322	Anjo	LBC#1	Santos	Ilha Barnabe	n-Propanol	400.00	Rodrigo Valente	Cindy Pecina	1	1
1	2022-06-25 17:11:41-03	Cindy Pecina	252525	525252	CQM	LBC#1	Santos	Alemoa	n-Propanol	400.00	Rodrigo Valente	Cindy Pecina	1	1
2	2022-06-25 17:11:41-03	Cindy Pecina	252526	PO656	Brenntag Brasil	LBC#1	Santos	Alemoa	n-Propanol	100.00	Rodrigo Valente	Cindy Pecina	1	1
3	2022-06-25 17:11:41-03	Cindy Pecina	252527	525253	CQM	LBC#1	Santos	Alemoa	n-Butyl Acetate	300.00	Rodrigo Valente	Cindy Pecina	1	1
4	2022-06-25 17:11:41-03	Erika Cantu	252528	SI-6566	Sucroal	LBC#1	Buenaventura	TBI	n-Propanol	900.00	Sharin Hernandez	Erika Cantu	2	1
5	2022-06-25 17:11:41-03	Erika Cantu	252529	34112	GTM Peru	LBC#1	Callao	TBI	n-Propanol	300.00	Sharin Hernandez	Erika Cantu	2	1
6	2022-06-25 17:11:41-03	Erika Cantu	252530	34113	GTM Peru	LBC#1	Callao	TBI	n-Propyl Acetate	252.50	Sharin Hernandez	Erika Cantu	2	1
\.


--
-- Data for Name: vessels; Type: TABLE DATA; Schema: public; Owner: bulkadmin
--

COPY public.vessels (id, created_at, created_by, name, voyage, service, status, tolerance, booking, internal_note, external_note, version) FROM stdin;
2	2022-06-25 17:11:41-03	Heather Valiant	Ginga Puma	124	WCSA	accepted	2% MOOLCO	81281233			1
3	2022-06-25 17:11:41-03	Chi Ho	Ginga Caracal	114	WCSA	accepted	2% MOOLCO	81281234			1
1	2022-06-25 17:11:41-03	Heather Valiant	Stolt Ocelot	123	ECSA	firmed	3% MOOLCO	81281232		Booking confirmed.	3
4	2022-06-26 15:52:19-03	Chi Ho	Stolt Teal	456	ECSA	accepted	3% MOOLCO	84838212	Hello!	GoodBye!	2
6	2022-06-27 18:24:04-03	Chi Ho	Stolt Vision	012	ECSA	loaded	2% MOOLCO	84838266			2
\.


--
-- Name: operations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bulkadmin
--

SELECT pg_catalog.setval('public.operations_id_seq', 11, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bulkadmin
--

SELECT pg_catalog.setval('public.orders_id_seq', 9, true);


--
-- Name: vessels_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bulkadmin
--

SELECT pg_catalog.setval('public.vessels_id_seq', 6, true);


--
-- Name: operations operations_pkey; Type: CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.operations
    ADD CONSTRAINT operations_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: vessels vessels_booking_key; Type: CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.vessels
    ADD CONSTRAINT vessels_booking_key UNIQUE (booking);


--
-- Name: vessels vessels_pkey; Type: CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.vessels
    ADD CONSTRAINT vessels_pkey PRIMARY KEY (id);


--
-- Name: operations operations_vessel_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.operations
    ADD CONSTRAINT operations_vessel_fkey FOREIGN KEY (vessel) REFERENCES public.vessels(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: orders orders_vessel_fkey; Type: FK CONSTRAINT; Schema: public; Owner: bulkadmin
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_vessel_fkey FOREIGN KEY (vessel) REFERENCES public.vessels(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

