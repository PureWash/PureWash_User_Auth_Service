CREATE TABLE "users" (
    "id" UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "full_name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(65) NOT NULL,
    "password_hash" VARCHAR(255) NOT NULL
    "role" VARCHAR(65) DEFAULT 'user'
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ
);