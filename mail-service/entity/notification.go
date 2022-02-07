package entity

type Notification interface {
	ToNotificate(mail Mail) string
}
