package types

//@see https://aria2.github.io/manual/en/html/aria2c.html#http-ftp-sftp-options
type Aria2HttpOptions struct {
	AllProxy               string `json:"all-proxy"`
	AllProxyPasswd         string `json:"all-proxy-passwd"`
	AllProxyUser           string `json:"all-proxy-user"`
	Checksum               string `json:"checksum"`                  //<TYPE>=<DIGEST>
	ConnectTimeout         int    `json:"connect-timeout"`           //@default 60
	DryRun                 bool   `json:"dry-run"`                   //@default false
	LowestSpeedLimit       int    `json:"lowest-speed-limit"`        //Close connection if download speed is lower than or equal to this value(bytes per sec). 0 means aria2 does not have a lowest speed limit. You can append K or M (1K = 1024, 1M = 1024K). This option does not affect BitTorrent downloads. @default 0
	MaxConnectionPerServer int    `json:"max-connection-per-server"` //@default 1
	MaxFileNotFound        int    `json:"max-file-not-found"`        //@default 0
	MaxTries               int    `json:"max-tries"`                 //@default 0
	MinSplitSize           string `json:"min-split-size"`            //aria2 does not split less than 2*SIZE byte range. For example, let's consider downloading 20MiB file. If SIZE is 10M, aria2 can split file into 2 range [0-10MiB) and [10MiB-20MiB) and download it using 2 sources(if --split >= 2, of course). If SIZE is 15M, since 2*15M > 20MiB, aria2 does not split file and download it using 1 source. You can append K or M (1K = 1024, 1M = 1024K). Possible Values: 1M -1024M @default 20M
	NetrcPath              string `json:"netrc-path"`                //@default $(HOME)/.netrc
	NoNetrc                bool   `json:"no-netrc"`
	NoProxy                bool   `json:"no-proxy"`
	Out                    string `json:"out"`
	ProxyMethod            string `json:"proxy-method"` //Set the method to use in proxy request. METHOD is either get or tunnel. HTTPS downloads always use tunnel regardless of this option. @default get @values('get' | 'tunnel')
	RemoteTime             bool   `json:"remote-time"`  //@default false
	ReuseURI               bool   `json:"reuse-uri"`    //Reuse already used URIs if no unused URIs are left. @default true
	RetryWait              int    `json:"retry-wait"`   //Set the seconds to wait between retries. When SEC > 0, aria2 will retry downloads when the HTTP server returns a 503 response. @default 0
	ServerStatOf           string `json:"server-stat-of"`
	ServerStatIf           string `json:"server-stat-if"`
	ServerStatTimeout      int    `json:"server-stat-timeout"`   //Specifies timeout in seconds to invalidate performance profile of the servers since the last contact to them. @default 86400 (24hours)
	Split                  int    `json:"split"`                 //@default 5
	StreamPieceSelector    string `json:"stream-piece-selector"` //@default default
	Timeout                int    `json:"timeout"`               //@default 60
	URISelector            string `json:"uri-selector"`          //@default feedback
}

type Aria2BaseOptions struct {
	Dir                    string `json:"dir"`
	InputFile              string `json:"input-file"`
	Log                    string `json:"log"`
	MaxConcurrentDownloads int    `json:"max-concurrent-downloads"` //@default 5
	CheckIntegrity         bool   `json:"check-integrity"`          //@default false
	Continue               bool   `json:"continue"`
	Help                   string `json:"help"`
}

type Aria2Options struct {
	Aria2HttpOptions
	Aria2BaseOptions
}

type GID = string

type GlobalStat struct {
	Downloadspeed   string `json:"downloadSpeed"`
	Uploadspeed     string `json:"uploadSpeed"`
	Numactive       string `json:"numActive"`
	Numwaiting      string `json:"numWaiting"`
	Numstopped      string `json:"numStopped"`
	Numstoppedtotal string `json:"numStoppedTotal"`
}

type AriaFile struct {
	Completedlength string    `json:"completedLength"`
	Index           string    `json:"index"`
	Length          string    `json:"length"`
	Path            string    `json:"path"`
	Selected        string    `json:"selected"` //@values('true' | 'false')
	Uris            []AriaUri `json:"uris"`
}
type AriaUri struct {
	URI    string `json:"uri"`
	Status string `json:"status"` //@values('used' | 'waiting')
}

type DownloadStatus struct {
	Gid                    string                 `json:"gid"`
	Status                 string                 `json:"status"` //@values('active' | 'waiting' | 'paused' | 'error' | 'complete' | 'removed')
	Totallength            string                 `json:"totalLength"`
	Completedlength        string                 `json:"completedLength"`
	Uploadlength           string                 `json:"uploadLength"`
	Bitfield               string                 `json:"bitfield"`
	Downloadspeed          string                 `json:"downloadSpeed"`
	Uploadspeed            string                 `json:"uploadSpeed"`
	Infohash               string                 `json:"infoHash"`
	Numseeders             string                 `json:"numSeeders"`
	Seeder                 bool                   `json:"seeder"`
	Piecelength            string                 `json:"pieceLength"`
	Numpieces              string                 `json:"numPieces"`
	Connections            string                 `json:"connections"`
	Errorcode              string                 `json:"errorCode"` //@see https://aria2.github.io/manual/en/html/aria2c.html#id1
	Errormessage           string                 `json:"errorMessage"`
	Followedby             string                 `json:"followedBy"`
	Following              string                 `json:"following"`
	Belongsto              string                 `json:"belongsTo"`
	Dir                    string                 `json:"dir"`
	Files                  []AriaFile             `json:"files"`
	Bittorrent             map[string]interface{} `json:"bittorrent"`
	Verifiedlength         string                 `json:"verifiedLength"`
	Verifyintegritypending bool                   `json:"verifyIntegrityPending"`
}

type AriaVersion struct {
	Version         string   `json:"version"`
	Enabledfeatures []string `json:"enabledFeatures"`
}
