package sherrytime
/*
   產生UUID
*/

import (
   "net"
   "bytes"
   "math/rand"
   "encoding/binary"
   "encoding/hex"
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

// NewUUID create Universally unique identifier
func (st *SherryTime) NewUUID() string {
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
