package common

import (
	"net"
	"os"
	"time"

	"github.com/cppsky/go-logger/logger"
	"github.com/go-errors/errors"
)

func LogError(e interface{}) {
	if e == nil {
		return
	}
	logger.ErrorN(1, errors.Wrap(e, 1).ErrorStack())
}

func CheckError(e error) {
	if e != nil {
		LogError(e)
		panic(e)
	}
}

func getIntefaceIps(i net.Interface) []net.IP {
	ips := []net.IP{}
	addrs, err := i.Addrs()
	if err != nil {
		logger.Error(err)
		return ips
	}
	// handle err
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			ips = append(ips, v.IP)
		case *net.IPAddr:
			ips = append(ips, v.IP)
		}
		// process IP address
	}
	return ips
}

func LocalIps() []net.IP {
	ifaces, err := net.Interfaces()
	CheckError(err)
	// handle err
	ips := []net.IP{}
	for _, i := range ifaces {

		if err != nil {
			logger.Error(err)
			continue
		}
		ips = append(ips, getIntefaceIps(i)...)

	}
	return ips
}

func StrToTime(t string) time.Time {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err == nil {
		return tm
	}
	tm, err = time.ParseInLocation("2006-01-02", t, time.Local)
	CheckError(err)

	return tm
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FileExists check the file if exists, if not exits, return false and fail error
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CopyMap ...
func CopyMap(src, dest map[string]interface{}) {
	for rk, rv := range src {
		dest[rk] = rv
	}
}
