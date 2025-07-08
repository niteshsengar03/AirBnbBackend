import { notificationDto } from "../dto/notification.dto";
import { mailerQueue } from "../queues/mailer.queue";

export const MAILER_PAYLOAD = "payload:mail";

export const addEmailToQueue = async (payload: notificationDto) => {
  await mailerQueue.add(MAILER_PAYLOAD, payload);
  console.log(`Email added to queue :${JSON.stringify(payload)}`);
};
