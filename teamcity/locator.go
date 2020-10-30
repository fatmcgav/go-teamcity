package teamcity

import (
	"fmt"
	"net/url"
)

//Locator represents a arbitraty locator to be used when querying resources, such as id:, type:, or key:
//These are used in GET requests within the URL so must be properly escaped
type Locator string

//LocatorID creates a locator for a Project/BuildType by Id
func LocatorID(id string) Locator {
	return Locator(url.QueryEscape("id:") + id)
}

//LocatorName creates a locator for Project/BuildType by Name
func LocatorName(name string) Locator {
	return Locator(url.QueryEscape("name:") + url.PathEscape(name))
}

//LocatorKey creates a locator for Group by Key
func LocatorKey(key string) Locator {
	return Locator(url.QueryEscape("key:") + url.PathEscape(key))
}

//LocatorType creates a locator for a Project Feature by Type
func LocatorType(id string) Locator {
	return Locator(url.QueryEscape("type:") + id)
}

//LocatorTypeProvider creates a locator for a Project Feature by Type and Provider
func LocatorTypeProvider(featureType string, featureProvider string) Locator {
	queryString := fmt.Sprintf("type:%s,property(name:providerType,value:%s)", featureType, featureProvider)
	return Locator(url.QueryEscape(queryString))
}

func (l Locator) String() string {
	return string(l)
}
