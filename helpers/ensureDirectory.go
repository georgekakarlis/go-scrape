package helpers

import "os"

// ensureDirectory checks if the directory exists, and if not, creates it.
func ensureDirectory(path string) error {
	// Check if the directory already exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}