package com.kurdi.inventory.domain.entities;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.io.Serializable;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "languages")
public class Language  implements Serializable {
    @Id
    @Column(name = "language_code")
    String languageCode;
    @Column(name = "language_name")
    String languageName;

}
