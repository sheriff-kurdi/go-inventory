package com.kurdi.inventory.product;


import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import com.kurdi.inventory.domain.enums.products.Seasons;



class SeasonsEnumTests {


	@Test
	void getSpringSeasonTest() {
		String expectedEnumValue = "SPR2020";
		String actualEnumValue = Seasons.SPRING.getSeason(2020);
		assertEquals(expectedEnumValue, actualEnumValue);
	}

	@Test
	void getWinterSeasonTest() {
		String expectedEnumValue = "WIN2020";
		String actualEnumValue = Seasons.WINTER.getSeason(2020);
		assertEquals(expectedEnumValue, actualEnumValue);
	}

	@Test
	void getAutumnSeasonTest() {
		String expectedEnumValue = "AUT2020";
		String actualEnumValue = Seasons.AUTUMN.getSeason(2020);
		assertEquals(expectedEnumValue, actualEnumValue);
	}

	@Test
	void getSummerSeasonTest() {
		String expectedEnumValue = "SUM2020";
		String actualEnumValue = Seasons.SUMMER.getSeason(2020);
		assertEquals(expectedEnumValue, actualEnumValue);
	}

}
