--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.9
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: city; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE city (
    city_code integer NOT NULL,
    city_name character varying(100) NOT NULL,
    province_code integer NOT NULL
);


ALTER TABLE city OWNER TO cakman;

--
-- Name: city_city_code_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE city_city_code_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE city_city_code_seq OWNER TO cakman;

--
-- Name: city_city_code_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE city_city_code_seq OWNED BY city.city_code;


--
-- Name: country; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE country (
    country_code integer NOT NULL,
    country_name character varying(100) NOT NULL
);


ALTER TABLE country OWNER TO cakman;

--
-- Name: country_country_code_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE country_country_code_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE country_country_code_seq OWNER TO cakman;

--
-- Name: country_country_code_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE country_country_code_seq OWNED BY country.country_code;


--
-- Name: courier; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE courier (
    id bigint NOT NULL,
    username character varying(256) NOT NULL,
    password character varying(256) NOT NULL,
    name character varying(256) NOT NULL,
    phone character varying(20) NOT NULL,
    email character varying(256) NOT NULL,
    create_by bigint DEFAULT 0,
    create_time timestamp with time zone DEFAULT '2017-12-10 15:34:09.772735+07'::timestamp with time zone,
    update_by bigint DEFAULT 0,
    update_time timestamp with time zone DEFAULT '2017-12-10 15:34:09.772735+07'::timestamp with time zone,
    lattitude double precision DEFAULT 0,
    longitude double precision DEFAULT 0,
    update_position_time timestamp with time zone DEFAULT '2017-12-10 15:34:09.772735+07'::timestamp with time zone
);


ALTER TABLE courier OWNER TO cakman;

--
-- Name: courier_id_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE courier_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE courier_id_seq OWNER TO cakman;

--
-- Name: courier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE courier_id_seq OWNED BY courier.id;


--
-- Name: delivery; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE delivery (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    courier_id bigint NOT NULL,
    status character varying(20) DEFAULT ''::character varying NOT NULL,
    note character varying(256) DEFAULT ''::character varying NOT NULL,
    create_by bigint DEFAULT 0 NOT NULL,
    create_time timestamp with time zone DEFAULT '2017-12-10 22:40:28.472763+07'::timestamp with time zone NOT NULL,
    update_by bigint DEFAULT 0 NOT NULL,
    update_time timestamp with time zone DEFAULT '2017-12-10 22:40:28.472763+07'::timestamp with time zone NOT NULL
);


ALTER TABLE delivery OWNER TO cakman;

--
-- Name: delivery_history; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE delivery_history (
    id bigint NOT NULL,
    delivery_id bigint NOT NULL,
    status character varying(20) DEFAULT ''::character varying NOT NULL,
    note character varying(256) DEFAULT ''::character varying NOT NULL,
    create_time timestamp with time zone NOT NULL
);


ALTER TABLE delivery_history OWNER TO cakman;

--
-- Name: delivery_history_id_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE delivery_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE delivery_history_id_seq OWNER TO cakman;

--
-- Name: delivery_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE delivery_history_id_seq OWNED BY delivery_history.id;


--
-- Name: delivery_id_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE delivery_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE delivery_id_seq OWNER TO cakman;

--
-- Name: delivery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE delivery_id_seq OWNED BY delivery.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE orders (
    id bigint NOT NULL,
    resi character varying(32) NOT NULL,
    consignee character varying(256) NOT NULL,
    phone character varying(20) NOT NULL,
    status character varying(20) DEFAULT 'WAREHOUSE'::character varying NOT NULL,
    address character varying(2044) NOT NULL,
    country_code bigint DEFAULT 0 NOT NULL,
    province_code bigint DEFAULT 0 NOT NULL,
    city_code bigint DEFAULT 0 NOT NULL,
    create_by bigint DEFAULT 0 NOT NULL,
    create_time timestamp with time zone DEFAULT '2017-12-10 19:41:15.125924+07'::timestamp with time zone NOT NULL,
    update_by character varying(2044) DEFAULT '0'::character varying NOT NULL,
    update_time timestamp with time zone DEFAULT '2017-12-10 19:41:15.125924+07'::timestamp with time zone NOT NULL
);


ALTER TABLE orders OWNER TO cakman;

--
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_id_seq OWNER TO cakman;

--
-- Name: order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE order_id_seq OWNED BY orders.id;


--
-- Name: province; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE province (
    province_code integer NOT NULL,
    province_name character varying(100) NOT NULL,
    country_code integer NOT NULL
);


ALTER TABLE province OWNER TO cakman;

--
-- Name: province_province_code_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE province_province_code_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE province_province_code_seq OWNER TO cakman;

--
-- Name: province_province_code_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE province_province_code_seq OWNED BY province.province_code;


--
-- Name: session; Type: TABLE; Schema: public; Owner: cakman
--

CREATE TABLE session (
    id bigint NOT NULL,
    courier_id bigint NOT NULL,
    token character varying(32) NOT NULL,
    expire_time timestamp with time zone,
    create_time timestamp with time zone
);


ALTER TABLE session OWNER TO cakman;

--
-- Name: session_id_seq; Type: SEQUENCE; Schema: public; Owner: cakman
--

CREATE SEQUENCE session_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE session_id_seq OWNER TO cakman;

--
-- Name: session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cakman
--

ALTER SEQUENCE session_id_seq OWNED BY session.id;


--
-- Name: city city_code; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY city ALTER COLUMN city_code SET DEFAULT nextval('city_city_code_seq'::regclass);


--
-- Name: country country_code; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY country ALTER COLUMN country_code SET DEFAULT nextval('country_country_code_seq'::regclass);


--
-- Name: courier id; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY courier ALTER COLUMN id SET DEFAULT nextval('courier_id_seq'::regclass);


--
-- Name: delivery id; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY delivery ALTER COLUMN id SET DEFAULT nextval('delivery_id_seq'::regclass);


--
-- Name: delivery_history id; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY delivery_history ALTER COLUMN id SET DEFAULT nextval('delivery_history_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY orders ALTER COLUMN id SET DEFAULT nextval('order_id_seq'::regclass);


--
-- Name: province province_code; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY province ALTER COLUMN province_code SET DEFAULT nextval('province_province_code_seq'::regclass);


--
-- Name: session id; Type: DEFAULT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY session ALTER COLUMN id SET DEFAULT nextval('session_id_seq'::regclass);


--
-- Data for Name: city; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO city VALUES (1, 'Jakarta Selatan', 1);
INSERT INTO city VALUES (2, 'Jakarta Utara', 1);
INSERT INTO city VALUES (3, 'Jakarta Timur', 1);
INSERT INTO city VALUES (4, 'Jakarta Barat', 1);
INSERT INTO city VALUES (5, 'Bandung', 2);
INSERT INTO city VALUES (6, 'Bekasi', 2);
INSERT INTO city VALUES (7, 'Bogor', 2);


--
-- Name: city_city_code_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('city_city_code_seq', 7, true);


--
-- Data for Name: country; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO country VALUES (2, 'Indonesia');
INSERT INTO country VALUES (3, 'Singapore');
INSERT INTO country VALUES (4, 'Malaysia');


--
-- Name: country_country_code_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('country_country_code_seq', 4, true);


--
-- Data for Name: courier; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO courier VALUES (1, 'admin', 'password', 'admin', '085434334343', 'admin@cakman.com', 1, '2017-12-10 15:14:51.460734+07', 1, '2017-12-10 17:33:45.752417+07', 46.3442999999999969, 123.123000000000005, '2017-12-10 17:35:00.579841+07');


--
-- Name: courier_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('courier_id_seq', 1, true);


--
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO delivery VALUES (6, 1, 1, 'CANCEL', 'kecelakaan', 1, '2017-12-10 23:04:31.904981+07', 1, '2017-12-10 23:15:04.730609+07');
INSERT INTO delivery VALUES (7, 1, 1, 'FINISH', '', 1, '2017-12-10 23:15:49.976153+07', 1, '2017-12-10 23:20:54.609666+07');


--
-- Data for Name: delivery_history; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO delivery_history VALUES (2, 6, 'PICKUP', '', '2017-12-10 23:04:31.905895+07');
INSERT INTO delivery_history VALUES (3, 6, 'CANCEL', 'kecelakaan', '2017-12-10 23:15:04.739135+07');
INSERT INTO delivery_history VALUES (4, 7, 'PICKUP', '', '2017-12-10 23:15:49.976817+07');
INSERT INTO delivery_history VALUES (6, 7, 'FINISH', '', '2017-12-10 23:17:22.312006+07');


--
-- Name: delivery_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('delivery_history_id_seq', 7, true);


--
-- Name: delivery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('delivery_id_seq', 7, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('order_id_seq', 3, true);


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO orders VALUES (2, 'CGK00000002', 'luthfi', '0811111111', 'WAREHOUSE', 'mangga dua', 2, 1, 1, 1, '2017-12-10 19:41:15.125924+07', '1', '2017-12-10 19:41:15.125924+07');
INSERT INTO orders VALUES (3, 'CGK00000003', 'fami', '08333333', 'WAREHOUSE', 'warung buncit', 2, 1, 1, 1, '2017-12-10 19:41:15.125924+07', '1', '2017-12-10 19:41:15.125924+07');
INSERT INTO orders VALUES (1, 'CGK00000001', 'fauzan', '081231312', 'FINISH', 'kuningan', 2, 1, 1, 1, '2017-12-10 19:41:15.125924+07', '1', '2017-12-10 23:20:54.618899+07');


--
-- Data for Name: province; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO province VALUES (1, 'DKI Jakarta', 2);
INSERT INTO province VALUES (2, 'Jawa Barat', 2);
INSERT INTO province VALUES (3, 'Jawa Tengah', 2);
INSERT INTO province VALUES (4, 'Jawa Timur', 2);


--
-- Name: province_province_code_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('province_province_code_seq', 4, true);


--
-- Data for Name: session; Type: TABLE DATA; Schema: public; Owner: cakman
--

INSERT INTO session VALUES (1, 1, 'f13b411acfe14af4a1694400292c990c', '2017-12-09 23:29:35.880249+07', '2017-12-09 23:29:35.880249+07');
INSERT INTO session VALUES (2, 1, '58c747e73f294f2081469d3046807525', '2017-12-10 00:50:32.374695+07', '2017-12-10 00:20:32.374695+07');
INSERT INTO session VALUES (3, 1, '05c8e638c37b4d2f9bd46783012258d9', '2017-12-10 00:52:40.133432+07', '2017-12-10 00:22:40.133432+07');
INSERT INTO session VALUES (4, 1, '09cbb7baf0aa4f2fa11453a26741b58a', '2017-12-10 00:52:55.683581+07', '2017-12-10 00:22:55.683581+07');
INSERT INTO session VALUES (5, 1, '71ccbe23054c487d84c2e2ab563ff0b3', '2017-12-10 00:57:52.077851+07', '2017-12-10 00:27:52.077851+07');
INSERT INTO session VALUES (6, 1, 'cb5ae8b82c3f4282bec37dfaa2fe446c', '2017-12-10 00:58:26.878676+07', '2017-12-10 00:28:26.878676+07');
INSERT INTO session VALUES (7, 1, '0bbbf9dc8fc541728e55ef26b4984bf9', '2017-12-10 00:58:45.726378+07', '2017-12-10 00:28:45.726378+07');
INSERT INTO session VALUES (8, 1, '07650b64b6b248058ed78d2404610d32', '2017-12-10 00:58:59.904871+07', '2017-12-10 00:28:59.904871+07');
INSERT INTO session VALUES (9, 1, '4c0324a64de24ae68c736b8a3a5773a3', '2017-12-10 00:59:19.23713+07', '2017-12-10 00:29:19.23713+07');
INSERT INTO session VALUES (10, 1, 'c69c0f72479142ef9b2c4d28997bd0c1', '2017-12-10 00:32:06.019616+07', '2017-12-10 00:31:06.019616+07');
INSERT INTO session VALUES (12, 1, 'd2dee6b1d8af4a2b8c511436c7cb2b5c', '2017-12-10 16:02:25.915316+07', '2017-12-10 15:32:25.915316+07');
INSERT INTO session VALUES (13, 1, 'e62756239ec04e59a531024debf336a1', '2017-12-10 16:06:29.707827+07', '2017-12-10 15:36:29.707827+07');
INSERT INTO session VALUES (14, 1, '34a524c3eb004e06be213602ea10df76', '2017-12-10 16:06:34.0873+07', '2017-12-10 15:36:34.0873+07');
INSERT INTO session VALUES (15, 1, 'ce316ee6b95d46619c8cb45d205c01a8', '2017-12-10 17:56:27.962147+07', '2017-12-10 17:26:27.962147+07');
INSERT INTO session VALUES (16, 1, 'f40489e351b2433296146917b23d9599', '2017-12-10 17:59:48.440722+07', '2017-12-10 17:29:48.440722+07');
INSERT INTO session VALUES (17, 1, 'e43f43928ab341eb8aa02cee48a7bfa7', '2017-12-10 23:15:30.592981+07', '2017-12-10 22:45:30.592981+07');
INSERT INTO session VALUES (18, 1, '2664e90c440a40a68a24eb9025f961a5', '2017-12-10 23:45:44.646358+07', '2017-12-10 23:15:44.646358+07');


--
-- Name: session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cakman
--

SELECT pg_catalog.setval('session_id_seq', 18, true);


--
-- Name: city city_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY city
    ADD CONSTRAINT city_pkey PRIMARY KEY (city_code);


--
-- Name: country country_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY country
    ADD CONSTRAINT country_pkey PRIMARY KEY (country_code);


--
-- Name: courier courier_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY courier
    ADD CONSTRAINT courier_pkey PRIMARY KEY (id);


--
-- Name: delivery_history delivery_history_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY delivery_history
    ADD CONSTRAINT delivery_history_pkey PRIMARY KEY (id);


--
-- Name: delivery delivery_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (id);


--
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- Name: province province_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY province
    ADD CONSTRAINT province_pkey PRIMARY KEY (province_code);


--
-- Name: session session_pkey; Type: CONSTRAINT; Schema: public; Owner: cakman
--

ALTER TABLE ONLY session
    ADD CONSTRAINT session_pkey PRIMARY KEY (id);


--
-- Name: index_country_code; Type: INDEX; Schema: public; Owner: cakman
--

CREATE INDEX index_country_code ON province USING btree (country_code);


--
-- Name: index_province_code; Type: INDEX; Schema: public; Owner: cakman
--

CREATE INDEX index_province_code ON city USING btree (province_code);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

