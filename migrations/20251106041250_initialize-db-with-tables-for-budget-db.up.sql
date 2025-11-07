-- Households table
CREATE TABLE IF NOT EXISTS households (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  username VARCHAR(255) NOT NULL,
  lastname VARCHAR(255) NOT NULL,
  gender VARCHAR(1) NOT NULL CHECK (gender IN ('M', 'F')),
  role VARCHAR(120) NOT NULL DEFAULT 'user' CHECK (role IN ('admin', 'user')),
  household_id INTEGER REFERENCES households(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Budget table
CREATE TABLE IF NOT EXISTS budgets (
  id SERIAL PRIMARY KEY,
  amount NUMERIC(10,2) NOT NULL DEFAULT 0,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_households_name ON households(name);
CREATE INDEX idx_users_created_at ON users(created_at DESC);
CREATE INDEX idx_users_household_id ON users(household_id);
