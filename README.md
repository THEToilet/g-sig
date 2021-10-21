# g-sig

project G-LocON using WebRTC

## プロトコル protocol

- 共通
    - ping
    - pong
- G-LocON
    - register
    - update
    - delete
    - search
        - static
        - dynamic
    - send
        - WebRTC接続後にP2P通信を行うのでいらないかも。。。？
        - DataProtocolでデータを送るのに使うかもしれない
- P2P
    - offer
        - SDPを送る。
        - WebRTCにおけるoffer側
    - answer
        - SDPを送る。
        - WebRTCにおけるanswer側
    - candidate
        - ICE candidateを受信
    - close

## 動作フロー

```text

  +----------------+ +----------------+
+----------------+ +----------------+
+----------------+
  |     Clinet     | |    FrontEnd    |
|    BackEnd     | |   STUN / TURN  |
|     Client     |
  |      -1-       | | ReactWebServer |
|SignalingServer | |                |
|      -2-       |
  +-------+--------+ +--------+-------+
+-------+--------+ +-------+--------+
+-------+--------+

         |   GET FrontPage   |                 |                  |   GET FrontPage  |
          +------------------>|<----------------+------------------+------------------+
          | 
                 |                 |                  |                  |
          |   register {}     |                 |       register {}|                  |
          +-------------------+---------------->|<-----------------+------------------+
          |                   |                 |                  |                  |
          |                   |                 |                  |                  |
          |   update {}       |           
          |       update {}  |                  |
          +-------------------+---------------->|<-----------------+------------------+
          |                   |           
     |                  |                  |
          |   search {}       |           
     |       search {}  |                  |
          +-------------------+---------------->|<-----------------+------------------+
          |                   |                 |                  |                  |
       xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxWebRTCxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
          |                   +                 |                  |                  |
          |   offer {type:offer,sdp:sdp , id : id}        |                  |                  |
          +-------------------+---------------->|       offer {type:offer, sdp:sdp, id: id}   |
          |                   |                 +------------------+----------------->|
          |                   |                 |                  |                  |
          |                   |                 |       answer {type:answer, sdp:sdp, id:id} |
          |                   |                 |<-----------------+------------------+
          |   answer {type:answer, sdp:sdp, id:id}     |                  |                  |
          |<------------------+-----------------+                  |                  |
          |                   |                 |                  |                  |
          |                   |                 |                  |                  |
      xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxVanillaxICExxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
          |   ice {type:ice,ice:ice}            |                  +                  |
          +-------------------+---------------->|         ice {type:ice,ice:ice}      |
          |                   |                 +------------------+----------------->|
          |                   |                 |                  |                  |
          |                   |                 |            ice {type:ice,ice:ice}   |
          |      ice {type:ice,ice:ice}         |<-----------------+------------------+
          |<------------------+-----------------+                  |                  |
          |                   |                 |                  |                  |
          |                   |                 |                  |                  |
      xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxP2Pxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
          |                   |                 |                  |                  |
          +-------------------+-----------------+------------------+----------------->|
          |                   |             P 2 P                  |                  |
          |<------------------+-----------------+------------------+------------------+
          |                   |                 |                  |                  |
          |                   |                 |                  |                  |
```

## プログラム構成

## データ定義

userIDはサーバ側で保持する方向にする コネクションが繋がっているから大丈夫

- client

```typescript
//--------response----------
type SearchResponse = {
    type: string
    message: string
    geoLocation: GeoLocation
    surroundingUserList: UserInfo[]
}

type RegisterResponse = {
    type: string
    message: string
    userID: string
}

type UpdateResponse = {
    type: string
    message: string
}

type DeleteResponse = {
    type: string
    message: string
}

type SendResponse = {
    type: string
    message: string
}

type JudgeMessageType = {
    type: string
}
//----------request----------

type PongReauest = {
    type: string
}

type RegisterReguest = {
    type: string
    geoLocation: GeoLocation
}

type UpdateRequest = {
    type: string
    geoLocation: GeoLocation
}

type SearchRequest = {
    type: string
    searchType: string
    searchDistance: number
}

type DeleteRequest = {
    type: string
}

type SendRequest = {
    type: string
    message: string
}
//--------------------------

type GeoLocation = {
    latitude: number
    longitude: number
}

type UserInfo = {
    userID: string
    geoLocation: GeoLocation
}

```

- server

typeにはやり取りをするプロトコルを書く messageはエラーがある場合のみ書く

```go


type RegisterResponse struct {
Type    string `json:"type"`
Message string `json:"message"`
UserID string `json:"userID"`
}
type UpdateResponse struct {
Type    string `json:"type"`
Message string `json:"message"`
}

type SearchResponse struct {
Type    string `json:"type"`
Message string `json:"message"`
SurroundingUserList []*model.UserInfo `json:"surroundingUserList"`
}

type DeleteResponse struct {
Type    string `json:"type"`
Message string `json:"message"`
}

type SendResponse struct {
Type    string `json:"type"`
Message string `json:"message"`
}


// --------------------------

type JudgeMessageType struct {
Type string `json:"type"`
}

type RegisterRequest struct {
Type     string   `json:"type"`
GeoLocation GeoLocation `json:"geoLocation"`
}

type UpdateRequest struct {
Type     string   `json:"type"`
GeoLocation GeoLocation `json:"geoLocation"`
}

type SearchRequest struct {
Type           string   `json:"type"`
SearchType     string   `json:"searchType"`
SearchDistance float64  `json:"searchDistance"`
}

type DeleteRequest struct {
Type     string   `json:"type"`
}

type SendRequest struct {
Type    string `json:"type"`
Message string `json:"message"`
}

// -------- model ---------------
// User 永続化するユーザ情報
type User struct {
UserID   string
UserName string
}

// UserInfo ユーザの頻繁に変わる情報
type UserInfo struct {
UserID      string  `json:"userID"`
GeoLocation GeoLocation `json:"geoLocation"`
}

// GeoLocation ユーザの位置情報
type GeoLocation struct {
Latitude  string `json:"latitude"`
Longitude string `json:"longitude"`
}
```

- message

- ping-pong
    - request
      ```json
        {"type": "ping"}    
      ```
    - response
      ```json
        {"type": "pong"} 
      ```
- register
    - request
      ```json
        {
          "type": "register", 
          "geoLocation": 
          {
            "latitude": 12.123,
            "longitude": 123.333
          }
        } 
      ```
    - response
      ```json
        {"type": "register", "userID": "123-21421-1241-24", "message": ""} 
      ```

