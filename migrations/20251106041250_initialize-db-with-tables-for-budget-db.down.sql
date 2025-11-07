-- Drop indexes first (optional since dropping tables will drop indexes automatically)
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_household_id;
DROP INDEX IF EXISTS idx_households_name;

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS budgets;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS households;
