-- name: CreateUser :one
INSERT INTO users (
  id, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUserOTP :one
INSERT INTO users_otps (
  id, user_id, otp, expires_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetLatestUserOTP :one
SELECT * FROM users_otps WHERE user_id = $1 AND used_at IS NULL 
ORDER BY created_at DESC LIMIT 1;

-- name: UpdateUserOTPUsedAt :exec
UPDATE users_otps SET used_at = NOW() WHERE id = $1;