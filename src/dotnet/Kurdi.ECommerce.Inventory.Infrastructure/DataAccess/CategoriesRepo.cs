using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Kurdi.ECommerce.Inventory.Core.Entities.ProductAggregate;
using Kurdi.ECommerce.Inventory.Core.Entities.CategoryAggregate;

namespace Kurdi.ECommerce.Inventory.Infrastructure.DataAccess
{
    public class CategoriesRepo : RepoBase<Category> , ICategoriesRepo
    {
        public CategoriesRepo(AppDbContext db) : base(db)
        {

        }
    }
}
