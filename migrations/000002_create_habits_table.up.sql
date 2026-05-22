CREATE TABLE IF NOT EXISTS habits(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    title TEXT NOT NULL

)