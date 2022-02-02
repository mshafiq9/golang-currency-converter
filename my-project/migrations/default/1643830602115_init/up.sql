SET check_function_bodies = false;
CREATE SCHEMA hasurapg;
CREATE TABLE hasurapg.articles (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    content text NOT NULL
);
CREATE SEQUENCE hasurapg.articles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE hasurapg.articles_id_seq OWNED BY hasurapg.articles.id;
CREATE TABLE hasurapg.transactions (
    id integer NOT NULL,
    date timestamp with time zone NOT NULL,
    from_currency_code text NOT NULL,
    to_currency_code text NOT NULL,
    amount numeric NOT NULL,
    conversion_rate numeric NOT NULL,
    converted_amount numeric NOT NULL
);
CREATE SEQUENCE hasurapg.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE hasurapg.transactions_id_seq OWNED BY hasurapg.transactions.id;
ALTER TABLE ONLY hasurapg.articles ALTER COLUMN id SET DEFAULT nextval('hasurapg.articles_id_seq'::regclass);
ALTER TABLE ONLY hasurapg.transactions ALTER COLUMN id SET DEFAULT nextval('hasurapg.transactions_id_seq'::regclass);
ALTER TABLE ONLY hasurapg.articles
    ADD CONSTRAINT articles_pkey PRIMARY KEY (id);
ALTER TABLE ONLY hasurapg.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);
