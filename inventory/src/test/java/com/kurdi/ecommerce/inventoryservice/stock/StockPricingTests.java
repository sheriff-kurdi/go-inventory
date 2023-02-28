package com.kurdi.ecommerce.inventoryservice.stock;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.item.StockItemPrices;

class StockPricingTests {


	@Test
	void sellingPriceWithDiscountStockTest() {
		StockItemPrices stockItemPricing = StockItemPrices.builder()
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
		StockItemPrices stockItemPricing = StockItemPrices.builder()
		.costPrice(80)
		.sellingPrice(100)
		.discount(10)
		.isDiscounted(false)
		.build();

		int actualSellingPrice = 100;
		
		assertEquals(stockItemPricing.getSellingPrice(), actualSellingPrice);
	}

}
