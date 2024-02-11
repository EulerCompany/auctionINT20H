USE auction;
INSERT INTO user (name, hashed_password) VALUES ('root', '$2a$10$k23wiFxJ7R18eISwfuxvK.g0rc/xQB7irU0XzC3EvxdOy6Nj/2SZq');
INSERT INTO user (name, hashed_password) VALUES ('max', '$2a$10$boJxBjdj1G/0VFbFN1HlCOuvIKXqdludClJ2TDUNacH7RW6sfbHGS');
INSERT INTO auction (author_id, title, description, start_price, status) VALUES (1, 'test1', 'test1 disc', 100,'active');