CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);
CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int REFERENCES users (id) ON DELETE CASCADE not null,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE not null
);

CREATE TABLE todo_items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    list_id int not null,
    done boolean not null default false
);
CREATE TABLE lists_items
(
    id serial,
    item_id int REFERENCES todo_items (id) ON DELETE CASCADE not null,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE not null
);