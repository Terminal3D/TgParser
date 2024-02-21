-- name: GetAllItems :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE chat_id = $1;

-- name: GetItemsByBrand :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1 AND chat_id = $2;

-- name: GetItemsByMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE price <= $1 AND chat_id = $2;

-- name: GetItemsByBrandAndMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1 AND price <= $2 AND chat_id = $3;

-- name: GetItemsByName :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE name = $1 AND chat_id = $2;

-- name: GetAllItemsWithoutSizes :many
SELECT * FROM item WHERE available = true;