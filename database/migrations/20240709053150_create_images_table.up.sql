CREATE TABLE images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    image VARCHAR(255),
    title VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
