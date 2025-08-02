-- +goose Up
-- +goose StatementBegin
CREATE TABLE reviews(
    id SERIAL PRIMARY KEY,
    userId int  NOT NULL,
    hotelId int NOT NULL,
    bookingId int NOT NULL,
    comments VARCHAR(255) NOT NULL,
    
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
