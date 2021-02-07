start-console:
	go run . ./main.go console
start-writer:
	go run . ./main.go writer
web-mailer:
	go run . mailer web
start-mailer-web:
	cd mailer && go run ./web/web.go
start-mailer-consumer:
	cd mailer && go run ./consumer/consumer.go
emit-log:
	go run ./queue.go ./utils.go ./logs.go ./emit_log_topic.go log.*
