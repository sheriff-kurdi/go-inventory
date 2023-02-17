package com.kurdi.ecommerce.inventoryservice.controllers;

import com.kurdi.ecommerce.inventoryservice.entities.categories.Category;
import com.kurdi.ecommerce.inventoryservice.repositories.CategoriesRepository;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@SuppressWarnings("unused")
@RestController
@RequestMapping("/categories/")
public class CategoriesController {
    final CategoriesRepository categoriesRepository;

    public CategoriesController(CategoriesRepository categoriesRepository) {
        this.categoriesRepository = categoriesRepository;
    }

    @GetMapping
    public List<Category> getAll()
    {

        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        return  categoriesRepository.findAll();
    }

    @PostMapping
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<Category> add(@RequestBody Category category)
    {

        return new ResponseEntity<>(categoriesRepository.save(category), HttpStatus.CREATED);
    }

    @PutMapping("{name}")
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<Category> update(@PathVariable String name, @RequestParam String newName)
    {
        if(!categoriesRepository.existsById(name))
        {
            return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);

        }
        categoriesRepository.delete(categoriesRepository.getById(name));
        return new ResponseEntity<>(categoriesRepository.save(Category.builder().name(newName).build()), HttpStatus.OK);
    }

    @DeleteMapping("{name}")
    @Operation(security = @SecurityRequirement(name = "bearerAuth"))
    public ResponseEntity<Category> delete(@PathVariable String name)
    {
        Category category = Category.builder().name(name).build();
        categoriesRepository.delete(category);
        return new ResponseEntity<>(category, HttpStatus.OK);
    }


}
