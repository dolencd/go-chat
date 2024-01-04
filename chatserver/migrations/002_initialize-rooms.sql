create table room(
  id uuid primary key,
  name varchar not null
);

create table user_room(
  id uuid primary key,
  user_id uuid not null,
  room_id uuid not null
);

---- create above / drop below ----

drop table room;
drop table user_room;
