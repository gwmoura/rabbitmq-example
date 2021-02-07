module g.com

go 1.12

replace g.com/queue => ./queue

replace g.com/utils => ./utils

replace g.com/mailer => ./mailer

replace g.com/logger => ./logger

require (
	g.com/logger v0.0.0-00010101000000-000000000000
	g.com/mailer v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.0.0
)
