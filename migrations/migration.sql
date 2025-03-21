CREATE TABLE IF NOT EXISTS entity (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    checksum VARCHAR(1080) NOT NULL,
    prompt VARCHAR(2000) NOT NULL,
    created DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS discussions (
    id SERIAL PRIMARY KEY,
    sender_user_id INT NOT NULL,
    receiver_user_id INT NOT NULL,
    message TEXT NOT NULL,
    is_new_message BOOLEAN NOT NULL DEFAULT TRUE,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_user_id) REFERENCES entity(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_user_id) REFERENCES entity(id) ON DELETE CASCADE
);
