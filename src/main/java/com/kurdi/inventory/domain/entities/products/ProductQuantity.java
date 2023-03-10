package com.kurdi.inventory.domain.entities.products;

import lombok.*;
import javax.persistence.Embeddable;
import java.io.Serializable;

@Embeddable
@Data
@AllArgsConstructor
@NoArgsConstructor
public class ProductQuantity implements Serializable {
    @Setter(AccessLevel.NONE)
    int totalStock = 0;
    @Setter(AccessLevel.NONE)
    int availableStock = 0;
    @Setter(AccessLevel.NONE)
    int reservedStock = 0;

    public void addStock(int quantity) {
        totalStock += quantity;
        availableStock += quantity;
    }

    public void reserveStock(int quantity) {
        availableStock -= quantity;
        reservedStock += quantity;
    }

    public void cancelReservation(int quantity) {
        availableStock += quantity;
        reservedStock -= quantity;
    }

    public void selling(int quantity) {
        availableStock -= quantity;
        totalStock -= quantity;
    }
}