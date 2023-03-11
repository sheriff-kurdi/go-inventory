package com.kurdi.inventory.core.models.categories;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.io.Serializable;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "category_item_details")
public class CategoryDetails implements Serializable {
    @Id
    @GeneratedValue
    Integer id;
    String name;
    String description;
    @JsonIgnore
    @ManyToOne
    @JoinColumn(name = "category_id")
    Category category;


}
