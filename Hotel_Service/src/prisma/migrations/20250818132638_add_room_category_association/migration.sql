/*
  Warnings:

  - Made the column `room_category_id` on table `rooms` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE `rooms` MODIFY `room_category_id` INTEGER NOT NULL;

-- AddForeignKey
ALTER TABLE `rooms` ADD CONSTRAINT `rooms_room_category_id_fkey` FOREIGN KEY (`room_category_id`) REFERENCES `room_categories`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;
