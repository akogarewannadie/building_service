CREATE TABLE IF NOT EXISTS buildings (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    floors INT NOT NULL
    );
