package com.kurdi.ecommerce.inventoryservice.domain.entities.categories;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.item.StockItem;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Set;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "categories")
public class Category implements Serializable {
    @Id
    String name;
    @ManyToOne
    @JoinColumn(name="parent_category")
    Category parentCategory;
    boolean isParent;
    @JsonIgnore
    @OneToMany(fetch = FetchType.LAZY, mappedBy = "category")
    Set<StockItem> stockItems;
    @OneToMany(fetch = FetchType.EAGER, mappedBy = "category")
    Set<CategoryDetails> categoryDetails;

}
