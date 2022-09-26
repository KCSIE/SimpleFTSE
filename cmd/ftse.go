package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"SimpleFTSE/ftse"
)

// ftseCmd represents the ftse command
var ftseCmd = &cobra.Command{
	Use:   "ftse",
	Short: "Indexing and searching",
	Long: `Input both documents' path and keywords to complete the whole search: $ ./SimpleFTSE ftse -p "DOCPATH" -q "KEYWORD"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("'ftse' command called")
		path,err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println("Please enter correct Doc path.")
			return
		}
		query,err := cmd.Flags().GetString("query")
		if err != nil {
			fmt.Println("Please enter correct query.")
			return
		}
		fmt.Println("Doc path:",path,"Keywords:",query)

		fmt.Println("Start Handling")

		// Load file
		start := time.Now()
		docs, err := ftse.LoadDoc(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Loaded %d documents in %v\n", len(docs), time.Since(start))

		// Create index
		start = time.Now()
		idx := ftse.NewIndex()
		idx.Add(docs)
		fmt.Printf("Indexed %d documents in %v\n", len(docs), time.Since(start))

		// Search word
		start = time.Now()
		matchedIDs := idx.Search(query)
		fmt.Printf("Search found %d documents in %v\n", len(matchedIDs), time.Since(start))
		for _, id := range matchedIDs {
			doc := docs[id]
			fmt.Printf("%d\t%s\n", id, doc.Text)
		}

	},
}

func init() {
	rootCmd.AddCommand(ftseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ftseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ftseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	ftseCmd.Flags().StringP("path","p","","Documents' path for indexing (required)")
	ftseCmd.Flags().StringP("query","q","","The keywords you are looking for (required)")
	ftseCmd.MarkFlagRequired("path")
	ftseCmd.MarkFlagRequired("query")
	ftseCmd.MarkFlagsRequiredTogether("path", "query")
}
