/*
  Warnings:

  - Made the column `hotel_id` on table `room_categories` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE `room_categories` MODIFY `hotel_id` INTEGER NOT NULL;

-- AddForeignKey
ALTER TABLE `room_categories` ADD CONSTRAINT `room_categories_hotel_id_fkey` FOREIGN KEY (`hotel_id`) REFERENCES `hotels`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;
