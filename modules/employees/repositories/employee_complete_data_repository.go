package repositories

import (
    "context"

    "github.com/ortizdavid/golang-modular-software/database"
    "github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeeCompleteDataRepository struct {
    db *database.Database
}

func NewEmployeeCompleteDataRepository(db *database.Database) *EmployeeCompleteDataRepository {
    return &EmployeeCompleteDataRepository{
        db: db,
    }
}

func (repo *EmployeeCompleteDataRepository) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeCompleteData, error) {
	var rawDataList []entities.EmployeeCompleteDataRaw
    var completeData entities.EmployeeCompleteData

    // Execute the query and retrieve all rows
    result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_complete_data WHERE unique_id=?", uniqueId).Scan(&rawDataList)
    if result.Error != nil {
        return completeData, result.Error
    }

	if len(rawDataList) == 0 {
        return completeData, nil // Return empty if no data is found
    }
    firstRow := rawDataList[0]
    completeData.EmployeeId = firstRow.EmployeeId
    completeData.UniqueId = firstRow.UniqueId
    completeData.CreatedAt = firstRow.CreatedAt
    completeData.UpdatedAt = firstRow.UpdatedAt

    completeData.PersonalInfo = entities.PersonalBasic{
        FirstName:            firstRow.FirstName,
        LastName:             firstRow.LastName,
        IdentificationNumber: firstRow.IdentificationNumber,
        Gender:               firstRow.Gender,
        DateOfBirth:          firstRow.DateOfBirth,
        IdentificationTypeName: firstRow.IdentificationTypeName,
        CountryName:          firstRow.CountryName,
        MaritalStatusName:    firstRow.MaritalStatusName,
    }

    completeData.ProfessionalInfo = entities.ProfessionalBasic{
        DepartmentName:       firstRow.DepartmentName,
        EmploymentStatusName: firstRow.EmploymentStatusName,
        JobTitleName:         firstRow.JobTitleName,
    }

    // Use map to avoid duplicates (e.g., phones, emails, documents)
    phoneMap := make(map[string]bool)
    emailMap := make(map[string]bool)
    documentMap := make(map[string]bool)
    addressMap := make(map[string]bool)

    // Iterate through all rows to build nested data
    for _, row := range rawDataList {
        // Add documents if not already in the map
        documentKey := row.DocumentName + row.DocumentNumber
        if !documentMap[documentKey] {
            completeData.Documents = append(completeData.Documents, entities.DocumentBasic{
                DocumentName:     row.DocumentName,
                DocumentNumber:   row.DocumentNumber,
                ExpirationDate:   row.ExpirationDate,
                FileName:         row.FileName,
                Status:           row.DocumentStatus,
                DocumentTypeName: row.DocumentTypeName,
            })
            documentMap[documentKey] = true
        }

        // Add phones if not already in the map
        if !phoneMap[row.PhoneNumber] {
            completeData.Phones = append(completeData.Phones, entities.PhoneBasic{
                PhoneNumber:     row.PhoneNumber,
                ContactTypeName: row.PhoneContactTypeName,
            })
            phoneMap[row.PhoneNumber] = true
        }

        // Add emails if not already in the map
        if !emailMap[row.EmailAddress] {
            completeData.Emails = append(completeData.Emails, entities.EmailBasic{
                EmailAddress:    row.EmailAddress,
                ContactTypeName: row.EmailContactTypeName,
            })
            emailMap[row.EmailAddress] = true
        }

        // Add addresses if not already in the map
        addressKey := row.State + row.City + row.Street
        if !addressMap[addressKey] {
            completeData.Addresses = append(completeData.Addresses, entities.AddressBasic{
                State:            row.State,
                City:             row.City,
                Neighborhood:     row.Neighborhood,
                Street:           row.Street,
                HouseNumber:      row.HouseNumber,
                PostalCode:       row.PostalCode,
                CountryCode:      row.CountryCode,
                AdditionalDetails: row.AdditionalDetails,
                IsCurrent:        row.IsCurrent,
            })
            addressMap[addressKey] = true
        }
    }

    // Add user account
    completeData.UserAccount = entities.AccountBasic{
        UserName: firstRow.UserName,
        UserEmail:    firstRow.UserEmail,
    }
    return completeData, nil
}


func (repo *EmployeeCompleteDataRepository) GetByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeCompleteData, error) {
    var rawDataList []entities.EmployeeCompleteDataRaw
    var completeData entities.EmployeeCompleteData

    // Execute the query and retrieve all rows
    result := repo.db.WithContext(ctx).Raw("SELECT * FROM employees.view_employee_complete_data WHERE identification_number=?", identNumber).Scan(&rawDataList)
    if result.Error != nil {
        return completeData, result.Error
    }

	if len(rawDataList) == 0 {
        return completeData, nil // Return empty if no data is found
    }
    firstRow := rawDataList[0]
    completeData.EmployeeId = firstRow.EmployeeId
    completeData.UniqueId = firstRow.UniqueId
    completeData.CreatedAt = firstRow.CreatedAt
    completeData.UpdatedAt = firstRow.UpdatedAt

    completeData.PersonalInfo = entities.PersonalBasic{
        FirstName:            firstRow.FirstName,
        LastName:             firstRow.LastName,
        IdentificationNumber: firstRow.IdentificationNumber,
        Gender:               firstRow.Gender,
        DateOfBirth:          firstRow.DateOfBirth,
        IdentificationTypeName: firstRow.IdentificationTypeName,
        CountryName:          firstRow.CountryName,
        MaritalStatusName:    firstRow.MaritalStatusName,
    }

    completeData.ProfessionalInfo = entities.ProfessionalBasic{
        DepartmentName:       firstRow.DepartmentName,
        EmploymentStatusName: firstRow.EmploymentStatusName,
        JobTitleName:         firstRow.JobTitleName,
    }

    // Use map to avoid duplicates (e.g., phones, emails, documents)
    phoneMap := make(map[string]bool)
    emailMap := make(map[string]bool)
    documentMap := make(map[string]bool)
    addressMap := make(map[string]bool)

    // Iterate through all rows to build nested data
    for _, row := range rawDataList {
        // Add documents if not already in the map
        documentKey := row.DocumentName + row.DocumentNumber
        if !documentMap[documentKey] {
            completeData.Documents = append(completeData.Documents, entities.DocumentBasic{
                DocumentName:     row.DocumentName,
                DocumentNumber:   row.DocumentNumber,
                ExpirationDate:   row.ExpirationDate,
                FileName:         row.FileName,
                Status:           row.DocumentStatus,
                DocumentTypeName: row.DocumentTypeName,
            })
            documentMap[documentKey] = true
        }

        // Add phones if not already in the map
        if !phoneMap[row.PhoneNumber] {
            completeData.Phones = append(completeData.Phones, entities.PhoneBasic{
                PhoneNumber:     row.PhoneNumber,
                ContactTypeName: row.PhoneContactTypeName,
            })
            phoneMap[row.PhoneNumber] = true
        }

        // Add emails if not already in the map
        if !emailMap[row.EmailAddress] {
            completeData.Emails = append(completeData.Emails, entities.EmailBasic{
                EmailAddress:    row.EmailAddress,
                ContactTypeName: row.EmailContactTypeName,
            })
            emailMap[row.EmailAddress] = true
        }

        // Add addresses if not already in the map
        addressKey := row.State + row.City + row.Street
        if !addressMap[addressKey] {
            completeData.Addresses = append(completeData.Addresses, entities.AddressBasic{
                State:            row.State,
                City:             row.City,
                Neighborhood:     row.Neighborhood,
                Street:           row.Street,
                HouseNumber:      row.HouseNumber,
                PostalCode:       row.PostalCode,
                CountryCode:      row.CountryCode,
                AdditionalDetails: row.AdditionalDetails,
                IsCurrent:        row.IsCurrent,
            })
            addressMap[addressKey] = true
        }
    }

    // Add user account
    completeData.UserAccount = entities.AccountBasic{
        UserName: firstRow.UserName,
        UserEmail:    firstRow.UserEmail,
    }
    return completeData, nil
}
