import { Request, Response } from "express";


export async function generateRoomHandler(req: Request, res: Response) {


    res.status(200).json({
        message: "Room generation job added to queue",
        success: true,
        data: {},
    })
}