-- +goose Up
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_answers_question_id ON answers(question_id);
CREATE INDEX idx_answers_deleted_at ON answers(deleted_at);
CREATE INDEX idx_answers_user_id ON answers(user_id);

-- +goose Down
DROP TABLE answers;