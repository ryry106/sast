/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sast/usecase"

	"github.com/spf13/cobra"
)

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "display burndownchart from csv files in target dir.",
	Long:  "display burndownchart from csv files in target dir.",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		templateName, _ := cmd.Flags().GetString("template-name")
		startDate, _ := cmd.Flags().GetString("start-date")
		usecase.NewPreviewOption(args[0], port, templateName, startDate).Do()
	},
}

func init() {
	rootCmd.AddCommand(previewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// previewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	previewCmd.Flags().IntP("port", "p", 8080, "preview server port.")
	previewCmd.Flags().StringP("template-name", "t", "plane", "html template select.")
	previewCmd.Flags().StringP("start-date", "s", "", "burndownchart start date. format 2006-01-02. automatic start date in case of parse error.")
}
