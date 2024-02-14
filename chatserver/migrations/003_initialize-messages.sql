create table message(
  id uuid primary key,
  created_at timestamp not null default current_timestamp,
  text text not null,
  room_id uuid not null references room(id) on delete cascade
);

create index idx_message_room_id_created_at on message(room_id, created_at desc);

---- create above / drop below ----

drop index idx_message_room_id_created_at;
drop table message;