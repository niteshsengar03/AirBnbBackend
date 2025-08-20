import { z } from "zod";

export const RoomGenerationRequestSchema = z.object({
  roomCategoryId: z.number().positive(),
  startDate: z.string().datetime(),
  endDate: z.string().datetime(),
  scheduleType: z.enum(["immediate", "scheduled"]).default("immediate"),
  scheduledAt: z.string().datetime().optional(),
  priceOverride: z.number().positive().optional(),
});

export const RoomGenerationJobSchema = z.object({
  roomCategoryId: z.number().positive(),
  startDate: z.string().datetime(),
  endDate: z.string().datetime(),
  priceOverride: z.number().positive().optional(),
  batchSize: z.number().positive().default(100),
});

// to make them a type in typscript(now we can use them as type also)
export type RoomGenerationRequest = z.infer<typeof RoomGenerationRequestSchema>;
export type RoomGenerationJob = z.infer<typeof RoomGenerationJobSchema>;
