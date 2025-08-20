import { getRoomCateogoryById } from "../repositories/roomCategory.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { RoomGenerationJob } from "../validator/roomGeneration.validator";
export async function generateRooms(jobData: RoomGenerationJob) {
  const roomCategoryId = getRoomCateogoryById(jobData.roomCategoryId);
  if (!roomCategoryId) {
    throw new NotFoundError(
      `Room Category with ${jobData.roomCategoryId} not found`
    );
  }

  const startDate = new Date(jobData.startDate);
  const endDate = new Date(jobData.endDate);
  if (startDate > endDate)
    throw new BadRequestError(`Start Date must be before end date`);
  if (startDate < new Date())
    throw new BadRequestError(`Start date cannot be in past`);

  const totalDays = Math.ceil((endDate.getTime() - startDate.getTime())/(1000*60*60*24))
  
  
}
