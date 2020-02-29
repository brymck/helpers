package secrets

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	"github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

// AccessSecret retrieves secret data, first by checking to see if an environment variable with the equivalent name
// exists (all caps with dashes replaced by underscores) and using its value, secondly by requesting the data from
// Secret Manager.
func AccessSecret(secretId string) (string, error) {
	log := logrus.WithField("secretId", secretId)
	envKey := strings.ReplaceAll(strings.ToUpper(secretId), "-", "_")
	apiKeyFromEnv := os.Getenv(envKey)
	if apiKeyFromEnv != "" {
		log.Infof("using value of environment variable %s for secret %s", envKey, secretId)
		return apiKeyFromEnv, nil
	}

	projectId, err := metadata.ProjectID()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve project ID from metadata server: %w", err)
	}
	log = log.WithField("project", projectId)

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to set up client: %w", err)
	}

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectId, secretId)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{Name: name}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", fmt.Errorf("failed to access secret: %v", err)
	}
	log.Infof("used Secrets Manager to retrieve secret %s", envKey)
	return string(result.Payload.Data), nil
}
