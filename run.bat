go build -o bookings.exe ./cmd/web/.
bookings.exe -dburl=postgresql://root:secret@localhost:5433/bookings?sslmode=disable