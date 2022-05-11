package util

import (
	"os"
	"poetry/global"
)

func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}


func CreateLogDir() {
	
    if ok, _ := PathExists(global.Settings.LogsAddress); !ok {
        _ = os.Mkdir(global.Settings.LogsAddress, os.ModePerm)
    }
}
