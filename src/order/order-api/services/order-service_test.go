package services_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	mock_publishers "github.com/omiga-group/omiga/src/order/order-api/publishers/mock"
	mock_orderrepositories "github.com/omiga-group/omiga/src/order/order-api/repositories/mock"
	"github.com/omiga-group/omiga/src/order/order-api/services"
	"github.com/omiga-group/omiga/src/order/shared/models"
	"github.com/omiga-group/omiga/src/order/shared/repositories"
	mock_repositories "github.com/omiga-group/omiga/src/order/shared/repositories/mock"
)

func TestOrderServiceSubmit(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	RegisterFailHandler(Fail)
	RunSpecs(t, "Order Service Submit Tests")
}

var _ = Describe("Order Service Submit Tests", func() {
	var (
		mockCtrl            *gomock.Controller
		fakeEntgoClient     *mock_repositories.MockEntgoClient
		fakeOrderRepository *mock_orderrepositories.MockOrderRepository
		fakeOrderPublisher  *mock_publishers.MockOrderPublisher
		sut                 services.OrderService
		ctx                 context.Context
		expectedTransaction *repositories.Tx
		expectedOrder       models.Order
		expectedError       error
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		logger, err := zap.NewDevelopment()
		Ω(err).Should(BeNil())

		sugarLogger := logger.Sugar()

		fakeEntgoClient = mock_repositories.NewMockEntgoClient(mockCtrl)
		fakeOrderRepository = mock_orderrepositories.NewMockOrderRepository(mockCtrl)
		fakeOrderPublisher = mock_publishers.NewMockOrderPublisher(mockCtrl)

		sut, err = services.NewOrderService(
			sugarLogger,
			fakeEntgoClient,
			fakeOrderRepository,
			fakeOrderPublisher)
		Ω(err).Should(BeNil())

		ctx = context.Background()
		expectedTransaction = &repositories.Tx{}
		expectedOrder = models.Order{}
		expectedError = errors.New(cuid.New())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to submit order", func() {
		When("entgoClient.CreateTransaction method returns error", func() {
			It("should return error", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(nil, expectedError)

				_, err := sut.Submit(ctx, expectedOrder)
				Ω(err).Should(Equal(expectedError))
			})
		})

		When("transaction successfully created, but orderRepository.CreateOrder method returns error", func() {
			It("should return error and rollback transaction", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(nil)

				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedOrder, expectedError)

				_, err := sut.Submit(ctx, expectedOrder)
				Ω(err).Should(Equal(expectedError))
			})

			It("should not return the same error if entgoClient.RollbackTransaction failed", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(errors.New(cuid.New()))

				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedOrder, expectedError)

				_, err := sut.Submit(ctx, expectedOrder)
				Ω(err).Should(Equal(expectedError))
			})
		})

		When("order successfully save in database, but orderPublisher.Publish method returns error", func() {
			It("should return error and rollback transaction", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(nil)

				expectedCreatedOrder := models.Order{}
				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedCreatedOrder, nil)

				fakeOrderPublisher.
					EXPECT().
					Publish(
						ctx,
						expectedTransaction,
						nil,
						expectedCreatedOrder).
					Return(expectedError)
				_, err := sut.Submit(ctx, expectedOrder)

				Ω(err).Should(Equal(expectedError))
			})

			It("should not return the same error if entgoClient.RollbackTransaction failed", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(errors.New(cuid.New()))

				expectedCreatedOrder := models.Order{}
				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedCreatedOrder, nil)

				fakeOrderPublisher.
					EXPECT().
					Publish(
						ctx,
						expectedTransaction,
						nil,
						expectedCreatedOrder).
					Return(expectedError)
				_, err := sut.Submit(ctx, expectedOrder)

				Ω(err).Should(Equal(expectedError))
			})
		})

		When("order successfully published, but entgoClient.CommitTransaction method returns error", func() {
			It("should return error and rollback transaction", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(nil)

				fakeEntgoClient.
					EXPECT().
					CommitTransaction(expectedTransaction).
					Return(expectedError)

				expectedCreatedOrder := models.Order{}
				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedCreatedOrder, nil)

				fakeOrderPublisher.
					EXPECT().
					Publish(
						ctx,
						expectedTransaction,
						nil,
						expectedCreatedOrder).
					Return(nil)
				_, err := sut.Submit(ctx, expectedOrder)

				Ω(err).Should(Equal(expectedError))
			})

			It("should not return the same error if entgoClient.RollbackTransaction failed", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					RollbackTransaction(expectedTransaction).
					Return(errors.New(cuid.New()))

				fakeEntgoClient.
					EXPECT().
					CommitTransaction(expectedTransaction).
					Return(expectedError)

				expectedCreatedOrder := models.Order{}
				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedCreatedOrder, nil)

				fakeOrderPublisher.
					EXPECT().
					Publish(
						ctx,
						expectedTransaction,
						nil,
						expectedCreatedOrder).
					Return(nil)
				_, err := sut.Submit(ctx, expectedOrder)

				Ω(err).Should(Equal(expectedError))
			})
		})

		When("order successfully save in database, published and transaction committed successfully", func() {
			It("should return the created order", func() {
				fakeEntgoClient.
					EXPECT().
					CreateTransaction(ctx).
					Return(expectedTransaction, nil)

				fakeEntgoClient.
					EXPECT().
					CommitTransaction(expectedTransaction).
					Return(nil)

				expectedCreatedOrder := models.Order{}
				fakeOrderRepository.
					EXPECT().
					CreateOrder(ctx, expectedTransaction, expectedOrder).
					Return(expectedCreatedOrder, nil)

				fakeOrderPublisher.
					EXPECT().
					Publish(
						ctx,
						expectedTransaction,
						nil,
						expectedCreatedOrder).
					Return(nil)
				submittedOrder, err := sut.Submit(ctx, expectedOrder)

				Ω(err).Should(BeNil())
				Ω(submittedOrder).Should(Equal(expectedCreatedOrder))
			})
		})
	})
})
