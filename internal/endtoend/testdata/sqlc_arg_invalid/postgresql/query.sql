CREATE TABLE users (
    id         serial,
    first_name text not null
);

-- name: WrongFunc :one
select id, first_name from users where id = sqlc.argh(target_id);

-- name: TooManyArgs :one
select id, first_name from users where id = sqlc