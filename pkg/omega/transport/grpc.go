package transport

import (
	"context"

	"github.com/pkg/errors"
	"github.com/saga-go/saga-go/pb/sagagrpc"
	"google.golang.org/grpc"
)

// SagaTransactionTransport Transaction Transport (Saga)
//
// The transaction transport module is responsible for communication between Omega and Alpha, and in the specific implementation process,
// Pack defines the transaction interaction methods of TCC and Saga
// by defining the relevant Grpc description interface file, as well as the events associated with the interaction.
// We enabled mutual calls between Omega and Alpha with the help of the two-way flow interface provided by Grpc.
//
// Omega and Alpha's transmissions are based on Grpc multilingual support and provide the foundation for a multilingual version of Omega.
type SagaTransactionTransport struct {
	cfg      sagagrpc.GrpcServiceConfig
	txc      sagagrpc.TxEventServiceClient
	txstream sagagrpc.TxEventService_OnConnectedClient
}
