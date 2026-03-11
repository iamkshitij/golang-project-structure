-- Oracle Database Setup Script for Go Fiber API
-- Run this script as SYSTEM or DBA user to properly configure SCOTT

-- 1. Grant necessary privileges to SCOTT user
ALTER USER SCOTT QUOTA UNLIMITED ON SYSTEM;
-- Alternative: Grant specific quota on SCOTT's default tablespace
-- ALTER USER SCOTT QUOTA 100M ON SYSTEM;

-- 2. Grant additional privileges if needed
GRANT CREATE TABLE TO SCOTT;
GRANT CREATE INDEX TO SCOTT;
GRANT CREATE SEQUENCE TO SCOTT;

-- 3. Connect as SCOTT and create table
-- You can now run the setup_oracle.sql script as SCOTT user
-- Or continue with the commands below:

-- Switch to SCOTT schema (if running as DBA)
-- ALTER SESSION SET CURRENT_SCHEMA = SCOTT;

-- 4. Verify SCOTT's privileges
SELECT 
    username,
    tablespace_name,
    bytes/1024/1024 as mb_used,
    max_bytes/1024/1024 as mb_max
FROM dba_ts_quotas 
WHERE username = 'SCOTT';

SELECT * FROM dba_sys_privs WHERE grantee = 'SCOTT';

-- 5. Show instructions
SELECT 'SCOTT user is now properly configured. Run setup_oracle.sql as SCOTT user.' as next_step FROM dual;