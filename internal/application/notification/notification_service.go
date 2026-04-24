package notification

import (
	"fmt"
	"time"

	bookingdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	tenantdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	servicedomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/email"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

type NotificationService interface {
	SendBookingConfirmation(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, service *servicedomain.Service) *rest_err.RestErr
	SendBookingReminder(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, service *servicedomain.Service, reminderType string) *rest_err.RestErr
	SendBookingCancellation(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, service *servicedomain.Service) *rest_err.RestErr
}

type notificationService struct {
	emailProvider email.EmailProvider
}

func NewNotificationService(provider email.EmailProvider) NotificationService {
	if provider == nil {
		provider = email.GetEmailProvider()
	}
	return &notificationService{
		emailProvider: provider,
	}
}

func (s *notificationService) SendBookingConfirmation(
	booking *bookingdomain.Booking,
	tenant *tenantdomain.Tenant,
	serviceItem *servicedomain.Service,
) *rest_err.RestErr {
	subject := fmt.Sprintf("Agendamento Confirmado - %s", tenant.Name)
	htmlBody := buildConfirmationEmail(booking, tenant, serviceItem)

	err := s.emailProvider.SendEmail(
		booking.CustomerEmail,
		booking.CustomerName,
		subject,
		htmlBody,
	)

	if err != nil {
		fmt.Printf("[Notification] Failed to send confirmation email: %v\n", err)
		return rest_err.NewInternalServerError("failed to send confirmation email")
	}

	return nil
}

func (s *notificationService) SendBookingReminder(
	booking *bookingdomain.Booking,
	tenant *tenantdomain.Tenant,
	serviceItem *servicedomain.Service,
	reminderType string,
) *rest_err.RestErr {
	var subject string
	var timeMsg string

	if reminderType == "24h" {
		subject = fmt.Sprintf("Lembrete: Agendamento Amanha - %s", tenant.Name)
		timeMsg = "amanha"
	} else {
		subject = fmt.Sprintf("Lembrete: Agendamento em 1 hora - %s", tenant.Name)
		timeMsg = "em 1 hora"
	}

	htmlBody := buildReminderEmail(booking, tenant, serviceItem, timeMsg)

	err := s.emailProvider.SendEmail(
		booking.CustomerEmail,
		booking.CustomerName,
		subject,
		htmlBody,
	)

	if err != nil {
		fmt.Printf("[Notification] Failed to send reminder email: %v\n", err)
		return rest_err.NewInternalServerError("failed to send reminder email")
	}

	return nil
}

func (s *notificationService) SendBookingCancellation(
	booking *bookingdomain.Booking,
	tenant *tenantdomain.Tenant,
	serviceItem *servicedomain.Service,
) *rest_err.RestErr {
	subject := fmt.Sprintf("Agendamento Cancelado - %s", tenant.Name)
	htmlBody := buildCancellationEmail(booking, tenant, serviceItem)

	err := s.emailProvider.SendEmail(
		booking.CustomerEmail,
		booking.CustomerName,
		subject,
		htmlBody,
	)

	if err != nil {
		fmt.Printf("[Notification] Failed to send cancellation email: %v\n", err)
		return rest_err.NewInternalServerError("failed to send cancellation email")
	}

	return nil
}

func buildConfirmationEmail(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, svc *servicedomain.Service) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Agendamento Confirmado</title>
</head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; background-color: #f5f5f5;">
    <div style="background-color: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
        <div style="text-align: center; margin-bottom: 30px;">
            <h1 style="color: #22c55e; margin: 0;">Agendamento Confirmado!</h1>
        </div>
        
        <p style="color: #333; font-size: 16px;">Ola <strong>%s</strong>!</p>
        
        <p style="color: #666; font-size: 14px;">Seu agendamento foi confirmado com sucesso. Aqui estao os detalhes:</p>
        
        <div style="background-color: #f9fafb; padding: 20px; border-radius: 8px; margin: 20px 0;">
            <table style="width: 100%%; border-collapse: collapse;">
                <tr>
                    <td style="padding: 10px 0; color: #666; font-size: 14px;">Servico:</td>
                    <td style="padding: 10px 0; color: #333; font-size: 14px; font-weight: bold;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 10px 0; color: #666; font-size: 14px;">Data:</td>
                    <td style="padding: 10px 0; color: #333; font-size: 14px; font-weight: bold;">%s</td>
                </tr>
                <tr>
                    <td style="padding: 10px 0; color: #666; font-size: 14px;">Horario:</td>
                    <td style="padding: 10px 0; color: #333; font-size: 14px; font-weight: bold;">%s</td>
                </tr>
            </table>
        </div>
        
        <p style="color: #666; font-size: 14px;">
            <strong>Estabelecimento:</strong> %s
        </p>
    </div>
</body>
</html>
	`,
		booking.CustomerName,
		svc.Name,
		formatDate(booking.Date),
		booking.TimeSlot,
		tenant.Name,
	)
}

func buildReminderEmail(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, svc *servicedomain.Service, timeMsg string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Lembrete de Agendamento</title>
</head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background-color: white; padding: 30px; border-radius: 10px;">
        <h1 style="color: #f59e0b;">Lembrete</h1>
        <p>Ola <strong>%s</strong>!</p>
        <p>Voce tem um agendamento <strong>%s</strong>.</p>
        <div style="background-color: #fef3c7; padding: 15px; border-radius: 8px;">
            <p><strong>Servico:</strong> %s</p>
            <p><strong>Data/Hora:</strong> %s as %s</p>
        </div>
    </div>
</body>
</html>
	`,
		booking.CustomerName,
		timeMsg,
		svc.Name,
		formatDate(booking.Date),
		booking.TimeSlot,
	)
}

func buildCancellationEmail(booking *bookingdomain.Booking, tenant *tenantdomain.Tenant, svc *servicedomain.Service) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Agendamento Cancelado</title>
</head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
    <div style="background-color: white; padding: 30px; border-radius: 10px;">
        <h1 style="color: #ef4444;">Agendamento Cancelado</h1>
        <p>Ola <strong>%s</strong>!</p>
        <p>Seu agendamento foicancelado.</p>
        <div style="background-color: #fee2e2; padding: 15px; border-radius: 8px;">
            <p><strong>Servico:</strong> %s</p>
            <p><strong>Data/Hora:</strong> %s as %s</p>
        </div>
    </div>
</body>
</html>
	`,
		booking.CustomerName,
		svc.Name,
		formatDate(booking.Date),
		booking.TimeSlot,
	)
}

func formatDate(dateStr string) string {
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	months := []string{
		"Janeiro", "Fevereiro", "Marco", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}
	return fmt.Sprintf("%d de %s de %d", parsed.Day(), months[parsed.Month()-1], parsed.Year())
}