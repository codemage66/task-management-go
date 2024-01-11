
CREATE TABLE IF NOT EXISTS "tasks" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "priority" bigint NOT NULL,
  "duedate" timestamp NOT NULL DEFAULT (now())
);

