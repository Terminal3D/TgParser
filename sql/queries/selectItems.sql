-- name: GetAllItems :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id;

-- name: GetItemsByBrand :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1;

-- name: GetItemsByMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE price <= $1;

-- name: GetItemsByBrandAndMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1 AND price <= $2;

-- name: GetItemsByName :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE name = $1;