package model


type Room_ListRoomsRes struct {
	Rooms []struct {
		CreatedAt        string `json:"createdAt"`
		CustomProperties struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"customProperties"`
		ExpirationTime      string `json:"expirationTime"`
		ExternalUniqueID    string `json:"externalUniqueId"`
		Filter              string `json:"filter"`
		ID                  int    `json:"id"`
		JoinCode            string `json:"joinCode"`
		MaxPlayers          int    `json:"maxPlayers"`
		MultiverseProfileID string `json:"multiverseProfileId"`
		Name                string `json:"name"`
		Namespace           string `json:"namespace"`
		OwnerID             string `json:"ownerId"`
		PlayerCount         int    `json:"playerCount"`
		RoomTTLInMinutes    int    `json:"roomTTLInMinutes"`
		RoomUUID            string `json:"roomUUID"`
		Status              string `json:"status"`
		Type                string `json:"type"`
		UosAppID            string `json:"uosAppId"`
		UpdatedAt           string `json:"updatedAt"`
		Visibility          string `json:"visibility"`
	} `json:"rooms"`
	TotalCount int `json:"totalCount"`
}

type Room_CreateRoomReq struct {
	AllocationEnvs struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"allocationEnvs"`
	CustomProperties struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"customProperties"`
	Filter              string `json:"filter"`
	GameRegionID        string `json:"gameRegionId"`
	JoinCode            string `json:"joinCode"`
	MaxPlayers          int    `json:"maxPlayers"`
	MultiverseProfileID string `json:"multiverseProfileId"`
	Name                string `json:"name"`
	Namespace           string `json:"namespace"`
	PlayerID            string `json:"playerId"`
	RoomTTLInMinutes    int    `json:"roomTTLInMinutes"`
	Visibility          string `json:"visibility"`
}

type Room_CreateRoomRes Room
// struct {
// 	AllocationInfo struct {
// 		AllocationTTL string `json:"allocationTTL"`
// 		CreatedAt     struct {
// 			Nanos   int `json:"nanos"`
// 			Seconds int `json:"seconds"`
// 		} `json:"createdAt"`
// 		CreatedByUser string `json:"createdByUser"`
// 		DeletedAt     struct {
// 			Nanos   int `json:"nanos"`
// 			Seconds int `json:"seconds"`
// 		} `json:"deletedAt"`
// 		DeletedByUser string `json:"deletedByUser"`
// 		FulfilledAt   struct {
// 			Nanos   int `json:"nanos"`
// 			Seconds int `json:"seconds"`
// 		} `json:"fulfilledAt"`
// 		GameID          string `json:"gameId"`
// 		GameServerName  string `json:"gameServerName"`
// 		GameServerPorts []struct {
// 			Name       string `json:"name"`
// 			Port       int    `json:"port"`
// 			PortPolicy int    `json:"portPolicy"`
// 			Protocol   string `json:"protocol"`
// 		} `json:"gameServerPorts"`
// 		IP         string `json:"ip"`
// 		ModifiedAt struct {
// 			Nanos   int `json:"nanos"`
// 			Seconds int `json:"seconds"`
// 		} `json:"modifiedAt"`
// 		ModifiedByUser  string `json:"modifiedByUser"`
// 		Msg             string `json:"msg"`
// 		ProfileID       string `json:"profileId"`
// 		ProfileName     string `json:"profileName"`
// 		ProfileRevision struct {
// 			Annotations struct {
// 				AdditionalProp1 string `json:"additionalProp1"`
// 				AdditionalProp2 string `json:"additionalProp2"`
// 				AdditionalProp3 string `json:"additionalProp3"`
// 			} `json:"annotations"`
// 			CPULimit   string `json:"cpuLimit"`
// 			CPURequest string `json:"cpuRequest"`
// 			CreatedAt  struct {
// 				Nanos   int `json:"nanos"`
// 				Seconds int `json:"seconds"`
// 			} `json:"createdAt"`
// 			EnvironmentVariables struct {
// 				AdditionalProp1 string `json:"additionalProp1"`
// 				AdditionalProp2 string `json:"additionalProp2"`
// 				AdditionalProp3 string `json:"additionalProp3"`
// 			} `json:"environmentVariables"`
// 			FileConfigs []struct {
// 				Content   string `json:"content"`
// 				Filename  string `json:"filename"`
// 				MountPath string `json:"mountPath"`
// 			} `json:"fileConfigs"`
// 			GameImage struct {
// 				CreatedAt struct {
// 					Nanos   int `json:"nanos"`
// 					Seconds int `json:"seconds"`
// 				} `json:"createdAt"`
// 				CreatedByUser string `json:"createdByUser"`
// 				GameImageID   string `json:"gameImageId"`
// 				GameImageTag  string `json:"gameImageTag"`
// 				GameImageURL  string `json:"gameImageUrl"`
// 			} `json:"gameImage"`
// 			GameServerEntryPoint []string `json:"gameServerEntryPoint"`
// 			GameServerPorts      []struct {
// 				Name       string `json:"name"`
// 				Port       int    `json:"port"`
// 				PortPolicy int    `json:"portPolicy"`
// 				Protocol   string `json:"protocol"`
// 			} `json:"gameServerPorts"`
// 			GameStartDuration string `json:"gameStartDuration"`
// 			MemoryLimit       string `json:"memoryLimit"`
// 			MemoryRequest     string `json:"memoryRequest"`
// 			NodeSelectors     struct {
// 				AdditionalProp1 string `json:"additionalProp1"`
// 				AdditionalProp2 string `json:"additionalProp2"`
// 				AdditionalProp3 string `json:"additionalProp3"`
// 			} `json:"nodeSelectors"`
// 			ProfileID   string `json:"profileId"`
// 			RevisionID  string `json:"revisionId"`
// 			Tolerations []struct {
// 				Effect   string `json:"effect"`
// 				Key      string `json:"key"`
// 				Operator string `json:"operator"`
// 				Value    string `json:"value"`
// 			} `json:"tolerations"`
// 		} `json:"profileRevision"`
// 		RegionID   string `json:"regionId"`
// 		RegionName string `json:"regionName"`
// 		Status     string `json:"status"`
// 		UUID       string `json:"uuid"`
// 		WsProxy    string `json:"wsProxy"`
// 	} `json:"allocationInfo"`
// 	Players  []string `json:"players"`
// 	RoomInfo struct {
// 		CreatedAt        string `json:"createdAt"`
// 		CustomProperties struct {
// 			AdditionalProp1 string `json:"additionalProp1"`
// 			AdditionalProp2 string `json:"additionalProp2"`
// 			AdditionalProp3 string `json:"additionalProp3"`
// 		} `json:"customProperties"`
// 		ExpirationTime      string `json:"expirationTime"`
// 		ExternalUniqueID    string `json:"externalUniqueId"`
// 		Filter              string `json:"filter"`
// 		ID                  int    `json:"id"`
// 		JoinCode            string `json:"joinCode"`
// 		MaxPlayers          int    `json:"maxPlayers"`
// 		MultiverseProfileID string `json:"multiverseProfileId"`
// 		Name                string `json:"name"`
// 		Namespace           string `json:"namespace"`
// 		OwnerID             string `json:"ownerId"`
// 		PlayerCount         int    `json:"playerCount"`
// 		RoomTTLInMinutes    int    `json:"roomTTLInMinutes"`
// 		RoomUUID            string `json:"roomUUID"`
// 		Status              string `json:"status"`
// 		Type                string `json:"type"`
// 		UosAppID            string `json:"uosAppId"`
// 		UpdatedAt           string `json:"updatedAt"`
// 		Visibility          string `json:"visibility"`
// 	} `json:"roomInfo"`
// }


type Room struct {
	AllocationInfo struct {
		AllocationTTL string `json:"allocationTTL"`
		CreatedAt     struct {
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"createdAt"`
		CreatedByUser string `json:"createdByUser"`
		DeletedAt     struct {
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"deletedAt"`
		DeletedByUser string `json:"deletedByUser"`
		FulfilledAt   struct {
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"fulfilledAt"`
		GameID          string `json:"gameId"`
		GameServerName  string `json:"gameServerName"`
		GameServerPorts []struct {
			Name       string `json:"name"`
			Port       int    `json:"port"`
			PortPolicy int    `json:"portPolicy"`
			Protocol   string `json:"protocol"`
		} `json:"gameServerPorts"`
		IP         string `json:"ip"`
		ModifiedAt struct {
			Nanos   int `json:"nanos"`
			Seconds int `json:"seconds"`
		} `json:"modifiedAt"`
		ModifiedByUser  string `json:"modifiedByUser"`
		Msg             string `json:"msg"`
		ProfileID       string `json:"profileId"`
		ProfileName     string `json:"profileName"`
		ProfileRevision struct {
			Annotations struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"annotations"`
			CPULimit   string `json:"cpuLimit"`
			CPURequest string `json:"cpuRequest"`
			CreatedAt  struct {
				Nanos   int `json:"nanos"`
				Seconds int `json:"seconds"`
			} `json:"createdAt"`
			EnvironmentVariables struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"environmentVariables"`
			FileConfigs []struct {
				Content   string `json:"content"`
				Filename  string `json:"filename"`
				MountPath string `json:"mountPath"`
			} `json:"fileConfigs"`
			GameImage struct {
				CreatedAt struct {
					Nanos   int `json:"nanos"`
					Seconds int `json:"seconds"`
				} `json:"createdAt"`
				CreatedByUser string `json:"createdByUser"`
				GameImageID   string `json:"gameImageId"`
				GameImageTag  string `json:"gameImageTag"`
				GameImageURL  string `json:"gameImageUrl"`
			} `json:"gameImage"`
			GameServerEntryPoint []string `json:"gameServerEntryPoint"`
			GameServerPorts      []struct {
				Name       string `json:"name"`
				Port       int    `json:"port"`
				PortPolicy int    `json:"portPolicy"`
				Protocol   string `json:"protocol"`
			} `json:"gameServerPorts"`
			GameStartDuration string `json:"gameStartDuration"`
			MemoryLimit       string `json:"memoryLimit"`
			MemoryRequest     string `json:"memoryRequest"`
			NodeSelectors     struct {
				AdditionalProp1 string `json:"additionalProp1"`
				AdditionalProp2 string `json:"additionalProp2"`
				AdditionalProp3 string `json:"additionalProp3"`
			} `json:"nodeSelectors"`
			ProfileID   string `json:"profileId"`
			RevisionID  string `json:"revisionId"`
			Tolerations []struct {
				Effect   string `json:"effect"`
				Key      string `json:"key"`
				Operator string `json:"operator"`
				Value    string `json:"value"`
			} `json:"tolerations"`
		} `json:"profileRevision"`
		RegionID   string `json:"regionId"`
		RegionName string `json:"regionName"`
		Status     string `json:"status"`
		UUID       string `json:"uuid"`
		WsProxy    string `json:"wsProxy"`
	} `json:"allocationInfo"`
	Players  []string `json:"players"`
	RoomInfo struct {
		CreatedAt        string `json:"createdAt"`
		CustomProperties struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"customProperties"`
		ExpirationTime      string `json:"expirationTime"`
		ExternalUniqueID    string `json:"externalUniqueId"`
		Filter              string `json:"filter"`
		ID                  int    `json:"id"`
		JoinCode            string `json:"joinCode"`
		MaxPlayers          int    `json:"maxPlayers"`
		MultiverseProfileID string `json:"multiverseProfileId"`
		Name                string `json:"name"`
		Namespace           string `json:"namespace"`
		OwnerID             string `json:"ownerId"`
		PlayerCount         int    `json:"playerCount"`
		RoomTTLInMinutes    int    `json:"roomTTLInMinutes"`
		RoomUUID            string `json:"roomUUID"`
		Status              string `json:"status"`
		Type                string `json:"type"`
		UosAppID            string `json:"uosAppId"`
		UpdatedAt           string `json:"updatedAt"`
		Visibility          string `json:"visibility"`
	} `json:"roomInfo"`
}