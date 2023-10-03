CREATE TABLE IF NOT EXISTS "user" (
    "id" INTEGER NOT NULL,
    "username" TEXT NOT NULL,

    PRIMARY KEY("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "user_username_index" ON "user" ("username");

CREATE TABLE IF NOT EXISTS "employee" (
    "id" INTEGER NOT NULL,
    "type" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,

    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "user"("id")
);

CREATE TABLE IF NOT EXISTS "time" (
    "id" INTEGER NOT NULL,
    "start" TEXT NOT NULL,
    "end" TEXT NOT NULL,
    "repeat" INTEGER NOT NULL,
    "employee_id" INTEGER NOT NULL,

    PRIMARY KEY("id"),
    FOREIGN KEY("employee_id") REFERENCES "employee"("id")
);
CREATE INDEX "time_repeat_idx" ON "time"("repeat");
CREATE INDEX "time_repeat&32_idx" ON "time"("repeat" & 32);
CREATE INDEX "time_repeat&64_idx" ON "time"("repeat" & 64);
CREATE INDEX "time_start_strftime_idx" ON "time" (strftime('%w', "start"));

CREATE TABLE IF NOT EXISTS "scheduling" (
    "id" INTEGER NOT NULL,
    "start" TEXT NOT NULL,
    "end" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,
    "time_id" INTEGER NOT NULL,

    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "user"("id"),
    FOREIGN KEY("time_id") REFERENCES "time"("id")
);
