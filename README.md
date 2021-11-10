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
        {
          "type": "ping"
        }    
      ```
    - response
      ```json
        {
          "type": "pong"
        } 
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
        {
          "type": "register",
          "userID": "c51f3e57-3769-11ec-ae49-8e4c1c8516c7",
          "message": ""
        } 
      ```

- update
    - request
      ```json
      {
        "type":"update",
        "userInfo":
        {
          "userID":"",
          "geoLocation":
          {
            "latitude":35.943218,
            "longitude":139.621248
          }
        }
      }
      ```
    - response
      ```json
      {
        "type":"update",
        "message": ""
      }
      ```
- search
    - request
      ```json
       {
         "type":"search",
         "searchType":"static",
         "searchDistance":100,
         "geoLocation": 
         {
           "latitude":35.943218,
           "longitude":139.621248
         }
       }
      ```
    - response
      ```json
      {
        "type":"search",
        "message":"",
        "surroundingUserList":
        [
          {
            "userID":"",
            "geoLocation":
            {
              "latitude":0,
              "longitude":0
            }
          }
        ]
      }
      ```
- delete
    - request
      ```json
       {
         "type":"delete"
       }
      ```
    - response
      ```json
       {
          "type":"delete",
          "message": ""
        }
      ```
- offer
    - toSignalingServer
      ```json
      {
        "type":"offer",
        "sdp":"v=0\r\no=- 9207231071010309254 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\na=group:BUNDLE 0 1\r\na=extmap-allow-mixed\r\na=msid-semantic: WMS JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot\r\nm=video 9 UDP/TLS/RTP/SAVPF 96 97 98 99 100 101 102 121 127 120 125 107 108 109 35 36 124 119 123 118 114 115 116\r\nc=IN IP4 0.0.0.0\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=ice-ufrag:Hp59\r\na=ice-pwd:m1SyttUrVOgHGfMH4xESw6X5\r\na=ice-options:trickle\r\na=fingerprint:sha-256 1A:D7:AB:E6:FD:CA:14:11:08:E9:41:22:20:9F:3A:F7:8A:B2:76:45:AB:C8:BB:FA:F9:C4:2D:1D:6B:35:84:EB\r\na=setup:actpass\r\na=mid:0\r\na=extmap:1 urn:ietf:params:rtp-hdrext:toffset\r\na=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\na=extmap:3 urn:3gpp:video-orientation\r\na=extmap:4 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=extmap:5 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay\r\na=extmap:6 http://www.webrtc.org/experiments/rtp-hdrext/video-content-type\r\na=extmap:7 http://www.webrtc.org/experiments/rtp-hdrext/video-timing\r\na=extmap:8 http://www.webrtc.org/experiments/rtp-hdrext/color-space\r\na=extmap:9 urn:ietf:params:rtp-hdrext:sdes:mid\r\na=extmap:10 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id\r\na=extmap:11 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id\r\na=sendrecv\r\na=msid:JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot 06ae7aa4-24ef-44c2-9593-294ef0c8ce5d\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=rtpmap:96 VP8/90000\r\na=rtcp-fb:96 goog-remb\r\na=rtcp-fb:96 transport-cc\r\na=rtcp-fb:96 ccm fir\r\na=rtcp-fb:96 nack\r\na=rtcp-fb:96 nack pli\r\na=rtpmap:97 rtx/90000\r\na=fmtp:97 apt=96\r\na=rtpmap:98 VP9/90000\r\na=rtcp-fb:98 goog-remb\r\na=rtcp-fb:98 transport-cc\r\na=rtcp-fb:98 ccm fir\r\na=rtcp-fb:98 nack\r\na=rtcp-fb:98 nack pli\r\na=fmtp:98 profile-id=0\r\na=rtpmap:99 rtx/90000\r\na=fmtp:99 apt=98\r\na=rtpmap:100 VP9/90000\r\na=rtcp-fb:100 goog-remb\r\na=rtcp-fb:100 transport-cc\r\na=rtcp-fb:100 ccm fir\r\na=rtcp-fb:100 nack\r\na=rtcp-fb:100 nack pli\r\na=fmtp:100 profile-id=2\r\na=rtpmap:101 rtx/90000\r\na=fmtp:101 apt=100\r\na=rtpmap:102 H264/90000\r\na=rtcp-fb:102 goog-remb\r\na=rtcp-fb:102 transport-cc\r\na=rtcp-fb:102 ccm fir\r\na=rtcp-fb:102 nack\r\na=rtcp-fb:102 nack pli\r\na=fmtp:102 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f\r\na=rtpmap:121 rtx/90000\r\na=fmtp:121 apt=102\r\na=rtpmap:127 H264/90000\r\na=rtcp-fb:127 goog-remb\r\na=rtcp-fb:127 transport-cc\r\na=rtcp-fb:127 ccm fir\r\na=rtcp-fb:127 nack\r\na=rtcp-fb:127 nack pli\r\na=fmtp:127 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f\r\na=rtpmap:120 rtx/90000\r\na=fmtp:120 apt=127\r\na=rtpmap:125 H264/90000\r\na=rtcp-fb:125 goog-remb\r\na=rtcp-fb:125 transport-cc\r\na=rtcp-fb:125 ccm fir\r\na=rtcp-fb:125 nack\r\na=rtcp-fb:125 nack pli\r\na=fmtp:125 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f\r\na=rtpmap:107 rtx/90000\r\na=fmtp:107 apt=125\r\na=rtpmap:108 H264/90000\r\na=rtcp-fb:108 goog-remb\r\na=rtcp-fb:108 transport-cc\r\na=rtcp-fb:108 ccm fir\r\na=rtcp-fb:108 nack\r\na=rtcp-fb:108 nack pli\r\na=fmtp:108 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f\r\na=rtpmap:109 rtx/90000\r\na=fmtp:109 apt=108\r\na=rtpmap:35 AV1X/90000\r\na=rtcp-fb:35 goog-remb\r\na=rtcp-fb:35 transport-cc\r\na=rtcp-fb:35 ccm fir\r\na=rtcp-fb:35 nack\r\na=rtcp-fb:35 nack pli\r\na=rtpmap:36 rtx/90000\r\na=fmtp:36 apt=35\r\na=rtpmap:124 H264/90000\r\na=rtcp-fb:124 goog-remb\r\na=rtcp-fb:124 transport-cc\r\na=rtcp-fb:124 ccm fir\r\na=rtcp-fb:124 nack\r\na=rtcp-fb:124 nack pli\r\na=fmtp:124 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=4d001f\r\na=rtpmap:119 rtx/90000\r\na=fmtp:119 apt=124\r\na=rtpmap:123 H264/90000\r\na=rtcp-fb:123 goog-remb\r\na=rtcp-fb:123 transport-cc\r\na=rtcp-fb:123 ccm fir\r\na=rtcp-fb:123 nack\r\na=rtcp-fb:123 nack pli\r\na=fmtp:123 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=64001f\r\na=rtpmap:118 rtx/90000\r\na=fmtp:118 apt=123\r\na=rtpmap:114 red/90000\r\na=rtpmap:115 rtx/90000\r\na=fmtp:115 apt=114\r\na=rtpmap:116 ulpfec/90000\r\na=ssrc-group:FID 3984706797 1490309860\r\na=ssrc:3984706797 cname:qC2uzZXQI5/76NQz\r\na=ssrc:3984706797 msid:JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot 06ae7aa4-24ef-44c2-9593-294ef0c8ce5d\r\na=ssrc:3984706797 mslabel:JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot\r\na=ssrc:3984706797 label:06ae7aa4-24ef-44c2-9593-294ef0c8ce5d\r\na=ssrc:1490309860 cname:qC2uzZXQI5/76NQz\r\na=ssrc:1490309860 msid:JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot 06ae7aa4-24ef-44c2-9593-294ef0c8ce5d\r\na=ssrc:1490309860 mslabel:JJQPCPdSMaeZ3u8BiQTF23MvIhOMe9soH7Ot\r\na=ssrc:1490309860 label:06ae7aa4-24ef-44c2-9593-294ef0c8ce5d\r\nm=application 9 UDP/DTLS/SCTP webrtc-datachannel\r\nc=IN IP4 0.0.0.0\r\na=ice-ufrag:Hp59\r\na=ice-pwd:m1SyttUrVOgHGfMH4xESw6X5\r\na=ice-options:trickle\r\na=fingerprint:sha-256 1A:D7:AB:E6:FD:CA:14:11:08:E9:41:22:20:9F:3A:F7:8A:B2:76:45:AB:C8:BB:FA:F9:C4:2D:1D:6B:35:84:EB\r\na=setup:actpass\r\na=mid:1\r\na=sctp-port:5000\r\na=max-message-size:262144\r\n",
        "destination":"eadb3d1e-3769-11ec-ae49-8e4c1c8516c7"
      }
      ```
    - toDestination
      ```json
      ```
- answer
    - toSignalingServer
        ```json
      {
        "type":"answer",
        "sdp":"v=0\r\no=- 5642943775146002561 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\na=group:BUNDLE 0 1\r\na=extmap-allow-mixed\r\na=msid-semantic: WMS 80hGY6EgcCgXC2Avez47sdbJi2iD6RyUzU3Z\r\nm=video 9 UDP/TLS/RTP/SAVPF 96 97 98 99 100 101 102 121 127 120 125 107 108 109 35 36 124 119 123 118 114 115 116\r\nc=IN IP4 0.0.0.0\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=ice-ufrag:NfQw\r\na=ice-pwd:D/L0Xo7HdxZZvvyEsTr8UJSn\r\na=ice-options:trickle\r\na=fingerprint:sha-256 62:3C:78:66:94:4D:87:0D:07:32:BC:5F:A4:23:32:60:8D:97:31:F9:9F:AC:E4:01:9E:17:7B:37:60:51:37:67\r\na=setup:active\r\na=mid:0\r\na=extmap:1 urn:ietf:params:rtp-hdrext:toffset\r\na=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\na=extmap:3 urn:3gpp:video-orientation\r\na=extmap:4 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=extmap:5 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay\r\na=extmap:6 http://www.webrtc.org/experiments/rtp-hdrext/video-content-type\r\na=extmap:7 http://www.webrtc.org/experiments/rtp-hdrext/video-timing\r\na=extmap:8 http://www.webrtc.org/experiments/rtp-hdrext/color-space\r\na=extmap:9 urn:ietf:params:rtp-hdrext:sdes:mid\r\na=extmap:10 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id\r\na=extmap:11 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id\r\na=sendrecv\r\na=msid:80hGY6EgcCgXC2Avez47sdbJi2iD6RyUzU3Z 19c18e17-1199-4dad-8520-cdab2ff2e7b0\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=rtpmap:96 VP8/90000\r\na=rtcp-fb:96 goog-remb\r\na=rtcp-fb:96 transport-cc\r\na=rtcp-fb:96 ccm fir\r\na=rtcp-fb:96 nack\r\na=rtcp-fb:96 nack pli\r\na=rtpmap:97 rtx/90000\r\na=fmtp:97 apt=96\r\na=rtpmap:98 VP9/90000\r\na=rtcp-fb:98 goog-remb\r\na=rtcp-fb:98 transport-cc\r\na=rtcp-fb:98 ccm fir\r\na=rtcp-fb:98 nack\r\na=rtcp-fb:98 nack pli\r\na=fmtp:98 profile-id=0\r\na=rtpmap:99 rtx/90000\r\na=fmtp:99 apt=98\r\na=rtpmap:100 VP9/90000\r\na=rtcp-fb:100 goog-remb\r\na=rtcp-fb:100 transport-cc\r\na=rtcp-fb:100 ccm fir\r\na=rtcp-fb:100 nack\r\na=rtcp-fb:100 nack pli\r\na=fmtp:100 profile-id=2\r\na=rtpmap:101 rtx/90000\r\na=fmtp:101 apt=100\r\na=rtpmap:102 H264/90000\r\na=rtcp-fb:102 goog-remb\r\na=rtcp-fb:102 transport-cc\r\na=rtcp-fb:102 ccm fir\r\na=rtcp-fb:102 nack\r\na=rtcp-fb:102 nack pli\r\na=fmtp:102 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f\r\na=rtpmap:121 rtx/90000\r\na=fmtp:121 apt=102\r\na=rtpmap:127 H264/90000\r\na=rtcp-fb:127 goog-remb\r\na=rtcp-fb:127 transport-cc\r\na=rtcp-fb:127 ccm fir\r\na=rtcp-fb:127 nack\r\na=rtcp-fb:127 nack pli\r\na=fmtp:127 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f\r\na=rtpmap:120 rtx/90000\r\na=fmtp:120 apt=127\r\na=rtpmap:125 H264/90000\r\na=rtcp-fb:125 goog-remb\r\na=rtcp-fb:125 transport-cc\r\na=rtcp-fb:125 ccm fir\r\na=rtcp-fb:125 nack\r\na=rtcp-fb:125 nack pli\r\na=fmtp:125 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f\r\na=rtpmap:107 rtx/90000\r\na=fmtp:107 apt=125\r\na=rtpmap:108 H264/90000\r\na=rtcp-fb:108 goog-remb\r\na=rtcp-fb:108 transport-cc\r\na=rtcp-fb:108 ccm fir\r\na=rtcp-fb:108 nack\r\na=rtcp-fb:108 nack pli\r\na=fmtp:108 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f\r\na=rtpmap:109 rtx/90000\r\na=fmtp:109 apt=108\r\na=rtpmap:35 AV1X/90000\r\na=rtcp-fb:35 goog-remb\r\na=rtcp-fb:35 transport-cc\r\na=rtcp-fb:35 ccm fir\r\na=rtcp-fb:35 nack\r\na=rtcp-fb:35 nack pli\r\na=rtpmap:36 rtx/90000\r\na=fmtp:36 apt=35\r\na=rtpmap:124 H264/90000\r\na=rtcp-fb:124 goog-remb\r\na=rtcp-fb:124 transport-cc\r\na=rtcp-fb:124 ccm fir\r\na=rtcp-fb:124 nack\r\na=rtcp-fb:124 nack pli\r\na=fmtp:124 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=4d001f\r\na=rtpmap:119 rtx/90000\r\na=fmtp:119 apt=124\r\na=rtpmap:123 H264/90000\r\na=rtcp-fb:123 goog-remb\r\na=rtcp-fb:123 transport-cc\r\na=rtcp-fb:123 ccm fir\r\na=rtcp-fb:123 nack\r\na=rtcp-fb:123 nack pli\r\na=fmtp:123 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=64001f\r\na=rtpmap:118 rtx/90000\r\na=fmtp:118 apt=123\r\na=rtpmap:114 red/90000\r\na=rtpmap:115 rtx/90000\r\na=fmtp:115 apt=114\r\na=rtpmap:116 ulpfec/90000\r\na=ssrc-group:FID 3121046868 4259058578\r\na=ssrc:3121046868 cname:ekyk8r/SCwZW9o6z\r\na=ssrc:4259058578 cname:ekyk8r/SCwZW9o6z\r\nm=application 9 UDP/DTLS/SCTP webrtc-datachannel\r\nc=IN IP4 0.0.0.0\r\na=ice-ufrag:NfQw\r\na=ice-pwd:D/L0Xo7HdxZZvvyEsTr8UJSn\r\na=ice-options:trickle\r\na=fingerprint:sha-256 62:3C:78:66:94:4D:87:0D:07:32:BC:5F:A4:23:32:60:8D:97:31:F9:9F:AC:E4:01:9E:17:7B:37:60:51:37:67\r\na=setup:active\r\na=mid:1\r\na=sctp-port:5000\r\na=max-message-size:262144\r\n",
        "destination":"eaf56485-3769-11ec-ae49-8e4c1c8516c7"
      }
        ```
    - toDestination
      ```json
      ```
- ice
    - message
      ```json
      {
        "type":"ice",
        "ice":
        {
          "candidate":"candidate:3415414722 1 udp 2121998079 172.29.0.1 61433 typ host generation 0 ufrag NfQw network-id 3",
          "sdpMid":"0",
          "sdpMLineIndex":0
        }
      }
      ```
- close
    - request
      ```json
       {
         "type":"close"
       }
      ```
    - response
      ```json
       {
         "type":"close",
         "message": ""
       }
      ```


