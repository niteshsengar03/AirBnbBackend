import express from "express";
import HotelRouter from "./hotel.router";
import RoomGenerationRouter from "./roomGeneration.router";

const V1Router = express.Router();

V1Router.use("/hotels", HotelRouter);
V1Router.use("/room-generate", RoomGenerationRouter);

export default V1Router;
