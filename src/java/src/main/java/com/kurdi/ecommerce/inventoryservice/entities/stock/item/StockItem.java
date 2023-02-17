package com.kurdi.ecommerce.inventoryservice.entities.stock.item;

import com.kurdi.ecommerce.inventoryservice.entities.categories.Category;
import com.kurdi.ecommerce.inventoryservice.entities.stock.details.StockItemDetails;
import lombok.*;
import org.hibernate.Hibernate;
import org.hibernate.id.GUIDGenerator;

import javax.persistence.*;
import java.io.Serializable;
import java.util.*;


@Getter
@Setter
@ToString
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "stock_items")
public class StockItem implements Serializable {
    @Id
    String sku = new GUIDGenerator().toString();
    Integer SupplierIdentity;
    @Embedded
    StockItemPrices stockItemPrices;
    @Embedded
    StockItemQuantity stockItemQuantity;
    @ManyToOne(fetch = FetchType.EAGER)
    Category category = new Category();
    @OneToMany(mappedBy = "stockItemDetailsId.stockItem", fetch = FetchType.EAGER, cascade = CascadeType.PERSIST)
    List<StockItemDetails> stockItemDetails = new ArrayList<>();


    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        StockItem stockItem = (StockItem) o;
        return sku != null && Objects.equals(sku, stockItem.sku);
    }
    @Override
    public int hashCode() {
        return getClass().hashCode();
    }
}
