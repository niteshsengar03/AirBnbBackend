import logger from "../config/logger.config";
import { createHotelDTO, updateHotelDTO } from "../DTO/hotel.dto";
import prisma from "../prisma/client";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData: createHotelDTO) {
  const hotel = await prisma.hotels.create({
    data: hotelData,
  });
  logger.info(`Hotel created: ${hotel.id}`);
  return hotel;
}

export async function getHotelById(id: number) {
  const hotel = await prisma.hotels.findUnique({
    where: { id },
  });
  if (!hotel) {
    logger.info(`No Hotel found of id: ${id}`);
    throw new NotFoundError(`No Hotel found of id: ${id}`);
  }
  console.info(`Hotel found: ${JSON.stringify(hotel)}`);
  return hotel;
}

export async function getAllHotels() {
  const hotels = await prisma.hotels.findMany();
  if (!hotels.length) {
    throw new NotFoundError("No Hotel found");
  }
  return hotels;
}

export async function deleteHotel(id: number) {
  await getHotelById(id); // error will be handle in this fn()
  const hotel = await prisma.hotels.delete({
    where: {
      id,
    },
  });
  if (!hotel) {
    throw new NotFoundError(`Hotel with ${id} cannot be deleted`);
  }
  return hotel;
}

export async function updateHotel(id: number, hotelData: updateHotelDTO) {
  await getHotelById(id);
  const hotel = await prisma.hotels.update({
    where: {
      id,
    },
    data: {
      name: hotelData.name,
      address: hotelData.address,
      location: hotelData.location,
    },
  });
  return hotel;
}
