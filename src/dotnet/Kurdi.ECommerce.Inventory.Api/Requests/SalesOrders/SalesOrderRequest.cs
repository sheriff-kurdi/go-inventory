using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Kurdi.ECommerce.Inventory.Core.DTOs.SalesOrders;

namespace Kurdi.ECommerce.Inventory.Api.Requests.SalesOrders
{
    public class SalesOrderRequest
    {
        public List<SalesOrderItemDTO> SalesOrderItems { get; set; } = new List<SalesOrderItemDTO>();
    }
}