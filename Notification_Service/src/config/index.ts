import dotenv from 'dotenv';

type serverconfig = {
    PORT:number;
    REDIS_PORT:number;
    REDIS_HOST:string;
    MAIL_PASSWORD:string;
    MAIL_USER:string;
}

// load the env when server is running on machine
// It reads the .env file (a simple text file with key=value pairs) and loads those values into process.env — which is just a JavaScript object holding environment variables.
//Gets "unloaded" (really just disappears) when the Node.js app stops
function loadEnv(){ 
    dotenv.config();
}

loadEnv();
 const serverConfig:serverconfig = {
    PORT: Number(process.env.PORT) || 3001,
    REDIS_PORT:Number(process.env.REDIS_PORT)|| 6379,
    REDIS_HOST:process.env.REDIS_HOST || "localhost",
    MAIL_PASSWORD:process.env.MAIL_PASSWORD || "you gmail app password",
    MAIL_USER:process.env.MAIL_USER || "user@gmail.com"
}

export default serverConfig;


