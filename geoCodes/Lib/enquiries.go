package geoCodes


import (
    Structs "github.com/alibe1971/go-geoCodes/geoCodes/Structs"
)

func setIndex(settings Structs.SettingsType, reference Structs.GeoCodeReference, index string) {
    if !isInSlice(settings.Indexes, index) {
        logPanicWithStackTrace("Property `" + index + "` not existent or not usable as index")
    }
    geocodesMap[reference].SetEnquiries.Index = &index
}

func setSelect(settings Structs.SettingsType, reference Structs.GeoCodeReference, selectable ...string) {
    for _, value := range selectable {
        if !isInSlice(settings.Public, value) {
            logPanicWithStackTrace("Property `" + value + "` not existent or not usable as selectable")
        }
        geocodesMap[reference].SetEnquiries.Select = append(geocodesMap[reference].SetEnquiries.Select, value)
    }
}

func setOrderBy(settings Structs.SettingsType, reference Structs.GeoCodeReference, props ...string) {
    if !isInSlice(settings.Indexes, props[0]) {
        logPanicWithStackTrace("Attribute `orderBy`.`property` must be usable as index. `" + props[0] + "` isn't valid")
    }
    orderType := strToUpper(props[1])
    if orderType == "" {
        orderType = "ASC"
    }
    if !isInSlice([]string{"ASC", "DESC"}, orderType) {
        logPanicWithStackTrace("Attribute `orderBy`.`property` must be `ASC` (default if empty string - ``) or `DESC` (case insensitive). `" + orderType + "` isn't valid")
    }
    geocodesMap[reference].SetEnquiries.OrderBy.Property = props[0]
    geocodesMap[reference].SetEnquiries.OrderBy.OrderType = orderType
}

func setInterval(reference Structs.GeoCodeReference, method string, value int) {
    message := "Attribute `" + method + "` cannot be less than 0"
    if value < 0 {
        logPanicWithStackTrace(message)
    }
    switch method {
        case "offset":
            geocodesMap[reference].SetEnquiries.Interval.Offset = value
        case "limit":
            geocodesMap[reference].SetEnquiries.Interval.Limit = value
    }
}
