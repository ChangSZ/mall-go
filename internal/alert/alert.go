package alert

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/pkg/errors"

	"github.com/ChangSZ/golib/log"
	"github.com/ChangSZ/golib/mail"
	"github.com/spf13/cast"
)

// NotifyHandler 告警通知
func NotifyHandler() func(msg *proposal.AlertMessage) {
	return func(msg *proposal.AlertMessage) {
		cfg := configs.Get().Mail
		if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
			log.Error("Mail config error")
			return
		}

		subject, body, err := newHTMLEmail(
			msg.Method,
			msg.HOST,
			msg.URI,
			msg.TraceID,
			msg.ErrorMessage,
			msg.ErrorStack,
		)
		if err != nil {
			log.Error("email template error: ", err)
			return
		}

		m, err := mail.Init(mail.WithHost(cfg.Host),
			mail.WithPort(cast.ToInt(cfg.Port)),
			mail.WithUser(cfg.User),
			mail.WithPwd(cfg.Pass),
		)
		if err != nil {
			log.Error(err)
			return
		}
		m.SetTo([]string{cfg.To})
		m.SetSubject(subject)
		m.SetBody(body)
		if err := m.Send(); err != nil {
			log.Error("发送告警通知邮件失败: ", errors.WithStack(err))
			return
		}
		log.Infof("告警通知邮件发送成功, To: %v", cfg.To)
	}
}
