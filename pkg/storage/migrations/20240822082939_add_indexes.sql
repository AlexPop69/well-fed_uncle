-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_orders_client_id ON orders(client_id);
CREATE INDEX idx_orders_shop_id ON orders(shop_id);
CREATE INDEX idx_clients_name ON clients(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_orders_client_id;
DROP INDEX IF EXISTS idx_orders_shop_id;
DROP INDEX IF EXISTS idx_clients_name;
-- +goose StatementEnd
