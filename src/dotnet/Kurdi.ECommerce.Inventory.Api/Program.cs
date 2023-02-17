
using Kurdi.ECommerce.Inventory.Api.Helpers;
using Kurdi.ECommerce.Inventory.Api.Middleware;
using Kurdi.ECommerce.Inventory.Api.Routes;
using Kurdi.ECommerce.Inventory.Core.Contracts;
using Kurdi.ECommerce.Inventory.Infrastructure.Data;
using Kurdi.ECommerce.Inventory.Infrastructure.DataAccess;
using Kurdi.ECommerce.Inventory.Services;


var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddCors();

builder.Services.AddDbContext<AppDbContext>();
builder.Services.AddScoped<IProductsRepo, ProductsRepo>();
builder.Services.AddScoped<ICategoriesRepo, CategoriesRepo>();
builder.Services.AddScoped<ISalesOrdersRepo, SalesOrdersRepo>();
builder.Services.AddScoped<ISalesOrderProductsRepo, SalesOrderProductsRepo>();
builder.Services.AddScoped<ProductsService>();
builder.Services.AddScoped<SalesOrdersService>();
builder.Services.AddLocalization();

var app = builder.Build();


app.UseSwagger();
app.UseSwaggerUI();
app.UseCors(cors => { cors.AllowAnyOrigin().AllowAnyMethod().AllowAnyHeader(); });

app.UseLanguageMiddleware();

app.UseStockEndPoints();
app.UseCategoriesEndPoints();
app.UseSalesOrdersEndPoints();

app.MapGet("/", () =>
{    return Translator.Translate("VALIDATION:NOT_VALID_LANGUAGE");
});
app.Run();