drop table if exists tags_videos;
drop table if exists tags;
drop table if exists videos;
drop table if exists users;

create table users (
  id serial not null primary key,
  name text not null
);

create table videos (
  id serial not null primary key,
  user_id int not null,
  name text not null,
  deleted bool not null default false,

  foreign key (user_id) references users (id)
);

create table tags (
  id serial not null primary key,
  name text not null,
  deleted bool not null default false
);

create table tags_videos (
  video_id int not null,
  tag_id int not null,

  primary key (video_id, tag_id),
  foreign key (video_id) references videos (id),
  foreign key (tag_id) references tags (id)
);

insert into users (name) values ('john');
insert into users (name) values ('wayne');

insert into videos (user_id, name) values (1, 'batman begins');
insert into videos (user_id, name) values (1, 'dark knight');
insert into videos (user_id, name) values (1, 'dark knight rises');
insert into videos (user_id, name) values (1, 'xmen');
insert into videos (user_id, name) values (1, 'matrix');
insert into videos (user_id, name) values (2, 'starship troopers');
insert into videos (user_id, name) values (2, 'other');

insert into tags (name, deleted) values ('action', false);
insert into tags (name, deleted) values ('drama', false);

insert into tags_videos (video_id, tag_id) values (1, 1);
insert into tags_videos (video_id, tag_id) values (1, 2);
insert into tags_videos (video_id, tag_id) values (2, 2);
