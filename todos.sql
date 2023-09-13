drop table if exists owners;
drop table if exists todos;

create table owners (
    id varchar(50) primary key,
    name varchar(255) not null
);

create table todos (
    id varchar(50) primary key,
    title varchar(255) not null,
    completed boolean default false,
    created_at timestamp not null,
    updated_at timestamp not null default now(),
    owner_id varchar(50) not null references owners(id) on delete cascade
);

insert into owners (id, name) values ('1', 'Alice');
insert into owners (id, name) values ('2', 'Bob');
insert into owners (id, name) values ('3', 'Carol');

insert into todos (id, title, created_at, owner_id) values ('1', 'Buy milk', now(), '1');
insert into todos (id, title, created_at, owner_id) values ('2', 'Buy eggs', now(), '1');
insert into todos (id, title, created_at, owner_id) values ('3', 'Buy bread', now(), '2');
insert into todos (id, title, created_at, owner_id) values ('4', 'Buy butter', now(), '3');

