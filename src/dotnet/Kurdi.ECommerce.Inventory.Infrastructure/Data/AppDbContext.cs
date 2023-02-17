using Microsoft.EntityFrameworkCore;
using Kurdi.ECommerce.Inventory.Core.Entities;
using Kurdi.ECommerce.Inventory.Core.Entities.ProductAggregate;
using Microsoft.Extensions.Configuration;
using Kurdi.ECommerce.Inventory.Core.Entities.CategoryAggregate;
using Kurdi.ECommerce.Inventory.Core.Entities.SalesOrderAggregate;

namespace Kurdi.ECommerce.Inventory.Infrastructure.Data
{
    public class AppDbContext : DbContext
    {
        private readonly IConfiguration _configuration;

        public AppDbContext(DbContextOptions<AppDbContext> options, IConfiguration configuration) : base(options)
        {
            _configuration = configuration;
        }
        protected override void OnConfiguring(DbContextOptionsBuilder options)
        {
            // connect to postgres with connection string from app settings
            options.UseNpgsql(_configuration.GetConnectionString("postgresDatabase"));

        }
        protected override void OnModelCreating(ModelBuilder builder)
        {
            builder.Entity<ProductDetails>().HasKey(details => new { details.LanguageCode, details.Sku });
            builder.Entity<CategoryDetails>().HasKey(details => new { details.LanguageCode, details.TranslatedName });
        }
        //protected override void OnConfiguring(DbContextOptionsBuilder options)
        //       => options.UseMySQL(configuration["db_conn"]);
        public DbSet<Language> Languages { get; set; }
        public DbSet<Product> Products { get; set; }
        public DbSet<ProductDetails> ProductDetails { get; set; }
        public DbSet<Category> Categories { get; set; }
        public DbSet<CategoryDetails> CategoriesDetails { get; set; }
        public DbSet<SalesOrder> SalesOrders { get; set; }
        public DbSet<SalesOrderProduct> SalesOrderProducts { get; set; }
        public DbSet<SalesOrderStatus> SalesOrderStatuses { get; set; }

    }
}
/***
     dotnet ef migrations add SalesOrders-naming --context AppDbContext -p ../Kurdi.ECommerce.Inventory.Infrastructure/Kurdi.ECommerce.Inventory.Infrastructure.csproj -o Data/Migrations
     dotnet ef database update  --context AppDbContext -p ../Kurdi.ECommerce.Inventory.Infrastructure/Kurdi.ECommerce.Inventory.Infrastructure.csproj 
**/

