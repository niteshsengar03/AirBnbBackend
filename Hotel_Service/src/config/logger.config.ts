import winston from "winston";
import { getCorrelationId } from "../utils/helpers/request.helpers";
import DailyRotateFile from "winston-daily-rotate-file"

const logger = winston.createLogger({
    format:winston.format.combine(
        winston.format.timestamp({format:"DD-MM-YYYY HH:mm:ss"}),
        winston.format.json(),
        //define custom print
        // data is for extra information you want to print
        winston.format.printf(({level,message,timestamp,...data})=>{
            const output = {
                level,
                message,
                timestamp,
                correlationId:getCorrelationId(),
                data
            };
        
            return JSON.stringify(output);
        })
    ),
    transports:[
        new winston.transports.Console(),
        new DailyRotateFile ({
            filename:"logs/%DATE%-app.log",
            datePattern:"DD-MM-YYYY",
            maxSize:"20m",
            maxFiles:"14d"//max number of log files to keep
        })
    ]
})

export default logger;