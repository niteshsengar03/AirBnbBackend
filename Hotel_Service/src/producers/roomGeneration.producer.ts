import { roomGenerationQueue } from "../queues/roomGeneration.queue";
import { RoomGenerationJob } from "../validator/roomGeneration.validator";

export const ROOM_GENERATION_PAYLOAD = "payload:room-generation";

export const addRoomGernerationJobToQueue = async (
  payload: RoomGenerationJob
) => {
  await roomGenerationQueue.add(ROOM_GENERATION_PAYLOAD, payload);
  console.log(`Room Genration job added to queue :${JSON.stringify(payload)}`);
};
