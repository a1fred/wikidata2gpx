// +build NOCOMPLETION

package cmd

func init() {
	// Disable cobra completion command on build
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
