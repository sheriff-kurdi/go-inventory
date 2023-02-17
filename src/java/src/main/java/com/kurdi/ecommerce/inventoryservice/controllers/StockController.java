package com.kurdi.ecommerce.inventoryservice.controllers;

import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItem;
import com.kurdi.ecommerce.inventoryservice.projections.StockItemProjection;
import com.kurdi.ecommerce.inventoryservice.requests.stock.CreateStockItemRequest;
import com.kurdi.ecommerce.inventoryservice.responses.CustomResponse;
import com.kurdi.ecommerce.inventoryservice.services.StockItemsService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;


@SuppressWarnings("unused")
@RestController
@RequestMapping("/stock/")
public class StockController {
    final StockItemsService stockItemsService;

    public StockController(StockItemsService stockItemsService) {
        this.stockItemsService = stockItemsService;
    }

    @GetMapping("/")
    public ResponseEntity<HashMap<String, Object>> getAll(@RequestHeader(value="language", required = false) String languageCode
            , @RequestParam(value = "page") int page, @RequestParam(value = "pageSize") int pageSize)
    {
        Page<StockItemProjection> stockItemProjections = stockItemsService.getAll(languageCode, page, pageSize);
        return CustomResponse.Ok(stockItemProjections);
    }

    @GetMapping("{sku}")
    public ResponseEntity<HashMap<String, Object>> getBySKU( @PathVariable String sku, @RequestHeader(value="language", required = false) String languageCode)
    {
        StockItemProjection stockItemProjection = stockItemsService.getBySKU(sku, languageCode);
        if(stockItemProjection == null)
        {
            return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);
        }
        return CustomResponse.Ok(stockItemProjection);
    }

    @PostMapping
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<StockItem> create(@RequestBody CreateStockItemRequest stockItemRequest)
    {
        StockItem stockItem = stockItemsService.create(stockItemRequest);
        return new ResponseEntity<>(stockItem, HttpStatus.CREATED);
    }

    @PutMapping("{sku}")
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<StockItem> edit(@RequestBody StockItem stockItem, @PathVariable String sku)
    {
        if(!stockItem.getSku().equals(sku))
        {
            return new ResponseEntity<>(stockItem, HttpStatus.BAD_REQUEST);
        }
        stockItem = stockItemsService.edit(stockItem, sku);
        if (stockItem == null)
        {
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
        }
        return new ResponseEntity<>(stockItem, HttpStatus.OK);
    }

    @DeleteMapping("{sku}")
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<StockItem> delete(@PathVariable String sku)
    {
        StockItem stockItem = stockItemsService.delete(sku);
        if (stockItem == null)
        {
            return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(stockItem, HttpStatus.OK);
    }

}
