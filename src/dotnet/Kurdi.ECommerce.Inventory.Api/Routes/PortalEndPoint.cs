
using Kurdi.ECommerce.Inventory.Api.Requests.SalesOrders;
using Kurdi.ECommerce.Inventory.Api.Requests.Stock;
using Kurdi.ECommerce.Inventory.Api.Responses;
using Kurdi.ECommerce.Inventory.Api.Responses.Categories;
using Kurdi.ECommerce.Inventory.Api.Responses.Stock;
using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Core.DTOs.SalesOrders;
using Kurdi.ECommerce.Inventory.Core.Entities.CategoryAggregate;
using Kurdi.ECommerce.Inventory.Core.Entities.ProductAggregate;
using Kurdi.ECommerce.Inventory.Core.Entities.SalesOrderAggregate;
using Kurdi.ECommerce.Inventory.Services;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace Kurdi.ECommerce.Inventory.Api.Routes
{
    public static class PortalEndPoint
    {
        public static void UseStockEndPoints(this WebApplication app)
        {
            RouteGroupBuilder stockGroup = app.MapGroup("/api/stock").WithTags("Stock");


            stockGroup.MapGet("/", async ([AsParameters] StockItemsRequestParameters requestParameters, [FromServices] IProductsRepo stockItemsRepository) =>
            {
                var stockItems = stockItemsRepository.FindAll()
                .Include(s => s.ProductDetails)
                .Skip(requestParameters.PageNumber - 1)
                .Take(requestParameters.PageSize)
                .Select(s => new StockItemResponse
                {
                    Sku = s.SKU,
                    Name = s.ProductDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Name,
                    Description = s.ProductDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Description,
                    Category = s.CategoryName,
                    ProductPrices = s.ProductPrices,
                    ProductQuantity = s.ProductQuantity,


                });


                if (!string.IsNullOrEmpty(requestParameters.Sku))
                {
                    stockItems = stockItems.Where(s => s.Sku == requestParameters.Sku);
                }
                if (!string.IsNullOrEmpty(requestParameters.Category))
                {
                    stockItems = stockItems.Where(s => s.Category == requestParameters.Category);
                }
                if (!string.IsNullOrEmpty(requestParameters.Name))
                {
                    stockItems = stockItems.Where(s => s.Name == requestParameters.Name);
                }
                if (!string.IsNullOrEmpty(requestParameters.Query))
                {

                    stockItems = stockItems.Where(s => s.Sku == requestParameters.Query || s.Name.Contains(requestParameters.Query));
                }

                return Results.Ok(new Responses.PaginatedResponse<List<StockItemResponse>>(await stockItems.ToListAsync(), requestParameters.PageNumber, requestParameters.PageSize, stockItemsRepository.Count()));
            });

            stockGroup.MapGet("/{sku}", async (string sku, [FromServices] IProductsRepo stockItemsRepository) =>
            {
                Product? stockItem = await stockItemsRepository.Find(stock => stock.SKU == sku).Include(stock => stock.ProductDetails).FirstOrDefaultAsync();
                if (stockItem == null) return Results.NotFound();
                return Results.Ok(new Responses.BaseResponse<Product>(stockItem));
            });

            stockGroup.MapPut("/{sku}", async (string sku, [FromBody] Product updatedStockItem, [FromServices] IProductsRepo stockItemsRepository) =>
            {
                if (sku != updatedStockItem.SKU) return Results.BadRequest();
                Product? actualStock = await stockItemsRepository.Find(stock => stock.SKU == sku).FirstOrDefaultAsync();
                if (actualStock == null) return Results.NotFound();
                updatedStockItem.ProductQuantity = actualStock.ProductQuantity;
                stockItemsRepository.Update(updatedStockItem);
                await stockItemsRepository.SaveChanges();
                return Results.Ok(new Responses.BaseResponse<Product>(updatedStockItem));
            });

            stockGroup.MapPost("/", async ([FromBody] Product stockItem, [FromServices] IProductsRepo stockItemsRepository) =>
            {
                await stockItemsRepository.Create(stockItem);
                return Results.Created(stockItem.SKU, stockItem);
            });

            stockGroup.MapPost("/reserve", ([FromBody] ReserveStockRequest request, [FromServices] ProductsService stockService) =>
            {
                try
                {
                    stockService.Reserve(request.SKU, request.Quantity);
                    return Results.Ok(new Responses.BaseResponse<string>(""));
                }
                catch (Exception ex)
                {
                    string[] errors = { ex.Message };
                    return Results.BadRequest(new Responses.BaseResponse<string>(data: "", errors: errors, succeeded: false));

                }
            });

            stockGroup.MapPost("/cancel-reservation", ([FromBody] CancelStockReservationRequest request, [FromServices] ProductsService stockService) =>
            {
                try
                {
                    stockService.CancelReservation(request.SKU, request.Quantity);
                    return Results.Ok(new Responses.BaseResponse<string>(""));
                }
                catch (Exception ex)
                {
                    string[] errors = { ex.Message };
                    return Results.BadRequest(new Responses.BaseResponse<string>(data: "", errors: errors, succeeded: false));

                }
            });

            stockGroup.MapPost("/add-stock", ([FromBody] AddStockRequest request, [FromServices] ProductsService stockService) =>
            {
                try
                {
                    stockService.AddStock(request.SKU, request.Quantity, request.TotalCost);
                    return Results.Ok(new Responses.BaseResponse<string>(""));
                }
                catch (Exception ex)
                {
                    string[] errors = { ex.Message };
                    return Results.BadRequest(new Responses.BaseResponse<string>(data: "", errors: errors, succeeded: false));

                }
            });

            stockGroup.MapDelete("/{sku}", (string sku, [FromServices] ProductsService stockService) =>
            {
                stockService.Delete(sku);
                return Results.NoContent();
            });

        }

        public static void UseCategoriesEndPoints(this WebApplication app)
        {
            RouteGroupBuilder categoriesGroup = app.MapGroup("/api/categories").WithTags("Categories");


            categoriesGroup.MapGet("/", async ([AsParameters] CategoriesRequestParameters requestParameters, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                var categories = categoriesRepository.FindAll()
                .Include(c => c.CategoryDetails)
                .Skip(requestParameters.PageNumber - 1)
                .Take(requestParameters.PageSize)
                .Select(c => new CategoryResponse()
                {
                    Name = c.Name,
                    TranslatedName = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.TranslatedName,
                    Description = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Description,
                    IsParent = c.IsParent,
                    Parent = c.ParentName
                });


                if (!string.IsNullOrEmpty(requestParameters.Query))
                {
                    categories = categories.Where(c => c.Name == requestParameters.Query.ToUpper().Trim() || c.TranslatedName.Contains(requestParameters.Query));
                }

                return Results.Ok(new Responses.PaginatedResponse<List<CategoryResponse>>(await categories.ToListAsync(), requestParameters.PageNumber, requestParameters.PageSize, categoriesRepository.Count()));
            });

            categoriesGroup.MapGet("/active", async ([FromServices] ICategoriesRepo categoriesRepository) =>
            {
                var categories = categoriesRepository.FindAll()
                .Include(c => c.CategoryDetails)
                .Where(c => c.Activation)
                .Select(c => new CategoryResponse()
                {
                    Name = c.Name,
                    TranslatedName = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.TranslatedName,
                    Description = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Description,
                    IsParent = c.IsParent,
                    Parent = c.ParentName
                });

                return Results.Ok(new Responses.BaseResponse<List<CategoryResponse>>(await categories.ToListAsync()));
            });

            categoriesGroup.MapGet("/{name}", async (string name, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                CategoryResponse? category = await categoriesRepository.Find(category => category.Name == name)
                    .Include(category => category.CategoryDetails)
                    .Select(c => new CategoryResponse()
                    {
                        Name = c.Name,
                        TranslatedName = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.TranslatedName,
                        Description = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Description,
                        IsParent = c.IsParent,
                        Parent = c.ParentName
                    }).FirstOrDefaultAsync();
                if (category == null) return Results.NotFound();

                return Results.Ok(new Responses.BaseResponse<CategoryResponse>(category));
            });

            categoriesGroup.MapPut("/{name}", async (string name, [FromBody] Category category, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                if (name != category.Name) return Results.BadRequest();
                Category? actualCategory = await categoriesRepository.Find(stock => stock.Name == name).FirstOrDefaultAsync();
                if (actualCategory == null) return Results.NotFound();
                categoriesRepository.Update(category);
                await categoriesRepository.SaveChanges();
                return Results.Ok(new Responses.BaseResponse<Category>(category));
            });

            categoriesGroup.MapPost("/", async ([FromBody] Category category, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                await categoriesRepository.Create(category);
                return Results.Created(category.Name, category);
            });

            categoriesGroup.MapDelete("/{name}", async (string name, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                await categoriesRepository.Find(stock => stock.Name == name).ExecuteDeleteAsync();
                return Results.NoContent();
            });

        }

        public static void UseSalesOrdersEndPoints(this WebApplication app)
        {
            RouteGroupBuilder salesOrdersGroup = app.MapGroup("/api/salesOrders").WithTags("Sales Orders");


            salesOrdersGroup.MapGet("/", async ([AsParameters] CategoriesRequestParameters requestParameters, [FromServices] ICategoriesRepo categoriesRepository) =>
            {
                var categories = categoriesRepository.FindAll()
                .Include(c => c.CategoryDetails)
                .Skip(requestParameters.PageNumber - 1)
                .Take(requestParameters.PageSize)
                .Select(c => new CategoryResponse()
                {
                    Name = c.Name,
                    TranslatedName = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.TranslatedName,
                    Description = c.CategoryDetails.FirstOrDefault(d => d.LanguageCode == "ar")!.Description,
                    IsParent = c.IsParent,
                    Parent = c.ParentName
                });


                if (!string.IsNullOrEmpty(requestParameters.Query))
                {
                    categories = categories.Where(c => c.Name == requestParameters.Query.ToUpper().Trim() || c.TranslatedName.Contains(requestParameters.Query));
                }

                return Results.Ok(new Responses.PaginatedResponse<List<CategoryResponse>>(await categories.ToListAsync(), requestParameters.PageNumber, requestParameters.PageSize, categoriesRepository.Count()));
            });


            salesOrdersGroup.MapPost("/", ([FromBody] SalesOrderRequest salesOrderRequest, [FromServices] SalesOrdersService salesOrdersService) =>
            {
                SalesOrder salesOrder = salesOrdersService.CreateOrder(new SalesOrderDTO(salesOrderRequest.SalesOrderItems));
                return new BaseResponse<SalesOrder>(salesOrder);
            });

     

        }

    }
}