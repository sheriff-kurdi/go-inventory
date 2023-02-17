using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Kurdi.ECommerce.Inventory.Infrastructure.Data.Migrations
{
    /// <inheritdoc />
    public partial class SalesOrdersnaming : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_sales_order_sales_order_status_status_id",
                table: "sales_order");

            migrationBuilder.DropForeignKey(
                name: "FK_sales_order_product_products_sku",
                table: "sales_order_product");

            migrationBuilder.DropForeignKey(
                name: "FK_sales_order_product_sales_order_sales_order_id",
                table: "sales_order_product");

            migrationBuilder.DropPrimaryKey(
                name: "PK_sales_order_product",
                table: "sales_order_product");

            migrationBuilder.DropPrimaryKey(
                name: "PK_sales_order",
                table: "sales_order");

            migrationBuilder.RenameTable(
                name: "sales_order_product",
                newName: "sales_order_products");

            migrationBuilder.RenameTable(
                name: "sales_order",
                newName: "sales_orders");

            migrationBuilder.RenameIndex(
                name: "IX_sales_order_product_sku",
                table: "sales_order_products",
                newName: "IX_sales_order_products_sku");

            migrationBuilder.RenameIndex(
                name: "IX_sales_order_product_sales_order_id",
                table: "sales_order_products",
                newName: "IX_sales_order_products_sales_order_id");

            migrationBuilder.RenameIndex(
                name: "IX_sales_order_status_id",
                table: "sales_orders",
                newName: "IX_sales_orders_status_id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_sales_order_products",
                table: "sales_order_products",
                column: "id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_sales_orders",
                table: "sales_orders",
                column: "id");

            migrationBuilder.AddForeignKey(
                name: "FK_sales_order_products_products_sku",
                table: "sales_order_products",
                column: "sku",
                principalTable: "products",
                principalColumn: "sku");

            migrationBuilder.AddForeignKey(
                name: "FK_sales_order_products_sales_orders_sales_order_id",
                table: "sales_order_products",
                column: "sales_order_id",
                principalTable: "sales_orders",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_sales_orders_sales_order_status_status_id",
                table: "sales_orders",
                column: "status_id",
                principalTable: "sales_order_status",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_sales_order_products_products_sku",
                table: "sales_order_products");

            migrationBuilder.DropForeignKey(
                name: "FK_sales_order_products_sales_orders_sales_order_id",
                table: "sales_order_products");

            migrationBuilder.DropForeignKey(
                name: "FK_sales_orders_sales_order_status_status_id",
                table: "sales_orders");

            migrationBuilder.DropPrimaryKey(
                name: "PK_sales_orders",
                table: "sales_orders");

            migrationBuilder.DropPrimaryKey(
                name: "PK_sales_order_products",
                table: "sales_order_products");

            migrationBuilder.RenameTable(
                name: "sales_orders",
                newName: "sales_order");

            migrationBuilder.RenameTable(
                name: "sales_order_products",
                newName: "sales_order_product");

            migrationBuilder.RenameIndex(
                name: "IX_sales_orders_status_id",
                table: "sales_order",
                newName: "IX_sales_order_status_id");

            migrationBuilder.RenameIndex(
                name: "IX_sales_order_products_sku",
                table: "sales_order_product",
                newName: "IX_sales_order_product_sku");

            migrationBuilder.RenameIndex(
                name: "IX_sales_order_products_sales_order_id",
                table: "sales_order_product",
                newName: "IX_sales_order_product_sales_order_id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_sales_order",
                table: "sales_order",
                column: "id");

            migrationBuilder.AddPrimaryKey(
                name: "PK_sales_order_product",
                table: "sales_order_product",
                column: "id");

            migrationBuilder.AddForeignKey(
                name: "FK_sales_order_sales_order_status_status_id",
                table: "sales_order",
                column: "status_id",
                principalTable: "sales_order_status",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_sales_order_product_products_sku",
                table: "sales_order_product",
                column: "sku",
                principalTable: "products",
                principalColumn: "sku");

            migrationBuilder.AddForeignKey(
                name: "FK_sales_order_product_sales_order_sales_order_id",
                table: "sales_order_product",
                column: "sales_order_id",
                principalTable: "sales_order",
                principalColumn: "id",
                onDelete: ReferentialAction.Cascade);
        }
    }
}
