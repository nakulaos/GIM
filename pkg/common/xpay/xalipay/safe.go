package xalipay

import (
	"GIM/pkg/conf"
	"github.com/smartwalle/alipay/v3"
	"log/slog"
)

func NewSafeAlipay(cfg *conf.Alipay) *Alipay {
	var (
		client *alipay.Client
		err    error
	)
	if client, err = alipay.New(cfg.Appid, cfg.AppPrivateKey, cfg.Release); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAppCertPublicKey(cfg.AppCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAliPayRootCert(cfg.AlipayRootCert); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	if err = client.LoadAlipayCertPublicKey(cfg.AlipayCertPublicKey); err != nil {
		slog.Warn(err.Error())
		return nil
	}
	//if err = client.SetEncryptKey(cfg.EncryptKey); err != nil {
	//	slog.Warn(err.Error())
	//	return nil
	//}
	pay = &Alipay{
		client: client,
		cfg:    cfg,
	}
	return pay
}
