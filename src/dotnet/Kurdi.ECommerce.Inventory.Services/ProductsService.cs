using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Core.Entities.ProductAggregate;

namespace Kurdi.ECommerce.Inventory.Services
{
    public class ProductsService
    {
        private readonly IProductsRepo _productsRepo;
        public ProductsService(IProductsRepo productsRepo)
        {
            this._productsRepo = productsRepo;

        }

        public void Reserve(string sku, int quantity)
        {
            Product product = _productsRepo.Find(s => s.SKU == sku).FirstOrDefault();
            product.ProductQuantity.ReserveStock(quantity);
            this._productsRepo.Update(product);
            this._productsRepo.SaveChanges();
        }

        public void CancelReservation(string sku, int quantity)
        {
            Product product = _productsRepo.Find(s => s.SKU == sku).FirstOrDefault();
            product.ProductQuantity.CancelReservation(quantity);
            this._productsRepo.Update(product);
            this._productsRepo.SaveChanges();
        }

        public void AddStock(string sku, int quantity, double addedItemsCost)
        {
            Product product = _productsRepo.Find(s => s.SKU == sku).FirstOrDefault();

            //update quantities
            product.ProductQuantity.AddStock(quantity);

            //update stock cost
            double costOfAllProductStock = product.ProductQuantity.TotalStock * product.ProductPrices.CostPrice;
            double costOfAllProductStockAfterAdding = costOfAllProductStock + addedItemsCost;
            product.ProductPrices.CostPrice = costOfAllProductStockAfterAdding / product.ProductQuantity.TotalStock;

            this._productsRepo.Update(product);
            this._productsRepo.SaveChanges();
        }

        public void Update(Product product)
        {
            this._productsRepo.Update(product);
            this._productsRepo.SaveChanges();
        }

        public void Create(Product product)
        {
            this._productsRepo.Create(product);
            this._productsRepo.SaveChanges();
        }

        public void Delete(string sku)
        {
            Product product = this._productsRepo.Find(stock => stock.SKU == sku).FirstOrDefault();
            _productsRepo.Delete(product);
            this._productsRepo.SaveChanges();
        }

    }
}