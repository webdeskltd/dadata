package dadata

// Определите, нужна ли дополнительная проверка оператором, используя код качества (qc):
const (
	QC_SUCCESS = 0 // Исходное значение распознано уверенно. Не требуется ручная проверка.
	QC_FAILURE = 1 // Исходное значение распознано с допущениями или не распознано. Требуется ручная проверка.
)

// Определите пригодность к рассылке, используя код полноты адреса (qc_complete):
const (
	QC_COMPLETE_SUITABLE        = 0  // Пригоден для почтовой рассылки
	QC_COMPLETE_NO_REGION       = 1  // Не пригоден, нет региона
	QC_COMPLETE_NO_CITY         = 2  // Не пригоден, нет города
	QC_COMPLETE_NO_STREET       = 3  // Не пригоден, нет улицы
	QC_COMPLETE_NOT_HOME        = 4  // Не пригоден, нет дома
	QC_COMPLETE_NO_APARTMENT    = 5  // Пригоден для юридических лиц или частных владений (нет квартиры)
	QC_COMPLETE_NOT_SUITABLE    = 6  // Не пригоден
	QC_COMPLETE_FOREIGN_ADDRESS = 7  // Иностранный адрес
	QC_COMPLETE_NO_KLADR        = 10 // Пригоден, но низкая вероятность успешной доставки (дом не найден в КЛАДР)
)

// Определите вероятность успешной доставки письма по адресу, используя код проверки дома (qc_house):
const (
	QC_HOUSE_EXACT_MATCH         = 2  // Дом найден по точному совпадению (КЛАДР)	Высокая
	QC_HOUSE_NOT_EXPANSION_MATCH = 3  // Различие в расширении дома (КЛАДР)	Средняя
	QC_HOUSE_RANGE_MATCH         = 4  // Дом найден по диапазону (КЛАДР)	Средняя
	QC_HOUSE_NOT_FOUND           = 10 // Дом не найден (КЛАДР)	Низкая
)

// Определите точность координат адреса доставки с помощью кода qc_geo:
const (
	QC_GEO_EXACT_COORDINATES = 0 // Точные координаты
	QC_GEO_NEAREST_HOUSE     = 1 // Ближайший дом
	QC_GEO_STREET            = 2 // Улица
	QC_GEO_LOCALITY          = 3 // Населенный пункт
	QC_GEO_CITY              = 4 // Город
	QC_GEO_NOT_DETERMINED    = 5 // Координаты не определены
)

// Проверьте, указал ли клиент телефон, соответствующий его адресу, с помощью кода qc_conflict (удобно для проверки уровня риска):
const (
	QC_CONFLICT_FULL_MATH   = 0 // Телефон соответствует адресу
	QC_CONFLICT_CITY_MATH   = 2 // Города адреса и телефона отличаются
	QC_CONFLICT_REGION_MATH = 3 // Регионы адреса и телефона отличаются
)

// BoundValue type wrapper for suggest bounds
// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=222888017
type BoundValue string

// const for SuggestBound
const (
	SuggestBoundRegion     BoundValue = "region"     // Регион
	SuggestBoundArea       BoundValue = "area"       // Район
	SuggestBoundCity       BoundValue = "city"       // Город
	SuggestBoundSettlement BoundValue = "settlement" // Населенный пункт
	SuggestBoundStreet     BoundValue = "street"     // Улица
	SuggestBoundHouse      BoundValue = "house"      // Дом
)
