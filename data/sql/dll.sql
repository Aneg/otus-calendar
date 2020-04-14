CREATE TABLE events (
    id             serial primary key,
    user_id        integer NOT NULL,
    description    varchar(1024) NOT NULL,
    date_time_from timestamp NOT NULL,
    date_time_to   timestamp NOT NULL
);

create index events_user_id_idx
    on events (user_id);

create index events_date_time_from_idx
    on events (date_time_from);

create index events_date_time_to_idx
    on events (date_time_to);