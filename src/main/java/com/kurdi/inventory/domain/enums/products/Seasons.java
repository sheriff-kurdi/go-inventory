package com.kurdi.inventory.domain.enums.products;

public enum Seasons {
    // enum constants calling the enum constructors
    SPRING("SPR"),
    WINTER("WIN"),
    SUMMER("SUM"),
    AUTUMN("AUT");

    private final String season;

    private Seasons(String season) {
        this.season = season;
    }

    public String getSeason(int year) {
        return season + year;
    }

}
