package blnk2

import (
	"github.com/ottojo/blnk2/client"
	"github.com/ottojo/blnk2/protocol"
	"github.com/ottojo/blnk2/storage"
	"log"
	"net"
	"strings"
)

type BlnkSystem struct {
	Leds    client.LedList
	Stage   client.LedList
	Clients []*client.Client
}

func CreateFromFile(filename string) BlnkSystem {
	var b BlnkSystem
	var err error
	var leds [][]client.Led
	var cs []client.Client
	cs, leds, err = storage.ReadClientsFromFile(filename)
	for i := 0; i < len(cs); i++ {
		b.Clients = append(b.Clients, &cs[i])
	}
	LogError(err)

	for clientId := range b.Clients {
		b.Clients[clientId].FirstLed = b.Leds.Add(leds[clientId][0])
		b.Clients[clientId].FirstStageLed = b.Stage.Add(leds[clientId][0])
		for i := 1; i < len(leds[clientId]); i++ {
			b.Leds.Add(leds[clientId][i])
			b.Stage.Add(leds[clientId][i])
		}
		b.Clients[clientId].Opt = client.NONE
	}
	return b
}

func (b *BlnkSystem) Commit() {
	for i := 0; i < len(b.Clients); i++ {
		c := b.Clients[i]

		if c.Socket == nil {
			var err error
			c.Socket, err = net.Dial("tcp", c.Address)
			LogError(err)
			continue
		}

		switch c.Opt {
		case client.NONE:

			// Copy entire stage to leds
			sled := c.FirstStageLed
			led := c.FirstLed
			for i := 0; i < c.Length; i++ {
				led.Data = sled.Data
				sled = sled.Next
				led = led.Next
			}
			if c.Socket != nil {
				log.Println("Sending data to " + c.Id)
				c.Socket.Write(protocol.SetPixelsMessage(c.FirstLed, c.Length))
			}
			break
		case client.PREVENTDUPLICATE:
			// Copy entire stage to leds
			sled := c.FirstStageLed
			led := c.FirstLed
			changesPresent := false
			for i := 0; i < c.Length; i++ {
				if !led.Data.Color.AlmostEqualRgb(sled.Data.Color) { //TODO: Compare colors better
					changesPresent = true
				}
				led.Data = sled.Data
				sled = sled.Next
				led = led.Next
			}
			// Only send if data has changed
			if changesPresent {
				if c.Socket != nil {
					c.Socket.Write(protocol.SetPixelsMessage(c.FirstLed, c.Length))
				}
			}
			break
		case client.PACKETSIZE:

			sled := c.FirstStageLed
			led := c.FirstLed
			firstChangedLed := c.Length - 1
			lastChangedLed := 0
			for i := 0; i < c.Length; i++ {
				if !led.Data.Color.AlmostEqualRgb(sled.Data.Color) { //TODO: Compare colors better
					if i < firstChangedLed {
						firstChangedLed = i
					}
					if i > lastChangedLed {
						lastChangedLed = i
					}
				}
				led.Data = sled.Data
				sled = sled.Next
				led = led.Next
			}
			// Only send smallest interval containing all changes
			if firstChangedLed == c.Length-1 && lastChangedLed == 0 {
				if c.Socket != nil {
					c.Socket.Write(protocol.SetIntervalMessage(c.FirstLed.GetOffset(firstChangedLed),
						lastChangedLed-firstChangedLed+1,
						firstChangedLed))
				}
			}

			break
		}
	}
}

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (bl *BlnkSystem) Discovery() {
	addr, err := net.ResolveUDPAddr("udp", "239.1.3.37:1337")
	LogError(err)

	l, err := net.ListenMulticastUDP("udp", nil, addr)
	LogError(err)

	l.SetReadBuffer(1024)
	for {
		b := make([]byte, 1024)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		message := string(b[:n])
		log.Printf("Message from %s: %s\n", src, message)
		m := strings.Split(message, " ")
		for clientIndex := range bl.Clients {
			if m[0] == bl.Clients[clientIndex].Id {
				newAddress := src.IP.String() + ":" + m[1]
				if bl.Clients[clientIndex].Address != newAddress {
					log.Printf("%s now has ip %s\n", bl.Clients[clientIndex].Id, newAddress)
					bl.Clients[clientIndex].Address = newAddress
					bl.Clients[clientIndex].Socket = nil
					bl.Connect()
				}
			}
		}
	}
}

func (bl *BlnkSystem) Connect() {
	for _, c := range bl.Clients {
		c.Connect()
	}
}

func (bl *BlnkSystem) Disconnect() {
	for _, c := range bl.Clients {
		if c.Socket != nil {
			var err error
			err = c.Socket.Close()
			LogError(err)
		}
	}
}
