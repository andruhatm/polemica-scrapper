CREATE TABLE "player"
(
    "player_id"   integer      NOT NULL,
    "username"    varchar(255) NOT NULL,
    "games_count" integer      NOT NULL,
    "score"       float      NOT NULL,
    "avg_score"   float      NOT NULL,
    CONSTRAINT "player_pk" PRIMARY KEY ("player_id")
) WITH (
      OIDS = FALSE
);

CREATE TABLE "operation_log"
(
    "operation_id" VARCHAR(255) NOT NULL,
    "time"         VARCHAR(255) NOT NULL,
    "meta"         VARCHAR(255) NOT NULL,
    CONSTRAINT "operation_log_pk" PRIMARY KEY ("operation_id")
) WITH (
      OIDS = FALSE
);






