package commands

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	pulsaradminv2 "github.com/omiga-group/omiga/src/shared/clients/openapi/pulsar/admin/v2"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/omigactl/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type setupPulsarOptions struct {
	tenant         string
	namespace      string
	partitionCount int
}

func setupPulsarCommand() *cobra.Command {
	opt := setupPulsarOptions{}
	cmd := &cobra.Command{
		Use:   "setup-pulsar",
		Short: "Setup pulsar",
		Long:  "Setup pulsar",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			pulsarAdminClient, err := pulsaradminv2.NewClientWithResponses(config.Pulsar.HttpUrl)
			if err != nil {
				sugarLogger.Fatalf("Failed to create pulsar admin client. Error: %v", err)
				return
			}

			if err := setupOrderBookTopic(ctx, pulsarAdminClient, opt); err != nil {
				sugarLogger.Fatalf(
					"Failed to setup Order Book topic. Error: %v",
					err)
				return
			}

			if err := setupOrderTopic(ctx, pulsarAdminClient, opt); err != nil {
				sugarLogger.Fatalf(
					"Failed to setup Order topic. Error: %v",
					err)
				return
			}

			if err := setupSyntheticOrderTopic(ctx, pulsarAdminClient, opt); err != nil {
				sugarLogger.Fatalf(
					"Failed to setup Synthetic Order topic. Error: %v",
					err)
				return
			}
		},
	}

	cmd.Flags().StringVar(&opt.tenant, "tenant", "public", "Specify the pulsar tenant name, default is public")
	cmd.Flags().StringVar(&opt.namespace, "namespace", "default", "Specify the pulsar namespace name, default is default")
	cmd.Flags().IntVar(&opt.partitionCount, "partition-count", 12, "Specify the number of partition, default is 12")

	return cmd
}

func setupOrderBookTopic(
	ctx context.Context,
	pulsarAdminClient pulsaradminv2.ClientWithResponsesInterface,
	options setupPulsarOptions,
) error {
	if err := createTopic(
		ctx,
		pulsarAdminClient,
		options,
		orderbookv1.TopicName); err != nil {
		return err
	}

	if err := createSchema(
		ctx,
		pulsarAdminClient,
		options,
		orderbookv1.TopicName,
		orderbookv1.GetAvroSchema()); err != nil {
		return err
	}

	return nil
}

func setupOrderTopic(
	ctx context.Context,
	pulsarAdminClient pulsaradminv2.ClientWithResponsesInterface,
	options setupPulsarOptions,
) error {
	if err := createTopic(
		ctx,
		pulsarAdminClient,
		options,
		orderv1.TopicName); err != nil {
		return err
	}

	if err := createSchema(
		ctx,
		pulsarAdminClient,
		options,
		orderv1.TopicName,
		orderv1.GetAvroSchema()); err != nil {
		return err
	}

	return nil
}

func setupSyntheticOrderTopic(
	ctx context.Context,
	pulsarAdminClient pulsaradminv2.ClientWithResponsesInterface,
	options setupPulsarOptions,
) error {
	if err := createTopic(
		ctx,
		pulsarAdminClient,
		options,
		syntheticorderv1.TopicName); err != nil {
		return err
	}

	if err := createSchema(
		ctx,
		pulsarAdminClient,
		options,
		syntheticorderv1.TopicName,
		syntheticorderv1.GetAvroSchema()); err != nil {
		return err
	}

	return nil
}

func createTopic(
	ctx context.Context,
	pulsarAdminClient pulsaradminv2.ClientWithResponsesInterface,
	options setupPulsarOptions,
	topicName string) error {
	response, err := pulsarAdminClient.CreatePersistentTopicPartitionedTopicWithBodyWithResponse(
		ctx,
		options.tenant,
		options.namespace,
		topicName,
		&pulsaradminv2.CreatePersistentTopicPartitionedTopicParams{},
		"application/json",
		strings.NewReader(strconv.Itoa(options.partitionCount)))
	if err != nil {
		return fmt.Errorf("failed to create topic %s. Error: %v",
			topicName,
			err)
	}

	if response.HTTPResponse.StatusCode != 204 &&
		response.HTTPResponse.StatusCode != 409 {
		return fmt.Errorf("failed to create topic %s. Return status code is %d",
			topicName,
			response.HTTPResponse.StatusCode)
	}

	return nil
}

func createSchema(
	ctx context.Context,
	pulsarAdminClient pulsaradminv2.ClientWithResponsesInterface,
	options setupPulsarOptions,
	topicName string,
	schema string) error {
	schemaType := "JSON"

	response, err := pulsarAdminClient.PostSchemaWithResponse(
		ctx,
		options.tenant,
		options.namespace,
		syntheticorderv1.TopicName,
		&pulsaradminv2.PostSchemaParams{},
		pulsaradminv2.PostSchemaPayload{
			Schema: &schema,
			Type:   &schemaType,
		},
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update schema on topic %s. Error: %v",
			topicName,
			err)
	}

	if response.HTTPResponse.StatusCode != 202 &&
		response.HTTPResponse.StatusCode != 409 {
		return fmt.Errorf("failed to update schema on topic %s. Return status code is %d",
			topicName,
			response.HTTPResponse.StatusCode)
	}

	return nil
}
