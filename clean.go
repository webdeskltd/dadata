package dadata

func (daData *DaData) sendCleanRequest(lastUrlPart string, source, result interface{}) error {
	return daData.sendRequest("clean/"+lastUrlPart, source, result)
}

type Address struct {
	Source               string `json:"source"`                  // Исходный адрес одной строкой
	Result               string `json:"result"`                  // Стандартизованный адрес одной строкой
	PostalCode           string `json:"postal_code"`             // Индекс
	Country              string `json:"country"`                 // Страна
	RegionFiasId         string `json:"region_fias_id"`          // Код ФИАС региона
	RegionKladrId        string `json:"region_kladr_id"`         // Код КЛАДР региона
	RegionWithType       string `json:"region_with_type"`        // Регион с типом
	RegionType           string `json:"region_type"`             // Тип региона (сокращенный)
	RegionTypeFull       string `json:"region_type_full"`        // Тип региона
	Region               string `json:"region"`                  // Регион
	AreaFiasId           string `json:"area_fias_id"`            // Код ФИАС района в регионе
	AreaKladrId          string `json:"area_kladr_id"`           // Код КЛАДР района в регионе
	AreaWithType         string `json:"area_with_type"`          // Район в регионе с типом
	AreaType             string `json:"area_type"`               // Тип района в регионе (сокращенный)
	AreaTypeFull         string `json:"area_type_full"`          // Тип района в регионе
	Area                 string `json:"area"`                    // Район в регионе
	CityFiasId           string `json:"city_fias_id"`            // Код ФИАС города
	CityKladrId          string `json:"city_kladr_id"`           // Код КЛАДР города
	CityWithType         string `json:"city_with_type"`          // Город с типом
	CityType             string `json:"city_type"`               // Тип города (сокращенный)
	CityTypeFull         string `json:"city_type_full"`          // Тип города
	City                 string `json:"city"`                    // Город
	CityArea             string `json:"city_area"`               // Административный округ (только для Москвы)
	CityDistrictFiasId   string `json:"city_district_fias_id"`   // Код ФИАС района города (заполняется, только если район есть в ФИАС)
	CityDistrictKladrId  string `json:"city_district_kladr_id"`  // Код КЛАДР района города (не заполняется)
	CityDistrictWithType string `json:"city_district_with_type"` // Район города с типом
	CityDistrictType     string `json:"city_district_type"`      // Тип района города (сокращенный)
	CityDistrictTypeFull string `json:"city_district_type_full"` // Тип района города
	CityDistrict         string `json:"city_district"`           // Район города
	SettlementFiasId     string `json:"settlement_fias_id"`      // Код ФИАС нас. пункта
	SettlementKladrId    string `json:"settlement_kladr_id"`     // Код КЛАДР нас. пункта
	SettlementWithType   string `json:"settlement_with_type"`    // Населенный пункт с типом
	SettlementType       string `json:"settlement_type"`         // Тип населенного пункта (сокращенный)
	SettlementTypeFull   string `json:"settlement_type_full"`    // Тип населенного пункта
	Settlement           string `json:"settlement"`              // Населенный пункт
	StreetFiasId         string `json:"street_fias_id"`          // Код ФИАС улицы
	StreetKladrId        string `json:"street_kladr_id"`         // Код КЛАДР улицы
	StreetWithType       string `json:"street_with_type"`        // Улица с типом
	StreetType           string `json:"street_type"`             // Тип улицы (сокращенный)
	StreetTypeFull       string `json:"street_type_full"`        // Тип улицы
	Street               string `json:"street"`                  // Улица
	HouseFiasId          string `json:"house_fias_id"`           // Код ФИАС дома
	HouseKladrId         string `json:"house_kladr_id"`          // Код КЛАДР дома
	HouseType            string `json:"house_type"`              // Тип дома (сокращенный)
	HouseTypeFull        string `json:"house_type_full"`         // Тип дома
	House                string `json:"house"`                   // Дом
	BlockType            string `json:"block_type"`              // Тип корпуса/строения (сокращенный)
	BlockTypeFull        string `json:"block_type_full"`         // Тип корпуса/строения
	Block                string `json:"block"`                   // Корпус/строение
	FlatType             string `json:"flat_type"`               // Тип квартиры (сокращенный)
	FlatTypeFull         string `json:"flat_type_full"`          // Тип квартиры
	Flat                 string `json:"flat"`                    // Квартира
	FlatArea             string `json:"flat_area"`               // Площадь квартиры
	SquareMeterPrice     string `json:"square_meter_price"`      // Рыночная стоимость м²
	FlatPrice            string `json:"flat_price"`              // Рыночная стоимость квартиры
	PostalBox            string `json:"postal_box"`              // Абонентский ящик
	FiasId               string `json:"fias_id"`                 // Код ФИАС
	FiasLevel            string `json:"fias_level"`              // Уровень детализации, до которого адрес найден в ФИАС
	KladrId              string `json:"kladr_id"`                // Код КЛАДР
	CapitalMarker        string `json:"capital_marker"`          // Статус центра
	Okato                string `json:"okato"`                   // Код ОКАТО
	Oktmo                string `json:"oktmo"`                   // Код ОКТМО
	TaxOffice            string `json:"tax_office"`              // Код ИФНС для физических лиц
	Timezone             string `json:"timezone"`                // Часовой пояс
	GeoLat               string `json:"geo_lat"`                 // Координаты: широта
	GeoLon               string `json:"geo_lon"`                 // Координаты: долгота
	BeltwayHit           string `json:"beltway_hit"`             // Внутри кольцевой?
	BeltwayDistance      string `json:"beltway_distance"`        // Расстояние от кольцевой в км.
	QualityCodeGeo       int    `json:"qc_geo"`                  // Код точности координат
	QualityCodeComplete  int    `json:"qc_complete"`             // Код полноты
	QualityCodeHouse     int    `json:"qc_house"`                // Код проверки дома
	QualityCode          int    `json:"qc"`                      // Код качества
	UnparsedParts        string `json:"unparsed_parts"`          // Нераспознанная часть адреса. Для адреса
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
	Source         string `json:"source"`          // Исходное ФИО одной строкой
	Result         string `json:"result"`          // Стандартизованное ФИО одной строкой
	ResultGenitive string `json:"result_genitive"` // ФИО в родительном падеже (кого?)
	ResultDative   string `json:"result_dative"`   // ФИО в дательном падеже (кому?)
	ResultAblative string `json:"result_ablative"` // ФИО в творительном падеже (кем?)
	Surname        string `json:"surname"`         // Фамилия
	Name           string `json:"name"`            // Имя
	Patronymic     string `json:"patronymic"`      // Отчество
	Gender         string `json:"gender"`          // Пол
	QualityCode    int    `json:"qc"`              // Код качества
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

type Vehicle struct {
	Source      string `json:"source"` // Исходное значение
	Result      string `json:"result"` // Стандартизованное значение
	Brand       string `json:"brand"`  // Марка
	Model       string `json:"model"`  // Модель
	QualityCode int    `json:"qc"`     // Код проверки
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
