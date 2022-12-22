/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Flags().GetString("")
		// if fileStat.IsDir() {
		// 	filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		// 		if err != nil {
		// 			return err
		// 		}
		// 		if !info.IsDir() {
		// 			ext := extensionCleaned(path)
		// 			extensionsMap[ext] = append(extensionsMap[ext], path)
		// 			extensionCounterStats[ext] += 1
		// 		}
		// 		return nil
		// 	})
		// }

	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	transferCmd.Flags().StringP("", "", "", "")
}
