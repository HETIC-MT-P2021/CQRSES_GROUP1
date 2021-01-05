CREATE TABLE profile (
    profile_id bigserial PRIMARY KEY,
    label VARCHAR (128) NOT NULL
);

CREATE TABLE user_account (
    user_id bigserial PRIMARY KEY,
    email VARCHAR (128) NOT NULL,
    password bytea NOT NULL,

    profile_id integer references profile(profile_id)
);
