using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Kurdi.ECommerce.Inventory.Infrastructure.Data.Migrations
{
    /// <inheritdoc />
    public partial class SalesOrders : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "sales_order_status",
                columns: table => new
                {
                    id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name = table.Column<string>(type: "text", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_sales_order_status", x => x.id);
                });

            migrationBuilder.CreateTable(
                name: "sales_order",
                columns: table => new
                {
                    id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    totalprice = table.Column<double>(name: "total_price", type: "double precision", nullable: false),
                    discount = table.Column<double>(type: "double precision", nullable: false),
                    statusid = table.Column<int>(name: "status_id", type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_sales_order", x => x.id);
                    table.ForeignKey(
                        name: "FK_sales_order_sales_order_status_status_id",
                        column: x => x.statusid,
                        principalTable: "sales_order_status",
                        principalColumn: "id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "sales_order_product",
                columns: table => new
                {
                    id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    salesorderid = table.Column<int>(name: "sales_order_id", type: "integer", nullable: false),
                    sku = table.Column<string>(type: "text", nullable: true),
                    costpriceperitem = table.Column<double>(name: "cost_price_per_item", type: "double precision", nullable: false),
                    sellingpriceperitem = table.Column<double>(name: "selling_price_per_item", type: "double precision", nullable: false),
                    discountperitem = table.Column<double>(name: "discount_per_item", type: "double precision", nullable: false),
                    sellingpriceperitembeforediscount = table.Column<double>(name: "selling_price_per_item_before_discount", type: "double precision", nullable: false),
                    quantity = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_sales_order_product", x => x.id);
                    table.ForeignKey(
                        name: "FK_sales_order_product_products_sku",
                        column: x => x.sku,
                        principalTable: "products",
                        principalColumn: "sku");
                    table.ForeignKey(
                        name: "FK_sales_order_product_sales_order_sales_order_id",
                        column: x => x.salesorderid,
                        principalTable: "sales_order",
                        principalColumn: "id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_sales_order_status_id",
                table: "sales_order",
                column: "status_id");

            migrationBuilder.CreateIndex(
                name: "IX_sales_order_product_sales_order_id",
                table: "sales_order_product",
                column: "sales_order_id");

            migrationBuilder.CreateIndex(
                name: "IX_sales_order_product_sku",
                table: "sales_order_product",
                column: "sku");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "sales_order_product");

            migrationBuilder.DropTable(
                name: "sales_order");

            migrationBuilder.DropTable(
                name: "sales_order_status");
        }
    }
}
