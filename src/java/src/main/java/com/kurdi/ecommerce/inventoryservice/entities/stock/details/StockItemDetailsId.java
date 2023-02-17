package com.kurdi.ecommerce.inventoryservice.entities.stock.details;

import com.kurdi.ecommerce.inventoryservice.entities.Language;
import com.kurdi.ecommerce.inventoryservice.entities.stock.item.StockItem;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.Hibernate;

import javax.persistence.Embeddable;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import java.io.Serializable;
import java.util.Objects;

@Builder
@AllArgsConstructor
@NoArgsConstructor
@Embeddable
@Data
public class StockItemDetailsId implements Serializable {
    @ManyToOne
    @JoinColumn(name = "sku", nullable = false)
    private StockItem stockItem;
    @ManyToOne
    @JoinColumn(name = "language_code")
    private Language language;


    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        StockItemDetailsId that = (StockItemDetailsId) o;
        return stockItem != null && Objects.equals(stockItem, that.stockItem)
                && language != null && Objects.equals(language, that.language);
    }

    @Override
    public int hashCode() {
        return Objects.hash(stockItem, language);
    }
}

