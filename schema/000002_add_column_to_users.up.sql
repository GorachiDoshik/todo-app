ALTER TABLE users ADD COLUMN email VARCHAR(255);

UPDATE users SET email = 'default@example.com';

ALTER TABLE users ALTER COLUMN email SET NOT NULL;
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);