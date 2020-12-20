
CREATE TABLE transactions (
    uri text NOT NULL,
    program_id text NOT NULL,
    data text NOT NULL
);

/* name: GetTransaction :many */
SELECT
	json_extract(transactions.data, '$.transaction.signatures[0]'),
	json_group_array(instructions.value)
FROM