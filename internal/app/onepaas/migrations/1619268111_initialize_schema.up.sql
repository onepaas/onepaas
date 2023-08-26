CREATE TABLE users (
       id varchar(26) PRIMARY KEY,
       name varchar(255) NULL,
       email varchar(255) NOT NULL,
       meta jsonb DEFAULT '{}',
       created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
       modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX users_email_key ON users (email);
CREATE INDEX users_created_at_idx ON users USING brin(created_at);
CREATE INDEX users_modified_at_idx ON users USING brin(modified_at);

-- CREATE TABLE identities (
--         id varchar(26) PRIMARY KEY,
--         user_id varchar(26) NOT NULL,
--         subject text NOT NULL,
--         provider text NOT NULL,
--         meta jsonb DEFAULT '{}',
--         created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
--         modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE UNIQUE INDEX identities_subject_provider_key ON identities (subject, provider);
-- CREATE INDEX identities_created_at_idx ON identities USING brin(created_at);
-- CREATE INDEX identities_modified_at_idx ON identities USING brin(modified_at);

CREATE TABLE projects (
        id varchar(26) PRIMARY KEY,
        name varchar(253) NOT NULL,
        slug varchar(253) NOT NULL,
        description text NOT NULL,
        meta jsonb DEFAULT '{}',
        created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
        modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX projects_created_at_idx ON projects USING brin(created_at);
CREATE INDEX projects_modified_at_idx ON projects USING brin(modified_at);

CREATE TABLE applications (
    id varchar(26) PRIMARY KEY,
    name varchar(253) NULL,
    repository_url text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX applications_created_at_idx ON applications USING brin(created_at);
CREATE INDEX applications_modified_at_idx ON applications USING brin(modified_at);

CREATE TABLE registries (
    id varchar(26) PRIMARY KEY,
    url text NOT NULL,
    username varchar(253) NULL,
    secret text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX registries_created_at_idx ON registries USING brin(created_at);
CREATE INDEX registries_modified_at_idx ON registries USING brin(modified_at);

CREATE TABLE infrastructures (
    id varchar(26) PRIMARY KEY,
    type varchar(20) NOT NULL,
    properties json NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX infrastructures_created_at_idx ON registries USING brin(created_at);
CREATE INDEX infrastructures_modified_at_idx ON registries USING brin(modified_at);
