package com.kurdi.ecommerce.inventoryservice.responses;

import org.springframework.http.ResponseEntity;

import java.util.HashMap;

public class CustomResponse {

    public static ResponseEntity<HashMap<String, Object>> Ok(Object data){
        HashMap<String, Object> response = new HashMap<>();
        response.put("data", data);
        response.put("success",  true);
        return ResponseEntity.ok (response);
    }

}
