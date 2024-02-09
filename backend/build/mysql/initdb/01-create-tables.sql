CREATE DATABASE IF NOT EXISTS auction;
USE auction;
-- TODO: change to singular
CREATE TABLE IF NOT EXISTS user (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN NOT NULL DEFAULT TRUE
);
ALTER TABLE user ADD CONSTRAINT users_uc_name UNIQUE (name);


CREATE TABLE IF NOT EXISTS auction (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    author_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1024) NOT NULL,
    start_price BIGINT UNSIGNED NOT NULL,
    current_price BIGINT UNSIGNED NOT NULL DEFAULT (`start_price`),
    status VARCHAR(50),
    start_date DATETIME,
    end_date DATETIME,
    FOREIGN KEY (author_id)
        REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS auction_user (
    auction_id INTEGER NOT NULL ,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (auction_id)
        REFERENCES auction(id)
);

CREATE TABLE IF NOT EXISTS auction_bet (
    auction_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    bet DECIMAL(13, 4) NOT NULL,
    FOREIGN KEY (auction_id)
        REFERENCES auction(id),
    FOREIGN KEY (user_id)
        REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS auction_image (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    auction_id INTEGER NOT NULL,
    img LONGBLOB NOT NULL,
    FOREIGN KEY (auction_id)
        REFERENCES auction(id)
        ON DELETE CASCADE
);
