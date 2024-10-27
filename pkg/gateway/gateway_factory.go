package gateway

type PaymentGatewayFactory struct{}

func NewPaymentGatewayFactory() *PaymentGatewayFactory {
	return &PaymentGatewayFactory{}
}

func (pgf *PaymentGatewayFactory) GetPaymentGateway(providerHandle string) PaymentGateway {
	switch providerHandle {
	case "gatewayA":
		return &GatewayA{}
	case "gatewayB":
		return &GatewayB{}
	default:
		return nil
	}
}
