package pb

//go:generate mockgen -destination sagagrpc/sagagrpcmock/tx.go -package sagagrpcmock github.com/saga-go/saga-go/pb/sagagrpc TxEventServiceClient
//go:generate mockgen -destination sagagrpc/sagagrpcmock/tx_stream.go -package sagagrpcmock github.com/saga-go/saga-go/pb/sagagrpc TxEventService_OnConnectedClient
