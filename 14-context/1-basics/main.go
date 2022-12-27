package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel Booking Canceled! Timeout Reached!")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel Booked Successfully")
	}
}
