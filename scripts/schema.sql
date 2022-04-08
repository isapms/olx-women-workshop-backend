CREATE DATABASE mini_olx;

USE mini_olx;

CREATE TABLE advert (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description text NOT NULL,
    price FLOAT not null,
    image_path VARCHAR(255) not null
);