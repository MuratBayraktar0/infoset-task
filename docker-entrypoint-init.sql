CREATE DATABASE IF NOT EXISTS restaurant;

USE restaurant;

CREATE TABLE IF NOT EXISTS restaurant_branches (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100),
    latitude DOUBLE,
    longitude DOUBLE
);

INSERT INTO restaurant_branches (id, name, latitude, longitude)
VALUES
   ('1', 'Restaurant 1', 80.0963, 90.0689),
    ('2', 'Restaurant 2', 80.0346, 89.9865),
    ('3', 'Restaurant 3', 80.0132, 90.0573),
    ('4', 'Restaurant 4', 80.1069, 90.0814),
    ('5', 'Restaurant 5', 80.0448, 90.1462),
    ('6', 'Restaurant 6', 79.9997, 89.9912),
    ('7', 'Restaurant 7', 80.0235, 90.1623),
    ('8', 'Restaurant 8', 80.0676, 89.9314),
    ('9', 'Restaurant 9', 79.9679, 90.1396),
    ('10', 'Restaurant 10', 80.0059, 90.0565),
    ('11', 'Restaurant 11', 80.1001, 90.1248),
    ('12', 'Restaurant 12', 80.0593, 89.9671),
    ('13', 'Restaurant 13', 79.9835, 89.9614),
    ('14', 'Restaurant 14', 79.9684, 89.9887),
    ('15', 'Restaurant 15', 80.0842, 90.0246),
    ('16', 'Restaurant 16', 80.0537, 90.0598),
    ('17', 'Restaurant 17', 80.0665, 90.0524),
    ('18', 'Restaurant 18', 80.0231, 90.0099),
    ('19', 'Restaurant 19', 80.0946, 89.9491),
    ('20', 'Restaurant 20', 80.0519, 90.1318);