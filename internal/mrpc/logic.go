package mrpc

import (
	"gfAdmin/internal/client"
	"gfAdmin/internal/model"
	// "gfAdmin/internal/dbase"
	"fmt"
	"gfAdmin/internal/protorpc"
	"math/rand"
	"sync"
	"time"
)

type MatchResult struct {
	Roomkey   string
	Result    bool
	Iswiating bool
	Ip        string
}

const (
	playertype_ren  = 0
	playertype_Guai = 1
)

// /B
type CreareRoomResult struct {
	Roomkey   string
	Result    bool
	Iswiating bool
	Ip        string
	Port      int
}

type GameRoom struct {
	Key                       string
	Players/*[5]*/ [2]*Player //不直接使用Client 防止离线丢失
	Turn                      int
	Finished                  bool
	Matched                   bool
	GameType                  int
	SubServer                 *client.Client //room from this server
	Me                        *client.Client //dbase.UserInfo
}

const (
	GameType_Diw_Classic = 0 //第五人格_经典
)

type Player struct {
	UID         uint
	Client_     *client.Client //can nil! disconnect
	Player_type int
	// Userinfo_   dbase.UserInfo
	Userinfo_ *model.ContextUser
	ctx *model.Context
}
type room_control struct {
	lock  sync.Mutex
	rooms map[string]*GameRoom
}

var (
	gameRoomsCtrol room_control
	// create_room_result chan CreareRoomResult
)

func init() {
	fmt.Println("GameRoomsCtrol init()...")
	gameRoomsCtrol.rooms = make(map[string]*GameRoom)
	// create_room_result = make(chan CreareRoomResult)
}

func match_normal(c client.Client, player_type int) {
	me := createPlayer(&c, player_type)
	if me == nil {
		fmt.Println("createPlayer() failed")
		return
	}
	room, err := gameRoomsCtrol.join(me)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(room) //print for debug

}

func generateRoomID() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// /B
func createRoomW() (*GameRoom, error) {
	gameRoomsCtrol.lock.Lock()
	defer gameRoomsCtrol.lock.Unlock()

	room_key := generateRoomID()
	gameR := &GameRoom{
		Key: room_key,
		Players: /*[5]*/ [2]*Player{
			{UID: 0, Player_type: playertype_Guai},
			{UID: 0, Player_type: playertype_ren},
			///&Player{UID: 0, Player_type: playertype_ren},
			/*&Player{UID: 0, Player_type: playertype_ren},
			&Player{UID: 0, Player_type: playertype_ren},*/
		},
		Turn:     0,
		Finished: false,
		Matched:  false,
		GameType: GameType_Diw_Classic,
		// SubServer: randServer,
	}
	//games = append(games, game)
	gameRoomsCtrol.rooms[room_key] = gameR
	//CALL TO SERVER TO CREATE ROOM! AND WAIT IT
	//----->subServer
	room, err := gameR.CreateOnlineRoomWaitResult()
	if err != nil {
		return nil, err
	}
	//Mester <-- subserver backed
	fmt.Println(room) //print for debug
	//nextGameID++
	return gameR, nil
}

func createRoom() (*GameRoom, error) {
	//gameRoomsCtrol.lock.Lock()
	//defer gameRoomsCtrol.lock.Unlock()
	room_key := generateRoomID()
	gameR := &GameRoom{
		Key: room_key,
		Players: /*[5]*/ [2]*Player{
			{UID: 0, Player_type: playertype_Guai},
			{UID: 0, Player_type: playertype_ren},
			////{UID: 0, Player_type: playertype_ren},
			//{UID: 0, Player_type: playertype_ren},
			//{UID: 0, Player_type: playertype_ren},*/
		},
		Turn:     0,
		Finished: false,
		Matched:  false,
		GameType: GameType_Diw_Classic,
	}
	gameRoomsCtrol.rooms[room_key] = gameR
	return gameR, nil
}

func createPlayer(c *client.Client, player_type int) *Player {
	fmt.Println("createPlayer()")
	if c.User == nil {
		return nil
	}
	player := &Player{
		UID:         c.User.Id,
		Client_:     c,
		Player_type: player_type,
		Userinfo_:   c.User,
	}
	return player
}

func (game *GameRoom) CheckRoomFull() bool {
	for _, player := range game.Players {
		fmt.Println(player.Client_)
		if /*player.UID == 0 || */ player.Client_ == nil {
			return false
		}
	}
	return true
}

func (roomsctr *room_control) join(player *Player) (*GameRoom, error) {
	roomsctr.lock.Lock()
	defer roomsctr.lock.Unlock()
	fmt.Println(roomsctr.rooms)
	for _, room_ := range roomsctr.rooms {
		if room_.Matched {
			continue
		}
		fmt.Println("LOOP START")
		for _, p := range room_.Players {
			fmt.Println("LOOP START...")
			if p.Client_ == nil && /*p.UID == 0 &&*/ p.Player_type == player.Player_type {
				p.UID = player.UID
				p.Client_ = player.Client_
				p.Player_type = player.Player_type
				p.Userinfo_ = player.Userinfo_
				//check room full
				if room_.CheckRoomFull() {
					fmt.Println("room full ! CreateOnlineRoomWaitResult()!")
					// Room is full
					//人满了再->在服务器创建房间
					_, err := room_.CreateOnlineRoomWaitResult()
					if err != nil {
						return nil, err
					}
					//通知所有人，房间已满开始链接服务器？
					room_.Matched = true
					room_.SentToJoinRoom()
				}
				return room_, nil
			}

		}
	}
	//no active room for this type player
	//create room and join it
	fmt.Println("createRoom()")
	new_room_, err := createRoom() //这里只是MasterServer创建房间
	if err != nil {
		return nil, err
	}
	//join it
	for _, _p := range new_room_.Players {
		if _p.UID == 0 && _p.Player_type == player.Player_type {
			_p.UID = player.UID
			_p.Client_ = player.Client_
			_p.Player_type = player.Player_type
			_p.Userinfo_ = player.Userinfo_
			//check room full
			return new_room_, nil
		}
	}
	return nil, fmt.Errorf("kunknown error")
}

//func (game *GameRoom) leave(player *Player) {
//}

// ALL Player JoinRoom to SubServer
// (Send message to all player in room)
func (game *GameRoom) SentToJoinRoom() {
	fmt.Println("SentToJoinRoom()")
	for _, player := range game.Players {
		go Rpc_client_match_ok(player.Client_,
			&protorpc.MatchOKParam{
				Result:  true,
				RoomKey: game.Key,
				MatchServer: &protorpc.Server{
					Ip:   game.SubServer.Service.IP,
					Port: int32(game.SubServer.Service.Port),
					Name: "TestSubServer",
				},
				Room: &protorpc.GameRoom{
					Key:      game.Key,
					Ip:       game.SubServer.Service.IP,
					Port:     int32(game.SubServer.Service.Port),
					RoomName: "TestRoomName",
					Finished: false,
					Playeras: []*protorpc.Player{
						{
							Uid: int32(game.Players[0].UID),
						},
						{
							Uid: int32(game.Players[1].UID),
						},
					},
				},
			})
	}
}

func (game *GameRoom) CreateOnlineRoomWaitResult() (*protorpc.CreateRoomResult, error) {
	//find sub server client
	LIST := control.GetSubServerClientList()
	if LIST == nil || len(*LIST) == 0 {
		return nil, fmt.Errorf("no sub server client found")
	}
	sub_server_cli := (*LIST)[0]
	if sub_server_cli == nil {
		return nil, fmt.Errorf("sub server client is nil")
	}
	game.SubServer = sub_server_cli //set sub server client to game room
	//Create room req to sub server
	fmt.Println("CreateOnlineRoomWaitResult()")
	result, err := Rpc_client_create_room(sub_server_cli, &protorpc.CreateRoomParam{
		Key: game.Key,
		Room: &protorpc.GameRoom{
			Key:      game.Key,
			RoomName: game.Key,
		},
	})
	if err != nil || result == nil {
		return nil, fmt.Errorf("create room failed: %v", err)
	}
	return result, nil
}
