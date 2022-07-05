CREATE TABLE vessels (
    id bigint NOT NULL UNIQUE,
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

CREATE TABLE operations (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    created_by character varying(50) NOT NULL,
    type character varying(20) NOT NULL,
    port character varying(50) NOT NULL,
    startop timestamp(0) with time zone NOT NULL,
    endop timestamp(0) with time zone NOT NULL,
    vessel bigint REFERENCES vessels(id) ON UPDATE CASCADE ON DELETE CASCADE,
    version integer DEFAULT 1 NOT NULL
);

CREATE TABLE orders (
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
    vessel bigint REFERENCES vessels(id) ON UPDATE CASCADE ON DELETE CASCADE,
    version integer DEFAULT 1 NOT NULL
);

