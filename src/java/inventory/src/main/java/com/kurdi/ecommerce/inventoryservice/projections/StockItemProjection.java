package com.kurdi.ecommerce.inventoryservice.projections;

public interface StockItemProjection {


    String getSku();
    String getName();
    String getCategoryName();
    String getDescription();
    Integer getSupplierIdentity();
    Integer getAvailableStock();
    Integer getReservedStock();
    Integer getTotalStock();
    Integer getCostPrice();
    Integer getSellingPrice();
    Integer getDiscount();
    boolean getIsDiscounted();



}