import express from 'express';
import pingRouter from './ping.router';
import HotelRouter from './hotel.router';

const V1Router = express.Router();

// V1Router.use('/', pingRouter);
V1Router.use('/',HotelRouter);

export default V1Router;