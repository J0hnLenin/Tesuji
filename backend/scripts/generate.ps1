cd $PSScriptRoot/..

# Генерация gRPC кода
protoc -I ./api `
  -I ./api/google/api `
  --go_out=./internal/pb --go_opt=paths=source_relative `
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative `
  ./api/games_api/games.proto ./api/models/game_model.proto

# Генерация gRPC-Gateway
protoc -I ./api `
  -I ./api/google/api `
  --grpc-gateway_out=./internal/pb `
  --grpc-gateway_opt paths=source_relative `
  --grpc-gateway_opt logtostderr=true `
  ./api/games_api/games.proto

# Генерация OpenAPI
protoc -I ./api `
  -I ./api/google/api `
  --openapiv2_out=./internal/pb/swagger `
  --openapiv2_opt logtostderr=true `
  ./api/games_api/games.proto