import dotenv from 'dotenv';

type serverconfig = {
    PORT:number;
}

// load the env when server is running on machine
// It reads the .env file (a simple text file with key=value pairs) and loads those values into process.env — which is just a JavaScript object holding environment variables.
//Gets "unloaded" (really just disappears) when the Node.js app stops
function loadEnv(){ 
    dotenv.config();
}

loadEnv();
 const serverConfig:serverconfig = {
    PORT: Number(process.env.PORT) || 3001
}

export default serverConfig;


