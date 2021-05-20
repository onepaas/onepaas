;CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
       id varchar(26) PRIMARY KEY,
       email varchar(255) NOT NULL,
       password text NULL,
       name varchar(255) NULL,
       meta jsonb DEFAULT '{}',
       created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
       modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX users_email_key ON users (email);
CREATE INDEX users_created_at_idx ON users USING brin(created_at);
CREATE INDEX users_modified_at_idx ON users USING brin(modified_at);
