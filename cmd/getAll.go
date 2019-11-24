package cmd

import (
	"log"

	"context"

	"github.com/spf13/cobra"
	pb "github.com/tombrookson/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

// getAllCmd represents the getAll command
var getAllCmd = &cobra.Command{
	Use:   "getAll",
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

		getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
		if err != nil {
			log.Fatalf("Could not list consignments: %v", err)
		}
		for _, v := range getAll.Consignments {
			log.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}
