ALTER TABLE item
    ADD COLUMN chat_id INTEGER,
    ADD CONSTRAINT fk_item_user_id
        FOREIGN KEY (chat_id)
            REFERENCES bot_user(chat_id)
            ON DELETE CASCADE;

ALTER TABLE item
    DROP CONSTRAINT product_data_name_brand_unique;

ALTER TABLE item
    ADD CONSTRAINT product_data_name_brand_user_id_unique
        UNIQUE (name, brand, chat_id);