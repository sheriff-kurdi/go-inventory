package seeders

import (
	"gorm.io/gorm"
)

func LanguagesSeed(connection *gorm.DB) {
	productQuery := `
			INSERT INTO "products" ("id", "cost_price", "selling_price", "discount", "is_discounted", "created_at", "updated_at", "deleted_at", "total_stock", "available_stock", "reserved_stock") VALUES
		(1,	10,	20,	5,	't',	'2023-04-23 12:10:52.39673',	'2023-04-23 12:10:52.396739',	'0001-01-01 00:00:00',	100,	100,	0);
		`

	productDetailsQuery := `
		INSERT INTO "product_details" ("name", "description", "language_code", "product_id", "created_at", "updated_at", "deleted_at") VALUES
		('تيشرت ',	'تيشرت اسود',	'ar',	1,	'2023-04-23 12:10:52.398698',	'2023-04-23 12:10:52.398709',	'0001-01-01 00:00:00'),
		('T-Shirt',	'Black T-Shirt',	'en',	1,	'2023-04-23 12:10:52.398698',	'2023-04-23 12:10:52.398709',	'0001-01-01 00:00:00');
		`
	connection.Exec(productQuery)
	connection.Exec(productDetailsQuery)

}
