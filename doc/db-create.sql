drop table players;
drop table rounds;
drop table throws;
drop table games;

CREATE table players(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  Name TEXT,
  created DATETIME
);

CREATE TABLE IF NOT EXISTS rounds(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  game_uid text,
  player_uid text,
  created DATETIME
);

CREATE TABLE IF NOT EXISTS throws(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  round_uid text,
  score int,
  modifier int,
  x int,
  y int,
  cam1img text,
  cam2img text,
  cam1x int,
  cam2x int,
  created DATETIME,
  edited_count int,
  orig_score int,
  orig_modifier int,
  modified DATETIME
);


CREATE TABLE IF NOT EXISTS games(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  type text,
  subtype text,
  created DATETIME
);
