import { Prisma, room_categories } from "@prisma/client";
import { bulkCreate, findByRoomCategoryIdAndDate } from "../repositories/room.repository";
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
// process date in batches
export async function processDateBatch(roomCategory:room_categories,startDate:Date,endDate:Date,priceOverride?:number){
    let roomCreated = 0;
    let dateProcessed = 0;
    const roomsToCreate: Prisma.roomsCreateManyInput[] = [];
    const currentDate = new Date(startDate);
    // N+1 querry problem 
    // quering n number of times calling db 
    // instead use sql query
    //kind of select * from roomCategory where id = ? and dateOfavailability in [?,?,?] <- array of dates
    while(currentDate<=endDate){
      const existingRoom = await findByRoomCategoryIdAndDate(roomCategory.id,currentDate);
      
      if(!existingRoom){
        roomsToCreate.push({
          hotel_id:roomCategory.hotel_id,
          room_category_id:roomCategory.id,
          date_of_availability:currentDate,
          price: priceOverride || roomCategory.price,
        })
      }

      currentDate.setDate(currentDate.getDate()+1);
      dateProcessed++;
    }

    if (roomsToCreate.length>0)
        await bulkCreate(roomsToCreate)
      roomCreated +=  roomsToCreate.length;
    
    return {
      roomCreated,
      dateProcessed
    }
}
