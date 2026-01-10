-- Create the automation function
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Households
CREATE TABLE IF NOT EXISTS households (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  address TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users Table
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

-- 4. Bookings Table (Soft Delete)
CREATE TABLE IF NOT EXISTS bookings (
  id SERIAL PRIMARY KEY,
  amount NUMERIC(10,2) NOT NULL DEFAULT 0,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  booking_platform VARCHAR(255) NOT NULL,
  free_cancel_before TIMESTAMP NOT NULL,
  booking_start TIMESTAMP NOT NULL,
  booking_end TIMESTAMP NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

---
-- INDEXES
---
CREATE INDEX idx_users_created_at ON users(created_at DESC);
CREATE INDEX idx_active_bookings_user ON bookings(user_id)
WHERE deleted_at IS NULL;

---
-- TRIGGERS (Automating updated_at)
---
CREATE TRIGGER trg_households_updated_at
  BEFORE UPDATE ON households FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_users_updated_at
  BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_bookings_updated_at
  BEFORE UPDATE ON bookings FOR EACH ROW EXECUTE FUNCTION set_updated_at();
