package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/rs/zerolog/log"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Info().Msg("Creating table users ...")
		_, err := db.Exec(`
			;CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

			CREATE TABLE users (
				id varchar(50) PRIMARY KEY,
				username varchar(255) NOT NULL,
				password text NULL,
				email varchar(255) NOT NULL,
				name varchar(255) NULL,
				meta jsonb DEFAULT '{}',
				created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
				modified_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
			);

			CREATE UNIQUE INDEX users_username_key ON users (username);
			CREATE UNIQUE INDEX users_email_key ON users (email);
			CREATE INDEX users_created_at_idx ON users USING brin(created_at);
			CREATE INDEX users_modified_at_idx ON users USING brin(modified_at);
		`)

		return err
	}, func(db migrations.DB) error {
		log.Info().Msg("Dropping table users ...")
		_, err := db.Exec(`
			DROP TABLE users;
		`)

		return err
	})
}
