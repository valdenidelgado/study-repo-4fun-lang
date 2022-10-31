package com.project.demoproject.model;

import jakarta.persistence.*;

import java.io.Serializable;

@Entity
@Table(name = "permission")
public class Permission implements Serializable {
    private static final long serialVersionUID = 1L;

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @Column
    private String description;
}
