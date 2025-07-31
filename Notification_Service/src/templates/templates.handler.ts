import Handlebars from "handlebars";
import path from "path";
import fs from "fs/promises";
import { InternalServerError } from "../utils/errors/app.error";

export async function renderMailTemplate(
  templateId: string,
  params: Record<string, any>
): Promise<string> {
  const templatePath = path.join(__dirname, "mailer", `${templateId}.hbs`);
  // console.log(templatePath)

  try {
    const content = await fs.readFile(templatePath, "utf-8");
    const finalTemplate = Handlebars.compile(content);
    const result = finalTemplate(params);
    return result;
  } catch (error) {
    throw new InternalServerError(`Template not found ${templateId}`);
  }
}
