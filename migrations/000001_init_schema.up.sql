CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR NOT NULL
);

CREATE TABLE employees
(
    id SERIAL PRIMARY KEY,
    name varchar(75) NOT NULL UNIQUE,
    birthday DATE NOT NULL
);

CREATE TABLE subscriptions (
    user_id INT NOT NULL,
    employee_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id),
   PRIMARY KEY (user_id, employee_id)
);