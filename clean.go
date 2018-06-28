package dadata

func (daData *DaData) sendCleanRequest(lastUrlPart string, source, result interface{}) error {
	return daData.sendRequest("clean/"+lastUrlPart, source, result)
}

/*
Call https://dadata.ru/api/v2/clean/address
*/
func (daData *DaData) CleanAddresses(sourceAddresses ...string) ([]Address, error) {
	var addresses []Address
	if sendErr := daData.sendCleanRequest("address", &sourceAddresses, &addresses); nil == sendErr {
		return addresses, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/phone
*/
func (daData *DaData) CleanPhones(sourcePhones ...string) ([]Phone, error) {
	var phones []Phone
	if sendErr := daData.sendCleanRequest("phone", &sourcePhones, &phones); nil == sendErr {
		return phones, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/name
*/
func (daData *DaData) CleanNames(sourceNames ...string) ([]Name, error) {
	var names []Name
	if sendErr := daData.sendCleanRequest("name", &sourceNames, &names); nil == sendErr {
		return names, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/email
*/
func (daData *DaData) CleanEmails(sourceEmails ...string) ([]Email, error) {
	var emails []Email
	if sendErr := daData.sendCleanRequest("email", &sourceEmails, &emails); nil == sendErr {
		return emails, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/birthdate
*/
func (daData *DaData) CleanBirthdates(sourceBirthdates ...string) ([]Birthdate, error) {
	var birthdates []Birthdate
	if sendErr := daData.sendCleanRequest("birthdate", &sourceBirthdates, &birthdates); nil == sendErr {
		return birthdates, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/vehicle
*/
func (daData *DaData) CleanVehicles(sourceVehicles ...string) ([]Vehicle, error) {
	var vehicles []Vehicle
	if sendErr := daData.sendCleanRequest("vehicle", &sourceVehicles, &vehicles); nil == sendErr {
		return vehicles, nil
	} else {
		return nil, sendErr
	}
}

/*
Call https://dadata.ru/api/v2/clean/passport
*/
func (daData *DaData) CleanPassports(sourcePassports ...string) ([]Passport, error) {
	var passports []Passport
	if sendErr := daData.sendCleanRequest("passport", &sourcePassports, &passports); nil == sendErr {
		return passports, nil
	} else {
		return nil, sendErr
	}
}
