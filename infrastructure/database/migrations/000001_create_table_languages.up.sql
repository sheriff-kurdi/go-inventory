DROP TABLE IF EXISTS "languages";
CREATE TABLE "public"."languages" (
    "language_code" text NOT NULL,
    "name" text,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    CONSTRAINT "languages_pkey" PRIMARY KEY ("language_code")
);