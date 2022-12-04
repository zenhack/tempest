package sessioncookies

import (
	"crypto/rand"
	"io/ioutil"
	"os"

	"zenhack.net/go/sandstorm-next/go/internal/config"
)

func GetKeys() ([][]byte, error) {
	var (
		data []byte
		err  error
	)
	const path = config.Localstatedir + "/sandstorm/session-key"
	data, err = ioutil.ReadFile(config.Localstatedir + "/sandstorm/session-key")
	if os.IsNotExist(err) {
		data = make([]byte, 64)
		rand.Read(data)
		err = os.WriteFile(path, data, 0600)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return [][]byte{data[:32], data[32:]}, nil
}
