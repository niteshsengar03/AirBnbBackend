export type createRoomCateogoryDTO = {
  room_type: "SINGLE" | "DOUBLE" | "FAMILY" | "DELUX" | "SUITE";
  price: number;
  room_count: number;
  hotel_id: number;
};

export type updateRoomCateogoryDTO = {
  room_type?: "SINGLE" | "DOUBLE" | "FAMILY" | "DELUX" | "SUITE";
  price?: number;
  room_count?: number;
  hotel_id?: number;
};
