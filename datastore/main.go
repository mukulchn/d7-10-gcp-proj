ckage main

import (
	    "context"
	        "fmt"
		    "log"
		        "os"

			    "cloud.google.com/go/datastore"
		    )

		    type Person struct {
			        Name string
			}

			func main() {
				    ctx := context.Background()

				        projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

					    // Creates a client.
					        client, err := datastore.NewClient(ctx, projectID)
						    if err != nil {
							            log.Fatalf("Failed to create client: %v", err)
								        }
									    defer client.Close()

									        kind := "Person"
										    name := "henryk"
										        personKey := datastore.NameKey(kind, name, nil)

											    person := Person{
												            Name: name,
													        }

														    // Saves the new entity.
														        if _, err := client.Put(ctx, personKey, &person); err != nil {
																        log.Fatalf("Failed to save person: %v", err)
																	    }

																	        fmt.Printf("Saved %v: %v\n", personKey, person.Name)
																	}
