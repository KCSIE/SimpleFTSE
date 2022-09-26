package cmd

import (
	"SimpleFTSE/ftse"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searching in documents' index",
	Long: `Use search command to find the words in the index: $ ./SimpleFTSE search -i "INDEXPATH" -p "DOCPATH" -q "KEYWORD"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'search' command called")
		indexp,err := cmd.Flags().GetString("index")
		if err != nil {
			fmt.Println("Please enter correct index path.")
			return
		}
		path,err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("Please enter correct doc path.")
			return
		}
		query,err := cmd.Flags().GetString("query")
		if err != nil {
			fmt.Println("Please enter correct query.")
			return
		}
		fmt.Println("Index path:",indexp,"Doc path:",path,"Keywords:",query)

		fmt.Println("Start Searching")

		// Load file
		start := time.Now()
		index, err := ftse.LoadIndex(indexp)
		if err != nil {
			fmt.Println(err)
		}
		docs, err := ftse.LoadDoc(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Loaded doc and index files in %v\n", time.Since(start))

		// Search word
		start = time.Now()
		matchedCnt := ftse.Qsearch(index,docs,query)
		fmt.Printf("Search found %d documents in %v\n", matchedCnt, time.Since(start))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().StringP("index","i","","Index's path (required)")
	searchCmd.Flags().StringP("path","p","","Documents' path for indexing (required)")
	searchCmd.Flags().StringP("query","q","","The keywords you are looking for (required)")
	searchCmd.MarkFlagRequired("index")
	searchCmd.MarkFlagRequired("path")
	searchCmd.MarkFlagRequired("query")
	searchCmd.MarkFlagsRequiredTogether("index", "path", "query")
}
