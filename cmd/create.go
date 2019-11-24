package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/tombrookson/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

// createCmd represents the create command
var (
	consignmentFile string
	createCmd       = &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Set up a connection to the server.
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewShippingServiceClient(conn)

			// Contact the server and print out its response.
			consignment, err := parseFile(consignmentFile)
			log.Printf("File: %v", consignment)

			r, err := client.CreateConsignment(context.Background(), consignment)
			if err != nil {
				log.Fatalf("Could not greet: %v", err)
			}
			log.Printf("Created: %t", r.Created)
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&consignmentFile, "file", "f", "file for consignment")
	createCmd.MarkFlagRequired("file")
}
