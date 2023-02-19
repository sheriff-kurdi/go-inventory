package com.kurdi.ecommerce.inventoryservice.domain.entities.stock.details;

import lombok.*;
import org.hibernate.Hibernate;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Objects;

@Getter
@Setter
@ToString
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "stock_items_details")
public class StockItemDetails implements Serializable {
    @EmbeddedId
    StockItemDetailsId stockItemDetailsId;
    String name;
    String description;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        StockItemDetails that = (StockItemDetails) o;
        return stockItemDetailsId != null && Objects.equals(stockItemDetailsId, that.stockItemDetailsId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(stockItemDetailsId);
    }
}
