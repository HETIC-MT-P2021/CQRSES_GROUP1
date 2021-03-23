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

DROP TABLE IF EXISTS "users";
CREATE TABLE "post_next_id" (
    "id" bigserial,
    PRIMARY KEY ("id")
);
