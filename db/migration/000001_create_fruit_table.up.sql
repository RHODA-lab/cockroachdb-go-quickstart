CREATE TABLE IF NOT EXISTS Fruit(id varchar(100) PRIMARY KEY, name varchar(100), description varchar(200));

INSERT INTO Fruit (id, name, description) VALUES ('d37f4fae-b572-47b3-93e0-17daf798f9d5', 'Banana', 'Good for health') ON CONFLICT (ID) DO NOTHING;
