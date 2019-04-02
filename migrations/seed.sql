use blog;

insert into users (display_name, avatar, email, username) values ("Michael Helvey", null, "michael.helvey1@gmail.com", "helvetici");

insert into posts (title, body, author_id) values ("Cool New Post", "first post body", 1);
insert into posts (title, body, author_id) values ("Second Post", "second post body", 1);

insert into tags (title) values ("first_tag");
insert into posts_tags (post_id, tag_id) values (1, 1);