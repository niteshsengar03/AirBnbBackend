-- CreateTable
CREATE TABLE `room_categories` (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `room_type` ENUM('SINGLE', 'DOUBLE', 'FAMILY', 'DELUX', 'SUITE') NOT NULL,
    `price` INTEGER NOT NULL,
    `hotel_id` INTEGER NULL,
    `room_count` INTEGER NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,
    `deleted_at` DATETIME(3) NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
