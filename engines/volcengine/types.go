package volcengine

type translationData struct {
	TranslationList []struct {
		Translation            string `json:"Translation"`
		DetectedSourceLanguage string `json:"DetectedSourceLanguage"`
	} `json:"TranslationList"`
	ResponseMetadata struct {
		RequestID string `json:"RequestId"`
		Action    string `json:"Action"`
		Version   string `json:"Version"`
		Service   string `json:"Service"`
		Region    string `json:"Region"`
		Error     struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error"`
	} `json:"ResponseMetadata"`
}
