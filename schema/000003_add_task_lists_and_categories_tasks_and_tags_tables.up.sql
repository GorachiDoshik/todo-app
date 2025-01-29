CREATE TABLE tasks(
	id serial not null unique,
	user_id int references users (id) on delete cascade not null,
	title varchar(255) not null,
	description TEXT
);

CREATE TABLE categories(
	id serial not null unique,
	name varchar(255) not null
);

CREATE TABLE tags(
	id serial not null unique,
	name varchar(255) not null
);

CREATE TABLE task_categories(
	id serial not null unique,
	task_id int references tasks (id) on delete cascade not null,
	category_id int references categories (id) on delete cascade not null
);

CREATE TABLE task_tags(
	id serial not null unique,
	task_id int references tasks (id) on delete cascade not null,
	tag_id int references tags (id) on delete cascade not null
);
