package notifiable

type Notifiable interface {
	Notify(parkingLotId string, status bool)
}
