create table
    users (
        id serial primary key,
        email text UNIQUE default '',
        password text default '',
        name text default ''
    );

create table
    category (
        id serial primary key,
        name text default '',
        created_at text default ''
    );

create table
    news (
        id serial primary key,
        title text default '',
        content text default '',
        category_id integer REFERENCES category (id) on delete cascade on update cascade,
        thumbnail text default '',
        views integer default 0,
        created_at text default '',
        author_id integer references users (id) on delete cascade on update cascade
    );

alter table users
add column role_id integer default 2;

INSERT INTO
    users (name, password, role_id)
VALUES
    (
        'admin@admin.com',
        '$2a$10$DMphGc0NQ1MJZCD6tyNeBOOrpP6REzj/t.iCwr9HCNwbZ4TN7xE8S',
        1
    );

alter table category
add column author_id integer default 0;