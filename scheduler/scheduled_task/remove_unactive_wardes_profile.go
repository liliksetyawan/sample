package scheduled_task

import (
	"time"

	"github.com/nexsoft-git/nexlogger/log"
)

type removeUnactiveWardesProfileTask struct {
}

func (s *removeUnactiveWardesProfileTask) RunType() string {
	return "remove_unactive_wardes_profile"
}

func (s *removeUnactiveWardesProfileTask) Start() {
	log.Info().
		Msgf("Starting remove unactive wardes profile task at: %s", time.Now().Format(time.RFC3339))

	// Remove wardes profile that have status unactive
	// Assuming there's a service or repository method to handle this
	// Customary pattern: return value, error
	err := removeUnactiveWardesProfile()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to remove unactive wardes profiles")
		return
	}

	log.Info().
		Msg("Successfully removed unactive wardes profiles")
}

// removeUnactiveWardesProfile is a placeholder for the actual implementation
// that removes wardes profiles with unactive status
func removeUnactiveWardesProfile() error {
	// Implementation would typically involve:
	// 1. Querying the database for wardes profiles with unactive status
	// 2. Removing those profiles
	// 3. Returning any error that occurs
	
	// For now, this is a stub implementation
	// In real code, this would contain the actual removal logic
	return nil
}
