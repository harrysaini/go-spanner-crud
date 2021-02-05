package test

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// [START spanner_read_data]

func read() error {
	ctx := context.Background()

	db := "projects/sharechat-development/instances/ritu-test-instance/databases/ritu-test-db"

	fmt.Println(db)
	client, err := spanner.NewClient(ctx, db, option.WithCredentialsFile("/Users/harishkumar/gcp_auth.json"))
	if err != nil {
		fmt.Println("a", err)
		return err
	}
	defer client.Close()

	iter := client.Single().Read(ctx, "Students", spanner.AllKeys(),
		[]string{"UUID", "RollNumber", "FirstName"})
	defer iter.Stop()

	fmt.Println(iter)

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			fmt.Println("c", err)
			return err
		}
		var rollNumber int64
		var UUID, firstName string
		if err := row.Columns(&UUID, &rollNumber, &firstName); err != nil {
			fmt.Println("b", err)
			return err
		}
		fmt.Printf("%s %d %s\n", UUID, rollNumber, firstName)
	}
}

// [END spanner_read_data]

func main() {
	err := read()
	log.Fatalln(err)
}
