-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    username varchar(100) NOT NULL,
    password varchar(200) NOT NULL
);

INSERT INTO users
    (username, password) 
    VALUES ('user', 'ee11cbb19052e40b07aac0ca060c23ee');

CREATE TABLE IF NOT EXISTS product (
    id SERIAL NOT NULL ,
    name varchar(100) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO product
    (id, name) 
    VALUES (1, 'Books');
INSERT INTO product(
    id, name) 
    VALUES (2, 'Pen');
INSERT INTO product(
    id, name) 
    VALUES (3, 'Pencil');

CREATE TABLE IF NOT EXISTS distributor (
    id SERIAL NOT NULL ,
    name character varying(100) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO distributor(
    id, name) 
    VALUES (1, 'Classmates');
INSERT INTO distributor(
    id, name) 
    VALUES (2, 'Camlin');
INSERT INTO distributor(
    id, name) 
    VALUES (3, 'Nataraj');

CREATE TABLE IF NOT EXISTS distribution (
    id SERIAL NOT NULL ,
    product_id INT NOT NULL,
    distributor_id INT NOT NULL,
    place varchar(100) NOT NULL,
    quantity_sold INT NOT NULL,
    year INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(distributor_id) REFERENCES distributor(id),
    FOREIGN KEY(product_id) REFERENCES product(id)
);

INSERT INTO distribution
    (product_id, distributor_id, place, quantity_sold, year)
	VALUES (1, 1, 'Ernakulam', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (2, 1, 'Ernakulam', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (3, 1, 'Kottayam', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (1, 2, 'Ernakulam', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (2, 2, 'Ernakulam', 200, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (3, 2, 'Kottayam', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (1, 3, 'Thrissur', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (2, 3, 'Thrissur', 100, 2019);
INSERT INTO distribution(
    product_id, distributor_id, place, quantity_sold, year)
	VALUES (3, 3, 'Thrissur', 500, 2019);

-- +goose Down
DROP TABLE distribution;
DROP TABLE product;
DROP TABLE distributor;
DROP TABLE users;