
CREATE TABLE IF NOT EXISTS "users"(
    "user_id" INTEGER GENERATED BY DEFAULT AS IDENTITY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "dob" DATE NOT NULL,
    "sex" VARCHAR(255) NULL,
    "address" VARCHAR(255) NULL,
    "phone" VARCHAR(255) NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
CREATE TABLE IF NOT EXISTS "tasks"(
    "task_id" INTEGER GENERATED BY DEFAULT AS IDENTITY,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "user_id" INTEGER NOT NULL
);
ALTER TABLE
    "tasks" ADD PRIMARY KEY("task_id");
ALTER TABLE
    "tasks" ADD CONSTRAINT "task_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
