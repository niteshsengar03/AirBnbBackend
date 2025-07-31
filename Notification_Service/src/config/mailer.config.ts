import nodemailer from "nodemailer"
import serverConfig from ".";

const transporter = nodemailer.createTransport({
   service: 'gmail',
  auth: {
    user: serverConfig.MAIL_USER,
    pass: serverConfig.MAIL_PASSWORD
  }
});

export default transporter;

