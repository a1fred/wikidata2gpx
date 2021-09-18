package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/cmd/filterCmd"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/cmd/wikidataCmd"
	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/utils"
)

var rootCmd = &cobra.Command{
	Use:   "wikidata2gpx",
	Short: "Wikidata pois exporter",
	Long:  `Wikidata pois exporter`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		utils.ErrCheck(err)
	},
}

func init() {
	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(wikidataCmd.WikidataCmd)
	rootCmd.AddCommand(filterCmd.FilterCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
