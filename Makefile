## TRAIN TASKS
train-order:
	go run main.go --train -sentence="Hello, I would like to order a pizza with capers, mushrooms, and onions on a think crust, medium size."

train-hours:
	go run main.go --train -sentence="What time do you open today?"

train-unknown:
	go run main.go --train -sentence="I really think I need to go to the doctor tomorrow!"

## CLASSIFY TASKS
classify-order:
	go run main.go --classify -sentence="Hello, I would like to order a pizza with capers, mushrooms, and onions on a think crust, medium size."

classify-hours:
	go run main.go --classify -sentence="What time do you open today?"

classify-unknown:
	go run main.go --classify -sentence="I really think I need to go to the doctor tomorrow!"
