package models

import "llm_training_management_system/internal/router"

type LLMOrder struct {
	Name     string
	Args     []string
	Parallel bool
	IsShell  bool
	Data     router.Request
}
