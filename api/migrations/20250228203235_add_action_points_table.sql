-- +goose Up
-- +goose StatementBegin
CREATE TABLE action_points (
  id uuid DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,

  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE action_points;
-- +goose StatementEnd
