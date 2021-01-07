package ipquery

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	ipRangeFieldCount = 3
	ipRange           = 4
)

type IpRange struct {
	Begin uint32
	End   uint32
	Data  string
}

type ipData []*IpRange

//func NewIpData() *IpData {
//	return &IpData{}
//}
//
var IpData = &ipData{}

//
//func init() {
//	defaultIpData = NewIpData()
//}

var ErrorIpRangeNotFound = errors.New("ip range not found")

func (id *ipData) City(ip string) (city string, err error) {
	//df := "/home/website/src/111.txt"
	df := "/home/website/src/github.com/tabalt/ipquery/testdata/ip_chunzhen.txt"
	err = id.Load(df)
	if err != nil {
		return
	}
	dt, err := id.Find(ip)
	if err != nil {
		return
	} else {
		return dt, nil
	}
}

func (id *ipData) Load(df string) error {
	reader, err := os.Open(df)
	if err != nil {
		return err
	}
	return id.Loader(reader)
}

func (id *ipData) Find(ip string) (string, error) {
	ir, err := id.Finder(ip)
	if err != nil {
		return "", err
	} else {
		return ir.Data, nil
	}
}

func (id *ipData) Loader(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		item := strings.SplitN(line, "\t", ipRangeFieldCount)
		if len(item) != ipRangeFieldCount {
			continue
		}
		//itembegin := strings.SplitN(item[0], ".", ipRange)
		//itemend := strings.SplitN(item[1], ".", ipRange)
		//itembeginint := [4]int{}
		//itemendint := [4]int{}
		//for i := 0; i < 4; i++ {
		//	itembeginint[i], _ = strconv.Atoi(itembegin[i])
		//	itemendint[i], _ = strconv.Atoi(itemend[i])
		//}
		//
		//begin := itembeginint[0]<<24 | itembeginint[1]<<16 | itembeginint[2]<<8 | itembeginint[3]
		//end := itemendint[0]<<24 | itemendint[1]<<16 | itemendint[2]<<8 | itemendint[3]
		begin, _ := strconv.Atoi(item[0])
		end, _ := strconv.Atoi(item[1])
		if begin > end {
			continue
		}

		ir := &IpRange{
			Begin: uint32(begin),
			End:   uint32(end),
			Data:  item[2],
		}

		*id = append(*id, ir)
	}

	return scanner.Err()
}

func (id *ipData) Finder(ip string) (*IpRange, error) {
	ir, err := id.getIpRange(ip)
	if err != nil {
		return nil, err
	}

	return ir, nil
}

func (id *ipData) Length() int {
	return len(*id)
}

func (id *ipData) getIpRange(ip string) (*IpRange, error) {
	var low, high int = 0, (id.Length() - 1)

	ipdt := *id
	il := ip2Long(ip)
	if il <= 0 {
		return nil, ErrorIpRangeNotFound
	}

	for low <= high {
		var middle int = (high-low)/2 + low

		ir := ipdt[middle]

		if il >= ir.Begin && il <= ir.End {
			return ir, nil
		} else if il < ir.Begin {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return nil, ErrorIpRangeNotFound
}

func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	fmt.Println("iplong: ", long)
	return long
}
