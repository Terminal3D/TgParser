CREATE TABLE bot_user (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    username VARCHAR(255),
    subscribed BOOLEAN NOT NULL,
    CONSTRAINT unique_chat_id UNIQUE (chat_id)
);

