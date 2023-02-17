
using System.Threading.Tasks;
using Kurdi.ECommerce.Inventory.Core.DTOs.SalesOrders;
using Kurdi.ECommerce.Inventory.Core.Entities.SalesOrderAggregate;
namespace Kurdi.ECommerce.Inventory.Core.Contracts
{
    public interface ISalesOrdersRepo : IRepoBase<SalesOrder>
    {
        SalesOrder CreateOrder(SalesOrderDTO salesOrderDTO);
    }
}
