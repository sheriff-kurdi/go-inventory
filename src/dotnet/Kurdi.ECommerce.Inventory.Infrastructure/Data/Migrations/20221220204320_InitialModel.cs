using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Kurdi.ECommerce.Inventory.Infrastructure.Data.Migrations
{
    /// <inheritdoc />
    public partial class InitialModel : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "categories",
                columns: table => new
                {
                    name = table.Column<string>(type: "text", nullable: false),
                    isparent = table.Column<bool>(name: "is_parent", type: "boolean", nullable: false),
                    parent = table.Column<string>(type: "text", nullable: true),
                    activation = table.Column<bool>(type: "boolean", nullable: false),
                    createdat = table.Column<DateTime>(name: "created_at", type: "timestamp with time zone", nullable: true),
                    updatedat = table.Column<DateTime>(name: "updated_at", type: "timestamp with time zone", nullable: true),
                    deletedat = table.Column<DateTime>(name: "deleted_at", type: "timestamp with time zone", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_categories", x => x.name);
                    table.ForeignKey(
                        name: "FK_categories_categories_parent",
                        column: x => x.parent,
                        principalTable: "categories",
                        principalColumn: "name");
                });

            migrationBuilder.CreateTable(
                name: "languages",
                columns: table => new
                {
                    languagecode = table.Column<string>(name: "language_code", type: "text", nullable: false),
                    languagename = table.Column<string>(name: "language_name", type: "text", nullable: true),
                    activation = table.Column<bool>(type: "boolean", nullable: false),
                    createdat = table.Column<DateTime>(name: "created_at", type: "timestamp with time zone", nullable: true),
                    updatedat = table.Column<DateTime>(name: "updated_at", type: "timestamp with time zone", nullable: true),
                    deletedat = table.Column<DateTime>(name: "deleted_at", type: "timestamp with time zone", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_languages", x => x.languagecode);
                });

            migrationBuilder.CreateTable(
                name: "products",
                columns: table => new
                {
                    sku = table.Column<string>(type: "text", nullable: false),
                    supplieridentity = table.Column<int>(name: "supplier_identity", type: "integer", nullable: false),
                    sellingprice = table.Column<double>(name: "selling_price", type: "double precision", nullable: true),
                    costprice = table.Column<double>(name: "cost_price", type: "double precision", nullable: true),
                    discount = table.Column<double>(type: "double precision", nullable: true),
                    isdiscounted = table.Column<bool>(name: "is_discounted", type: "boolean", nullable: true),
                    totalstock = table.Column<int>(name: "total_stock", type: "integer", nullable: true),
                    availablestock = table.Column<int>(name: "available_stock", type: "integer", nullable: true),
                    reservedstock = table.Column<int>(name: "reserved_stock", type: "integer", nullable: true),
                    category = table.Column<string>(type: "text", nullable: true),
                    activation = table.Column<bool>(type: "boolean", nullable: false),
                    createdat = table.Column<DateTime>(name: "created_at", type: "timestamp with time zone", nullable: true),
                    updatedat = table.Column<DateTime>(name: "updated_at", type: "timestamp with time zone", nullable: true),
                    deletedat = table.Column<DateTime>(name: "deleted_at", type: "timestamp with time zone", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_products", x => x.sku);
                    table.ForeignKey(
                        name: "FK_products_categories_category",
                        column: x => x.category,
                        principalTable: "categories",
                        principalColumn: "name");
                });

            migrationBuilder.CreateTable(
                name: "categories_details",
                columns: table => new
                {
                    languagecode = table.Column<string>(name: "language_code", type: "text", nullable: false),
                    translatedname = table.Column<string>(name: "translated_name", type: "text", nullable: false),
                    name = table.Column<string>(type: "text", nullable: true),
                    description = table.Column<string>(type: "text", nullable: true),
                    createdat = table.Column<DateTime>(name: "created_at", type: "timestamp with time zone", nullable: true),
                    updatedat = table.Column<DateTime>(name: "updated_at", type: "timestamp with time zone", nullable: true),
                    deletedat = table.Column<DateTime>(name: "deleted_at", type: "timestamp with time zone", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_categories_details", x => new { x.languagecode, x.translatedname });
                    table.ForeignKey(
                        name: "FK_categories_details_categories_name",
                        column: x => x.name,
                        principalTable: "categories",
                        principalColumn: "name");
                    table.ForeignKey(
                        name: "FK_categories_details_languages_language_code",
                        column: x => x.languagecode,
                        principalTable: "languages",
                        principalColumn: "language_code",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "product_details",
                columns: table => new
                {
                    languagecode = table.Column<string>(name: "language_code", type: "text", nullable: false),
                    sku = table.Column<string>(type: "text", nullable: false),
                    name = table.Column<string>(type: "text", nullable: true),
                    description = table.Column<string>(type: "text", nullable: true),
                    createdat = table.Column<DateTime>(name: "created_at", type: "timestamp with time zone", nullable: true),
                    updatedat = table.Column<DateTime>(name: "updated_at", type: "timestamp with time zone", nullable: true),
                    deletedat = table.Column<DateTime>(name: "deleted_at", type: "timestamp with time zone", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_product_details", x => new { x.languagecode, x.sku });
                    table.ForeignKey(
                        name: "FK_product_details_languages_language_code",
                        column: x => x.languagecode,
                        principalTable: "languages",
                        principalColumn: "language_code",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_product_details_products_sku",
                        column: x => x.sku,
                        principalTable: "products",
                        principalColumn: "sku",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_categories_parent",
                table: "categories",
                column: "parent");

            migrationBuilder.CreateIndex(
                name: "IX_categories_details_name",
                table: "categories_details",
                column: "name");

            migrationBuilder.CreateIndex(
                name: "IX_product_details_sku",
                table: "product_details",
                column: "sku");

            migrationBuilder.CreateIndex(
                name: "IX_products_category",
                table: "products",
                column: "category");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "categories_details");

            migrationBuilder.DropTable(
                name: "product_details");

            migrationBuilder.DropTable(
                name: "languages");

            migrationBuilder.DropTable(
                name: "products");

            migrationBuilder.DropTable(
                name: "categories");
        }
    }
}
