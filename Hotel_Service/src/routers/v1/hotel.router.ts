import { Router } from "express";
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelByIdHandler, updateHotelHandler } from "../../controllers/hotel.controller";
import { validateBody } from "../../validator";
import { hotelSchema } from "../../validator/hotel.validator";

const HotelRouter = Router();

HotelRouter.post('/',validateBody(hotelSchema),createHotelHandler);
HotelRouter.get('/:id',getHotelByIdHandler);
HotelRouter.get('/',getAllHotelsHandler);
HotelRouter.delete('/:id',deleteHotelHandler);
HotelRouter.put('/:id',updateHotelHandler);

export default HotelRouter; 