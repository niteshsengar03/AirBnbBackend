import { createHotelDTO } from "../DTO/hotel.dto";
import { createHotel, getHotelById } from "../repositories/hotel.repositories";

export async function createHotelService(hotelData:createHotelDTO){
    const hotel =  await createHotel(hotelData)
    return hotel;
}

export async function getHotelByIdService(id:number){
    const hotel = await getHotelById(id)
    return hotel;
}