package dadata

func (daData *DaData) sendCleanRequest(lastURLPart string, source, result interface{}) error {
	return daData.sendRequest("clean/"+lastURLPart, source, result)
}

// CleanAddresses clean all provided addresses
// Call https://dadata.ru/api/v2/clean/address
func (daData *DaData) CleanAddresses(sourceAddresses ...string) ([]Address, error) {
	var addresses []Address
	if err := daData.sendCleanRequest("address", &sourceAddresses, &addresses); err != nil {
		return nil, err
	}
	return addresses, nil
}

// CleanPhones clean all provided phones
// Call https://dadata.ru/api/v2/clean/phone
func (daData *DaData) CleanPhones(sourcePhones ...string) ([]Phone, error) {
	var phones []Phone
	if err := daData.sendCleanRequest("phone", &sourcePhones, &phones); err != nil {
		return nil, err
	}
	return phones, nil
}

// CleanNames clean all provided names
// Call https://dadata.ru/api/v2/clean/name
func (daData *DaData) CleanNames(sourceNames ...string) ([]Name, error) {
	var names []Name
	if err := daData.sendCleanRequest("name", &sourceNames, &names); err != nil {
		return nil, err
	}
	return names, nil
}

// CleanEmails clean all provided emails
// Call https://dadata.ru/api/v2/clean/email
func (daData *DaData) CleanEmails(sourceEmails ...string) ([]Email, error) {
	var emails []Email
	if err := daData.sendCleanRequest("email", &sourceEmails, &emails); err != nil {
		return nil, err
	}
	return emails, nil
}

// CleanBirthdates clean all provided birthdates
// Call https://dadata.ru/api/v2/clean/birthdate
func (daData *DaData) CleanBirthdates(sourceBirthdates ...string) ([]Birthdate, error) {
	var birthdates []Birthdate
	if err := daData.sendCleanRequest("birthdate", &sourceBirthdates, &birthdates); err != nil {
		return nil, err
	}
	return birthdates, nil
}

// CleanVehicles clean all provided vehicles
// Call https://dadata.ru/api/v2/clean/vehicle
func (daData *DaData) CleanVehicles(sourceVehicles ...string) ([]Vehicle, error) {
	var vehicles []Vehicle
	if err := daData.sendCleanRequest("vehicle", &sourceVehicles, &vehicles); err != nil {
		return nil, err
	}
	return vehicles, nil
}

// CleanPassports clean all provided passports
// Call https://dadata.ru/api/v2/clean/passport
func (daData *DaData) CleanPassports(sourcePassports ...string) ([]Passport, error) {
	var passports []Passport
	if err := daData.sendCleanRequest("passport", &sourcePassports, &passports); err != nil {
		return nil, err
	}
	return passports, nil
}
