import express from 'express';  
import bookingRouter from './booking.router';

const V1Router = express.Router();

V1Router.use('/bookings', bookingRouter);


export default V1Router;