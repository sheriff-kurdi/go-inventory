package com.kurdi.inventory.domain.contracts.products;

import java.io.Serializable;

public interface SKU extends Serializable {
    void skuDecode(String sku);
    public String generateSKU();
}
