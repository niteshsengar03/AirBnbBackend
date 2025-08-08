import express from 'express';
import pingHander from '../../controllers/ping.controllers';
import { validateBody } from '../../validator';
import { pingSchema } from '../../validator/ping.validator';

const pingRouter = express.Router();

pingRouter.get('/ping',validateBody(pingSchema),pingHander);

export default pingRouter;