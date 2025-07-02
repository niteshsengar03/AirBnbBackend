import express from 'express';
import { validateBody } from '../../validator';
import {createBookingSchema} from '../../validator/booking.validator'
import { confrimBookingHandler, createBookingHandler } from '../../controllers/booking.controllers';

const bookingRouter = express.Router();

bookingRouter.post('/',validateBody(createBookingSchema),createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey',confrimBookingHandler);

export default bookingRouter;