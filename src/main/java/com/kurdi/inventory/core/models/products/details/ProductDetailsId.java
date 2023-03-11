package com.kurdi.inventory.core.models.products.details;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.Hibernate;

import com.kurdi.inventory.core.models.Language;
import com.kurdi.inventory.core.models.products.Product;

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
public class ProductDetailsId implements Serializable {
    @ManyToOne
    @JoinColumn(name = "sku", nullable = false)
    private Product stockItem;
    @ManyToOne
    @JoinColumn(name = "language_code")
    private Language language;


    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        ProductDetailsId that = (ProductDetailsId) o;
        return stockItem != null && Objects.equals(stockItem, that.stockItem)
                && language != null && Objects.equals(language, that.language);
    }

    @Override
    public int hashCode() {
        return Objects.hash(stockItem, language);
    }
}

