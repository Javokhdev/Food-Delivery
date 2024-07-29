-- Drop the password_resets table
DROP TABLE IF EXISTS password_resets;

-- Drop the tokens table
DROP TABLE IF EXISTS tokens;

-- Drop the trigger and function for updating the updated_at timestamp
DROP TRIGGER IF EXISTS update_timestamp_trigger ON users;
DROP FUNCTION IF EXISTS update_timestamp;

-- Drop the users table
DROP TABLE IF EXISTS users;
