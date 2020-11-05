package helpers

// GetEnv returns the environment file to use for bootstrapping the config
func GetEnv(isTesting bool) string {
	if isTesting {
		return ".env.test"
	}

	return ".env"
}
