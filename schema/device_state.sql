CREATE TABLE devices_state (
    id SERIAL,
    uid varchar(20),
    data TEXT,
    created_at timestamp with time zone
)