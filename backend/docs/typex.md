```
├── github.com
│   ├── J0hnLenin
│   │   └── Tesuji
│   │       ├── config
│   │       │   ├── Config struct {
│   │       │   │       Database config.DatabaseConfig          `yaml:"database"`
│   │       │   │       Kafka config.KafkaConfig                `yaml:"kafka"`
│   │       │   │       Redis config.RedisConfig                `yaml:"redis"`
│   │       │   │       GameServers config.GameServersConfig            `yaml:"gameServersConfig"`
│   │       │   │   }
│   │       │   ├── DatabaseConfig struct {
│   │       │   │       Host string             `yaml:"host"`
│   │       │   │       Port int                `yaml:"port"`
│   │       │   │       Username string         `yaml:"username"`
│   │       │   │       Password string         `yaml:"password"`
│   │       │   │       DBName string           `yaml:"name"`
│   │       │   │       SSLMode string          `yaml:"ssl_mode"`
│   │       │   │       BucketQuantity uint16           `yaml:"bucket_quantity"`
│   │       │   │   }
│   │       │   ├── GameServersConfig struct {
│   │       │   │       KGS config.KGSConfig            `yaml:"kgs"`
│   │       │   │   }
│   │       │   ├── KGSConfig struct {
│   │       │   │       URL string              `yaml:"url"`
│   │       │   │       User string             `yaml:"user"`
│   │       │   │       Password string         `yaml:"password"`
│   │       │   │       Timeout uint16          `yaml:"timeout"`
│   │       │   │       SrvID uint8             `yaml:"srv_id"`
│   │       │   │   }
│   │       │   ├── KafkaConfig struct {
│   │       │   │       Host string             `yaml:"host"`
│   │       │   │       Port int                `yaml:"port"`
│   │       │   │       GameTopicName string            `yaml:"game_topic_name"`
│   │       │   │   }
│   │       │   └── RedisConfig struct {
│   │       │           Host string             `yaml:"host"`
│   │       │           Port int                `yaml:"port"`
│   │       │           Password string         `yaml:"password"`
│   │       │           DBStr string            `yaml:"dbStr"`
│   │       │       }
│   │       └── internal
│   │           ├── api
│   │           │   └── gameserviceapi
│   │           │       ├── GameServiceAPI struct {
│   │           │       │       games_api.UnimplementedGamesServiceServer
│   │           │       │   }
│   │           │       └── gameService interface {
│   │           │               GetGameByID(ctx context.Context, ID int64) (*models.Game, error)
│   │           │               GetGamesSummaryList(ctx context.Context, page uint64, limit uint64) ([]*models.GameSummary, error)
│   │           │           }
│   │           ├── models
│   │           │   ├── Color bool
│   │           │   ├── Game struct {
│   │           │   │       models.GameSummary
│   │           │   │       GameData []models.GameState
│   │           │   │   }
│   │           │   ├── GameState struct {
│   │           │   │       LastMove *models.Move
│   │           │   │       Position models.Position
│   │           │   │   }
│   │           │   ├── GameSummary struct {
│   │           │   │       ID uint64
│   │           │   │       BoardSize uint8
│   │           │   │       Title string
│   │           │   │       Komi float32
│   │           │   │       Date time.Time
│   │           │   │       IsFinished bool
│   │           │   │       Rules models.Rules
│   │           │   │       BlackPlayer *models.Player
│   │           │   │       WhitePlayer *models.Player
│   │           │   │   }
│   │           │   ├── Move struct {
│   │           │   │       Point *models.Point
│   │           │   │       IsPass bool
│   │           │   │       Color models.Color
│   │           │   │   }
│   │           │   ├── Player struct {
│   │           │   │       Name string
│   │           │   │       Rank uint8
│   │           │   │   }
│   │           │   ├── Point struct {
│   │           │   │       X uint8
│   │           │   │       Y uint8
│   │           │   │   }
│   │           │   ├── PointState uint8
│   │           │   ├── Position [][]models.PointState
│   │           │   └── Rules uint8
│   │           ├── pb
│   │           │   ├── games_api
│   │           │   │   ├── GamesServiceClient interface {
│   │           │   │   │       GetGameByID(ctx context.Context, in *games_api.GetGameByIDRequest, opts ...grpc.CallOption) (*games_api.GetGameByIDResponse, error)
│   │           │   │   │       GetGamesSummaryList(ctx context.Context, in *games_api.GetGamesSummaryListRequest, opts ...grpc.CallOption) (*games_api.GetGamesSummaryListResponse, error)
│   │           │   │   │   }
│   │           │   │   ├── GamesServiceServer interface {
│   │           │   │   │       GetGameByID(context.Context, *games_api.GetGameByIDRequest) (*games_api.GetGameByIDResponse, error)
│   │           │   │   │       GetGamesSummaryList(context.Context, *games_api.GetGamesSummaryListRequest) (*games_api.GetGamesSummaryListResponse, error)
│   │           │   │   │   }
│   │           │   │   ├── GetGameByIDRequest struct {
│   │           │   │   │       Id int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
│   │           │   │   │   }
│   │           │   │   ├── GetGameByIDResponse struct {
│   │           │   │   │       Game *models.Game               `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
│   │           │   │   │   }
│   │           │   │   ├── GetGamesSummaryListRequest struct {
│   │           │   │   │       Page int64              `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
│   │           │   │   │       Limit int64             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
│   │           │   │   │   }
│   │           │   │   ├── GetGamesSummaryListResponse struct {
│   │           │   │   │       Games []*models.GameSummary             `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
│   │           │   │   │   }
│   │           │   │   ├── UnimplementedGamesServiceServer struct {}
│   │           │   │   └── UnsafeGamesServiceServer interface {}
│   │           │   └── models
│   │           │       ├── Color int32
│   │           │       ├── Game struct {
│   │           │       │       Summary *models.GameSummary             `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
│   │           │       │       GameData []*models.GameState            `protobuf:"bytes,2,rep,name=game_data,json=gameData,proto3" json:"game_data,omitempty"`
│   │           │       │   }
│   │           │       ├── GameState struct {
│   │           │       │       LastMove *models.Move           `protobuf:"bytes,1,opt,name=last_move,json=lastMove,proto3" json:"last_move,omitempty"`
│   │           │       │       Position *models.Position               `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
│   │           │       │   }
│   │           │       ├── GameSummary struct {
│   │           │       │       Id int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
│   │           │       │       BoardSize int32         `protobuf:"varint,2,opt,name=board_size,json=boardSize,proto3" json:"board_size,omitempty"`
│   │           │       │       Title string            `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
│   │           │       │       Komi float32            `protobuf:"fixed32,4,opt,name=komi,proto3" json:"komi,omitempty"`
│   │           │       │       Date string             `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
│   │           │       │       IsFinished bool         `protobuf:"varint,6,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
│   │           │       │       Rules models.Rules              `protobuf:"varint,7,opt,name=rules,proto3,enum=tesuji.models.v1.Rules" json:"rules,omitempty"`
│   │           │       │       BlackPlayer *models.Player              `protobuf:"bytes,8,opt,name=black_player,json=blackPlayer,proto3" json:"black_player,omitempty"`
│   │           │       │       WhitePlayer *models.Player              `protobuf:"bytes,9,opt,name=white_player,json=whitePlayer,proto3" json:"white_player,omitempty"`
│   │           │       │   }
│   │           │       ├── Move struct {
│   │           │       │       Point *models.Point             `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
│   │           │       │       IsPass bool             `protobuf:"varint,2,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
│   │           │       │       Color models.Color              `protobuf:"varint,3,opt,name=color,proto3,enum=tesuji.models.v1.Color" json:"color,omitempty"`
│   │           │       │   }
│   │           │       ├── Player struct {
│   │           │       │       Name string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
│   │           │       │       Rank int32              `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
│   │           │       │   }
│   │           │       ├── Point struct {
│   │           │       │       X int32         `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
│   │           │       │       Y int32         `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
│   │           │       │   }
│   │           │       ├── Position struct {
│   │           │       │       Data []int32            `protobuf:"varint,1,rep,packed,name=data,proto3" json:"data,omitempty"`
│   │           │       │   }
│   │           │       └── Rules int32
│   │           ├── producer
│   │           │   └── gameprod
│   │           │       └── GameProducer struct {}
│   │           ├── services
│   │           │   ├── gamelogic
│   │           │   │   └── GameLogic struct {}
│   │           │   ├── gameservers
│   │           │   │   └── kgsservice
│   │           │   │       ├── GameService interface {
│   │           │   │       │       InitGameData(ctx context.Context, game *models.Game) error
│   │           │   │       │       MakeAMove(ctx context.Context, game *models.Game, move *models.Move) error
│   │           │   │       │       UpsertGamesData(ctx context.Context, games []*models.Game) error
│   │           │   │       │   }
│   │           │   │       └── KGSService struct {}
│   │           │   ├── gameservice
│   │           │   │   ├── GameCache interface {
│   │           │   │   │       GetGameByID(ctx context.Context, id int64) (*models.Game, error)
│   │           │   │   │       UpsertGamesData(ctx context.Context, games []*models.Game) error
│   │           │   │   │   }
│   │           │   │   ├── GameProducer interface {
│   │           │   │   │       ProduceGames(ctx context.Context, games []*models.Game) error
│   │           │   │   │   }
│   │           │   │   ├── GameService struct {}
│   │           │   │   ├── GameStorage interface {
│   │           │   │   │       GetGameByID(ctx context.Context, id int64) (*models.Game, error)
│   │           │   │   │       GetGamesSummaryList(ctx context.Context, page uint64, limit uint64) ([]*models.GameSummary, error)
│   │           │   │   │       UpsertGamesData(ctx context.Context, games []*models.Game) error
│   │           │   │   │   }
│   │           │   │   ├── gameLogic interface {
│   │           │   │   │       InitGameData(ctx context.Context, game *models.Game) error
│   │           │   │   │       MakeAMove(ctx context.Context, game *models.Game, move *models.Move) error
│   │           │   │   │   }
│   │           │   │   ├── mocks
│   │           │   │   │   ├── GameCache struct {
│   │           │   │   │   │       mock.Mock
│   │           │   │   │   │   }
│   │           │   │   │   ├── GameProducer struct {
│   │           │   │   │   │       mock.Mock
│   │           │   │   │   │   }
│   │           │   │   │   └── GameStorage struct {
│   │           │   │   │           mock.Mock
│   │           │   │   │       }
│   │           │   │   └── protoConvert interface {
│   │           │   │           MapGamesToProto(games []*models.Game) []*models.Game
│   │           │   │       }
│   │           │   └── protoconvert
│   │           │       └── ProtoConvert struct {}
│   │           └── storage
│   │               ├── pgstorage
│   │               │   └── PGstorage struct {}
│   │               └── redisstorage
│   │                   └── RedisStorage struct {}
```