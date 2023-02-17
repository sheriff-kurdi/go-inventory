using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Kurdi.ECommerce.Inventory.Core.Entities.ProductAggregate;

namespace Kurdi.ECommerce.Inventory.Infrastructure.DataAccess
{
    public class ProductsRepo : RepoBase<Product> , IProductsRepo
    {
        public ProductsRepo(AppDbContext db) : base(db)
        {

        }
    }
}
