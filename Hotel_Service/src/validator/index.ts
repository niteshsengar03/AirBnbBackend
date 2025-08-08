import { NextFunction, Request, Response } from "express";
import { AnyZodObject } from "zod";
import logger from "../config/logger.config";


// this function will expect the schema 
// and return a function or middleware which will check  
export const validateBody =(schema:AnyZodObject) =>{
    return  async (req:Request, res:Response, next:NextFunction) =>{
        try{
            
            logger.info("Validating request body");
            await schema.parseAsync(req.body);
            logger.info("Requet body is valid");
            next();
        } catch(err){
            logger.info("Request body is invaid");
            res.status(400).json({message:"Invalid",success:false,error:err});
        }
    }
}