-- postgreSQL

CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$ 
BEGIN 
  NEW.updated_at = now();
  RETURN NEW;
END; $$ language 'plpgsql';

CREATE TABLE Users (
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT NOT NULL CHECK (char_length(name) <= 45),
	user_name TEXT NOT NULL CHECK (char_length(user_name) <= 45),
	password TEXT NOT NULL CHECK (char_length(password) <= 225),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by INTEGER NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by TEXT
);

CREATE TRIGGER Users BEFORE
UPDATE
	ON Users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
 

CREATE TABLE Merchants (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id INTEGER NOT NULL REFERENCES Users(id),
	merchant_name TEXT NOT NULL CHECK (char_length(merchant_name) <= 40),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by INTEGER NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by INTEGER NOT NULL
);

CREATE TRIGGER Merchants BEFORE
UPDATE
	ON Merchants FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE Outlets (
	id SERIAL PRIMARY KEY NOT NULL,
	merchant_id INTEGER NOT NULL REFERENCES Merchants(id),
	outlet_name TEXT NOT NULL CHECK (char_length(outlet_name) <= 40),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by INTEGER NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by INTEGER NOT NULL
);

CREATE TRIGGER Outlets BEFORE
UPDATE
	ON Outlets FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
 
 
CREATE TABLE Transactions (
	id SERIAL PRIMARY KEY NOT NULL,
	merchant_id INTEGER NOT NULL REFERENCES Merchants(id),
	outlet_id INTEGER NOT NULL REFERENCES Outlets(id),
	bill_total NUMERIC(20,6) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by INTEGER NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by INTEGER NOT NULL
);

CREATE TRIGGER Transactions BEFORE
UPDATE
	ON Transactions FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
 