package com.kurdi.inventory.domain.entities.products;

import lombok.*;
import javax.persistence.Embeddable;
import java.io.Serializable;

@Embeddable
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class ProductSKU implements Serializable {
    @Setter(AccessLevel.NONE)
    @Getter(AccessLevel.NONE)
    String sku;
    int modelId;
    String seasonCode;
    String typeCode;
    String genderCode;
    String ageCode;
    String colorCode;
    String sizeCode;
    int serialNumber;
    int purchaseReceiptId;

    public void skuDecode(String sku) {
        // sku bulder will be (-) seperated
        // modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
        // 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

        String[] splitedSKU = sku.split("-");
        modelId = Integer.parseInt(splitedSKU[0]);
        seasonCode = splitedSKU[1];
        typeCode = splitedSKU[2];
        genderCode = splitedSKU[3];
        ageCode = splitedSKU[4];
        colorCode = splitedSKU[5];
        sizeCode = splitedSKU[6];
        serialNumber = Integer.parseInt(splitedSKU[7]);
        purchaseReceiptId = Integer.parseInt(splitedSKU[8]);

    }

    public String generateSKU() {
        // sku bulder will be (-) seperated
        // modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
        // 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

        StringBuilder skuBuilder = new StringBuilder();
        skuBuilder.append(modelId);
        skuBuilder.append("-");
        skuBuilder.append(seasonCode);
        skuBuilder.append("-");
        skuBuilder.append(typeCode);
        skuBuilder.append("-");
        skuBuilder.append(genderCode);
        skuBuilder.append("-");
        skuBuilder.append(ageCode);
        skuBuilder.append("-");
        skuBuilder.append(colorCode);
        skuBuilder.append("-");
        skuBuilder.append(sizeCode);
        skuBuilder.append("-");
        skuBuilder.append(serialNumber);
        skuBuilder.append("-");
        skuBuilder.append(purchaseReceiptId);

        this.sku = skuBuilder.toString();

        return this.sku;
    }
}
