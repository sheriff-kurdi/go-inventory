using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Kurdi.ECommerce.Inventory.Core.Entities.SalesOrderAggregate;
using System.Threading.Tasks;
using Kurdi.ECommerce.Inventory.Core.DTOs.SalesOrders;

namespace Kurdi.ECommerce.Inventory.Infrastructure.DataAccess
{
    public class SalesOrdersRepo : RepoBase<SalesOrder> , ISalesOrdersRepo
    {
        public SalesOrdersRepo(AppDbContext db) : base(db)
        {

        }

        public SalesOrder CreateOrder(SalesOrderDTO salesOrderDTO)
        {
            return new SalesOrder();
        }
    }
}
