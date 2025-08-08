import { PrismaClient } from "@prisma/client";
import { boolean } from "zod";

const prisma = new PrismaClient();

// Test database connection without disconnecting
async function testConnection():Promise<boolean> {
  try {
    await prisma.$connect();
    console.log('✅ Successfully connected to the database');
    return true;
  } catch (error) {
    console.error('❌ Error connecting to the database:',error);
    return false;
  }
}

export { prisma as default, testConnection };
