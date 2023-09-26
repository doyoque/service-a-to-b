module github.com/doyoque/service_a

go 1.21.1

replace github.com/doyoque/service_util => /Users/doyoque/Documents/personal-work/projects/go-silly/service_util

require (
	github.com/doyoque/service_util v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.5
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/fx v1.20.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
