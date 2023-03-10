package com.kurdi.inventory.domain.entities.products.details;

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
@Table(name = "product_details")
public class ProductDetails implements Serializable {
    @EmbeddedId
    ProductDetailsId stockItemDetailsId;
    String name;
    String description;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        ProductDetails that = (ProductDetails) o;
        return stockItemDetailsId != null && Objects.equals(stockItemDetailsId, that.stockItemDetailsId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(stockItemDetailsId);
    }
}
