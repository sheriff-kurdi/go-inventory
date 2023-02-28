package com.kurdi.ecommerce.inventoryservice.infrastructure.repositories;

import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.details.StockItemDetails;
import com.kurdi.ecommerce.inventoryservice.domain.entities.stock.details.StockItemDetailsId;
import org.springframework.data.jpa.repository.JpaRepository;

public interface StockItemDetailsRepository extends JpaRepository<StockItemDetails, StockItemDetailsId> {

}
