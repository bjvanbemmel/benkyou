-- +goose Up
-- +goose StatementBegin
CREATE TABLE features (
  id uuid DEFAULT gen_random_uuid(),
  user_id uuid,
  sprint_id uuid,
  state INT NOT NULL DEFAULT 0,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,

  PRIMARY KEY (id),

  CONSTRAINT fk_feature_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT fk_feature_sprint FOREIGN KEY (sprint_id) REFERENCES sprints (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE features;
-- +goose StatementEnd
