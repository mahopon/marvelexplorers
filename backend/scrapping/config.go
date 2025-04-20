package scrapping

import (
	"os"
	"strconv"
	"time"
)

var (
	public_key  string = os.Getenv("PUBLIC_KEY")
	private_key string = os.Getenv("PRIVATE_KEY")
	ts          string = strconv.FormatInt(time.Now().Unix(), 10)
	hashString  string = getHash(ts, private_key, public_key)
)
