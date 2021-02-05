// // Copyright 2020 Google LLC
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     https://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

package main

// // [START spanner_read_data]

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func read() error {
	ctx := context.Background()

	db := "projects/sharechat-development/instances/ritu-test-instance/databases/ritu-test-db"

	fmt.Println(db)
	client, err := spanner.NewClient(ctx, db)
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
