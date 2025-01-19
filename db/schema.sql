CREATE TABLE app_user(
  id INT PRIMARY KEY,
  username VARCHAR NOT NULL,
  email VARCHAR NOT NULL
);

CREATE TABLE room(
  id INT PRIMARY KEY,
  name VARCHAR NOT NULL
);

CREATE TABLE user_room(
  id INT PRIMARY KEY,
  user_id INT NOT NULL REFERENCES app_user(id) ON DELETE CASCADE,
  room_id INT NOT NULL REFERENCES room(id) ON DELETE CASCADE,
  unique(user_id, room_id)
);

CREATE TABLE message(
  id INT PRIMARY KEY,
  CREATEd_at timestamp NOT NULL default current_timestamp,
  text text NOT NULL,
  room_id INT NOT NULL REFERENCES room(id) ON DELETE CASCADE,
  sender_user_id INT NOT NULL REFERENCES app_user(id) ON DELETE CASCADE
);

CREATE index idx_message_room_id_created_at on message(room_id, created_at desc);
