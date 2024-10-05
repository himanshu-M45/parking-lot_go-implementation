package receipt

import (
	"fmt"
	"math/rand"
	"time"
)

type Receipt struct {
	receiptId string
}

func Construct() *Receipt {
	ticketId := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000))
	return &Receipt{receiptId: ticketId}
}

func (ticket *Receipt) ValidateTicket(receivedTicket Receipt) bool {
	return ticket.receiptId == receivedTicket.receiptId
}
