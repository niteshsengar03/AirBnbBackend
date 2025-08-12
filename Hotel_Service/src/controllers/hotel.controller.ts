import { Request, Response } from "express";
import { createHotelService, getHotelByIdService } from "../service/hotel.service";

export async function createHotelHandler(req:Request,res:Response){
 
    const hotel = await createHotelService(req.body)

     res.status(200).json({
        hotel:hotel
     })
    
}

export async function getHotelByIdHandler(req:Request,res:Response){
   const hotel = await getHotelByIdService(req.body.id)

   res.status(200).json({
      hotel:hotel
   })
}