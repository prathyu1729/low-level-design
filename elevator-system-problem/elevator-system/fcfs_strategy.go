package elevator_system

type fcfsStrategy struct {
	requests []request
}

func (s *fcfsStrategy) processRequest(request request) {
	s.requests = append(s.requests, request)
}

func (s *fcfsStrategy) start() {
	for {
		if len(s.requests) != 0 {
			s.process()
		}
	}
}

func (s *fcfsStrategy) process() {
	request := s.requests[0]
	requests := s.requests
	s.requests = requests[1:]

	request.elevatorCar.move(request.floorNumber)
	//TODO: ask elevator car to go to particular floor

}
