-- Autogenerated by migration tool on 2022-04-04 11:17:00
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: storage.buckets
ALTER TABLE IF EXISTS "gcp_storage_buckets" DROP COLUMN IF EXISTS "encryption_type";
