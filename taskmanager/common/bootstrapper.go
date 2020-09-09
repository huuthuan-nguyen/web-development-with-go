package common

func StartUp() {
	// init AppConfig variable
	initConfig()
	// init private/public keys for JWT authentication
	initKeys()
	// start MongoDB session
	createDBSession()
	// add Indexes to MongoDB
	addIndexes()
}