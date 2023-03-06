package flags

import (
	"log"
	"os"
	"strconv"

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
	if *variable == "" {
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
	cmd.Flags().IntVarP(variable, flag.FullName, flag.ShortName, 0, flag.Usage)

	// Flag not set, try env var
	if *variable == 0 {
		tmp := os.Getenv(flag.EnvVar)

		// If it isn't set, that's ok.
		if tmp == "" {
			return
		}

		intVal, err := strconv.Atoi(tmp)
		if err != nil {
			log.Fatal(flag.FullName + " needs to be an integer")
		}

		*variable = intVal
	}

	// Flag and env var not set, look for default
	if *variable == 0 && flag.Default != nil {
		if intVal, ok := flag.Default.(int); ok {
			*variable = intVal
		}
	}
}
