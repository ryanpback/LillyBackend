package models

import bs "lillyAppBackend/bootstrap"

// AppConfig holds all application configuration
var AppConfig bs.Config

// Response is how all requests and responses are wrapped
type Response map[string]interface{}
