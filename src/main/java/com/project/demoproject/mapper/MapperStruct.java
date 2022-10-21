package com.project.demoproject.mapper;

import com.project.demoproject.model.Person;
import com.project.demoproject.model.dto.v1.PersonDTO;
import org.mapstruct.InheritInverseConfiguration;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import java.util.List;

@Mapper(componentModel = "spring")
public interface MapperStruct {

    MapperStruct INSTANCE = Mappers.getMapper(MapperStruct.class);

    PersonDTO toPersonDTO(Person person);

    List<PersonDTO> toPersonDTOs(List<Person> persons);

    @InheritInverseConfiguration
    Person toPerson(PersonDTO personDTO);
}
