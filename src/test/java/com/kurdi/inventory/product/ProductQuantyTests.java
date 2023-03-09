package com.kurdi.inventory.product;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.inventory.domain.entities.products.ProductQuantity;


class ProductQuantyTests {

	/*
	 * Add Stock
	 */
	@Test
	void addQuantityWithTotalStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 10;
		stockItemQuantity.addStock(actualAddedStock);

		assertEquals(stockItemQuantity.getTotalStock(), actualAddedStock);
	}

	@Test
	void addQuantityWithAvailableStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 10;
		stockItemQuantity.addStock(actualAddedStock);

		assertEquals(stockItemQuantity.getAvailableStock(), actualAddedStock);
	}

	/*
	 * Reserve Stock
	 */
	@Test
	void reserveQuantityWithTotalStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		assertEquals(stockItemQuantity.getTotalStock(), actualAddedStock);
	}

	@Test
	void reserveQuantityWithAvailableStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		int actualAvailableStock = actualAddedStock - actualReservedStock;

		assertEquals(stockItemQuantity.getAvailableStock(), actualAvailableStock);
	}

	@Test
	void reserveQuantityWithReservedStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		assertEquals(stockItemQuantity.getReservedStock(), actualReservedStock);
	}

	/*
	 * Cancel Reservation
	 */
	@Test
	void CancelReservationQuantityWithTotalStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		int actualUnReservedStock = 5;
		stockItemQuantity.cancelReservation(actualUnReservedStock);

		assertEquals(stockItemQuantity.getTotalStock(), actualAddedStock);
	}

	@Test
	void CancelReservationQuantityWithAvailableStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		int actualUnReservedStock = 5;
		stockItemQuantity.cancelReservation(actualUnReservedStock);

		int actualAvailableStock = actualAddedStock - actualReservedStock + actualUnReservedStock;

		assertEquals(stockItemQuantity.getAvailableStock(), actualAvailableStock);
	}

	@Test
	void CancelReservationWithReservedStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		int actualUnReservedStock = 5;
		stockItemQuantity.cancelReservation(actualUnReservedStock);

		actualReservedStock -= actualUnReservedStock;

		assertEquals(stockItemQuantity.getReservedStock(), actualReservedStock);
	}

	/*
	 * selling
	 */
	@Test
	void sellingWithTotalStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualsoldStock = 10;
		stockItemQuantity.selling(actualsoldStock);
		int actualStock = actualAddedStock - actualsoldStock;
		assertEquals(stockItemQuantity.getTotalStock(), actualStock);
	}

	@Test
	void sellingWithAvailableStockTest() {
		ProductQuantity stockItemQuantity = new ProductQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualsoldStock = 10;
		stockItemQuantity.selling(actualsoldStock);
		int actualStock = actualAddedStock - actualsoldStock;

		assertEquals(stockItemQuantity.getAvailableStock(), actualStock);
	}

}
