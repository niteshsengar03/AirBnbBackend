import { Router } from "express";
import { createHotelHandler, getAllHotelsHandler, getHotelByIdHandler, softDeleteHotelHandler, updateHotelHandler } from "../../controllers/hotel.controller";
import { validateBody } from "../../validator";
import { hotelSchema } from "../../validator/hotel.validator";

const HotelRouter = Router();

HotelRouter.post('/',validateBody(hotelSchema),createHotelHandler);
HotelRouter.get('/:id',getHotelByIdHandler);
HotelRouter.get('/',getAllHotelsHandler);
HotelRouter.put('/:id',updateHotelHandler);
HotelRouter.delete('/soft/:id',softDeleteHotelHandler);

export default HotelRouter; 