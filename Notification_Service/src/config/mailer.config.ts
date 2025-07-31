import nodemailer from "nodemailer"
import { renderMailTemplate } from "../templates/templates.handler";
import serverConfig from ".";

export const transporter = nodemailer.createTransport({
  port: 465,
  host: "smtp.gmail.com",
  auth: {
    user: serverConfig.MAIL_USER,
    pass: serverConfig.MAIL_PASSWORD
  },
  secure: true,
});

