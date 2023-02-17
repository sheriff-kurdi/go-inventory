package com.kurdi.ecommerce.inventoryservice.repositories;

import com.kurdi.ecommerce.inventoryservice.entities.categories.Category;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CategoriesRepository extends JpaRepository<Category, String> {
}
