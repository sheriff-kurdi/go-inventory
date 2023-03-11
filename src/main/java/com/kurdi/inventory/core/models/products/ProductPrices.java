package com.kurdi.inventory.core.models.products;

import lombok.*;

import javax.persistence.Embeddable;
import java.io.Serializable;

@Embeddable
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class ProductPrices implements Serializable {

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

    public void reCalculateCostPrice(double purchasinPrice, int purchasedQuantity, int totallStockQuantity) {
        costPrice = (costPrice + purchasinPrice) / (totallStockQuantity + purchasedQuantity);

    }

}
