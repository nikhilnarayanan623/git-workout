package main

funcd main(){

}

func CacheStore() (addValueToCache func(key, value int), getValuesFromCache func(key int) (value int, exist bool)) {

	cacheStorage := map[int]int{}

	addValueToCache = func(key, value int) {
		cacheStorage[key] = value
	}

	getValuesFromCache = func(key int) (value int, exist bool) {
		value, ok := cacheStorage[key]
		return value, ok
	}

	return addValueToCache, getValuesFromCache
}
