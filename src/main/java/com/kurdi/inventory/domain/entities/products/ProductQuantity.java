package com.kurdi.inventory.domain.entities.products;

import lombok.*;
import javax.persistence.Embeddable;
import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

@Embeddable
@Data
@AllArgsConstructor
@NoArgsConstructor
public class ProductQuantity implements Serializable {
    private List<StockItem> stock = new ArrayList<>();

    @Setter(AccessLevel.NONE)
    int totalStock = 0;
    @Setter(AccessLevel.NONE)
    int availableStock = 0;
    @Setter(AccessLevel.NONE)
    int reservedStock = 0;

    
    public void addStock(List<StockItem> stock) {

        //add to stock list
        this.stock.addAll(stock);
        //generate sku for all of them
        totalStock += stock.size();
        availableStock += stock.size();
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
