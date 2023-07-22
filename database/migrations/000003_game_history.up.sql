CREATE TABLE "game_history"
(
    "game_id"    INTEGER      NOT NULL,
    "players"    json         NOT NULL,
    "first_kill" INTEGER,
    "win"        bool         NOT NULL,
    CONSTRAINT "game_history_pk" PRIMARY KEY ("game_id")
) WITH (
      OIDS = FALSE
    );



