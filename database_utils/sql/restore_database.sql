--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.15
-- Dumped by pg_dump version 9.6.15

-- Started on 2019-11-12 21:22:19 MST

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

--
-- TOC entry 1 (class 3079 OID 12657)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2418 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 481 (class 1247 OID 148071)
-- Name: gender; Type: TYPE; Schema: public; Owner: rkmac
--

CREATE TYPE public.gender AS ENUM (
    'M',
    'F',
    'O'
);


ALTER TYPE public.gender OWNER TO rkmac;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 185 (class 1259 OID 148077)
-- Name: posts; Type: TABLE; Schema: public; Owner: rkmac
--

CREATE TABLE public.posts (
    id integer NOT NULL,
    userid integer,
    content character varying(240),
    upvotes integer,
    downvotes integer,
    deleted boolean
);


ALTER TABLE public.posts OWNER TO rkmac;

--
-- TOC entry 186 (class 1259 OID 148080)
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: rkmac
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.posts_id_seq OWNER TO rkmac;

--
-- TOC entry 2419 (class 0 OID 0)
-- Dependencies: 186
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rkmac
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- TOC entry 187 (class 1259 OID 148082)
-- Name: users; Type: TABLE; Schema: public; Owner: rkmac
--

CREATE TABLE public.users (
    id integer NOT NULL,
    age integer NOT NULL,
    firstname text NOT NULL,
    lastname text NOT NULL,
    email text NOT NULL,
    username text NOT NULL,
    public boolean NOT NULL,
    joindate timestamp without time zone NOT NULL,
    active boolean NOT NULL,
    password text NOT NULL,
    gender public.gender NOT NULL
);


ALTER TABLE public.users OWNER TO rkmac;

--
-- TOC entry 188 (class 1259 OID 148088)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: rkmac
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO rkmac;

--
-- TOC entry 2420 (class 0 OID 0)
-- Dependencies: 188
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rkmac
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2279 (class 2604 OID 148090)
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- TOC entry 2280 (class 2604 OID 148091)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2407 (class 0 OID 148077)
-- Dependencies: 185
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: rkmac
--

COPY public.posts (id, userid, content, upvotes, downvotes, deleted) FROM stdin;
\.


--
-- TOC entry 2421 (class 0 OID 0)
-- Dependencies: 186
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rkmac
--

SELECT pg_catalog.setval('public.posts_id_seq', 1, false);


--
-- TOC entry 2409 (class 0 OID 148082)
-- Dependencies: 187
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: rkmac
--

COPY public.users (id, age, firstname, lastname, email, username, public, joindate, active, password, gender) FROM stdin;
1	35	Reagan	Karnes	reagan.karnes@colorado.edu	reka18	t	2019-11-12 15:59:43.930777	t	1234	M
2	42	Rigo	Garcia	rigo.garcia@colorado.edu	riga69	t	2019-11-12 16:00:42.337296	t	1234	M
3	26	Graham	Dominick	graham.dominick@colorado.edu	grdo42	t	2019-11-12 16:01:27.827039	t	1234	M
4	1	Alex	Karnes	alex.kar@fam.com	alex01	t	2019-11-12 21:16:23.930953	t	1234	F
5	5	Wes	Karnes	wes.kar@fam.com	wes05	t	2019-11-12 21:16:52.265622	t	1234	M
6	37	Diyou	Karnes	di.kar@fam.com	dika14	t	2019-11-12 21:18:02.566631	t	1234	F
7	10	Beverly	Pudin	dog.mail@fam.com	dog10	t	2019-11-12 21:18:59.927224	t	1234	F
8	3	Luka	Mao	cat.mail@fam.com	cat03	t	2019-11-12 21:19:34.513244	t	1234	M
\.


--
-- TOC entry 2422 (class 0 OID 0)
-- Dependencies: 188
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rkmac
--

SELECT pg_catalog.setval('public.users_id_seq', 8, true);


--
-- TOC entry 2282 (class 2606 OID 148093)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- TOC entry 2284 (class 2606 OID 148095)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 2286 (class 2606 OID 148097)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2288 (class 2606 OID 148099)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 2289 (class 2606 OID 148100)
-- Name: posts posts_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id);


-- Completed on 2019-11-12 21:22:19 MST

--
-- PostgreSQL database dump complete
--

