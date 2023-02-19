package com.kurdi.ecommerce.inventoryservice.infrastructure.repositories;

import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.item.StockItem;
import com.kurdi.ecommerce.inventoryservice.infrastructure.projections.StockItemProjection;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.repository.query.Param;

public interface StockItemsRepository extends PagingAndSortingRepository<StockItem, String> {
    @Query(value =
            "select stock.sku as sku, stock.available_stock as availableStock, stock.reserved_stock as ReservedStock, stock.cost_price as CostPrice" +
                    ", stock.selling_price as SellingPrice, stock.is_discounted as IsDiscounted ,stock.total_stock as TotalStock" +
                    ", stock.discount, stock.supplier_identity as SupplierIdentity, details.name, details.description, stock.category_name as CategoryName " +
                    "FROM  stock_items as stock " +
                    "join stock_items_details as details on details.sku =  stock.sku and details.language_code = :languageCode ", nativeQuery = true)
    Page<StockItemProjection> getAll(@Param("languageCode") String languageCode, Pageable pageable);

    @Query(value =
            "select stock.sku as sku, stock.available_stock as availableStock, stock.reserved_stock as ReservedStock, stock.cost_price as CostPrice" +
                    ", stock.discount, stock.supplier_identity as SupplierIdentity, details.name, details.description, stock.category_name as CategoryName" +
                    ", stock.selling_price as SellingPrice, stock.is_discounted as IsDiscounted ,stock.total_stock as TotalStock" +
                    " FROM  stock_items as stock " +
                    " join stock_items_details as details on details.sku =  stock.sku and details.language_code = :languageCode" +
                    " where stock.sku =  :SKU ;", nativeQuery = true)
    StockItemProjection getBySKU(@Param("SKU") String SKU, @Param("languageCode") String languageCode);




}
