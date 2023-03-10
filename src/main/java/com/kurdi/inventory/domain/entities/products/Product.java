package com.kurdi.inventory.domain.entities.products;
import com.kurdi.inventory.domain.entities.categories.Category;
import com.kurdi.inventory.domain.entities.products.details.ProductDetails;
import lombok.*;
import org.hibernate.Hibernate;
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
@Table(name = "products")
public class Product implements Serializable {
    @Id
    String id;
    @Builder.Default
    private List<Stock> stock = new ArrayList<>();
    Integer supplierIdentity;
    @Embedded
    ProductPrices stockItemPrices;
    @Embedded
    ProductQuantity stockItemQuantity;
    @ManyToOne(fetch = FetchType.EAGER)
    @Builder.Default
    Category category = new Category();
    @OneToMany(mappedBy = "stockItemDetailsId.stockItem", fetch = FetchType.EAGER, cascade = CascadeType.PERSIST)
    @Builder.Default
    private List<ProductDetails> productDetails = new ArrayList<>();

    @Override
    public boolean equals(Object o) {
        if (this == o)
            return true;
        if (o == null || Hibernate.getClass(this) != Hibernate.getClass(o))
            return false;
        Product product = (Product) o;

        for (Stock stockItem : product.stock) {
            if (this.stock.stream().anyMatch(s -> s.sku.generateSKU() == stockItem.sku.generateSKU())) {
                return true;
            }
        }
        return false;
    }

    @Override
    public int hashCode() {
        return getClass().hashCode();
    }
}
