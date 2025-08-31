import { Job, Worker } from "bullmq";
import { getRedisConnObject } from "../config/redis.config";
import { BadRequestError } from "../utils/errors/app.error";
import { RoomGenerationJob } from "../validator/roomGeneration.validator";
import { ROOM_GENERATION_QUEUE } from "../queues/roomGeneration.queue";
import { ROOM_GENERATION_PAYLOAD } from "../producers/roomGeneration.producer";
import { generateRooms } from "../service/roomGeneration.service";
import logger from "../config/logger.config";

export const setupRoomGererationWorker = () => {
  const roomGenerationProcessor = new Worker<RoomGenerationJob>(
    ROOM_GENERATION_QUEUE, // name of queue

    // process function
    async (job: Job) => {
      if (job.name !== ROOM_GENERATION_PAYLOAD) {
        throw new BadRequestError("Invalid job name");
      }
      // call the service layer for business logic
      const payload = job.data;
      console.log(`Processing email for ${JSON.stringify(payload)}`);
      await generateRooms(payload);
      logger.info(`Room generation completted for this ${payload}`)
    },

    // connection of redis instance
    {
      connection: getRedisConnObject(),
    }
  );

  roomGenerationProcessor.on("failed", () => {
    console.log("Room Generation failed");
  });

  roomGenerationProcessor.on("completed", () => {
    console.log("Room Generation completed sucesfully");
  });
};
