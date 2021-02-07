module g.com/mailer

go 1.12

replace g.com/mailer => ./

replace g.com/utils => ../utils

replace g.com/queue => ../queue

replace g.com/logger => ../logger

require (
	g.com/logger v0.0.0-00010101000000-000000000000 // indirect
	g.com/queue v0.0.0-00010101000000-000000000000
	g.com/utils v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.0.0 // indirect
)
