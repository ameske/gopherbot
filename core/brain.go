package core

// A Brain is capapble of providing persistent storage for a gopherbot
type Brain interface {
	remember(key string, value interface{}) error
	recall(key string) (string, error)
	rememberHash(hash string, key string, value interface{}) error
	recallHash(hash string, key string) (string, error)
	recallHashAll(key string) ([]string, error)
}

var (
	brain        Brain
	RedisBrain   = newRedisBrain()
	DefaultBrain = RedisBrain
)

func init() {
	brain = DefaultBrain
}

func SetDefaultBrain(desired Brain) {
	brain = desired
}

func Remember(key string, value interface{}) error {
	return brain.remember(key, value)
}

func Recall(key string) (string, error) {
	return brain.recall(key)
}

func RememberHash(hash string, key string, value interface{}) error {
	return brain.rememberHash(hash, key, value)
}

func RecallHash(hash string, key string) (string, error) {
	return brain.recallHash(hash, key)
}

func RecallHashAll(hash string) ([]string, error) {
	return brain.recallHashAll(hash)
}
