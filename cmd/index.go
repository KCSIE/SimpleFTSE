package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"SimpleFTSE/ftse"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Indexing the document",
	Long: `Use index command to transfer the text to tokens and build the index: $ ./SimpleFTSE index -p "DOCPATH"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'index' command called")
		path,err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("Please enter correct doc path.")
			return
		}
		fmt.Println("Doc path:",path)

		fmt.Println("Start Indexing")

		// Load file
		start := time.Now()
		docs, err := ftse.LoadDoc(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Loaded %d documents in %v\n", len(docs), time.Since(start))

		// Create index
		start = time.Now()
		ftse.GenIndex(docs)
		fmt.Printf("Indexed %d documents in %v\n", len(docs), time.Since(start))

	},
}

func init() {
	rootCmd.AddCommand(indexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// indexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// indexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	indexCmd.Flags().StringP("path","p","","Documents' path for indexing (required)")
	indexCmd.MarkFlagRequired("path")
}
