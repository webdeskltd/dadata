package dadata

func (daData *DaData) sendCleanRequest(lastUrlPart string, source, result interface{}) error {
	return daData.sendRequest("clean/"+lastUrlPart, source, result)
}

type Address struct {
	Source              string `json:"source"`               // Исходный адрес одной строкой
	PostalCode          string `json:"postal_code"`          // Индекс
	Country             string `json:"country"`              // Страна
	RegionType          string `json:"region_type"`          // Тип региона (сокращенный)
	RegionTypeFull      string `json:"region_type_full"`     // Тип региона
	Region              string `json:"region"`               // Регион
	AreaType            string `json:"area_type"`            // Тип района в регионе (сокращенный)
	AreaTypeFull        string `json:"area_type_full"`       // Тип района в регионе
	Area                string `json:"area"`                 // Район в регионе
	CityType            string `json:"city_type"`            // Тип города (сокращенный)
	CityTypeFull        string `json:"city_type_full"`       // Тип города
	City                string `json:"city"`                 // Город
	SettlementType      string `json:"settlement_type"`      // Тип населенного пункта (сокращенный)
	SettlementTypeFull  string `json:"settlement_type_full"` // Тип населенного пункта
	Settlement          string `json:"settlement"`           // Населенный пункт
	StreetType          string `json:"street_type"`          // Тип улицы (сокращенный)
	StreetTypeFull      string `json:"street_type_full"`     // Тип улицы
	Street              string `json:"street"`               // Улица
	HouseType           string `json:"house_type"`           // Тип дома (сокращенный)
	HouseTypeFull       string `json:"house_type_full"`      // Тип дома
	House               string `json:"house"`                // Дом
	BlockType           string `json:"block_type"`           // Тип корпуса/строения
	Block               string `json:"block"`                // Корпус/строение
	FlatType            string `json:"flat_type"`            // Тип квартиры
	Flat                string `json:"flat"`                 // Квартира
	PostalBox           string `json:"postal_box"`           // Абонентский ящик
	KladrId             string `json:"kladr_id"`             // Код КЛАДР
	Okato               string `json:"okato"`                // Код ОКАТО
	Oktmo               string `json:"oktmo"`                // Код ОКТМО
	TaxOffice           string `json:"tax_office"`           // Код ИФНС
	FlatArea            string `json:"flat_area"`            // Площадь квартиры
	Timezone            string `json:"timezone"`             // Часовой пояс
	GeoLat              string `json:"geo_lat"`              // Координаты: широта
	GeoLon              string `json:"geo_lon"`              // Координаты: долгота
	QualityCodeGeo      int    `json:"qc_geo"`               // Код точности координат
	QualityCodeComplete int    `json:"qc_complete"`          // Код полноты
	QualityCodeHouse    int    `json:"qc_house"`             // Код проверки дома
	QualityCode         int    `json:"qc"`                   // Код качества
	UnparsedParts       string `json:"unparsed_parts"`       // Нераспознанная часть адреса. Для адреса
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

type Phone struct {
	Source              string `json:"source"`       // Исходный телефон одной строкой
	Type                string `json:"type"`         // Тип телефона
	Phone               string `json:"phone"`        // Стандартизованный телефон одной строкой
	CountryCode         string `json:"country_code"` // Код страны
	CityCode            string `json:"city_code"`    // Код города / DEF-код
	Number              string `json:"number"`       // Локальный номер телефона
	Extension           string `json:"extension"`    // Добавочный номер
	Provider            string `json:"provider"`     // Оператор связи
	Region              string `json:"region"`       // Регион
	Timezone            string `json:"timezone"`     // Часовой пояс
	QualityCodeConflict int    `json:"qc_conflict"`  // Признак конфликта телефона с адресом
	QualityCode         int    `json:"qc"`           // Код качества
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

type Name struct {
	Source      string `json:"source"`     // Исходное ФИО одной строкой
	Surname     string `json:"surname"`    // Фамилия
	Name        string `json:"name"`       // Имя
	Patronymic  string `json:"patronymic"` // Отчество
	Gender      string `json:"gender"`     // Пол
	QualityCode int    `json:"qc"`         // Код качества
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

type Email struct {
	Source      string `json:"source"` // Исходный e-mail
	Email       string `json:"email"`  // Стандартизованный e-mail
	QualityCode int    `json:"qc"`     // Код качества
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

type Birthdate struct {
	Source      string `json:"source"`    // Исходная дата
	Birthdate   string `json:"birthdate"` // Стандартизованная дата
	QualityCode int    `json:"qc"`        // Код качества
}

/*
Call https://dadata.ru/api/v2/cleanbirthdate
*/
func (daData *DaData) CleanBirthdates(sourceBirthdates ...string) ([]Birthdate, error) {
	var birthdates []Birthdate
	if sendErr := daData.sendCleanRequest("birthdate", &sourceBirthdates, &birthdates); nil == sendErr {
		return birthdates, nil
	} else {
		return nil, sendErr
	}
}
