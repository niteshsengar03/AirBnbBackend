import { log } from "console";
import logger from "../config/logger.config";
import { createHotelDTO } from "../DTO/hotel.dto";
import prisma from "../prisma/client";
import { threadId } from "worker_threads";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData: createHotelDTO) {
  const hotel = await prisma.hotels.create({
    data: hotelData,
  });
  logger.info(`Hotel created: ${hotel.id}`)
  return hotel;
   
}

export async function getHotelById(id:number){
     const hotel = await prisma.hotels.findUnique({
        where:{id}
     })
     if(!hotel){
        logger.info(`No Hotel found of id: ${id}`)
        throw new NotFoundError(`No Hotel found of id: ${id}`)
     }
    console.info(`Hotel found: ${JSON.stringify(hotel)}`)
    return hotel;
}  