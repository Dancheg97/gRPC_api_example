# gRPC_api_example

Extract following command examples:

- Get description of services `grpcurl -proto api.proto describe`
- Get description of message `grpcurl -proto api.proto describe pb.Message`
- Send grpcurl message to service `grpcurl -proto api.proto -d '{\"message\":12.123}' -plaintext localhost:12201 pb.Basics/DoubleCall`



grpcurl -import-path . -proto C:\Users\dangd\OneDrive\Документы\gRPC_api_example\api.proto describe

grpcurl -import-path / -proto c:\Users\dangd\OneDrive\Документы\grpcclicker_flutter\api.proto -d '{\"message\": false}' -plaintext  localhost:12201 pb.v1.Basics.BoolCall