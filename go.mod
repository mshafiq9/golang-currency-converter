module currency-converter

go 1.17

require (
	currency-converter/database v0.0.0-00010101000000-000000000000 // indirect
	currency-converter/tradermade v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
)

replace currency-converter/database => ./database

replace currency-converter/tradermade => ./tradermade
