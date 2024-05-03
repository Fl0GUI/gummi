package throne

import (
	"errors"
	"fmt"
)

func Extract(url string) (creatorId string, err error) {
	var n int
	n, err = fmt.Sscanf(url, "https://throne.com/stream-alerts/%s", &creatorId)
	if n != 1 {
		err = errors.New("did not extract the expected amount of symbols from url.")
	}
	return
}
