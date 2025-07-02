import { NextFunction, Request,Response } from "express"
import fs from "fs/promises";
import { NotFoundError } from "../utils/errors/app.error";
import logger from "../config/logger.config";

 const pingHander =  async (req:Request,res:Response,next:NextFunction) =>{
  
 try{
   //  await fs.readFile("sample");
    logger.info("Ping request recieved");
    res.status(200).json({message:"Pong"});

 }catch(error){
   logger.info("Ping request falied");
   // console.log(error);
    next(new NotFoundError("File not found"));
    // or 
   //  throw (new NotFoundError("File not found"));
 } 
}

export default pingHander; 