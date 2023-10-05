-- +goose Up
-- +goose StatementBegin
INSERT INTO tenor (tenor, value) VALUES (1, 0.3),(2,0.5),(3,0.6),(6,0.9);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
