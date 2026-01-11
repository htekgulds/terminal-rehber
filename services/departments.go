package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GetDepartments reads and returns all departments from the JSON file
func GetDepartments() ([]Department, error) {
	// Get the path to the data directory relative to the project root
	dataPath := filepath.Join("data", "departments.json")

	// Read the file
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read departments.json: %w", err)
	}

	// Unmarshal JSON data
	var departments []Department
	if err := json.Unmarshal(data, &departments); err != nil {
		return nil, fmt.Errorf("failed to unmarshal departments.json: %w", err)
	}

	return departments, nil
}

// GetDepartmentById finds a department by its Id
func GetDepartmentById(id string) (*Department, error) {
	departments, err := GetDepartments()
	if err != nil {
		return nil, err
	}

	for i := range departments {
		if departments[i].Id == id {
			return &departments[i], nil
		}
	}

	return nil, fmt.Errorf("department with Id %s not found", id)
}

// GetDepartmentsByParentId returns all departments with a specific parent department Id
func GetDepartmentsByParentId(parentId string) ([]Department, error) {
	departments, err := GetDepartments()
	if err != nil {
		return nil, err
	}

	var result []Department
	for i := range departments {
		if departments[i].ParentDepartmentId != nil && *departments[i].ParentDepartmentId == parentId {
			result = append(result, departments[i])
		}
	}

	return result, nil
}

// GetTopLevelDepartments returns all departments without a parent (top-level departments)
func GetTopLevelDepartments() ([]Department, error) {
	departments, err := GetDepartments()
	if err != nil {
		return nil, err
	}

	var result []Department
	for i := range departments {
		if departments[i].ParentDepartmentId == nil {
			result = append(result, departments[i])
		}
	}

	return result, nil
}
