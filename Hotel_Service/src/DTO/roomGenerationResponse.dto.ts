export interface RoomGenerationResponse {
  success: boolean;
  totalRoomsCreated: number;
  totalDatesProcessed: number;
  errors: string[];
  jobId: string;
}
