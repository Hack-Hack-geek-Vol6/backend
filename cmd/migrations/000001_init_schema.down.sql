ALTER TABLE "likes" DROP CONSTRAINT "likes_account_id_fkey";
ALTER TABLE "likes" DROP CONSTRAINT "likes_opus_fkey";
ALTER TABLE "rate_entities" DROP CONSTRAINT "rate_entities_account_id_fkey";
ALTER TABLE "account_tags" DROP CONSTRAINT "account_tags_account_id_fkey";
ALTER TABLE "account_tags" DROP CONSTRAINT "account_tags_tech_tag_id_fkey";
ALTER TABLE "accounts" DROP CONSTRAINT "accounts_locate_id_fkey";
ALTER TABLE "accounts" DROP CONSTRAINT "accounts_user_id_fkey";
ALTER TABLE "accounts_achievments" DROP CONSTRAINT "accounts_achievments_account_id_fkey";
ALTER TABLE "accounts_achievments" DROP CONSTRAINT "accounts_achievments_achievment_id_fkey";
ALTER TABLE "account_frameworks" DROP CONSTRAINT "account_frameworks_account_id_fkey";
ALTER TABLE "account_frameworks" DROP CONSTRAINT "account_frameworks_framework_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_account_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_room_id_fkey";
ALTER TABLE "rooms_accounts" DROP CONSTRAINT "rooms_accounts_role_fkey";
ALTER TABLE "follows" DROP CONSTRAINT "follows_from_account_id_fkey";
ALTER TABLE "follows" DROP CONSTRAINT "follows_to_account_id_fkey";
ALTER TABLE "award_data" DROP CONSTRAINT "award_data_award_id_fkey";
ALTER TABLE "award_data" DROP CONSTRAINT "award_data_hackathon_id_fkey";
ALTER TABLE "past_works" DROP CONSTRAINT "past_works_award_data_id_fkey";
ALTER TABLE "past_work_frameworks" DROP CONSTRAINT "past_work_frameworks_framework_id_fkey";
ALTER TABLE "past_work_frameworks" DROP CONSTRAINT "past_work_frameworks_opus_fkey";
ALTER TABLE "hackathon_status_tags" DROP CONSTRAINT "hackathon_status_tags_hackathon_id_fkey";
ALTER TABLE "hackathon_status_tags" DROP CONSTRAINT "hackathon_status_tags_status_id_fkey";
ALTER TABLE "rooms" DROP CONSTRAINT "rooms_hackathon_id_fkey";
ALTER TABLE "past_work_tags" DROP CONSTRAINT "past_work_tags_opus_fkey";
ALTER TABLE "past_work_tags" DROP CONSTRAINT "past_work_tags_tech_tag_id_fkey";
ALTER TABLE "frameworks" DROP CONSTRAINT "frameworks_tech_tag_id_fkey";
ALTER TABLE "frameworks" DROP CONSTRAINT "frameworks_pkey";
ALTER TABLE "tech_tags" DROP CONSTRAINT "tech_tags_pkey";
ALTER TABLE "roles" DROP CONSTRAINT "roles_pkey";
ALTER TABLE "status_tags" DROP CONSTRAINT "status_tags_pkey";
ALTER TABLE "locates" DROP CONSTRAINT "locates_pkey";
ALTER TABLE "users" DROP CONSTRAINT "users_pkey";
ALTER TABLE "tutor" DROP CONSTRAINT "tutor_pkey";


-- テーブルを削除
DROP TABLE IF EXISTS "likes";
DROP TABLE IF EXISTS "rate_entities";
DROP TABLE IF EXISTS "account_tags";
DROP TABLE IF EXISTS "accounts_achievments";
DROP TABLE IF EXISTS "account_frameworks";
DROP TABLE IF EXISTS "rooms_accounts";
DROP TABLE IF EXISTS "follows";
DROP TABLE IF EXISTS "award_data";
DROP TABLE IF EXISTS "awards";
DROP TABLE IF EXISTS "past_work_frameworks";
DROP TABLE IF EXISTS "hackathon_status_tags";
DROP TABLE IF EXISTS "status_tags";
DROP TABLE IF EXISTS "rooms";
DROP TABLE IF EXISTS "frameworks";
DROP TABLE IF EXISTS "locates";
DROP TABLE IF EXISTS "hackathons";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "tutor";
DROP TABLE IF EXISTS "past_work_tags";
DROP TABLE IF EXISTS "tech_tags";
DROP TABLE IF EXISTS "accounts";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "achievments";
DROP TABLE IF EXISTS "past_works";
DROP TABLE IF EXISTS "account_past_works";