module senyas.tour.com/tour/cmd

go 1.18

replace senyas.tour.com/tour/internal/word => ../internal/word

replace senyas.tour.com/tour/internal/timer => ../internal/timer

require (
	github.com/spf13/cobra v1.5.0
	senyas.tour.com/tour/internal/timer v0.0.0-00010101000000-000000000000
	senyas.tour.com/tour/internal/word v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.3.7 // indirect
)
