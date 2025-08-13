import logger from "../config/logger.config";
import { createHotelDTO, updateHotelDTO } from "../DTO/hotel.dto";
import prisma from "../prisma/client";
import { InternalServerError, NotFoundError } from "../utils/errors/app.error";

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
  return hotel;
}

export async function getAllHotels() {
  const hotels = await prisma.hotels.findMany({
    where: {
      deleted_at: null,
    },
  });
  if (!hotels.length) {
    throw new NotFoundError("No Hotel found");
  }
  return hotels;
}

// export async function deleteHotel(id: number) {
//   await getHotelById(id); // error will be handle in this fn()
//   const hotel = await prisma.hotels.delete({
//     where: {
//       id,
//     },
//   });
//   if (!hotel) {
//     throw new NotFoundError(`Hotel with ${id} cannot be deleted`);
//   }
//   return hotel;
// }

// not using update many because it throws error if not updated 
export async function softDeleteHotel(id: number) {
  const result = await prisma.hotels.updateMany({
    where: { id },
    data: { deleted_at: new Date() },
  });

  if (result.count === 0) {
    throw new InternalServerError("Update failed — no rows affected.");
  }
  return true;
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
