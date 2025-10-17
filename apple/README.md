# apple

The `apple` package facilitates integration with Apple's "Sign in with Apple" service. It provides functionalities for handling JSON Web Tokens (JWTs) issued by Apple, retrieving public keys for verification, generating client secrets, validating authorization and refresh tokens, and revoking tokens.

## Features

*   **JWT Parsing and Validation:** Parses and validates JWTs issued by Apple using their public keys.
*   **Public Key Retrieval:** Fetches Apple's public keys from their authentication server.
*   **Client Secret Generation:** Generates the client secret required for secure communication with Apple's authentication endpoints.
*   **Token Validation:** Supports validation of authorization codes for both web and native applications, as well as refreshing access tokens.
*   **Token Revocation:** Allows for the revocation of refresh or access tokens.

## Usage Example

Due to the complexity of setting up "Sign in with Apple" (requiring Apple Developer account, service IDs, keys, etc.), a full runnable example is not provided here. Instead, conceptual usage examples for key functionalities are outlined.

### 1. Generating a Client Secret

```go
package main

import (
    "fmt"
    "log"

    "github.com/skyrocket-qy/ciri/apple"
)

func main() {
	// Replace with your actual Apple Developer credentials
	signingKey := `-----BEGIN PRIVATE KEY-----
... your private key ...
-----END PRIVATE KEY-----`
	teamID := "YOUR_TEAM_ID"
	clientID := "YOUR_SERVICE_ID" // e.g., com.example.app
	keyID := "YOUR_KEY_ID"

	clientSecret, err := apple.GenerateClientSecret(signingKey, teamID, clientID, keyID)
	if err != nil {
		log.Fatalf("Error generating client secret: %v", err)
	}
	fmt.Printf("Generated Client Secret: %s\n", clientSecret)
}
```

### 2. Parsing and Validating an Apple JWT

```go
package main

import (
	"fmt"
	"log"

	"github.com/skyrocket-qy/ciri/apple"
)

func main() {
	// This is a placeholder JWT. In a real scenario, you would receive this from Apple.
	// You would need a valid JWT from Apple's authentication flow to test this.
	appleJWT := "eyJhbGciOiJFUzI1NiIsImtpZCI6IkFCRU5DRUwxMjM0In0.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLmV4YW1wbGUuYXBwIiwiaWF0IjoxNjcwMDAwMDAwLCJleHAiOjE2NzAwMDM2MDAsInN1YiI6IjAwMDAwMC4xMjM0NTY3ODkwYWJjZGVmIiwic2lnbmF0dXJlIjoiaW52YWxpZCJ9.INVALID_SIGNATURE"

	claims, err := apple.ParseAppleJWT(appleJWT)
	if err != nil {
		log.Fatalf("Error parsing Apple JWT: %v", err)
	}

	fmt.Printf("JWT Claims: %+v\n", claims)
	// You would then typically verify the 'iss', 'aud', 'exp', and 'sub' claims.
}
```

### 3. Validating an Authorization Code (Web)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/skyrocket-qy/ciri/apple"
)

func main() {
	client := apple.New()

	req := apple.WebValidationTokenRequest{
		ClientID:     "YOUR_SERVICE_ID",
		ClientSecret: "GENERATED_CLIENT_SECRET", // Use the secret generated above
		Code:         "AUTHORIZATION_CODE_FROM_APPLE",
		RedirectURI:  "YOUR_REDIRECT_URI",
	}

	var validationResponse apple.ValidationResponse
	err := client.VerifyWebToken(context.Background(), req, &validationResponse)
	if err != nil {
		log.Fatalf("Error validating web token: %v", err)
	}

	if validationResponse.Error != "" {
		log.Fatalf("Apple returned an error: %s - %s", validationResponse.Error, validationResponse.ErrorDescription)
	}

	fmt.Printf("Validation Successful! Access Token: %s, Refresh Token: %s\n",
		validationResponse.AccessToken, validationResponse.RefreshToken)
}
```

### 4. Revoking a Token

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/skyrocket-qy/ciri/apple"
)

func main() {
	revokeClient := apple.NewRVKClient()

	req := apple.RevokeTokenRequest{
		ClientID:     "YOUR_SERVICE_ID",
		ClientSecret: "GENERATED_CLIENT_SECRET", // Use the secret generated above
		Token:        "REFRESH_OR_ACCESS_TOKEN_TO_REVOKE",
		TokenTypeHint: apple.RefreshTokenTypeHint, // or apple.AccessTokenTypeHint
	}

	err := revokeClient.RevokeToken(context.Background(), req)
	if err != nil {
		log.Fatalf("Error revoking token: %v", err)
	}

	fmt.Println("Token revoked successfully.")
}
```