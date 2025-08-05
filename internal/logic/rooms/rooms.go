package rooms

import (
	"context"
	"encoding/json"
	"io"
	"time"

	// "encoding/json"
	// "time"

	// "fmt"
	// "gfAdmin/internal/cache"
	"gfAdmin/internal/model"
	"gfAdmin/internal/service"
	"gfAdmin/internal/http_client"
	// "io/ioutil"
	// "net/http"
	// "net/url"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	// "github.com/gogf/gf/v2/errors/gerror"
)

type (
	sRooms struct{}
)

var logger = g.Log("Rooms service")
var _client *http_client.HttpClient //*http_client.

const (
	ROOM_PUBLIC  = "public"
	ROOM_PRIVATE = "private"
	ROOM_MATCH   = "match"
)

func init() {
	service.RegisterRooms(New())
	_client = http_client.NewHttpClient("https://s.unity.cn", 10*time.Second)
}

func New() service.IRooms {
	return &sRooms{}
}

const Authorization = "Basic NDNjYTIzNjUtN2YxNy00NTJjLWEzZTgtNzYyZjRjNzIwYWYwOjE3ZWM0ZjdjMGUwMzRmYzJiOGQzOTRkMGUwYTIwMGNi"

// 43ca2365-7f17-452c-a3e8-762f4c720af0:17ec4f7c0e034fc2b8d394d0e0a200cb
type Room struct {
}

func (s *sRooms) CreateRoom(ctx context.Context, in *model.Room_CreateRoomReq) (out *model.Room_CreateRoomRes, err error) {
	logger.Println("CreateRoom()")
	if in == nil {
		return nil, gerror.New("CreateRoomReq is nil")
	}
	//post /service/rooms
	res, err := _client.Post("/service/rooms", &http_client.RequestOption{
		Headers: map[string]string{
			"Authorization": Authorization,
		},
		JsonData: in,
	})
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	logger.Println(string(body))
	err = json.Unmarshal(body, &out)
	return
}

func (s *sRooms) QueryRoom(ctx context.Context) (list *model.Room_ListRoomsRes, err error) {
	//get /service/rooms
	res, err := _client.Get("/service/rooms", &http_client.RequestOption{
		Headers: map[string]string{
			"Authorization": Authorization,
		},
		QueryParams: map[string]string{
			"name":      "",
			"namespace": "",
		},
	})
	if err != nil {
		logger.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &list)
	return
}

func (s *sRooms) GetRoomByUuid(ctx context.Context, uuid string) (r *model.Room, err error) {
	//get /service/rooms/:roomUuid
	res, err := _client.Get("/service/rooms/:roomUuid", &http_client.RequestOption{
		Headers: map[string]string{
			"Authorization": Authorization,
		},
		QueryParams: map[string]string{
			"roomUuid ": uuid,
		},
	})
	if err != nil {
		logger.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &r)
	return
}

//确认房间加入 /service/rooms/:roomUuid/confirm-join
