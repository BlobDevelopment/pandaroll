package flags

import (
	"os"
	"strconv"

	"blobdev.com/pandaroll/internal/logger"
	"github.com/spf13/cobra"
)

type Flag struct {
	FullName  string
	ShortName string
	Usage     string
	EnvVar    string
	Default   any
}

func NewStringFlag(cmd *cobra.Command, variable *string, flag Flag) {
	cmd.PersistentFlags().StringVarP(variable, flag.FullName, flag.ShortName, "", flag.Usage)

	// Flag not set, try env var
	if *variable == "" && flag.EnvVar != "" {
		*variable = os.Getenv(flag.EnvVar)
	}

	// Flag and env var not set, look for default
	if *variable == "" && flag.Default != nil {
		if str, ok := flag.Default.(string); ok {
			*variable = str
		}
	}
}

func NewIntFlag(cmd *cobra.Command, variable *int, flag Flag) {
	cmd.PersistentFlags().IntVarP(variable, flag.FullName, flag.ShortName, 0, flag.Usage)

	// Flag not set, try env var
	if *variable == 0 {
		tmp := os.Getenv(flag.EnvVar)

		// If the env var is set, try to use it
		if tmp != "" {
			intVal, err := strconv.Atoi(tmp)
			if err != nil {
				logger.Fatal(flag.FullName + " needs to be an integer")
				return
			}

			*variable = intVal
		}
	}

	// Flag and env var not set, look for default
	if *variable == 0 && flag.Default != nil {
		if intVal, ok := flag.Default.(int); ok {
			*variable = intVal
		}
	}
}
