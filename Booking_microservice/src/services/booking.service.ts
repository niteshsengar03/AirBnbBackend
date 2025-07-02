import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalzeIdempotencyKey,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from "../utils/errors/app.error";
import { generateIdempotencykey } from "../utils/generateIdempotencykey";
import { CreateBookingDTO } from "../dto/booking.dto";

import primsaClient from "../prisma/client";
import { redlock } from "../config/redis.config";
import serverConfig from "../config";




// using redis to apply redlock from multiple users while creating Booking for same room 
export async function createBookingService(createBookingDTO: CreateBookingDTO) {
  const ttl = serverConfig.LOCK_TTL; // n minutes lock
  const bookingResource = `hotel:${createBookingDTO.hotelId}`;
  let lock;
  try {
    lock = await redlock.acquire([bookingResource], ttl);
    const booking = await createBooking({
      userId: createBookingDTO.userId,
      hotelId: createBookingDTO.hotelId,
      totalGuests: createBookingDTO.totalGuests,
      bookingAmount: createBookingDTO.bookingAmount,
    });

    const idempotencyKey = generateIdempotencykey();

    await createIdempotencyKey(idempotencyKey, booking.id);

    return {
      bookingId: booking.id,
      idempotencyKey: idempotencyKey,
    };
  } catch (err) {
    throw new InternalServerError(
      "Hotel is being booked by someone else. Please try again."
    );
  }
}

// implementing transaction and pestimic lock on row when using getIdempotnecyKey
export async function confrimBookingService(idempotencyKey: string) {
  return primsaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey
    );
    // not found check is done in repositroy layer
    if (!idempotencyKeyData) {
      throw new NotFoundError("Idempotency key not found");
    }
    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("Idempotency key is already finalised");
    }

    if (idempotencyKeyData.bookingId === null) {
      throw new NotFoundError(
        "Booking ID not found for the given idempotency key"
      );
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);

    await finalzeIdempotencyKey(tx, idempotencyKey);

    return booking;
  });
}
