package komikcast

import "bacakomik/config"

func Start(size int, config *config.Config) {
	ProcessReadChapter(size, config)
}
