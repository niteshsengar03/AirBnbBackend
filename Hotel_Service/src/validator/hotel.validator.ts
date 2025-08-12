import { z } from "zod";

export const hotelSchema = z.object({
  name:z.string(),
  address:z.string(),
  location:z.string(),
  rating:z.number().optional(),
  rating_count:z.number().optional()
});