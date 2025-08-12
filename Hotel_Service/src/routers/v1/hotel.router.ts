import { Router } from "express";
import { createHotelHandler, getHotelByIdHandler } from "../../controllers/hotel.controller";
import { validateBody } from "../../validator";
import { hotelSchema } from "../../validator/hotel.validator";

const HotelRouter = Router();

HotelRouter.post('/hotel',validateBody(hotelSchema),createHotelHandler);
HotelRouter.get('/hotel',getHotelByIdHandler);

export default HotelRouter;