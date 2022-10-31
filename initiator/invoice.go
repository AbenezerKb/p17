package initiator

import (
	"context"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	persistiance_invoice "sms-gateway/internal/adapter/storage/persistance/invoice"
	const_init "sms-gateway/internal/constant/init"
	module_invoice "sms-gateway/internal/module/invoice"
)

func InvoiceDomainInit(common const_init.Utils) {

	invoiceStorage := persistiance_invoice.StorageInit(common)
	invoiceModule := module_invoice.InvoiceModule(common, invoiceStorage)

	_cron := cron.New()

	cronId, err := _cron.AddFunc("@every 15m", func() {
		err := invoiceModule.GenerateInvoice(context.Background())
		if err != nil {
			log.Panic(err)
		}
	})
	if err != nil {
		log.Panic(err)
	}
	log.Info(cronId)
}
