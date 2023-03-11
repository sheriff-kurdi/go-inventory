package com.kurdi.inventory.core.models.purchases;

import lombok.*;
import java.util.ArrayList;
import java.util.List;
import javax.persistence.*;

@Getter
@Setter
@ToString
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "purchase_reciepts")
public class PurchaseReciept {
    int id;
    @Builder.Default
    private List<PurchaseRecieptProduct> products = new ArrayList<>();
    double costPrice;
}
