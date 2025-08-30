import { Request, Response } from "express";
import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer";

export async function generateRoomHandler(req: Request, res: Response) {


    await addRoomGenerationJobToQueue(req.body);

    res.status(200).json({
        message: "Room generation job added to queue",
        success: true,
        data: {},
    })
}