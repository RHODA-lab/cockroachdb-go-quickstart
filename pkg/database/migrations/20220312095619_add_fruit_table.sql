-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS Fruit(id varchar(100) PRIMARY KEY , name varchar(100), quantity varchar(11) null, description varchar(200));

INSERT INTO Fruit (id, name, quantity, description) VALUES ('d37f4fae-b572-47b3-93e0-17daf798f9d5', 'Banana', '0', 'Good for health') ON CONFLICT (ID) DO NOTHING;
INSERT INTO Fruit (id, name, quantity, description) VALUES ('51661376-0a07-449b-a3bd-9db79cd4ead4', 'Apple', '0', 'Keeps the doctor away') ON CONFLICT (ID) DO NOTHING;
INSERT INTO Fruit (id, name, quantity, description) VALUES ('69f6cd81-59fc-493b-8ebf-1b9f150ecead', 'Blueberry','0', 'Antioxidant Superfood') ON CONFLICT (ID) DO NOTHING;;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS FRUIT;
-- +goose StatementEnd
