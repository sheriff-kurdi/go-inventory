package com.kurdi.inventory.product;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.inventory.domain.entities.products.ProductSKU;
import com.kurdi.inventory.domain.enums.products.Ages;
import com.kurdi.inventory.domain.enums.products.Genders;
import com.kurdi.inventory.domain.enums.products.Seasons;
import com.kurdi.inventory.domain.enums.products.Types;

class ProductSKUTests {

	final private String SKU = "115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120";
	@Test
	void SKUGeneratorTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		ProductSKU productSKU = ProductSKU.builder()
				.modelId(115)
				.seasonCode("SPR2020")
				.typeCode("SHIRT")
				.genderCode(Genders.MALE.name())
				.ageCode(Ages.ADULTS.name())
				.colorCode("RED")
				.sizeCode("S")
				.serialNumber(1)
				.purchaseReceiptId(120)
				.build();

		String actualSKU = productSKU.generateSKU();

		assertEquals(SKU, actualSKU);

	}

	@Test
	void SKUModelIdDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		int ExpectedModel = 115;
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(this.SKU);

		assertEquals(ExpectedModel, productSKU.getModelId());
	}

	@Test
	void SKUSeasonCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedSeason = "SPR2020";
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedSeason, productSKU.getSeasonCode());
	}

	@Test
	void SKUTypeCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedType = "SHIRT";
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedType, productSKU.getTypeCode());
	}

	@Test
	void SKUGynderCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedGender = "MALE";
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedGender, productSKU.getGenderCode());
	}

	@Test
	void SKUAgeCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedAge = "ADULTS";
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedAge, productSKU.getAgeCode());
	}

	@Test
	void SKUSizeCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedSize = "S";
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedSize, productSKU.getSizeCode());
	}

	@Test
	void SKUSerialNumberCodeDecoderTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		int ExpectedSerial = 1;
		ProductSKU productSKU = new ProductSKU();
		productSKU.skuDecode(SKU);

		assertEquals(ExpectedSerial, productSKU.getSerialNumber());
	}

	@Test
	void SKUGeneratingWithEnumsTest() {
		// modelCode-seasonCode-typeCode-genderCode-ageCode-colorCode-sizeCode-serialNumber-purchaseReceiptId
		// 115-SPR2020-SHIRT-MALE-ADULTS-RED-S-1-120

		String ExpectedSKU = "115-SPR2020-SHIRTS-MALE-ADULTS-RED-S-1-120";
		ProductSKU productSKU = ProductSKU.builder()
		.modelId(115)
		.seasonCode(Seasons.SPRING.getSeason(2020))
		.typeCode(Types.SHIRTS.name())
		.genderCode(Genders.MALE.name())
		.ageCode(Ages.ADULTS.name())
		.colorCode("RED")
		.sizeCode("S")
		.serialNumber(1)
		.purchaseReceiptId(120)
		.build();
		String actualSKU = productSKU.generateSKU();

		assertEquals(ExpectedSKU, actualSKU);
	}

}
