// For creating a room
export interface CreateRoomDTO {
  room_category_id: number;
  hotel_id: number;
  room_no: number;
  date_of_availability: Date;
  booking_id?: number | null;
}

// For updating a room
export interface UpdateRoomDTO {
  roomCategoryId?: number;
  hotelId?: number;
  roomNo?: number;
  dateOfAvailability?: Date;
  bookingId?: number | null;
  deletedAt?: Date | null;
}
