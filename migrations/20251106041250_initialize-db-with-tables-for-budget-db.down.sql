-- Drop Triggers first
DROP TRIGGER IF EXISTS trg_bookings_updated_at ON bookings;
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
DROP TRIGGER IF EXISTS trg_households_updated_at ON households;

-- Drop Tables in reverse order of creation
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS households;

-- Drop the Indexes (Optional, as DROP TABLE usually removes associated indexes)
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_active_bookings_user;

-- Drop the Function last
DROP FUNCTION IF EXISTS set_updated_at();
