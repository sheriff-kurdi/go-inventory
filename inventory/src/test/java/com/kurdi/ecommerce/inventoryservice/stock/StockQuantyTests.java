package com.kurdi.ecommerce.inventoryservice.stock;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.item.StockItemQuantity;

class StockQuantyTests {

	/*
	 * Add Stock
	 */
	@Test
	void addQuantityWithTotalStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 10;
		stockItemQuantity.addStock(actualAddedStock);

		assertEquals(stockItemQuantity.getTotalStock(), actualAddedStock);
	}

	@Test
	void addQuantityWithAvailableStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 10;
		stockItemQuantity.addStock(actualAddedStock);

		assertEquals(stockItemQuantity.getAvailableStock(), actualAddedStock);
	}

	/*
	 * Reserve Stock
	 */
	@Test
	void reserveQuantityWithTotalStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		assertEquals(stockItemQuantity.getTotalStock(), actualAddedStock);
	}

	@Test
	void reserveQuantityWithAvailableStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualReservedStock = 10;
		stockItemQuantity.reserveStock(actualReservedStock);

		int actualAvailableStock = actualAddedStock - actualReservedStock;

		assertEquals(stockItemQuantity.getAvailableStock(), actualAvailableStock);
	}

	@Test
	void reserveQuantityWithReservedStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

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
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

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
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

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
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

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
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualsoldStock = 10;
		stockItemQuantity.selling(actualsoldStock);
		int actualStock = actualAddedStock - actualsoldStock;
		assertEquals(stockItemQuantity.getTotalStock(), actualStock);
	}

	@Test
	void sellingWithAvailableStockTest() {
		StockItemQuantity stockItemQuantity = new StockItemQuantity();

		int actualAddedStock = 100;
		stockItemQuantity.addStock(actualAddedStock);

		int actualsoldStock = 10;
		stockItemQuantity.selling(actualsoldStock);
		int actualStock = actualAddedStock - actualsoldStock;

		assertEquals(stockItemQuantity.getAvailableStock(), actualStock);
	}

}
