package kgsservice

type response struct {
	Messages []message
}

type message struct {
	Type      string
	Games     []game
	Rooms     []roomDesc
	SgfEvents []sgfEvent
	RoomDesc  *roomDesc
	RoomNames *roomNames
}

type sgfEvent struct {
	Type        string
	NodeID      int
	ChildNodeID int
	Props       []property
}

type property struct {
	Name  string
	Color string
	Loc   interface{}
	Text  string
	Int   int
	Float float64

	ByoYomiPeriods int
	Komi           float32
	TimeSystem     string
	Size           uint8
	MainTime       uint16
	Rules          string
	ByoYomiTime    uint16
}

type moveLocation struct {
	X int
	Y int
}

type room struct {
	Category  string
	ChannelID int
	Name      string
}

type game struct {
	GameType  string
	Komi      float32
	Size      int
	Saved     bool
	Handicap  int
	MoveNum   int
	Global    bool
	ChannelID int
	RoomID    int
}

type user struct {
	Name  string
	Flags string
}

type roomDesc struct {
	Description string
	Owners      []user
	Type        string
	ChannelID   int
}

type roomNames struct {
	Rooms []room
	Type  string
}
