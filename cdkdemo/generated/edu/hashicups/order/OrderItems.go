package order


type OrderItems struct {
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs/resources/order#coffee Order#coffee}.
	Coffee *OrderItemsCoffee `field:"required" json:"coffee" yaml:"coffee"`
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs/resources/order#quantity Order#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

