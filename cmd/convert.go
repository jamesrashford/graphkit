/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	gio "github.com/jamesrashford/graphkit/io"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Used to convert different graph type files",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		inFormat, _ := cmd.Flags().GetString("if")
		outFormat, _ := cmd.Flags().GetString("of")
		directed, _ := cmd.Flags().GetBool("directed")
		comm, _ := cmd.Flags().GetString("comments")

		var read gio.GraphIO
		switch inFormat {
		case TypeName[EdgeListType]:
			read = gio.NewEdgeListIO(comm, "", directed)
		case TypeName[CSVType]:
			source, _ := cmd.Flags().GetString("source")
			target, _ := cmd.Flags().GetString("target")
			delim, _ := cmd.Flags().GetString("delimiter")
			read = gio.NewCSVIO(comm, delim, source, target, directed)
		case TypeName[GraphologyJSONType]:
			read = gio.NewGraphologyIO()
		case TypeName[NodeLinkJSONType]:
			read = gio.NewJSONIO()
		default:
			log.Fatalf("\"%s\" not valid (or supported) file type!\n", inFormat)
			os.Exit(0)
		}

		var write gio.GraphIO
		switch outFormat {
		case TypeName[EdgeListType]:
			write = gio.NewEdgeListIO(comm, "", directed)
		case TypeName[CSVType]:
			source, _ := cmd.Flags().GetString("source")
			target, _ := cmd.Flags().GetString("target")
			delim, _ := cmd.Flags().GetString("delimiter")
			write = gio.NewCSVIO(comm, delim, source, target, directed)
		case TypeName[GraphologyJSONType]:
			write = gio.NewGraphologyIO()
		case TypeName[NodeLinkJSONType]:
			write = gio.NewJSONIO()
		default:
			log.Fatalf("\"%s\" not valid (or supported) file type!\n", outFormat)
			os.Exit(0)
		}

		log.Printf("Reading graph '%s' of format '%s'...\n", input, inFormat)

		var inFile io.Reader

		if gio.IsURL(input) {
			fi, err := gio.ReadUrl(input)
			inFile = fi
			if err != nil {
				log.Fatal(err)
				os.Exit(0)
			}
		} else {
			fi, err := os.Open(input)
			inFile = fi
			if err != nil {
				log.Fatal(err)
				os.Exit(0)
			}
		}

		log.Println("Processing graph...")

		graph, err := read.ReadGraph(inFile)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		log.Printf("Writing graph to file '%s' of format '%s'...\n", output, outFormat)

		outFile, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		err = write.WriteGraph(graph, outFile)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}

		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	convertCmd.Flags().StringP("input", "i", "", "the input file")
	convertCmd.Flags().StringP("output", "o", "", "the output file")

	convertCmd.Flags().StringP("source", "s", "source", "the name of the source column (for CSV only)")
	convertCmd.Flags().StringP("target", "t", "target", "the name of the target column (for CSV only)")
	convertCmd.Flags().String("delimiter", ",", "the delimiter character used to separate values (for CSV only)")
	convertCmd.Flags().String("comments", "#", "the comments character used to indicate which rows are rto be ignored (for CSV and edgelist only)")

	description := "the file format of the input file. Supported file formats include:\n"
	for k, v := range TypeDescription {
		name := TypeName[k]
		line := fmt.Sprintf("- '%s': %s", name, v)
		description = fmt.Sprintf("%s%s\n", description, line)
	}
	convertCmd.Flags().String("if", "", description)

	description = "the file format of the output file. Supported file formats include:\n"
	for k, v := range TypeDescription {
		name := TypeName[k]
		line := fmt.Sprintf("- '%s': %s", name, v)
		description = fmt.Sprintf("%s%s\n", description, line)
	}
	convertCmd.Flags().String("of", "", description)

	convertCmd.Flags().BoolP("directed", "d", false, "directed graph")

	convertCmd.MarkFlagRequired("input")
	convertCmd.MarkFlagRequired("output")
	convertCmd.MarkFlagRequired("if")
	convertCmd.MarkFlagRequired("of")
}
