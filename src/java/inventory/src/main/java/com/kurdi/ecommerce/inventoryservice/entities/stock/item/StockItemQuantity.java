package com.kurdi.ecommerce.inventoryservice.entities.stock.item;

import lombok.*;
import javax.persistence.Embeddable;
import java.io.Serializable;

@Embeddable
@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class StockItemQuantity implements Serializable {
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
}
