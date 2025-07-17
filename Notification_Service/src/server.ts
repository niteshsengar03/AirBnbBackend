import express from "express";
import { Express } from "express";
import serverConfig from "./config/index";
import V1Router from "./routers/v1/index.router";
import { genericErrorHandler } from "./middlewares/error.middleware";
import logger from "./config/logger.config";
import { attachCorrelationIdMiddleware } from "./middlewares/correlation.middleware";
import { setupMailerWorker } from "./processors/email.processor";
import { notificationDto } from "./dto/notification.dto";
import { addEmailToQueue } from "./producers/email.producer";

// const app = express(); // implicit
const app: Express = express(); // explcit

// const port: number = 3000;
app.use(express.json());

app.use(attachCorrelationIdMiddleware);
app.use("/api/v1", V1Router);

app.use(genericErrorHandler);

// testing
app.listen(serverConfig.PORT, () => {
  logger.info(`Port is running on http://localhost:${serverConfig.PORT}`);
  logger.info(`Press Cnt+C to exist`, { server: "dev server" });
  setupMailerWorker();
  logger.info(`Mailer worker setup completed`);

  const sampleNotification: notificationDto = {
    to: "sample",
    subject: "Sample Email",
    templateId: "1",
    params: {
      name: "Alice",
      orderId: "123",
    },
  };

  addEmailToQueue(sampleNotification);
});