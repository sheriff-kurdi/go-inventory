package com.kurdi.inventory.domain.entities.products;

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

}
