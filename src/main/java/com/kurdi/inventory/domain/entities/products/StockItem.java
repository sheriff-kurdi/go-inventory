package com.kurdi.inventory.domain.entities.products;

import java.io.Serializable;

import javax.persistence.Entity;
import javax.persistence.Table;

import com.kurdi.inventory.domain.contracts.products.SKU;

import lombok.*;

@Getter
@Setter
@ToString
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "stocks")
public class StockItem implements Serializable {
    int id;
    SKU sku;
    int purchaseReciept;
    int purchasePrice;

}
