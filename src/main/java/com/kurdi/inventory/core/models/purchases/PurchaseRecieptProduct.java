package com.kurdi.inventory.core.models.purchases;

import lombok.*;
import javax.persistence.*;



@Getter
@Setter
@ToString
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "purchase_reciepts_products")
public class PurchaseRecieptProduct {
    int id;
    String sku;
    int quantity;
    double costPrice;
}
