create table promises
(
    id            uuid   not null
        constraint promises_pk
            primary key,
    user_id       bigint not null,
    description   text,
    status        text,
    date_deadline timestamp with time zone,
    crated_at     timestamp with time zone,
    updated_at    timestamp with time zone
);