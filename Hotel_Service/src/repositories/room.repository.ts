import { Prisma, rooms } from "@prisma/client";
import logger from "../config/logger.config";
import { CreateRoomDTO } from "../DTO/room.dto";
import prisma from "../prisma/client";
import { InternalServerError, NotFoundError } from "../utils/errors/app.error";

// Create Room
export async function createRoom(roomData: CreateRoomDTO) {
  const room = await prisma.rooms.create({
    data: roomData,
  });
  logger.info(`Room created: ${room.id}`);
  return room;
}

// Get Room by ID
export async function getRoomById(id: number) {
  const room = await prisma.rooms.findUnique({
    where: { id },
  });

  //   if (!room || room.deleted_at) {
  //     throw new NotFoundError(`Room with ID ${id} not found`);
  //   }

  return room;
}

// Get All Rooms
export async function getAllRooms() {
  const rooms = await prisma.rooms.findMany({
    where: {
      deleted_at: null,
    },
  });

  if (!rooms.length) {
    throw new NotFoundError("No Rooms found");
  }

  return rooms;
}

// Soft Delete Room
export async function softDeleteRoom(id: number) {
  const result = await prisma.rooms.updateMany({
    where: { id, deleted_at: null },
    data: { deleted_at: new Date() },
  });

  if (result.count === 0) {
    throw new InternalServerError("Soft delete failed — no rows affected.");
  }

  logger.info(`Room soft deleted: ${id}`);
  return true;
}

export async function findByRoomCategoryIdAndDate(
  roomCategoryId: number,
  currentDate: Date
) {
  return prisma.rooms.findFirst({
    where: {
      room_category_id: roomCategoryId,
      date_of_availability: currentDate,
      deleted_at: null,
    },
  });
}

export async function bulkCreate(rooms:Prisma.roomsCreateManyInput[]){
  return await prisma.rooms.createMany({data:rooms})
}