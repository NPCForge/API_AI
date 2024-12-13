CREATE TABLE IF NOT EXISTS entity (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    token VARCHAR(1080) NOT NULL,
    prompt VARCHAR(2000) NOT NULL,
    created DATE DEFAULT CURRENT_DATE
);
