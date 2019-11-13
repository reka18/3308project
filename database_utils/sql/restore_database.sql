--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.15
-- Dumped by pg_dump version 9.6.15

-- Started on 2019-11-13 09:47:03 MST

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
-- TOC entry 481 (class 1247 OID 148107)
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
-- TOC entry 188 (class 1259 OID 148130)
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
-- TOC entry 187 (class 1259 OID 148128)
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
-- Dependencies: 187
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rkmac
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- TOC entry 186 (class 1259 OID 148115)
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
-- TOC entry 185 (class 1259 OID 148113)
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
-- Dependencies: 185
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rkmac
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2280 (class 2604 OID 148133)
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- TOC entry 2279 (class 2604 OID 148118)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2410 (class 0 OID 148130)
-- Dependencies: 188
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: rkmac
--

COPY public.posts (id, userid, content, upvotes, downvotes, deleted) FROM stdin;
\.


--
-- TOC entry 2421 (class 0 OID 0)
-- Dependencies: 187
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rkmac
--

SELECT pg_catalog.setval('public.posts_id_seq', 1, false);


--
-- TOC entry 2408 (class 0 OID 148115)
-- Dependencies: 186
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: rkmac
--

COPY public.users (id, age, firstname, lastname, email, username, public, joindate, active, password, gender) FROM stdin;
1	35	Reagan	Karnes	reagan.karnes@colorado.edu	reka18	t	2019-11-13 09:43:55.646359	t	$2a$04$.iwEEhuoYAfnr8oeAU8X5eKP17DfrP1elZ9OmFT3PSYg7w.F/IiUS	M
2	36	Diyou	Karnes	diyoubolical@icloud.com	dika14	t	2019-11-13 09:45:07.883697	t	$2a$04$e6aMPeV/XrNYag5EcweyOOai9RjuSB482xnqZqghsARQ30UA3I5aS	F
3	5	Wesley	Karnes	weslogica@icloud.com	wes14	t	2019-11-13 09:45:50.043926	t	$2a$04$6q9dJq9UvZyOI/ydrs7bMegNxFP/FD/L3oEh4ydryEYgGnmqc/2rq	M
4	1	Alexandria	Karnes	reaganomica@icloud.com	alex18	t	2019-11-13 09:46:30.636529	t	$2a$04$zstL.wI/rbvC.WHn4mhjpuK2EWZ7G9Sgb9wis4Xi1FBZakhuhjivW	F
\.


--
-- TOC entry 2422 (class 0 OID 0)
-- Dependencies: 185
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rkmac
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- TOC entry 2288 (class 2606 OID 148135)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- TOC entry 2282 (class 2606 OID 148125)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 2284 (class 2606 OID 148123)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2286 (class 2606 OID 148127)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 2289 (class 2606 OID 148136)
-- Name: posts posts_userid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rkmac
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_userid_fkey FOREIGN KEY (userid) REFERENCES public.users(id);


-- Completed on 2019-11-13 09:47:03 MST

--
-- PostgreSQL database dump complete
--

