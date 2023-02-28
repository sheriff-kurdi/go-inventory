package com.kurdi.ecommerce.inventoryservice.domain.entities.products;

import com.kurdi.ecommerce.inventoryservice.domain.entities.categories.Category;
import com.kurdi.ecommerce.inventoryservice.domain.entities.products.details.ProductDetails;
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
public class Product implements Serializable {
    @Id
    @Builder.Default
    String sku = new GUIDGenerator().toString();
    Integer SupplierIdentity;
    @Embedded
    ProductPrices stockItemPrices;
    @Embedded
    ProductQuantity stockItemQuantity;
    @ManyToOne(fetch = FetchType.EAGER)
    @Builder.Default
    Category category = new Category();
    @OneToMany(mappedBy = "stockItemDetailsId.stockItem", fetch = FetchType.EAGER, cascade = CascadeType.PERSIST)
    @Builder.Default
    List<ProductDetails> stockItemDetails = new ArrayList<>();


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
