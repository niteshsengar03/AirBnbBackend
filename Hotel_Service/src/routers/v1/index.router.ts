import express from 'express';
import HotelRouter from './hotel.router';

const V1Router = express.Router();


V1Router.use('/hotels',HotelRouter);

export default V1Router;