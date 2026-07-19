-- Placeholder for a migration version that was applied directly against
-- the live database (recorded in hdb_catalog.schema_migrations) but
-- never had a corresponding local migration file. Whatever schema
-- change happened here is already captured redundantly by the later
-- full-schema sync migrations, so this is intentionally a no-op - it
-- exists only to make the CLI's local/remote bookkeeping consistent so
-- `migrate squash` can run.
SELECT 1;
