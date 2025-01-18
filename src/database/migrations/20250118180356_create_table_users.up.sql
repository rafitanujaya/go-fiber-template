CREATE TABLE IF NOT EXISTS users (
	id varchar(255) NOT NULL DEFAULT gen_random_uuid(),
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT user_email_key UNIQUE (email),
	CONSTRAINT user_pkey1 PRIMARY KEY (id)
);