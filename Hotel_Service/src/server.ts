import express from 'express';
import { Express } from 'express';
import serverConfig from './config/index';
import V1Router from './routers/v1/index.router';
import { genericErrorHandler } from './middlewares/error.middleware';
import logger from './config/logger.config';
import { attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';



// const app = express(); // implicit
const app:Express = express() // explcit

// const port: number = 3000;
app.use(express.json());


app.use(attachCorrelationIdMiddleware);
app.use('/api/v1',V1Router);

app.use(genericErrorHandler);


app.listen(serverConfig.PORT,()=>{
    logger.info(`Port is running on http://localhost:${serverConfig.PORT}`);
    logger.info(`Press Cnt+C to exist`,{"server":"dev server"});
})