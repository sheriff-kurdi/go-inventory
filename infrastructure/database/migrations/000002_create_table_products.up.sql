DROP TABLE IF EXISTS "products";

DROP SEQUENCE IF EXISTS products_id_seq;

CREATE SEQUENCE products_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."products" (
    "id" bigint DEFAULT nextval('products_id_seq') NOT NULL,
    "cost_price" numeric,
    "selling_price" numeric,
    "discount" numeric,
    "is_discounted" boolean,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    "total_stock" bigint,
    "available_stock" bigint,
    "reserved_stock" bigint,
    CONSTRAINT "products_pkey" PRIMARY KEY ("id")
);