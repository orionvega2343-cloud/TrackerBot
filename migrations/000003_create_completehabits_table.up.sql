CREATE TABLE IF NOT EXISTS complete_habits
(
    id SERIAL PRIMARY KEY,
    habit_id    INT REFERENCES habits(id),
    date DATE NOT NULL,
    is_complete BOOLEAN NOT NULL
)