package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout      int64
	size         int
	count        int
	typ          uint8 = 8
	code         uint8 = 0
	sendCount    int
	successCount int
	failCount    int
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	ID          uint16
	SequenceNum uint16
}

func main() {

	getCommandArgs()
	//fmt.Println(timeout, size, count)

	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	fmt.Printf("PING %s(%s): %d data bytes\n", desIp, conn.RemoteAddr(), size)

	for i := 0; i < count; i++ {
		start := time.Now()
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			Checksum:    0,
			ID:          1,
			SequenceNum: 1,
		}

		data := make([]byte, size)
		var buffer bytes.Buffer
		binary.Write(&buffer, binary.BigEndian, icmp)
		buffer.Write(data)
		data = buffer.Bytes()

		checkSum := checkSum(data)
		//fmt.Println(checkSum)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)

		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		_, err := conn.Write(data)
		if err != nil {
			log.Println(err)
			return
		}
		sendCount++

		buf := make([]byte, 65535)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			failCount++
			return
		}
		successCount++
		ts := time.Since(start).Milliseconds()
		fmt.Printf("%d bytes from %d.%d.%d.%d: icmp_seq=%d ttl=%d time=%d ms\n",
			n-28, buf[12], buf[13], buf[14], buf[15], i, buf[8], ts)
		// fmt.Println(n, buf)
		time.Sleep(time.Second)
	}

	fmt.Printf("--- %s ping statistics ---\n", desIp)
	fmt.Printf("%d packets transmitted, %d packets received, %d%% packet loss\n", sendCount, successCount, failCount/sendCount)
	fmt.Printf("round-trip min/avg/max/stddev = 7.813/24.631/89.450/32.411 ms")

}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长（毫秒）")
	flag.IntVar(&size, "l", 32, "请求发生内容大小（字节）")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()
}

// 这里的逻辑这要是协议定的，我们不关心具体怎么写，知道即可
func checkSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32 = 0
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}

	if length != 0 {
		sum += uint32(data[index])
	}

	hi16 := sum >> 16

	for hi16 != 0 {
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >> 16
	}
	return uint16(^sum)
}
