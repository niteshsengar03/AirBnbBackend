import { createHotelDTO, updateHotelDTO } from "../DTO/hotel.dto";
import {
  createHotel,
  getAllHotels,
  getHotelById,
  softDeleteHotel,
  updateHotel,
} from "../repositories/hotel.repositories";
import logger from "../config/logger.config";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotelService(hotelData: createHotelDTO) {
  const hotel = await createHotel(hotelData);
  return hotel;
}

export async function getHotelByIdService(id: number) {
  const hotel = await getHotelById(id);
  if (!hotel) {
    logger.info(`No Hotel found of id: ${id}`);
    throw new NotFoundError(`No Hotel found of id: ${id}`);
  }
  if (hotel.deleted_at!=null){
    throw new NotFoundError(`Hotel of id: ${id} is deleted`);
  }
  console.info(`Hotel found: ${JSON.stringify(hotel)}`);
  return hotel;
}

export async function getAllHotelsService() {
  const hotels = await getAllHotels();
  return hotels;
}

// export async function deleteHotelService(id: number) {
//   const hotel = await deleteHotel(id);
//   return hotel;
// }

export async function updateHotelService(
  id: number,
  hotelData: updateHotelDTO
) {
  await getHotelByIdService(id );
  const hotel = await updateHotel(id, hotelData);
  return hotel;
}

export async function softDeleteHotelService(id: number) {
  // check if hotel exists or if hotel is soft deleted
  const hotel = await getHotelByIdService(id); 

  const deleted = await softDeleteHotel(id);
  return deleted;
}
