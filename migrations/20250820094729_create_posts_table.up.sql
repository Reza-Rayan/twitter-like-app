CREATE TABLE IF NOT EXISTS roles (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     name TEXT NOT NULL UNIQUE
);

-- Insert default roles
INSERT OR IGNORE INTO roles (id, name) VALUES
    (1, 'user'),
    (2, 'admin');