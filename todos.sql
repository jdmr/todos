drop table if exists todos;
create table todos (
    id varchar(50) primary key,
    title varchar(255) not null,
    completed boolean default false,
    created_at timestamp not null,
    updated_at timestamp not null default now()
);