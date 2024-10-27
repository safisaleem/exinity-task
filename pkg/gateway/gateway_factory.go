package gateway

type PaymentGatewayFactory interface {
	GetPaymentGateway(providerHandle string) PaymentGateway
}

type PaymentGatewayFactoryImpl struct{}

func NewPaymentGatewayFactory() PaymentGatewayFactory {
	return &PaymentGatewayFactoryImpl{}
}

func (pgf *PaymentGatewayFactoryImpl) GetPaymentGateway(providerHandle string) PaymentGateway {
	switch providerHandle {
	case "gatewayA":
		return &GatewayA{}
	case "gatewayB":
		return &GatewayB{}
	default:
		return nil
	}
}
