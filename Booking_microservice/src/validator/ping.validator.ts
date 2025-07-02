import { z } from "zod";



export const pingSchema = z.object({
    message: z.string().min(1, "Message cannot be empty"),
});



// Example usage
// const result = pingSchema.safeParse({ message: "Hello"});