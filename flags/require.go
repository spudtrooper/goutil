package flags

import (
	"log"
	"time"
)

// RequireString fails via `log.Fatalf` if the flag hasn't been set
func RequireString(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}

// RequireBool c fails via `log.Fatalf` if the flag hasn't been set
func RequireBool(flag *bool, name string) {
	if !*flag {
		log.Fatalf("--%s required", name)
	}

}

// RequireInt fails via `log.Fatalf` if the flag hasn't been set
func RequireInt(flag *int, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}

// RequireInt64  fails via `log.Fatalf` if the flag hasn't been set
func RequireInt64(flag *int64, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}

// RequireUint64 fails via `log.Fatalf` if the flag hasn't been set
func RequireUint64(flag *uint64, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}

// RequireUint c fails via `log.Fatalf` if the flag hasn't been set
func RequireUint(flag *uint, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}

// RequireFloat6 fails via `log.Fatalf` if the flag hasn't been set
func RequireFloat64(flag *float64, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}

// RequireDurati fails via `log.Fatalf` if the flag hasn't been set
func RequireDuration(flag *time.Duration, name string) {
	if *flag == 0 {
		log.Fatalf("--%s required", name)
	}
}
