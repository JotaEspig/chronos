CREATE TABLE IF NOT EXISTS "user" (
    "id" INTEGER NOT NULL,
    "username" TEXT NOT NULL,

    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "employee" (
    "id" INTEGER NOT NULL,
    "type" TEXT NOT NULL,
    "user_id" INTEGER NOT NULL,

    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "user"("id")
);

CREATE TABLE IF NOT EXISTS "time" (
    "id" INTEGER NOT NULL,
    "start" TEXT NOT NULL,
    "end" TEXT NOT NULL,
    "repeat" TEXT NOT NULL,
    "employee_id" INTEGER NOT NULL,

    PRIMARY KEY("id"),
    FOREIGN KEY("employee_id") REFERENCES "employee"("id")
);

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

CREATE UNIQUE INDEX IF NOT EXISTS "username_index" ON "user" ("username");
