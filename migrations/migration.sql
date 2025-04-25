CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    permission INT DEFAULT 0,
    password_hash VARCHAR(1080) NOT NULL,
    created DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS entities (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    checksum VARCHAR(1080) NOT NULL UNIQUE,
    prompt VARCHAR(2000) NOT NULL,
    created DATE DEFAULT CURRENT_DATE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    sender_user_id INT NOT NULL,
    message TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_user_id) REFERENCES entities(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS message_receivers (
    message_id INT NOT NULL,
    receiver_user_id INT NOT NULL,
    is_new_message BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY (message_id, receiver_user_id),
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_user_id) REFERENCES entities(id) ON DELETE CASCADE
);
