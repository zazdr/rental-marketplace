-- +goose Up
CREATE TABLE "user" (
	id              SERIAL NOT NULL PRIMARY KEY,
	mail            TEXT   NOT NULL UNIQUE,
	hashed_password TEXT   NOT NULL
);

-- +goose Down
DROP TABLE "user";
