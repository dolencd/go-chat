create table app_user(
  id uuid primary key,
  username varchar not null,
  email varchar not null
);

---- create above / drop below ----

drop table app_user;