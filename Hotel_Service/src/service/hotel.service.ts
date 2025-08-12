import { number } from "zod";
import { createHotelDTO, updateHotelDTO } from "../DTO/hotel.dto";
import { createHotel, deleteHotel, getAllHotels, getHotelById, updateHotel } from "../repositories/hotel.repositories";

export async function createHotelService(hotelData:createHotelDTO){
    const hotel =  await createHotel(hotelData)
    return hotel;
}

export async function getHotelByIdService(id:number){
    const hotel = await getHotelById(id)
    return hotel;
}

export async function getAllHotelsService(){
    const hotels = await getAllHotels()
    return hotels;
}

export async function deleteHotelService(id:number){
    const hotel = await deleteHotel(id);
    return hotel;
}

export async function updateHotelService(id:number,hotelData:updateHotelDTO){
    const hotel = await updateHotel(id,hotelData);
    return hotel;
}