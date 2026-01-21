-- +goose Up
-- +goose StatementBegin
CREATE TABLE note (
	id serial primary key,
	title text not null,
	content text not null,
	created_at timestamp not null default now(),
	updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE note;
-- +goose StatementEnd
