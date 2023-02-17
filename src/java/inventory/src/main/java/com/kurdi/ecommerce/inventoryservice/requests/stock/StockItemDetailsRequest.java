package com.kurdi.ecommerce.inventoryservice.requests.stock;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class StockItemDetailsRequest {
    String languageCode;
    String name;
    String description;
}
