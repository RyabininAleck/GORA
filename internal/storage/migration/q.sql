CREATE TABLE IF NOT EXISTS photo (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     image_path TEXT NOT NULL,
                                     preview_path TEXT NOT NULL,
                                     name TEXT NOT NULL,
                                     type TEXT NOT NULL
);

