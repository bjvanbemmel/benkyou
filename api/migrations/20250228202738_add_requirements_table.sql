-- +goose Up
-- +goose StatementBegin
CREATE TABLE requirements (
  id uuid DEFAULT gen_random_uuid(),
  external_id SERIAL NOT NULL,
  user_id uuid,
  sprint_id uuid,
  feature_id uuid,
  state INT NOT NULL DEFAULT 0,
  title VARCHAR(255) NOT NULL,
  estimate INT,
  description TEXT,
  position INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,

  PRIMARY KEY (id),

  CONSTRAINT fk_requirement_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT fk_requirement_sprint FOREIGN KEY (sprint_id) REFERENCES sprints (id) ON DELETE CASCADE,
  CONSTRAINT fk_requirement_feature FOREIGN KEY (feature_id) REFERENCES features (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE requirements;
-- +goose StatementEnd
