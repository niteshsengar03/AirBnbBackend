import { Request, Response } from "express";
import {
  confrimBookingService,
  createBookingService,
} from "../services/booking.service";

export const createBookingHandler = async (req: Request, res: Response) => {
  // 1call service
  const booking = await createBookingService(req.body);
  // 2 send response
  res.status(201).json({
    bookingId: booking.bookingId,
    idempotencyKey: booking.idempotencyKey,
  });
};

export const confrimBookingHandler = async (req: Request, res: Response) => {
  const booking = await confrimBookingService(req.params.idempotencyKey);

  res.status(201).json({
    booking,
  });
};
