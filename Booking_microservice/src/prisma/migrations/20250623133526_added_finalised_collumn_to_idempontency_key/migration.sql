-- AlterTable
ALTER TABLE `IdempotencyKey` ADD COLUMN `finalized` BOOLEAN NOT NULL DEFAULT false;
