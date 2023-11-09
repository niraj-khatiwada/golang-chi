package models

func GetAllModels() map[string]interface{} {
	models := map[string]interface{}{
		"contact": &Contact{},
	}
	return models
}
