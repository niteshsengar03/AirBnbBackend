import { NextFunction, Request, Response } from 'express';
import { v4 as uuidv4 } from 'uuid';
import { asyncLocalStorage } from '../utils/helpers/request.helpers';
export const attachCorrelationIdMiddleware=(req:Request,res:Response,next:NextFunction)=>{
    const correlationId = uuidv4();
    req.correlationId = correlationId;
    // req.headers['X-Correlation-ID']=correlationId;
    
    // whatever internal function you want to run and want it to have a 
    // correlation id run in inside asyn.run function
    asyncLocalStorage.run({correlationId:correlationId},()=>{
        next();
    })
}