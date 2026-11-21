package config

// GetString gets a string json node from a json key.
// JSON key example: 'foo.bar'
func GetString(key string) string {
	conf := registeredConfigurator()
	if conf == nil {
		return ""
	}
	return conf.GetString(key)
}

// GetInt64 gets an integer 64 json node from a json key.
// JSON key example: 'foo.bar'
func GetInt64(key string) int64 {
	conf := registeredConfigurator()
	if conf == nil {
		return 0
	}
	return conf.GetInt64(key)
}

// GetInt gets an integer json node from a json key.
// JSON key example: 'foo.bar'
func GetInt(key string) int {
	conf := registeredConfigurator()
	if conf == nil {
		return 0
	}
	return conf.GetInt(key)
}

// GetFloat64 gets a float 64 json node from a json key.
// JSON key example: 'foo.bar'
func GetFloat64(key string) float64 {
	conf := registeredConfigurator()
	if conf == nil {
		return 0
	}
	return conf.GetFloat64(key)
}

// GetBool gets a bool json node from a json key.
// JSON key example: 'foo.bar'
func GetBool(key string) bool {
	conf := registeredConfigurator()
	if conf == nil {
		return false
	}
	return conf.GetBool(key)
}

// Get gets a json node from a json key.
// JSON key example: 'foo.bar'
func Get(key string) interface{} {
	conf := registeredConfigurator()
	if conf == nil {
		return nil
	}
	return conf.Get(key)
}

// GetArrayMaps gets a array json node from a json key. Returns
// a json array as map.
// JSON key example: 'foo.bar'
func GetArrayMaps(key string) []map[string]interface{} {
	conf := registeredConfigurator()
	if conf == nil {
		return nil
	}
	return conf.GetArrayMaps(key)
}

// LoadConfigFile Open file that contins a json object.
func LoadConfigFile(name string) error {
	conf := registeredConfigurator()
	if err := validateConfigurator(conf); err != nil {
		return err
	}
	return conf.Open(name)
}
