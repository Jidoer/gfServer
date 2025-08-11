package mrpc

import (
	// "context"
	"errors"
	"gfAdmin/internal/client"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
	"gfAdmin/internal/tool"
	"strings"

	// "gfAdmin/internal/service"

	// "gfAdmin/internal/dbase"
	"fmt"
	"gfAdmin/internal/notify"
	"gfAdmin/internal/protorpc"
	"math/rand"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
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
	// SubServer                 *client.Client //room from this server
	IP   string
	Port int
	Me   *client.Client //dbase.UserInfo
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
	ctx       *model.Context
	UUID      string //room key
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
					out, err := room_.CreateOnlineRoomWaitResult()
					if err != nil {
						return nil, err
					}
					fmt.Println("[OUT]:", out)
					ports := map[string]string{}
					for _, p := range out.GameServerPorts {
						parts := strings.Split(p, "/")
						if len(parts) >= 2 {
							ports[strings.ToUpper(parts[1])] = parts[0]
						}
					}
					fmt.Println("UDP:", ports["UDP"], "TCP:", ports["TCP"])
					port := tool.String2Int(ports["UDP"])
					room_.IP = out.IP
					room_.Port = port
					//通知所有人，房间已满开始链接服务器？ !需要服务器ready才行！！
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
					Ip:   game.IP,
					Port: int32(game.Port),
					Name: "TestSubServer",
				},
				Room: &protorpc.GameRoom{
					Key:      game.Key,
					Ip:       game.IP,
					Port:     int32(game.Port),
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

// func (game *GameRoom) CreateOnlineRoomWaitResult() (*model.AllocatedReq, error) {
// 	// //find sub server client
// 	// LIST := control.GetSubServerClientList()
// 	// if LIST == nil || len(*LIST) == 0 {
// 	// 	return nil, fmt.Errorf("no sub server client found")
// 	// }
// 	// sub_server_cli := (*LIST)[0]
// 	// if sub_server_cli == nil {
// 	// 	return nil, fmt.Errorf("sub server client is nil")
// 	// }
// 	// game.SubServer = sub_server_cli //set sub server client to game room
// 	// //Create room req to sub server
// 	// fmt.Println("CreateOnlineRoomWaitResult()")
// 	// result, err := Rpc_client_create_room(sub_server_cli, &protorpc.CreateRoomParam{
// 	// 	Key: game.Key,
// 	// 	Room: &protorpc.GameRoom{
// 	// 		Key:      game.Key,
// 	// 		RoomName: game.Key,
// 	// 	},
// 	// })
// 	// if err != nil || result == nil {
// 	// 	return nil, fmt.Errorf("create room failed: %v", err)
// 	// }
// 	// return result, nil

// 	// out, err := service.Rooms().CreateRoom(&model.Room_CreateRoomReq{
// 	// 	Key: game.Key,
// 	// 	Room: &model.Room_CreateRoomReq_Room{
// 	// 		Key:      game.Key,
// 	// 		RoomName: game.Key,
// 	// 	},
// 	// })

// 	uid := uuid.NewV4().String()

// 	out, err := service.Rooms().CreateRoom(&model.AllocationReq{
// 		UUID:      uid,
// 		ProfileID: "193ae4b9-6a73-4d3f-8203-7d5a6f67dd9c",
// 		RegionID:  "7bba5ff1-b4a3-498c-9c0b-24cad73cff54", //上海
// 		Envs: map[string]string{
// 			"UUID":             uid,
// 			"players_env_data": "env2",
// 		},
// 	})

// 	return out, err
// }

// func (game *GameRoom) CreateOnlineRoomWaitResult() (*model.AllocatedReq, error) {
//     uid := uuid.NewV4().String()
//     done := make(chan *model.AllocatedReq, 1)
//     errChan := make(chan error, 1)
//     go func() {
//         _, err := service.Rooms().CreateRoom(&model.AllocationReq{
//             UUID:      uid,
//             ProfileID: "193ae4b9-6a73-4d3f-8203-7d5a6f67dd9c",
//             RegionID:  "7bba5ff1-b4a3-498c-9c0b-24cad73cff54", // 上海
//             Envs: map[string]string{
//                 "UUID":             uid,
//                 "players_env_data": "env2",
//             },
//         })
//         if err != nil {
//             errChan <- err
//             return
//         }
//         // 这里假设 service.Rooms().CreateRoom 只是发起请求
//         // 真正的创建完成通知由另一个回调触发
//     }()

//     // 假设我们有一个全局或单例的监听器来等回调
//     // go func() {
//     //     // 等待回调，例如从消息队列/WebSocket收到
//     //     allocated := WaitForRoomAllocated(uid) // 自己实现
//     //     done <- allocated
//     // }()

//     select {
//     case result := <-done:
//         return result, nil
//     case err := <-errChan:
//         return nil, err
//     case <-time.After(30 * time.Second):
//         return nil, fmt.Errorf("等待房间创建超时")
//     }
// }

func (game *GameRoom) CreateOnlineRoomWaitResult() (*model.AllocatedReq, error) {
	var players []ReqPlayers
	for _, v := range game.Players {
		uid_key := generateId()
		v.UUID = uid_key
		players = append(players, ReqPlayers{
			Passport:   v.Userinfo_.Passport,
			UID:        int64(v.Userinfo_.Id),
			Key:        uid_key,
			Nickname:   v.Userinfo_.Nickname,
			PlayerType: string(rune(v.Player_type)),
			// CreatedAt: v.Userinfo_.CreatedAt,
			// Version: v.Userinfo_.Version,
		})
	}
	env_data, err := CreateMatchEnv(players)
	if err != nil {
		return nil, err
	}
	uid := uuid.NewV4().String()
	ch := notify.RegisterWait(uid)
	_, err = service.Rooms().CreateRoom(&model.AllocationReq{
		UUID:      uid,
		ProfileID: "193ae4b9-6a73-4d3f-8203-7d5a6f67dd9c",
		RegionID:  "7bba5ff1-b4a3-498c-9c0b-24cad73cff54", // 上海
		Envs: map[string]string{
			"UUID":             uid,
			"players_env_data": env_data,
		},
	})
	if err != nil {
		notify.Cancel(uid)
		return nil, err
	}
	select {
	case result := <-ch:
		if result == nil {
			return nil, errors.New("房间创建失败或被取消")
		}
		return result, nil
	case <-time.After(60 * time.Second):
		notify.Cancel(uid)
		return nil, errors.New("等待房间创建超时")
	}
}
