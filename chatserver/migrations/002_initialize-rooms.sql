create table room(
  id uuid primary key,
  name varchar not null
);

create table user_room(
  id uuid primary key,
  user_id uuid not null references app_user(id) on delete cascade,
  room_id uuid not null references room(id) on delete cascade,
  unique(user_id, room_id)
);

---- create above / drop below ----

drop table user_room;
drop table room;
