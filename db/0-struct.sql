DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
    "id" bigserial,
    "nickname" varchar(255) NOT NULL UNIQUE,
    "email" varchar(100) NOT NULL UNIQUE,
    "password" varchar(100) NOT NULL,
    "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
DROP TABLE IF EXISTS "posts";
CREATE TABLE "posts" (
    "id" bigserial,
    "title" varchar(255) NOT NULL UNIQUE,
    "content" varchar(255) NOT NULL,
    "author_id" bigint NOT NULL,
    "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);
ALTER TABLE "posts"
ADD CONSTRAINT posts_author_id_users_id_foreign FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE cascade ON UPDATE cascade;