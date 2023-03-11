package com.kurdi.inventory.core.models.products;

import lombok.*;
import org.hibernate.Hibernate;
import org.hibernate.id.GUIDGenerator;

import com.kurdi.inventory.core.models.categories.Category;
import com.kurdi.inventory.core.models.products.details.ProductDetails;

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
public class Product implements Serializable {
    @Id
    @Builder.Default
    String sku = new GUIDGenerator().toString();
    @Embedded
    ProductPrices stockItemPrices;
    @Embedded
    ProductQuantity stockItemQuantity;
    @ManyToOne(fetch = FetchType.EAGER)
    @Builder.Default
    Category category = new Category();
    @OneToMany(mappedBy = "stockItemDetailsId.stockItem", fetch = FetchType.EAGER, cascade = CascadeType.PERSIST)
    @Builder.Default
    private List<ProductDetails> stockItemDetails = new ArrayList<>();


    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o)) return false;
        Product stockItem = (Product) o;
        return sku != null && Objects.equals(sku, stockItem.sku);
    }
    @Override
    public int hashCode() {
        return getClass().hashCode();
    }
}
