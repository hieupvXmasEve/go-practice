package main

import (
	"fmt"
)

type Bidder interface {
	Bid(amount float64)
	GetBidAmount() float64
}

type IndividualBidder struct {
	Name      string
	BidAmount float64
}

type CompanyBidder struct {
	CompanyName string
	BidAmount   float64
}

// Phương thức String() cho IndividualBidder
func (i *IndividualBidder) String() string {
	return i.Name
}
func (i *IndividualBidder) Bid(amount float64) {
	i.BidAmount = amount
}

func (i *IndividualBidder) GetBidAmount() float64 {
	return i.BidAmount
}

func (c *CompanyBidder) Bid(amount float64) {
	c.BidAmount = amount
}

func (c *CompanyBidder) GetBidAmount() float64 {
	return c.BidAmount
}

func main() {
	bidders := []Bidder{
		&IndividualBidder{Name: "Alice"},
		&IndividualBidder{Name: "Bob"},
		&CompanyBidder{CompanyName: "TechCorp"},
	}
	fmt.Printf("123 %v\n", bidders) // In ra các giá trị của các phần tử trong slice

	// Đặt giá thầu
	bidders[0].Bid(100.0)
	bidders[1].Bid(150.0)
	bidders[2].Bid(200.0)

	// Tìm người thắng thầu
	var highestBidder Bidder
	var highestBid float64

	for _, bidder := range bidders {
		if bidder.GetBidAmount() > highestBid {
			highestBid = bidder.GetBidAmount()
			highestBidder = bidder
		}
	}

	// In kết quả
	fmt.Println("Người thắng thầu:")
	switch b := highestBidder.(type) {
	case *IndividualBidder:
		fmt.Printf("Cá nhân: %s với giá %.2f\n", b.Name, b.GetBidAmount())
	case *CompanyBidder:
		fmt.Printf("Công ty: %s với giá %.2f\n", b.CompanyName, b.GetBidAmount())
	}
}
