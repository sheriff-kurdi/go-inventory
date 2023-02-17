package com.kurdi.ecommerce.inventoryservice.entities.stock.item;

import lombok.*;

import javax.persistence.Embeddable;
import javax.persistence.Embedded;
import java.io.Serializable;

@Embeddable
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class StockItemPrices implements Serializable {

    double costPrice;
    @Getter(AccessLevel.NONE)
    double sellingPrice;
    boolean isDiscounted;
    double discount;

    public double getSellingPrice() {
        if (isDiscounted) {
            return sellingPrice - discount;
        }
        return sellingPrice;
    }

}
