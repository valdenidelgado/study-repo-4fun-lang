package com.project.demoproject.config;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.info.License;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class OpenApiConfig {

    @Bean
    public OpenAPI customOpenAPI() {
        return new OpenAPI()
                .info(new Info()
                        .title("Demo Restful API with spring boot 3")
                        .version("v1")
                        .description("Restful API with Java 17 and Spring Boot 3")
                        .termsOfService("https://github.com/valdenidelgado")
                        .license(new License().name("Apache 2.0").url("https://github.com/valdenidelgado")));
    }
}
