use blog;

create table users (
    id int not null AUTO_INCREMENT,
    display_name varchar(255) NOT NULL,
    avatar varchar(255) DEFAULT NULL,
    email varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
    primary key (id)
);

create table posts (
    id int not null AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    body text NOT NULL,
    author_id int NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
    primary key (id),
	foreign key (author_id) references users(id)
);

create table tags (
    id int not null AUTO_INCREMENT,
    title varchar(255) NOT NULL,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
    primary key (id)
);

create table posts_tags (
    post_id int NOT NULL,
    tag_id int NOT NULL,
    primary key (post_id, tag_id),
    foreign key (post_id) references posts(id) on delete cascade,
    foreign key (tag_id) references tags(id) on delete cascade
);

create table categories (
    id int not null AUTO_INCREMENT,
    title varchar(255) NOT NULL,
	created_at timestamp default CURRENT_TIMESTAMP,
	updated_at timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
    primary key (id)
);

create table posts_categories (
	post_id int NOT NULL,
	category_id int NOT NULL,
	primary key (post_id, category_id),
	foreign key (post_id) references posts(id) on delete cascade,
	foreign key (category_id) references tags(id) on delete cascade
);
