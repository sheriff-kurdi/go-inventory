package com.kurdi.inventory.product;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.inventory.domain.entities.products.ProductPrices;


class ProductPricingTests {


	@Test
	void sellingPriceWithDiscountStockTest() {
		ProductPrices stockItemPricing = ProductPrices.builder()
		.costPrice(80)
		.sellingPrice(100)
		.discount(10)
		.isDiscounted(true)
		.build();

		int actualSellingPrice = 90;
		
		assertEquals(stockItemPricing.getSellingPrice(), actualSellingPrice);
	}

	@Test
	void sellingPriceWithoutDiscountStockTest() {
		ProductPrices stockItemPricing = ProductPrices.builder()
		.costPrice(80)
		.sellingPrice(100)
		.discount(10)
		.isDiscounted(false)
		.build();

		int actualSellingPrice = 100;
		
		assertEquals(stockItemPricing.getSellingPrice(), actualSellingPrice);
	}

}
