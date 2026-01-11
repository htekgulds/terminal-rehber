package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GetPeople reads and returns all people from the JSON file
func GetPeople() ([]Person, error) {
	// Get the path to the data directory relative to the project root
	dataPath := filepath.Join("data", "people.json")

	// Read the file
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read people.json: %w", err)
	}

	// Unmarshal JSON data
	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		return nil, fmt.Errorf("failed to unmarshal people.json: %w", err)
	}

	return people, nil
}

// GetPersonById finds a person by their Id
func GetPersonById(id string) (*Person, error) {
	people, err := GetPeople()
	if err != nil {
		return nil, err
	}

	for i := range people {
		if people[i].Id == id {
			return &people[i], nil
		}
	}

	return nil, fmt.Errorf("person with Id %s not found", id)
}

// GetPeopleByDepartmentId returns all people in a specific department
func GetPeopleByDepartmentId(departmentId string) ([]Person, error) {
	people, err := GetPeople()
	if err != nil {
		return nil, err
	}

	var result []Person
	for i := range people {
		if people[i].DepartmentId == departmentId {
			result = append(result, people[i])
		}
	}

	return result, nil
}
