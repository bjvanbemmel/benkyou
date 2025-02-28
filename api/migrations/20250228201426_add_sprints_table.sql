-- +goose Up
-- +goose StatementBegin
CREATE TABLE sprints (
  id uuid DEFAULT gen_random_uuid(),
  title VARCHAR(80) NOT NULL,
  start_date TIMESTAMP NOT NULL DEFAULT now(),
  end_date TIMESTAMP NOT NULL DEFAULT now(),
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,

  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sprints;
-- +goose StatementEnd
