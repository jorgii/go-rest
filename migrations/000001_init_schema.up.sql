CREATE TABLE "users" (
    "id" bigserial,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "email" text NOT NULL,
    "created_at" timestamptz NOT NULL,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");

CREATE TABLE "books" (
    "id" bigserial,
    "user_id" bigint NOT NULL,
    "title" text NOT NULL,
    "author" text NOT NULL,
    "description" text,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_users_books" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);
