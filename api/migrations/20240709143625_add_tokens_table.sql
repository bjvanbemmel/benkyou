-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
  id uuid DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  value VARCHAR(255) UNIQUE NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,

  PRIMARY KEY (id),

  CONSTRAINT fk_token_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
-- +goose StatementEnd
