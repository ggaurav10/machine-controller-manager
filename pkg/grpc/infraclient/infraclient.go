/*
Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package infraclient

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/gardener/machine-controller-manager/pkg/grpc/infrapb"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// Address of grpc server
	addr = "127.0.0.1:50051"

	// API key
	key = ""

	// Authentication token
	token = ""

	// Path to a Google service account key file
	keyfile = ""

	// Audience
	audience = ""
)

// Create makes grpc call to create instance
func Create(name string) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInfragrpcClient(conn)

	if keyfile != "" {
		log.Printf("Authenticating using Google service account key in %s", keyfile)
		keyBytes, err := ioutil.ReadFile(keyfile)
		if err != nil {
			log.Fatalf("Unable to read service account key file %s: %v", keyfile, err)
		}

		tokenSource, err := google.JWTAccessTokenSourceFromJSON(keyBytes, audience)
		if err != nil {
			log.Fatalf("Error building JWT access token source: %v", err)
		}
		jwt, err := tokenSource.Token()
		if err != nil {
			log.Fatalf("Unable to generate JWT token: %v", err)
		}
		token = jwt.AccessToken
		// NOTE: the generated JWT token has a 1h TTL.
		// Make sure to refresh the token before it expires by calling TokenSource.Token() for each outgoing requests.
		// Calls to this particular implementation of TokenSource.Token() are cheap.
	}

	ctx := context.Background()
	if key != "" {
		log.Printf("Using API key: %s", key)
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", key))
	}
	if token != "" {
		log.Printf("Using authentication token: %s", token)
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", token)))
	}

	// Contact the server and print out its response.
	r, err := c.Create(ctx, &pb.CreateParams{Name: name})
	if err != nil {
		log.Fatalf("could not create infra: %v", err)
	}
	log.Printf("infra: %s", r.ProviderID)
}
