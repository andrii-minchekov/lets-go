CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE IF NOT EXISTS users
(
    id       bigint NOT NULL DEFAULT nextval('users_id_seq')
        constraint users_pk
            primary key,
    name     text   not null,
    email    text   not null,
    password text   not null,
    created  date   not null
);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_uindex
    on users (email);

CREATE SEQUENCE IF NOT EXISTS snippets_id_seq;
CREATE TABLE IF NOT EXISTS snippets
(
    id      bigint NOT NULL DEFAULT nextval('snippets_id_seq')
        constraint snippets_pk
            primary key,
    title   text   not null,
    content text   not null,
    created date   not null,
    expires date,
    user_id bigint references users (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS snippets_title_uindex
    on snippets (title);


