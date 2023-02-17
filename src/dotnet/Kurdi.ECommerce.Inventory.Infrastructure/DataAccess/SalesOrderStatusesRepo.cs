using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Kurdi.ECommerce.Inventory.Core.Entities.SalesOrderAggregate;

namespace Kurdi.ECommerce.Inventory.Infrastructure.DataAccess
{
    public class SalesOrderStatusesRepo : RepoBase<SalesOrderStatus> , ISalesOrderStatusesRepo
    {
        public SalesOrderStatusesRepo(AppDbContext db) : base(db)
        {

        }
    }
}
