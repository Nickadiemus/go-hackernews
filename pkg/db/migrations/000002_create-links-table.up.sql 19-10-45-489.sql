create table if not exists links (
    id serial not null unique,
    title varchar (255) ,
    address varchar (255) ,
    userId integer references users on delete cascade,
    deleted_at       timestamp with time zone,
    created_at       timestamp with time zone,
    updated_at       timestamp with time zone
)
