/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jamesrashford/graphkit/io"
	"github.com/jamesrashford/graphkit/webui"
	"github.com/spf13/cobra"
)

type FileTypeSelect struct {
	FlagName    string
	Description string
}

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Visualising a graph in the web browser from file.",
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		filename, _ := cmd.Flags().GetString("input")
		fileFormat, _ := cmd.Flags().GetString("format")
		directed, _ := cmd.Flags().GetBool("directed")
		comm, _ := cmd.Flags().GetString("comments")

		var rw io.GraphIO
		switch fileFormat {
		case TypeName[EdgeListType]:
			rw = io.NewEdgeListIO(comm, "", directed)
		case TypeName[CSVType]:
			source, _ := cmd.Flags().GetString("source")
			target, _ := cmd.Flags().GetString("target")
			delim, _ := cmd.Flags().GetString("delimiter")
			rw = io.NewCSVIO(comm, delim, source, target, directed)
		case TypeName[GraphologyJSONType]:
			rw = io.NewGraphologyIO()
		case TypeName[NodeLinkJSONType]:
			rw = io.NewJSONIO()
		default:
			log.Fatalf("\"%s\" not valid (or supported) file type!\n", fileFormat)
			os.Exit(0)
		}

		log.Printf("Reading graph '%s' of format '%s'...\n", filename, fileFormat)

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		log.Println("Processing graph...")

		graph, err := rw.ReadGraph(file)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		log.Println("Starting WebUI...")

		ui := webui.NewWebUI(addr, *graph)
		ui.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webCmd.Flags().StringP("addr", "a", "0.0.0.0:8080", "the address and quote for the web server")
	webCmd.Flags().StringP("input", "i", "", "the input file to visualise")

	webCmd.Flags().StringP("source", "s", "source", "the name of the source column (for CSV only)")
	webCmd.Flags().StringP("target", "t", "target", "the name of the target column (for CSV only)")
	webCmd.Flags().String("delimiter", ",", "the delimiter character used to separate values (for CSV only)")
	webCmd.Flags().String("comments", "#", "the comments character used to indicate which rows are rto be ignored (for CSV and edgelist only)")

	description := "the file format of the input file to visualise. Supported file formats include:\n"
	for k, v := range TypeDescription {
		name := TypeName[k]
		line := fmt.Sprintf("- '%s': %s", name, v)
		description = fmt.Sprintf("%s%s\n", description, line)
	}

	webCmd.Flags().StringP("format", "f", "", description)
	webCmd.Flags().BoolP("directed", "d", false, "directed graph")

	webCmd.MarkFlagRequired("input")
	webCmd.MarkFlagRequired("format")

}
