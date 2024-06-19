package sherrytime
/*
   產生UUID
*/

import (
   "net"
   "time"
   "bytes"
   "strings"
   "math/rand"
   "crypto/rand"
   "encoding/hex"
   "encoding/binary"
   "github.com/google/uuid"
)

func (st *SherryTime) clockSeq() uint16 {
   // 16383 is the max number of 14 bit
   return uint16(rand.Intn(16383))
}

func (st *SherryTime) getMacAddr() ([]byte, bool) {
   var addr []byte
   interfaces, err := net.Interfaces()
   if err == nil {
       for _, i := range interfaces {
           if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
               addr = i.HardwareAddr
               return addr, true
           }
       }
   }
   return addr, false
}

func (st *SherryTime) getNode() [6]byte {
   var nodeID [6]byte

   node, ok := st.getMacAddr()
   if !ok {
       return nodeID
   }

   copy(nodeID[:], node[:6])
   return nodeID
}

// UUID　Ｖ７
func(st *SherryTime) NewUUIDV7(version string)([16]byte, error) {
    var value [16]byte
    _, err := rand.Read(value[:])
    if err != nil {
        return value, err
    }
    timestamp := uint64(time.Now().UnixNano() / int64(time.Millisecond)) // timestamp

    value[0] = byte((timestamp >> 40) & 0xFF)
    value[1] = byte((timestamp >> 32) & 0xFF)
    value[2] = byte((timestamp >> 24) & 0xFF)
    value[3] = byte((timestamp >> 16) & 0xFF)
    value[4] = byte((timestamp >> 8) & 0xFF)
    value[5] = byte(timestamp & 0xFF)
    // version and variant
    value[6] = (value[6] & 0x0F) | 0x70
    value[8] = (value[8] & 0x3F) | 0x80

    return value, nil
}

// NewUUID create Universally unique identifier
func(st *SherryTime) NewUUIDByTime() string {
   var uuid [16]byte
   t := st.getTimeSince1582()
   cSeq := st.clockSeq()
   timeLow := uint32(t)
   timeMid := uint16((t >> 32))
   timeHi := uint16((t >> 48))
   timeHi += 0x1000

   node := st.getNode()

   binary.BigEndian.PutUint32(uuid[0:], timeLow)
   binary.BigEndian.PutUint16(uuid[4:], timeMid)
   binary.BigEndian.PutUint16(uuid[6:], timeHi)
   binary.BigEndian.PutUint16(uuid[6:], timeHi)
   binary.BigEndian.PutUint16(uuid[8:], cSeq)

   copy(uuid[10:], node[:6])  // uuid [16]byte

   dst := make([]byte, hex.EncodedLen(len(uuid)+3))

   hex.Encode(dst, uuid[0:4])
   dst[8] = '-'
   hex.Encode(dst[9:17], uuid[4:8])
   dst[17] = '-'
   hex.Encode(dst[18:26], uuid[8:12])
   dst[26] = '-'
   hex.Encode(dst[27:], uuid[12:])

   return string(dst[:])
}

func(st *SherryTime) NewUUID() string {
    uuidWithHyphen := uuid.New()
    return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}
