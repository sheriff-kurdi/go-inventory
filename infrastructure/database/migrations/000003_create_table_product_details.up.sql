DROP TABLE IF EXISTS "product_details";

CREATE TABLE "public"."product_details" (
    "name" text NOT NULL,
    "description" text,
    "language_code" text NOT NULL,
    "product_id" bigint NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    CONSTRAINT "product_details_language_code_name_product_id" PRIMARY KEY ("language_code", "name", "product_id") 
    CONSTRAINT "product_details_language_code_product_id" UNIQUE ("language_code", "product_id")
);