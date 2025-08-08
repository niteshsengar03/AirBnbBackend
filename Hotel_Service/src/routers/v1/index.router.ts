import express from 'express';
import pingRouter from './ping.router';

const V1Router = express.Router();

V1Router.use('/', pingRouter);


export default V1Router;