package test

const BliData = `{
  "code": 0,
  "message": "0",
  "ttl": 1,
  "data": {
    "from": "local",
    "result": "suee",
    "message": "",
    "quality": 64,
    "format": "flv720",
    "timelength": 242794,
    "accept_format": "flv,flv720,flv480,mp4",
    "accept_description": [
      "高清 1080P",
      "高清 720P",
      "清晰 480P",
      "流畅 360P"
    ],
    "accept_quality": [80, 64, 32, 16],
    "video_codecid": 7,
    "seek_param": "start",
    "seek_type": "offset",
    "dash": {
      "duration": 243,
      "minBufferTime": 1.5,
      "min_buffer_time": 1.5,
      "video": [
        {
          "id": 80,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30080.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&uipk=5&og=hw&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&os=cosovbv&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=59826e2324b976e3df2e1e50be7be8e4&uparams=e,mid,oi,uipk,og,platform,deadline,nbs,gen,os,trid&bvc=vod&nettype=0&bw=282689&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30080.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&uipk=5&og=hw&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&os=cosovbv&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=59826e2324b976e3df2e1e50be7be8e4&uparams=e,mid,oi,uipk,og,platform,deadline,nbs,gen,os,trid&bvc=vod&nettype=0&bw=282689&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30080.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&os=akam&og=hw&mid=386935072&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=be44113f3d839618151fd5167e0227b4&uparams=e,oi,platform,deadline,nbs,gen,os,og,mid,uipk,trid&hdnts=exp=1756544207~hmac=68922f273f0c30f61337d0585a2488b1cae8e0d40e1554396ca3ec9e2cbcd866&bvc=vod&nettype=0&bw=282689&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30080.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&os=akam&og=hw&mid=386935072&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=be44113f3d839618151fd5167e0227b4&uparams=e,oi,platform,deadline,nbs,gen,os,og,mid,uipk,trid&hdnts=exp=1756544207~hmac=68922f273f0c30f61337d0585a2488b1cae8e0d40e1554396ca3ec9e2cbcd866&bvc=vod&nettype=0&bw=282689&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "bandwidth": 282115,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "avc1.640033",
          "width": 1892,
          "height": 1028,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-929",
            "indexRange": "930-1549"
          },
          "segment_base": {
            "initialization": "0-929",
            "index_range": "930-1549"
          },
          "codecid": 7
        },
        {
          "id": 80,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30077.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&platform=pc&deadline=1756544207&uipk=5&mid=386935072&gen=playurlv3&os=cosovbv&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&upsig=9e663174ee02b54ddb23e308bee71d20&uparams=e,oi,platform,deadline,uipk,mid,gen,os,nbs,trid,og&bvc=vod&nettype=0&bw=316258&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30077.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&platform=pc&deadline=1756544207&uipk=5&mid=386935072&gen=playurlv3&os=cosovbv&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&upsig=9e663174ee02b54ddb23e308bee71d20&uparams=e,oi,platform,deadline,uipk,mid,gen,os,nbs,trid,og&bvc=vod&nettype=0&bw=316258&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30077.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&nbs=1&uipk=5&mid=386935072&oi=1823807241&os=akam&og=cos&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&gen=playurlv3&upsig=b098f077fa2b5f4164203debcdfbd86a&uparams=e,platform,nbs,uipk,mid,oi,os,og,deadline,trid,gen&hdnts=exp=1756544207~hmac=6482569bc47a0386378353f44b8d0058f5080a02193dba089b5251f04c1996ee&bvc=vod&nettype=0&bw=316258&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30077.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&nbs=1&uipk=5&mid=386935072&oi=1823807241&os=akam&og=cos&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&gen=playurlv3&upsig=b098f077fa2b5f4164203debcdfbd86a&uparams=e,platform,nbs,uipk,mid,oi,os,og,deadline,trid,gen&hdnts=exp=1756544207~hmac=6482569bc47a0386378353f44b8d0058f5080a02193dba089b5251f04c1996ee&bvc=vod&nettype=0&bw=316258&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 315618,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "hev1.1.6.L150.90",
          "width": 1892,
          "height": 1028,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-1090",
            "indexRange": "1091-1710"
          },
          "segment_base": {
            "initialization": "0-1090",
            "index_range": "1091-1710"
          },
          "codecid": 12
        },
        {
          "id": 80,
          "baseUrl": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100026.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&deadline=1756544207&nbs=1&og=ali&oi=1823807241&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&os=akam&upsig=1a065e808da3ba56755c815301f9a9d3&uparams=e,platform,deadline,nbs,og,oi,uipk,trid,mid,gen,os&hdnts=exp=1756544207~hmac=37c1bbe08bcd7403ec73902e69f3f332105adfdb6cbd45a9e002020b7c0fa85f&bvc=vod&nettype=0&bw=149091&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&agrr=1&orderid=0,2",
          "base_url": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100026.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&deadline=1756544207&nbs=1&og=ali&oi=1823807241&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&os=akam&upsig=1a065e808da3ba56755c815301f9a9d3&uparams=e,platform,deadline,nbs,og,oi,uipk,trid,mid,gen,os&hdnts=exp=1756544207~hmac=37c1bbe08bcd7403ec73902e69f3f332105adfdb6cbd45a9e002020b7c0fa85f&bvc=vod&nettype=0&bw=149091&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&agrr=1&orderid=0,2",
          "backupUrl": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100026.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&os=cosovbv&deadline=1756544207&gen=playurlv3&og=ali&mid=386935072&oi=1823807241&platform=pc&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=9b8ef783dcb49c15b211498df406f84c&uparams=e,os,deadline,gen,og,mid,oi,platform,nbs,uipk,trid&bvc=vod&nettype=0&bw=149091&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100026.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&os=cosovbv&deadline=1756544207&gen=playurlv3&og=ali&mid=386935072&oi=1823807241&platform=pc&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&upsig=9b8ef783dcb49c15b211498df406f84c&uparams=e,os,deadline,gen,og,mid,oi,platform,nbs,uipk,trid&bvc=vod&nettype=0&bw=149091&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=1,2"
          ],
          "bandwidth": 148763,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "av01.0.00M.10.0.110.01.01.01.0",
          "width": 1892,
          "height": 1028,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-993",
            "indexRange": "994-1613"
          },
          "segment_base": {
            "initialization": "0-993",
            "index_range": "994-1613"
          },
          "codecid": 13
        },
        {
          "id": 64,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30064.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&deadline=1756544207&nbs=1&uipk=5&gen=playurlv3&oi=1823807241&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&os=cosovbv&og=cos&upsig=39fdc3f11f461b1a3886330d6ad2edec&uparams=e,platform,deadline,nbs,uipk,gen,oi,trid,mid,os,og&bvc=vod&nettype=0&bw=146888&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&agrr=1&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30064.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&deadline=1756544207&nbs=1&uipk=5&gen=playurlv3&oi=1823807241&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&os=cosovbv&og=cos&upsig=39fdc3f11f461b1a3886330d6ad2edec&uparams=e,platform,deadline,nbs,uipk,gen,oi,trid,mid,os,og&bvc=vod&nettype=0&bw=146888&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&agrr=1&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30064.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&deadline=1756544207&uipk=5&mid=386935072&oi=1823807241&platform=pc&gen=playurlv3&os=akam&upsig=7e7723651382400d9323b8e45b9212fa&uparams=e,nbs,trid,og,deadline,uipk,mid,oi,platform,gen,os&hdnts=exp=1756544207~hmac=1e85ed388fdbf6ae35ecdb71a819b4ebe049fe68b429d1b78a0efd18ee983ec6&bvc=vod&nettype=0&bw=146888&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30064.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&deadline=1756544207&uipk=5&mid=386935072&oi=1823807241&platform=pc&gen=playurlv3&os=akam&upsig=7e7723651382400d9323b8e45b9212fa&uparams=e,nbs,trid,og,deadline,uipk,mid,oi,platform,gen,os&hdnts=exp=1756544207~hmac=1e85ed388fdbf6ae35ecdb71a819b4ebe049fe68b429d1b78a0efd18ee983ec6&bvc=vod&nettype=0&bw=146888&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 146565,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "avc1.640033",
          "width": 1280,
          "height": 694,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-928",
            "indexRange": "929-1548"
          },
          "segment_base": {
            "initialization": "0-928",
            "index_range": "929-1548"
          },
          "codecid": 7
        },
        {
          "id": 64,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30066.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&os=cosovbv&uipk=5&oi=1823807241&deadline=1756544207&gen=playurlv3&og=hw&nbs=1&upsig=01e95b7e2cea48e23c5bec3b494bdc40&uparams=e,trid,mid,platform,os,uipk,oi,deadline,gen,og,nbs&bvc=vod&nettype=0&bw=181923&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30066.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&os=cosovbv&uipk=5&oi=1823807241&deadline=1756544207&gen=playurlv3&og=hw&nbs=1&upsig=01e95b7e2cea48e23c5bec3b494bdc40&uparams=e,trid,mid,platform,os,uipk,oi,deadline,gen,og,nbs&bvc=vod&nettype=0&bw=181923&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30066.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&nbs=1&uipk=5&gen=playurlv3&os=akam&og=hw&upsig=632fa89601734f1190255bf4940ed013&uparams=e,oi,deadline,trid,mid,platform,nbs,uipk,gen,os,og&hdnts=exp=1756544207~hmac=49e024fc87b33026fedb99bb504c49683e04b8b4479d96c007e13e4ea0ecbd9f&bvc=vod&nettype=0&bw=181923&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30066.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&nbs=1&uipk=5&gen=playurlv3&os=akam&og=hw&upsig=632fa89601734f1190255bf4940ed013&uparams=e,oi,deadline,trid,mid,platform,nbs,uipk,gen,os,og&hdnts=exp=1756544207~hmac=49e024fc87b33026fedb99bb504c49683e04b8b4479d96c007e13e4ea0ecbd9f&bvc=vod&nettype=0&bw=181923&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 181531,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "hev1.1.6.L120.90",
          "width": 1280,
          "height": 694,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-1091",
            "indexRange": "1092-1711"
          },
          "segment_base": {
            "initialization": "0-1091",
            "index_range": "1092-1711"
          },
          "codecid": 12
        },
        {
          "id": 64,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100024.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&gen=playurlv3&os=cosovbv&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&platform=pc&uipk=5&deadline=1756544207&upsig=efb7d84348ed80f289fffb5057cd07e2&uparams=e,nbs,gen,os,og,trid,mid,oi,platform,uipk,deadline&bvc=vod&nettype=0&bw=97927&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100024.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&gen=playurlv3&os=cosovbv&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&platform=pc&uipk=5&deadline=1756544207&upsig=efb7d84348ed80f289fffb5057cd07e2&uparams=e,nbs,gen,os,og,trid,mid,oi,platform,uipk,deadline&bvc=vod&nettype=0&bw=97927&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100024.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&platform=pc&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=hw&oi=1823807241&deadline=1756544207&uipk=5&gen=playurlv3&os=akam&upsig=b4f7fe5e6c670745392d57f2ac17251f&uparams=e,mid,platform,nbs,trid,og,oi,deadline,uipk,gen,os&hdnts=exp=1756544207~hmac=fc68bae09175c05c33e37d457d9afd0255d5b19dbd1075fc83afd5d1f54bcb7c&bvc=vod&nettype=0&bw=97927&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100024.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&platform=pc&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=hw&oi=1823807241&deadline=1756544207&uipk=5&gen=playurlv3&os=akam&upsig=b4f7fe5e6c670745392d57f2ac17251f&uparams=e,mid,platform,nbs,trid,og,oi,deadline,uipk,gen,os&hdnts=exp=1756544207~hmac=fc68bae09175c05c33e37d457d9afd0255d5b19dbd1075fc83afd5d1f54bcb7c&bvc=vod&nettype=0&bw=97927&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 97693,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "av01.0.00M.10.0.110.01.01.01.0",
          "width": 1280,
          "height": 694,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-993",
            "indexRange": "994-1613"
          },
          "segment_base": {
            "initialization": "0-993",
            "index_range": "994-1613"
          },
          "codecid": 13
        },
        {
          "id": 32,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30032.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&uipk=5&gen=playurlv3&og=hw&platform=pc&deadline=1756544207&nbs=1&os=cosovbv&upsig=384c15c5dcf9b56d768dc92db14bd23d&uparams=e,trid,mid,oi,uipk,gen,og,platform,deadline,nbs,os&bvc=vod&nettype=0&bw=71302&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30032.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&uipk=5&gen=playurlv3&og=hw&platform=pc&deadline=1756544207&nbs=1&os=cosovbv&upsig=384c15c5dcf9b56d768dc92db14bd23d&uparams=e,trid,mid,oi,uipk,gen,og,platform,deadline,nbs,os&bvc=vod&nettype=0&bw=71302&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30032.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&os=akam&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&deadline=1756544207&gen=playurlv3&oi=1823807241&upsig=7768ed8d23b7f606abe1354b9a7d168a&uparams=e,nbs,uipk,os,og,trid,mid,platform,deadline,gen,oi&hdnts=exp=1756544207~hmac=1d0c8e22c2847d1c0a36c59674ab6114a998b49457bcc389d20af44fb808ac35&bvc=vod&nettype=0&bw=71302&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30032.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&os=akam&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&deadline=1756544207&gen=playurlv3&oi=1823807241&upsig=7768ed8d23b7f606abe1354b9a7d168a&uparams=e,nbs,uipk,os,og,trid,mid,platform,deadline,gen,oi&hdnts=exp=1756544207~hmac=1d0c8e22c2847d1c0a36c59674ab6114a998b49457bcc389d20af44fb808ac35&bvc=vod&nettype=0&bw=71302&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 71119,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "avc1.640033",
          "width": 854,
          "height": 464,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-928",
            "indexRange": "929-1548"
          },
          "segment_base": {
            "initialization": "0-928",
            "index_range": "929-1548"
          },
          "codecid": 7
        },
        {
          "id": 32,
          "baseUrl": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30033.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&deadline=1756544207&nbs=1&gen=playurlv3&mid=386935072&uipk=5&os=akam&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&platform=pc&upsig=1c7f5bd09e92773b13b132aad2cfbb59&uparams=e,deadline,nbs,gen,mid,uipk,os,og,trid,oi,platform&hdnts=exp=1756544207~hmac=df63bd0ef77c2d869eac66da964764af283e5bf04d89762447f174077acfb7b1&bvc=vod&nettype=0&bw=96778&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30033.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&deadline=1756544207&nbs=1&gen=playurlv3&mid=386935072&uipk=5&os=akam&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&platform=pc&upsig=1c7f5bd09e92773b13b132aad2cfbb59&uparams=e,deadline,nbs,gen,mid,uipk,os,og,trid,oi,platform&hdnts=exp=1756544207~hmac=df63bd0ef77c2d869eac66da964764af283e5bf04d89762447f174077acfb7b1&bvc=vod&nettype=0&bw=96778&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30033.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&deadline=1756544207&uipk=5&os=cosovbv&platform=pc&nbs=1&gen=playurlv3&og=hw&upsig=1edffd097bbf5fb428e24ed97ae598d9&uparams=e,trid,mid,oi,deadline,uipk,os,platform,nbs,gen,og&bvc=vod&nettype=0&bw=96778&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30033.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&deadline=1756544207&uipk=5&os=cosovbv&platform=pc&nbs=1&gen=playurlv3&og=hw&upsig=1edffd097bbf5fb428e24ed97ae598d9&uparams=e,trid,mid,oi,deadline,uipk,os,platform,nbs,gen,og&bvc=vod&nettype=0&bw=96778&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 96543,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "hev1.1.6.L120.90",
          "width": 854,
          "height": 464,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-1089",
            "indexRange": "1090-1709"
          },
          "segment_base": {
            "initialization": "0-1089",
            "index_range": "1090-1709"
          },
          "codecid": 12
        },
        {
          "id": 32,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100023.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&os=cosovbv&deadline=1756544207&oi=1823807241&og=hw&upsig=835d4554c3271b768e2f162836771819&uparams=e,platform,nbs,uipk,trid,mid,gen,os,deadline,oi,og&bvc=vod&nettype=0&bw=50604&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100023.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&os=cosovbv&deadline=1756544207&oi=1823807241&og=hw&upsig=835d4554c3271b768e2f162836771819&uparams=e,platform,nbs,uipk,trid,mid,gen,os,deadline,oi,og&bvc=vod&nettype=0&bw=50604&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100023.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&deadline=1756544207&nbs=1&gen=playurlv3&os=akam&mid=386935072&oi=1823807241&platform=pc&uipk=5&og=hw&upsig=4db0ff42baa0bc29eda51cc2700bb618&uparams=e,trid,deadline,nbs,gen,os,mid,oi,platform,uipk,og&hdnts=exp=1756544207~hmac=f67dcf4aab750081ce07c070e19e5845386b2c8356a0710578c749391943b3aa&bvc=vod&nettype=0&bw=50604&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100023.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&deadline=1756544207&nbs=1&gen=playurlv3&os=akam&mid=386935072&oi=1823807241&platform=pc&uipk=5&og=hw&upsig=4db0ff42baa0bc29eda51cc2700bb618&uparams=e,trid,deadline,nbs,gen,os,mid,oi,platform,uipk,og&hdnts=exp=1756544207~hmac=f67dcf4aab750081ce07c070e19e5845386b2c8356a0710578c749391943b3aa&bvc=vod&nettype=0&bw=50604&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 50458,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "av01.0.00M.10.0.110.01.01.01.0",
          "width": 854,
          "height": 464,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-993",
            "indexRange": "994-1613"
          },
          "segment_base": {
            "initialization": "0-993",
            "index_range": "994-1613"
          },
          "codecid": 13
        },
        {
          "id": 16,
          "baseUrl": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30011.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&nbs=1&gen=playurlv3&mid=386935072&platform=pc&deadline=1756544207&uipk=5&os=akam&og=ali&upsig=b4a9fe920c3188001c344f3183285503&uparams=e,trid,oi,nbs,gen,mid,platform,deadline,uipk,os,og&hdnts=exp=1756544207~hmac=0d94e71c93fedceb7c5e0cc61575d4835c7b80065a7ddf9302f8424fbbf5ce68&bvc=vod&nettype=0&bw=60683&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "base_url": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30011.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&nbs=1&gen=playurlv3&mid=386935072&platform=pc&deadline=1756544207&uipk=5&os=akam&og=ali&upsig=b4a9fe920c3188001c344f3183285503&uparams=e,trid,oi,nbs,gen,mid,platform,deadline,uipk,os,og&hdnts=exp=1756544207~hmac=0d94e71c93fedceb7c5e0cc61575d4835c7b80065a7ddf9302f8424fbbf5ce68&bvc=vod&nettype=0&bw=60683&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "backupUrl": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30011.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&os=cosovbv&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&og=ali&upsig=cc38ff04b21dc7154992029bd4763ef5&uparams=e,uipk,oi,platform,deadline,nbs,os,trid,mid,gen,og&bvc=vod&nettype=0&bw=60683&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30011.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&os=cosovbv&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&gen=playurlv3&og=ali&upsig=cc38ff04b21dc7154992029bd4763ef5&uparams=e,uipk,oi,platform,deadline,nbs,os,trid,mid,gen,og&bvc=vod&nettype=0&bw=60683&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 60515,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "hev1.1.6.L120.90",
          "width": 640,
          "height": 348,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-1089",
            "indexRange": "1090-1709"
          },
          "segment_base": {
            "initialization": "0-1089",
            "index_range": "1090-1709"
          },
          "codecid": 12
        },
        {
          "id": 16,
          "baseUrl": "https://upos-sz-mirrorawsov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30016.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&og=cos&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&os=awsovbv&mid=386935072&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&upsig=657ca54857411f9220f06627722d3a6f&uparams=e,og,uipk,trid,os,mid,oi,platform,deadline,nbs,gen&bvc=vod&nettype=0&bw=49445&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorawsov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30016.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&og=cos&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&os=awsovbv&mid=386935072&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&gen=playurlv3&upsig=657ca54857411f9220f06627722d3a6f&uparams=e,og,uipk,trid,os,mid,oi,platform,deadline,nbs,gen&bvc=vod&nettype=0&bw=49445&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30016.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&og=cos&deadline=1756544207&mid=386935072&os=akam&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&platform=pc&gen=playurlv3&upsig=e2a634e96a807a4580587aa2b799cf74&uparams=e,og,deadline,mid,os,nbs,uipk,trid,oi,platform,gen&hdnts=exp=1756544207~hmac=48a613eda78062fe33795524aff25a4e966a78c5e3a98d9c4aefd799835c07ea&bvc=vod&nettype=0&bw=49445&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30016.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&og=cos&deadline=1756544207&mid=386935072&os=akam&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&oi=1823807241&platform=pc&gen=playurlv3&upsig=e2a634e96a807a4580587aa2b799cf74&uparams=e,og,deadline,mid,os,nbs,uipk,trid,oi,platform,gen&hdnts=exp=1756544207~hmac=48a613eda78062fe33795524aff25a4e966a78c5e3a98d9c4aefd799835c07ea&bvc=vod&nettype=0&bw=49445&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 49302,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "avc1.640033",
          "width": 640,
          "height": 348,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-935",
            "indexRange": "936-1555"
          },
          "segment_base": {
            "initialization": "0-935",
            "index_range": "936-1555"
          },
          "codecid": 7
        },
        {
          "id": 16,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100022.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&platform=pc&gen=playurlv3&os=cosovbv&nbs=1&uipk=5&mid=386935072&oi=1823807241&og=hw&upsig=803b75bba36ae866c3cbec8a90b830f8&uparams=e,deadline,trid,platform,gen,os,nbs,uipk,mid,oi,og&bvc=vod&nettype=0&bw=31795&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-100022.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&platform=pc&gen=playurlv3&os=cosovbv&nbs=1&uipk=5&mid=386935072&oi=1823807241&og=hw&upsig=803b75bba36ae866c3cbec8a90b830f8&uparams=e,deadline,trid,platform,gen,os,nbs,uipk,mid,oi,og&bvc=vod&nettype=0&bw=31795&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100022.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&os=akam&og=hw&nbs=1&uipk=5&platform=pc&gen=playurlv3&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&upsig=0b096cf75d85f6c431d90470f26eb678&uparams=e,oi,os,og,nbs,uipk,platform,gen,deadline,trid,mid&hdnts=exp=1756544207~hmac=2f5a759083fc1428af40830eb86767f4c632922e5be94276f70351c8d12ac354&bvc=vod&nettype=0&bw=31795&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-100022.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&oi=1823807241&os=akam&og=hw&nbs=1&uipk=5&platform=pc&gen=playurlv3&deadline=1756544207&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&upsig=0b096cf75d85f6c431d90470f26eb678&uparams=e,oi,os,og,nbs,uipk,platform,gen,deadline,trid,mid&hdnts=exp=1756544207~hmac=2f5a759083fc1428af40830eb86767f4c632922e5be94276f70351c8d12ac354&bvc=vod&nettype=0&bw=31795&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "bandwidth": 31684,
          "mimeType": "video/mp4",
          "mime_type": "video/mp4",
          "codecs": "av01.0.00M.10.0.110.01.01.01.0",
          "width": 640,
          "height": 348,
          "frameRate": "9.994",
          "frame_rate": "9.994",
          "sar": "N/A",
          "startWithSap": 1,
          "start_with_sap": 1,
          "SegmentBase": {
            "Initialization": "0-993",
            "indexRange": "994-1613"
          },
          "segment_base": {
            "initialization": "0-993",
            "index_range": "994-1613"
          },
          "codecid": 13
        }
      ],
      "audio": [
        {
          "id": 30232,
          "baseUrl": "https://upos-sz-mirrorawsov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30232.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&gen=playurlv3&os=awsovbv&og=ali&oi=1823807241&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&deadline=1756544207&upsig=5933740fc246f4f0240e7c4a114d5d4f&uparams=e,nbs,uipk,gen,os,og,oi,trid,mid,platform,deadline&bvc=vod&nettype=0&bw=65837&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "base_url": "https://upos-sz-mirrorawsov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30232.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&gen=playurlv3&os=awsovbv&og=ali&oi=1823807241&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&platform=pc&deadline=1756544207&upsig=5933740fc246f4f0240e7c4a114d5d4f&uparams=e,nbs,uipk,gen,os,og,oi,trid,mid,platform,deadline&bvc=vod&nettype=0&bw=65837&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30232.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&platform=pc&os=akam&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&deadline=1756544207&nbs=1&uipk=5&gen=playurlv3&og=ali&upsig=07448d1dd4698cba2e5e9dbbef3fdc34&uparams=e,mid,oi,platform,os,trid,deadline,nbs,uipk,gen,og&hdnts=exp=1756544207~hmac=8affc9e36b6fb11372168b3c86947a9c58f8031518897aad469e51356248c8bd&bvc=vod&nettype=0&bw=65837&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30232.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&platform=pc&os=akam&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&deadline=1756544207&nbs=1&uipk=5&gen=playurlv3&og=ali&upsig=07448d1dd4698cba2e5e9dbbef3fdc34&uparams=e,mid,oi,platform,os,trid,deadline,nbs,uipk,gen,og&hdnts=exp=1756544207~hmac=8affc9e36b6fb11372168b3c86947a9c58f8031518897aad469e51356248c8bd&bvc=vod&nettype=0&bw=65837&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&orderid=1,2"
          ],
          "bandwidth": 65574,
          "mimeType": "audio/mp4",
          "mime_type": "audio/mp4",
          "codecs": "mp4a.40.2",
          "width": 0,
          "height": 0,
          "frameRate": "",
          "frame_rate": "",
          "sar": "",
          "startWithSap": 0,
          "start_with_sap": 0,
          "SegmentBase": {
            "Initialization": "0-817",
            "indexRange": "818-1437"
          },
          "segment_base": {
            "initialization": "0-817",
            "index_range": "818-1437"
          },
          "codecid": 0
        },
        {
          "id": 30216,
          "baseUrl": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30216.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&uipk=5&gen=playurlv3&os=akam&upsig=3c075ce314cbcb9b4122e9db7c3f4b55&uparams=e,mid,oi,platform,deadline,nbs,trid,og,uipk,gen,os&hdnts=exp=1756544207~hmac=0a6ca7d20986c64c8daddf6d7cfdf776cda1ab7a717febe8bfba55f03cdc9f87&bvc=vod&nettype=0&bw=65977&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "base_url": "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30216.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&mid=386935072&oi=1823807241&platform=pc&deadline=1756544207&nbs=1&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&og=cos&uipk=5&gen=playurlv3&os=akam&upsig=3c075ce314cbcb9b4122e9db7c3f4b55&uparams=e,mid,oi,platform,deadline,nbs,trid,og,uipk,gen,os&hdnts=exp=1756544207~hmac=0a6ca7d20986c64c8daddf6d7cfdf776cda1ab7a717febe8bfba55f03cdc9f87&bvc=vod&nettype=0&bw=65977&build=0&dl=0&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&orderid=0,2",
          "backupUrl": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30216.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&gen=playurlv3&mid=386935072&deadline=1756544207&os=cosovbv&og=cos&oi=1823807241&platform=pc&nbs=1&uipk=5&upsig=c2ca122f0f1cf401b8800c7af369b9b6&uparams=e,trid,gen,mid,deadline,os,og,oi,platform,nbs,uipk&bvc=vod&nettype=0&bw=65977&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30216.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&gen=playurlv3&mid=386935072&deadline=1756544207&os=cosovbv&og=cos&oi=1823807241&platform=pc&nbs=1&uipk=5&upsig=c2ca122f0f1cf401b8800c7af369b9b6&uparams=e,trid,gen,mid,deadline,os,og,oi,platform,nbs,uipk&bvc=vod&nettype=0&bw=65977&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=1,2"
          ],
          "bandwidth": 65713,
          "mimeType": "audio/mp4",
          "mime_type": "audio/mp4",
          "codecs": "mp4a.40.2",
          "width": 0,
          "height": 0,
          "frameRate": "",
          "frame_rate": "",
          "sar": "",
          "startWithSap": 0,
          "start_with_sap": 0,
          "SegmentBase": {
            "Initialization": "0-825",
            "indexRange": "826-1445"
          },
          "segment_base": {
            "initialization": "0-825",
            "index_range": "826-1445"
          },
          "codecid": 0
        },
        {
          "id": 30280,
          "baseUrl": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30280.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&mid=386935072&oi=1823807241&deadline=1756544207&gen=playurlv3&os=cosovbv&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&platform=pc&upsig=25d10fe30a3959261fb75090dc7ff72f&uparams=e,nbs,uipk,mid,oi,deadline,gen,os,og,trid,platform&bvc=vod&nettype=0&bw=65837&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "base_url": "https://upos-sz-mirrorcosov.bilivideo.com/upgcxcode/27/88/31649038827/31649038827-1-30280.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&nbs=1&uipk=5&mid=386935072&oi=1823807241&deadline=1756544207&gen=playurlv3&os=cosovbv&og=hw&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&platform=pc&upsig=25d10fe30a3959261fb75090dc7ff72f&uparams=e,nbs,uipk,mid,oi,deadline,gen,os,og,trid,platform&bvc=vod&nettype=0&bw=65837&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&f=u_0_0&orderid=0,2",
          "backupUrl": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30280.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&gen=playurlv3&og=hw&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&os=akam&deadline=1756544207&upsig=d4a45e16ef7f16f08ab5daab6a451374&uparams=e,platform,gen,og,nbs,uipk,trid,mid,oi,os,deadline&hdnts=exp=1756544207~hmac=4fc14cef527fe48ec10769e1915af34a6c65344793bdd5f3da5310d088c456f9&bvc=vod&nettype=0&bw=65837&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=1,2"
          ],
          "backup_url": [
            "https://upos-hz-mirrorakam.akamaized.net/upgcxcode/27/88/31649038827/31649038827-1-30280.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&platform=pc&gen=playurlv3&og=hw&nbs=1&uipk=5&trid=6fd49c876fb34aeaa2d9804a8f96ce8u&mid=386935072&oi=1823807241&os=akam&deadline=1756544207&upsig=d4a45e16ef7f16f08ab5daab6a451374&uparams=e,platform,gen,og,nbs,uipk,trid,mid,oi,os,deadline&hdnts=exp=1756544207~hmac=4fc14cef527fe48ec10769e1915af34a6c65344793bdd5f3da5310d088c456f9&bvc=vod&nettype=0&bw=65837&f=u_0_0&agrr=1&buvid=0F76B854-35CE-A328-3D9D-B7C768354FFF01167infoc&build=0&dl=0&orderid=1,2"
          ],
          "bandwidth": 65574,
          "mimeType": "audio/mp4",
          "mime_type": "audio/mp4",
          "codecs": "mp4a.40.2",
          "width": 0,
          "height": 0,
          "frameRate": "",
          "frame_rate": "",
          "sar": "",
          "startWithSap": 0,
          "start_with_sap": 0,
          "SegmentBase": {
            "Initialization": "0-817",
            "indexRange": "818-1437"
          },
          "segment_base": {
            "initialization": "0-817",
            "index_range": "818-1437"
          },
          "codecid": 0
        }
      ],
      "dolby": {
        "type": 0,
        "audio": null
      },
      "flac": null
    },
    "support_formats": [
      {
        "quality": 80,
        "format": "flv",
        "new_description": "1080P 高清",
        "display_desc": "1080P",
        "superscript": "",
        "codecs": [
          "av01.0.00M.10.0.110.01.01.01.0",
          "avc1.640033",
          "hev1.1.6.L150.90"
        ]
      },
      {
        "quality": 64,
        "format": "flv720",
        "new_description": "720P 准高清",
        "display_desc": "720P",
        "superscript": "",
        "codecs": [
          "av01.0.00M.10.0.110.01.01.01.0",
          "avc1.640033",
          "hev1.1.6.L120.90"
        ]
      },
      {
        "quality": 32,
        "format": "flv480",
        "new_description": "480P 标清",
        "display_desc": "480P",
        "superscript": "",
        "codecs": [
          "av01.0.00M.10.0.110.01.01.01.0",
          "avc1.640033",
          "hev1.1.6.L120.90"
        ]
      },
      {
        "quality": 16,
        "format": "mp4",
        "new_description": "360P 流畅",
        "display_desc": "360P",
        "superscript": "",
        "codecs": [
          "av01.0.00M.10.0.110.01.01.01.0",
          "avc1.640033",
          "hev1.1.6.L120.90"
        ]
      }
    ],
    "high_format": null,
    "last_play_time": 117000,
    "last_play_cid": 31649038827,
    "view_info": null,
    "play_conf": {
      "is_new_description": false
    }
  }
}`
