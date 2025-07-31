import { Job, Worker } from "bullmq";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { notificationDto } from "../dto/notification.dto";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import { BadRequestError } from "../utils/errors/app.error";
import { sendEmail } from "../service/mailer.service";
import { renderMailTemplate } from "../templates/templates.handler";

export const setupMailerWorker = () => {
  const emailProcessor = new Worker<notificationDto>(
    MAILER_QUEUE, // name of queue

    // process function
    async (job: Job) => {
      if (job.name !== MAILER_PAYLOAD) {
        throw new BadRequestError("Invalid job name");
      }
      // call the service layer for business logic
      const payload = job.data;
      console.log(`Processing email for ${JSON.stringify(payload)}`);
      const emailContent = await renderMailTemplate(
        payload.templateId,
        payload.params
      );

      await sendEmail(payload.to, payload.subject, emailContent);
    },

    // connection of redis instance
    {
      connection: getRedisConnObject(),
    }
  );

  emailProcessor.on("failed", () => {
    console.log("Email processing failed");
  });

  emailProcessor.on("completed", () => {
    console.log("Email processing completed sucesfully");
  });
};
