-- Characters Table
CREATE TABLE IF NOT EXISTS Characters (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    modified TIMESTAMPTZ,
    thumbnail_path VARCHAR(255),
    thumbnail_extension VARCHAR(10),
    resourceURI VARCHAR(255)
);

-- Comics Table
CREATE TABLE IF NOT EXISTS Comics (
    comic_id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE,
    collectionURI VARCHAR(255)
);

-- Series Table
CREATE TABLE IF NOT EXISTS Series (
    series_id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE,
    collectionURI VARCHAR(255)
);

-- Stories Table
CREATE TABLE IF NOT EXISTS Stories (
    story_id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE,
    type VARCHAR(50),
    collectionURI VARCHAR(255)
);

-- Events Table
CREATE TABLE IF NOT EXISTS Events (
    event_id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE,
    collectionURI VARCHAR(255)
);

-- URLs Table
CREATE TABLE IF NOT EXISTS URLs (
    url_id SERIAL PRIMARY KEY,
    character_id INT REFERENCES Characters(id),
    type VARCHAR(50),
    url VARCHAR(255)
);

-- Join table for Characters and Comics (many-to-many relationship)
CREATE TABLE IF NOT EXISTS Character_Comics (
    character_id INT REFERENCES Characters(id),
    comic_id INT REFERENCES Comics(comic_id),
    PRIMARY KEY (character_id, comic_id)
);

-- Join table for Characters and Series (many-to-many relationship)
CREATE TABLE IF NOT EXISTS Character_Series (
    character_id INT REFERENCES Characters(id),
    series_id INT REFERENCES Series(series_id),
    PRIMARY KEY (character_id, series_id)
);

-- Join table for Characters and Stories (many-to-many relationship)
CREATE TABLE IF NOT EXISTS Character_Stories (
    character_id INT REFERENCES Characters(id),
    story_id INT REFERENCES Stories(story_id),
    PRIMARY KEY (character_id, story_id)
);

-- Join table for Characters and Events (many-to-many relationship)
CREATE TABLE IF NOT EXISTS Character_Events (
    character_id INT REFERENCES Characters(id),
    event_id INT REFERENCES Events(event_id),
    PRIMARY KEY (character_id, event_id)
);
