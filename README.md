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
          |   offer {type:offer,sdp:sdp}        |                  |                  |
          +-------------------+---------------->|       offer {type:offer, sdp:sdp}   |
          |                   |                 +------------------+----------------->|
          |                   |                 |                  |                  |
          |                   |                 |       answer {type:answer, sdp:sdp} |
          |                   |                 |<-----------------+------------------+
          |   answer {type:answer, sdp:sdp}     |                  |                  |
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

userIDはサーバ側で保持する方向にする
コネクションが繋がっているから大丈夫

- client

```typescript
type Status = {
    code: string
    message: string
    type: string
}
type SearchResponse = {
    status: Status
    searchedUserList: BackEndUserInfo[]
}

type RegisterResponse = {
    status: Status
    userID: string
}

type StunResponse = {
    addr: Addr
}

type JudgeStatus = {
    status: Status
}

type BackEndUserInfo = {
    userID: string
    publicIP: string
    publicPort: number
    privateIP: string
    privatePort: number
    latitude: number
    longitude: number
}

```

- server

```go

type Status struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Type    string `json:"type"`
}

type SearchResponse struct {
    Status           Status            `json:"status"`
    SearchedUserList []*model.UserInfo `json:"searchedUserList"`
}
type RegisterResponse struct {
    Status Status `json:"status"`
    UserID string `json:"userID"`
}


// --------------------------

type Message struct {
    Type string `json:"type"`
}

type RegisterMessage struct {
    Type     string   `json:"type"`
    UserInfo UserInfo `json:"userInfo"`
}

type UpdateMessage struct {
    Type     string   `json:"type"`
    UserInfo UserInfo `json:"userInfo"`
}

type SearchMessage struct {
    Type           string   `json:"type"`
    UserInfo       UserInfo `json:"userInfo"`
    SearchType     string   `json:"searchType"`
    SearchDistance float64  `json:"searchDistance"`
}

type DeleteMessage struct {
    Type     string   `json:"type"`
    UserInfo UserInfo `json:"userInfo"`
}

type SendMessage struct {
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
    PublicIP    string  `json:"publicIP"`
    PublicPort  uint8   `json:"publicPort"`
    PrivateIP   string  `json:"privateIP"`
    PrivatePort uint8   `json:"privatePort"`
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
}

// GeoLocation ユーザの位置情報
type GeoLocation struct {
    Latitude  string `json:"latitude"`
    Longitude string `json:"longitude"`
}

type Addr struct {
    IP   string `json:"ip"`
    Port uint8  `json:"port"`
}

// AlterUserInfo 別の案
type AlterUserInfo struct {
    UserID      string      `json:"userID"`
    PublicAddr  Addr        `json:"public"`
    PrivateAddr Addr        `json:"private"`
    GeoLocation GeoLocation `json:"geoLocation"`
}
```

- message

```json

```
