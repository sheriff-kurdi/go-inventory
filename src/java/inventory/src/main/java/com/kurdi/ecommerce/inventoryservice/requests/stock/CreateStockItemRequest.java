package com.kurdi.ecommerce.inventoryservice.requests.stock;


import com.kurdi.ecommerce.inventoryservice.entities.Language;
import com.kurdi.ecommerce.inventoryservice.entities.stock.details.StockItemDetails;
import com.kurdi.ecommerce.inventoryservice.entities.stock.details.StockItemDetailsId;
import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItem;
import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItemPrices;
import lombok.Builder;
import lombok.Data;

import java.io.Serializable;
import java.util.List;
import java.util.stream.Collectors;

@Data
@Builder
public class CreateStockItemRequest implements Serializable {
    String sku;
    List<StockItemDetailsRequest> details;
    int stockQuantity;
    String category;
    StockItemPrices prices;

    public List<StockItemDetails> getStockItemDetails(){
        return details.stream().map(
                details -> StockItemDetails.builder()
                                            .stockItemDetailsId(StockItemDetailsId.builder()
                                                    .stockItem(StockItem.builder().sku(sku).build())
                                                    .language(Language.builder().languageCode(details.languageCode).build())
                                                    .build())
                                            .name(details.name)
                                            .description(details.description)
                                            .build()).collect(Collectors.toList());
    }
}

