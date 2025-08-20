import logger from "../config/logger.config";
import { createRoomCateogoryDTO, updateRoomCateogoryDTO } from "../DTO/roomCategoryDto";
import prisma from "../prisma/client";
import { InternalServerError, NotFoundError } from "../utils/errors/app.error";

export async function createRoomCateogory(
  RoomCateogoryData: createRoomCateogoryDTO
) {
  const roomCateogory = await prisma.room_categories.create({
    data: RoomCateogoryData,
  });
  logger.info(`Room Cateogory created: ${roomCateogory.id}`);
  return roomCateogory;
}

export async function getRoomCateogoryById(id: number) {
  const roomCateogory = await prisma.room_categories.findUnique({
    where: { id },
  });
  return roomCateogory;
}

export async function getAllRoomCateogory() {
  const roomCateogorys = await await prisma.room_categories.findMany({
    where: {
      deleted_at: null,
    },
  });
  if (!roomCateogorys.length) {
    throw new NotFoundError("No Hotel found");
  }
  return roomCateogorys;
}

export async function softDeleteRoomCateogory(id: number) {
  const result = await prisma.room_categories.updateMany({
    where: { id },
    data: { deleted_at: new Date() },
  });

  if (result.count === 0) {
    throw new InternalServerError("Update failed — no rows affected.");
  }
  return true;
}

