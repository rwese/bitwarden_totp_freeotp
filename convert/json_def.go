package convert

type JSONData struct {
	Encrypted bool `json:"encrypted"`
	Folders   []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"folders"`
	Items []Item `json:"items"`
}

// Item represents the JSON data structure of an item
type Item struct {
	ID           string      `json:"id"`
	Organization string      `json:"organizationId"`
	FolderID     string      `json:"folderId"`
	Type         int         `json:"type"`
	Reprompt     int         `json:"reprompt"`
	Name         string      `json:"name"`
	Notes        string      `json:"notes"`
	Favorite     bool        `json:"favorite"`
	Login        Login       `json:"login"`
	Collection   interface{} `json:"collectionIds"`
}

// Login represents the login part of the Item
type Login struct {
	URIs []URIs `json:"uris"`
	User string `json:"username"`
	Pass string `json:"password"`
	TOTP string `json:"totp"`
}

// URIs represents the URI part of the Login
type URIs struct {
	Match int    `json:"match"`
	URI   string `json:"uri"`
}
