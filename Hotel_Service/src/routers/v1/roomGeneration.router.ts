import { Router } from "express";

import { validateBody } from "../../validator";
import { RoomGenerationJobSchema } from "../../validator/roomGeneration.validator";
import { generateRoomHandler } from "../../controllers/roomGeneration.controller";

const RoomGenerationRouter = Router();

RoomGenerationRouter.post(
  "/",
  validateBody(RoomGenerationJobSchema),
  generateRoomHandler
);

export default RoomGenerationRouter;
