package com.kurdi.ecommerce.inventoryservice.services;

import com.kurdi.ecommerce.inventoryservice.entities.categories.Category;
import com.kurdi.ecommerce.inventoryservice.entities.stock.details.StockItemDetails;
import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItem;
import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItemQuantity;
import com.kurdi.ecommerce.inventoryservice.projections.StockItemProjection;
import com.kurdi.ecommerce.inventoryservice.repositories.StockItemDetailsRepository;
import com.kurdi.ecommerce.inventoryservice.repositories.StockItemsRepository;
import com.kurdi.ecommerce.inventoryservice.requests.stock.CreateStockItemRequest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;


import javax.transaction.Transactional;
import java.util.ArrayList;
import java.util.List;

@Service
public class StockItemsService {
    final StockItemsRepository stockItemsRepository;
    final StockItemDetailsRepository stockItemDetailsRepository;

    public StockItemsService(StockItemsRepository stockItemsRepository, StockItemDetailsRepository stockItemDetailsRepository) {
        this.stockItemsRepository = stockItemsRepository;
        this.stockItemDetailsRepository = stockItemDetailsRepository;
    }

    public Page<StockItemProjection> getAll(String languageCode, int page, int pageSize){

        Pageable sortedBySKUDesc =
                PageRequest.of(page, pageSize, Sort.by("sku").descending());

            return stockItemsRepository.getAll(languageCode == null ? "ar" : languageCode, sortedBySKUDesc);
    }

    public StockItemProjection getBySKU(String sku, String languageCode){

        if(!stockItemsRepository.existsById(sku))
        {
            return null;
        }
        return stockItemsRepository.getBySKU(sku, languageCode == null ? "ar" : languageCode);
    }

    @Transactional
    public StockItem create(CreateStockItemRequest stockItemRequest)
    {
        //TODO: implement authorization
        /*
            Authentication auth = SecurityContextHolder.getContext().getAuthentication();
            Integer identity = Integer.parseInt(auth.getPrincipal().toString());
        **/
        StockItemQuantity stockItemQuantity = new StockItemQuantity();
        stockItemQuantity.addStock(stockItemRequest.getStockQuantity());

        StockItem stockItem = StockItem.builder()
                //.SupplierIdentity(identity)
                .sku(stockItemRequest.getSku())
                //.stockItemDetails(stockItemRequest.getStockItemDetails())
                .stockItemQuantity(stockItemQuantity)
                .category(Category.builder().name(stockItemRequest.getCategory()).build())
                .stockItemPrices(stockItemRequest.getPrices())
                .build();

        stockItemsRepository.save(stockItem);
        stockItemDetailsRepository.saveAll(stockItemRequest.getStockItemDetails());
        stockItem.setStockItemDetails(stockItemRequest.getStockItemDetails());
        stockItemsRepository.save(stockItem);

        return stockItem;
    }

    public StockItem edit(StockItem stockItem, String sku)
    {
        if(!stockItem.getSku().equals(sku))
        {
            return null;
        }
        if(!stockItemsRepository.existsById(sku))
        {
            return null;
        }
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        Integer identity = Integer.parseInt(auth.getPrincipal().toString());
        if(!stockItemsRepository.getBySKU(sku,"ar").getSupplierIdentity().equals(identity))
        {
            return null;
        }
        return stockItemsRepository.save(stockItem);
    }

    public StockItem delete(String sku)
    {
        if(!stockItemsRepository.existsById(sku))
        {
            return null;

        }
        StockItem stockItem = stockItemsRepository.findById(sku).get();

        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        Integer identity = Integer.parseInt(auth.getPrincipal().toString());
        if(!stockItemsRepository.findById(sku).get().getSupplierIdentity().equals(identity))
        {
            //TODO: return un authorized exception error and then handle exception
            return null;

        }
        stockItemsRepository.delete(stockItem);

        return stockItem;
    }




}
