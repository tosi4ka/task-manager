CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID REFERENCES users(id),
	name VARCHAR(100) NOT NULL,
	description VARCHAR(255),
	category VARCHAR(50),
	status VARCHAR(50) NOT NULL  DEFAULT 'todo',
	priority VARCHAR(50) NOT NULL DEFAULT 'medium',
	deadline TIMESTAMP,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
)