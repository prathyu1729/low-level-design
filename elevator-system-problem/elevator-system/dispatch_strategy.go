package elevator_system

type dispatchStrategy interface {
	processRequest(request request)
	start()
}
