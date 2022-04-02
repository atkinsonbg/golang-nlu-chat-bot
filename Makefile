run:
	go run main.go

train-order:
	go run main.go --train -test="Hello, I would like to order a pizza with capers, mushrooms, and onions on a think crust, medium size."

train-hours:
	go run main.go --train -test="What time do you open today?"

train-unknown:
	go run main.go --train -test="I really think I need to go to the doctor tomorrow!"
