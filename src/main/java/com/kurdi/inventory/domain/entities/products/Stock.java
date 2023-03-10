package com.kurdi.inventory.domain.entities.products;

import java.io.Serializable;

import javax.persistence.Entity;
import javax.persistence.Table;

import lombok.*;

@Getter
@Setter
@ToString
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "stocks")
public class Stock implements Serializable {
    int id;
    @Builder.Default
    ProductSKU sku = new ProductSKU();
    int purchaseReciept;
    int purchasePrice;

}
