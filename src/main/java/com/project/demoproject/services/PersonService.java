package com.project.demoproject.services;

import com.project.demoproject.exceptions.ResourceNotFoundException;
import com.project.demoproject.mapper.MapperStruct;
import com.project.demoproject.model.Person;
import com.project.demoproject.model.dto.v1.PersonDTO;
import com.project.demoproject.repositories.PersonRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.logging.Logger;

@Service
public class PersonService {

    @Autowired
    private PersonRepository repository;

    private final Logger logger = Logger.getLogger(PersonService.class.getName());

    public List<PersonDTO> findAll() {
        logger.info("Find all persons");
        return MapperStruct.INSTANCE.toPersonDTOs(repository.findAll());
    }

    public PersonDTO findById(Long id) {
        logger.info("Findind one person");
        return MapperStruct.INSTANCE.toPersonDTO(repository.findById(id).orElseThrow(() -> new ResourceNotFoundException("No records found for this ID")));
    }

    public PersonDTO create(PersonDTO personDTO) {
        logger.info("Creating a new person");
        Person save =  repository.save(MapperStruct.INSTANCE.toPerson(personDTO));
        return MapperStruct.INSTANCE.toPersonDTO(save);
    }


    public PersonDTO update(PersonDTO personDTO) {
        logger.info("Update person");
        Person entity = repository.findById(personDTO.getId()).orElseThrow(() -> new ResourceNotFoundException("No records found for this ID"));
        entity.setFirstName(personDTO.getFirstName());
        entity.setLastName(personDTO.getLastName());
        entity.setAddress(personDTO.getAddress());
        entity.setGender(personDTO.getGender());
        return MapperStruct.INSTANCE.toPersonDTO(repository.save(entity));
    }

    public void delete(Long id) {
        logger.info("Delete person");
        Person entity = repository.findById(id).orElseThrow(() -> new ResourceNotFoundException("No records found for this ID"));
        repository.delete(entity);
    }
}
