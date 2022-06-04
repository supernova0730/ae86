package enums

type OrderState string

const (
	OrderStatePending            OrderState = "PENDING"
	OrderStateCanceled           OrderState = "CANCELED"
	OrderStateAccepted           OrderState = "ACCEPTED"
	OrderStateDeliveryInProgress OrderState = "DELIVERY_IN_PROGRESS"
	OrderStateDelivered          OrderState = "DELIVERED"
)
