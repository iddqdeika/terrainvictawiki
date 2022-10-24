package localization

import (
	"os"
	"strings"
)

func NewStorage(filename string) (*Storage, error) {
	data, err:= load(filename)
	if err != nil{
		return nil, err
	}
	return &Storage{
		data: data,
	}, nil
}

func load(filename string) (map[string]string, error){
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(data), "\r\n")
	res := make(map[string]string)
	for _, row := range rows {
		parts := strings.SplitN(row, "=", 2)
		if len(parts)==2{
			res[parts[0]] = parts[1]
		}
	}
	return res, nil
}

type Storage struct {
	data map[string]string
}

func (s *Storage) Get(root, group, name string) string {
	return s.data[key(root, group, name)]
}

func key(root, group, name string) string {
	return root + "." + group + "." + name
}
