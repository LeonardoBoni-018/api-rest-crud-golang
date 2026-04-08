# ✅ Checklist de Implementação - Semana a Semana

## 🎯 Como usar este checklist

- Marque com `[x]` conforme concluir cada item
- Commit ao final de cada sessão de trabalho
- Teste cada feature antes de marcar como completa
- Documente problemas encontrados

---

## 📅 SEMANA 1-2: Fundação Multi-tenant

### Setup Inicial
- [ ] Criar branch `scheduler-saas`
- [ ] Configurar nova estrutura de pastas
- [ ] Atualizar go.mod com novas dependências
- [ ] Criar .env.example e .env
- [ ] Configurar Docker Compose
- [ ] Testar conexão MongoDB
- [ ] Testar conexão Redis (opcional para fase 1)

### Modelo Tenant
- [ ] Criar `internal/domain/tenant/entity.go`
- [ ] Criar `internal/domain/tenant/repository.go`
- [ ] Implementar MongoRepository para Tenant
- [ ] Criar validações de negócio (IsActive, CanCreateService)
- [ ] Escrever testes unitários para Tenant entity

### Atualizar User Model
- [ ] Adicionar campo `tenant_id` em User
- [ ] Modificar UserRepository para incluir tenant
- [ ] Atualizar métodos de autenticação
- [ ] Criar roles: owner, employee, customer
- [ ] Testar isolamento de dados por tenant

### Casos de Uso - Tenant
- [ ] Implementar CreateTenant use case
- [ ] Implementar UpdateTenantSettings use case
- [ ] Implementar GetTenantBySlug use case
- [ ] Validar unicidade de slug
- [ ] Gerar slug automaticamente a partir do nome

### HTTP Layer - Tenant
- [ ] Criar TenantHandler
- [ ] Criar routes para tenant
- [ ] Implementar middleware de autenticação JWT
- [ ] Implementar middleware de isolamento por tenant
- [ ] Testar endpoints com Postman/Insomnia

### Endpoints Tenant
```bash
- [ ] POST /api/v1/tenants/register     # Cadastro
- [ ] GET  /api/v1/tenants/me           # Info da empresa
- [ ] PATCH /api/v1/tenants/settings    # Atualizar config
```

### Testes Semana 1-2
- [ ] Criar tenant via API
- [ ] Login com usuário owner
- [ ] Atualizar configurações do tenant
- [ ] Validar que tenants não vejam dados uns dos outros

---

## 📅 SEMANA 3-4: Sistema de Serviços e Agendamentos

### Modelo Service
- [ ] Criar `internal/domain/service/entity.go`
- [ ] Criar `internal/domain/service/repository.go`
- [ ] Implementar CRUD de serviços
- [ ] Adicionar validações (preço, duração, etc)
- [ ] Permitir upload de imagem do serviço (opcional)

### Modelo Booking
- [ ] Criar `internal/domain/booking/entity.go`
- [ ] Criar `internal/domain/booking/repository.go`
- [ ] Implementar FindByDateRange
- [ ] Implementar FindByStatus
- [ ] Criar índices no MongoDB para performance

### Lógica de Disponibilidade
- [ ] Criar AvailabilityService
- [ ] Implementar geração de time slots
- [ ] Considerar working_days do tenant
- [ ] Considerar working_hours do tenant
- [ ] Respeitar booking_interval
- [ ] Marcar slots ocupados
- [ ] Validar max_advance_days

### Casos de Uso - Booking
- [ ] Implementar GetAvailableSlots use case
- [ ] Implementar CreateBooking use case
- [ ] Implementar CancelBooking use case
- [ ] Implementar ConfirmBooking use case
- [ ] Implementar CompleteBooking use case
- [ ] Adicionar validação de conflitos de horário

### HTTP Layer - Services
- [ ] Criar ServiceHandler
- [ ] Proteger rotas (apenas owner pode criar)
- [ ] Implementar CRUD completo

### Endpoints Services
```bash
- [ ] POST   /api/v1/services            # Criar
- [ ] GET    /api/v1/services            # Listar
- [ ] GET    /api/v1/services/:id        # Ver
- [ ] PATCH  /api/v1/services/:id        # Editar
- [ ] DELETE /api/v1/services/:id        # Deletar
```

### HTTP Layer - Bookings Públicas
- [ ] Criar BookingHandler
- [ ] Implementar busca por slug (público)
- [ ] Permitir agendamento sem login
- [ ] Validar dados do cliente

### Endpoints Bookings Públicas
```bash
- [ ] GET  /api/v1/booking/:slug/info
- [ ] GET  /api/v1/booking/:slug/services
- [ ] GET  /api/v1/booking/:slug/availability
- [ ] POST /api/v1/booking/:slug/appointments
```

### HTTP Layer - Gestão de Bookings
- [ ] Proteger com autenticação
- [ ] Filtrar por tenant_id
- [ ] Permitir apenas owner/employee gerenciar

### Endpoints Gestão
```bash
- [ ] GET   /api/v1/appointments
- [ ] GET   /api/v1/appointments/:id
- [ ] PATCH /api/v1/appointments/:id/confirm
- [ ] PATCH /api/v1/appointments/:id/cancel
- [ ] PATCH /api/v1/appointments/:id/complete
```

### Testes Semana 3-4
- [ ] Criar serviço como owner
- [ ] Listar serviços via slug público
- [ ] Verificar disponibilidade de um dia
- [ ] Criar agendamento como cliente (sem login)
- [ ] Confirmar agendamento como owner
- [ ] Cancelar agendamento
- [ ] Validar que horários ocupados não aparecem
- [ ] Testar edge cases (horários inválidos, datas passadas)

---

## 📅 SEMANA 5: Dashboard e Métricas

### Modelos de Métricas
- [ ] Criar DashboardMetrics struct
- [ ] Criar ServiceMetric struct
- [ ] Criar RevenueReport struct

### Casos de Uso - Dashboard
- [ ] Implementar GetDashboardMetrics use case
- [ ] Calcular total de agendamentos
- [ ] Calcular taxa de confirmação
- [ ] Calcular taxa de cancelamento
- [ ] Calcular faturamento total
- [ ] Identificar top serviços
- [ ] Listar próximos agendamentos

### Visualização de Calendário
- [ ] Criar endpoint de calendar view
- [ ] Agrupar bookings por dia
- [ ] Retornar status de cada slot
- [ ] Permitir navegação mês a mês

### HTTP Layer - Dashboard
- [ ] Criar DashboardHandler
- [ ] Proteger com autenticação
- [ ] Filtrar por período (hoje, semana, mês)

### Endpoints Dashboard
```bash
- [ ] GET /api/v1/dashboard/metrics
- [ ] GET /api/v1/dashboard/calendar?month=2024-03
- [ ] GET /api/v1/dashboard/revenue
```

### Relatórios Básicos
- [ ] Exportar agendamentos como CSV
- [ ] Relatório de faturamento por período
- [ ] Relatório de serviços mais vendidos

### Testes Semana 5
- [ ] Acessar dashboard e ver métricas corretas
- [ ] Visualizar calendário do mês
- [ ] Exportar relatório CSV
- [ ] Validar cálculos de faturamento

---

## 📅 SEMANA 6-7: Chat com IA

### Setup OpenAI
- [ ] Criar conta OpenAI
- [ ] Configurar API key
- [ ] Criar OpenAIClient em infrastructure
- [ ] Testar chamada básica à API

### Modelo Conversation
- [ ] Criar `internal/domain/chat/entity.go`
- [ ] Criar Conversation com histórico de mensagens
- [ ] Criar Message struct
- [ ] Implementar repository para chat

### System Prompt Dinâmico
- [ ] Criar builder de system prompt
- [ ] Incluir informações do tenant
- [ ] Incluir lista de serviços e preços
- [ ] Incluir horários de funcionamento
- [ ] Definir regras de comportamento da IA

### Function Calling
- [ ] Implementar function `check_availability`
- [ ] Implementar function `create_booking`
- [ ] Implementar function `get_services`
- [ ] Parser de function calls do OpenAI

### Casos de Uso - Chat
- [ ] Implementar SendMessage use case
- [ ] Armazenar histórico de conversa
- [ ] Processar resposta da IA
- [ ] Executar funções quando necessário
- [ ] Retornar resposta formatada

### HTTP Layer - Chat
- [ ] Criar ChatHandler
- [ ] Rota pública por slug
- [ ] Gerenciar sessões de conversa
- [ ] Rate limiting para evitar abuso

### Endpoints Chat
```bash
- [ ] POST /api/v1/chat/:slug/conversations
- [ ] POST /api/v1/chat/:slug/messages
- [ ] GET  /api/v1/chat/:slug/conversations/:id
```

### WebSocket (opcional)
- [ ] Implementar WebSocket handler
- [ ] Stream de respostas da IA
- [ ] Notificação em tempo real

### Testes Semana 6-7
- [ ] Iniciar conversa com IA
- [ ] Perguntar "Quais serviços vocês têm?"
- [ ] Perguntar "Tem horário hoje às 14h?"
- [ ] Solicitar agendamento via chat
- [ ] Validar que IA usa function calling corretamente
- [ ] Testar múltiplas conversas simultâneas

---

## 📅 SEMANA 8: Autenticação Pública e Onboarding

### Agendamento sem Cadastro
- [ ] Permitir POST sem JWT token
- [ ] Validar apenas nome e telefone
- [ ] Criar token único para cancelamento
- [ ] Enviar link de cancelamento por email/SMS

### Link Único de Cancelamento
- [ ] Gerar token com UUID
- [ ] Criar rota pública de cancelamento
- [ ] Validar token antes de cancelar
- [ ] Expirar token após uso

### Onboarding de Empresas
- [ ] Criar fluxo multi-step
- [ ] Step 1: Dados básicos (nome, email, senha)
- [ ] Step 2: Tipo de negócio
- [ ] Step 3: Configurações de horário
- [ ] Step 4: Primeiro serviço
- [ ] Step 5: URL pública gerada

### Endpoints Onboarding
```bash
- [ ] POST /api/v1/onboarding/start
- [ ] POST /api/v1/onboarding/business-info
- [ ] POST /api/v1/onboarding/settings
- [ ] POST /api/v1/onboarding/first-service
- [ ] POST /api/v1/onboarding/complete
```

### Landing Page Pública
- [ ] Criar endpoint de info pública
- [ ] Retornar nome, tipo de negócio, foto
- [ ] Listar serviços ativos
- [ ] Mostrar avaliações (futuro)

### Testes Semana 8
- [ ] Completar onboarding de nova empresa
- [ ] Verificar URL pública funcionando
- [ ] Agendar sem login
- [ ] Receber link de cancelamento
- [ ] Cancelar via link único

---

## 📅 SEMANA 9: Notificações e Lembretes

### Modelo de Notificação
- [ ] Criar Notification entity
- [ ] Definir tipos: email, sms, whatsapp
- [ ] Criar templates de mensagens

### Integração SendGrid (Email)
- [ ] Criar conta SendGrid
- [ ] Configurar API key
- [ ] Criar EmailService
- [ ] Implementar templates HTML

### Templates de Email
- [ ] Template: Confirmação de agendamento
- [ ] Template: Lembrete 24h antes
- [ ] Template: Lembrete 1h antes
- [ ] Template: Cancelamento
- [ ] Template: Reagendamento

### Integração Twilio (SMS) - opcional
- [ ] Criar conta Twilio
- [ ] Configurar credentials
- [ ] Criar SMSService
- [ ] Implementar envio de SMS

### Sistema de Jobs (Background)
- [ ] Configurar Redis como queue
- [ ] Criar worker para processar jobs
- [ ] Job: Enviar lembrete 24h antes
- [ ] Job: Enviar lembrete 1h antes
- [ ] Scheduler para rodar jobs periodicamente

### Casos de Uso - Notificações
- [ ] SendBookingConfirmation
- [ ] SendReminder24h
- [ ] SendReminder1h
- [ ] SendCancellationNotice

### Testes Semana 9
- [ ] Criar agendamento e receber email de confirmação
- [ ] Simular job de lembrete 24h
- [ ] Simular job de lembrete 1h
- [ ] Cancelar e receber notificação
- [ ] Testar rate limiting de emails

---

## 📅 SEMANA 10-11: Planos e Pagamentos

### Modelo de Planos
- [ ] Criar Plan entity
- [ ] Definir limites por plano
- [ ] Free: 50 bookings/mês, 1 serviço
- [ ] Basic: 200 bookings/mês, ilimitado
- [ ] Pro: ilimitado + features extras

### Middleware de Validação de Limites
- [ ] Verificar bookings do mês
- [ ] Bloquear se excedeu limite
- [ ] Retornar mensagem de upgrade

### Integração Stripe
- [ ] Criar conta Stripe
- [ ] Configurar API keys
- [ ] Criar StripeService
- [ ] Implementar checkout session
- [ ] Implementar webhooks

### Webhooks Stripe
- [ ] Webhook: payment_intent.succeeded
- [ ] Webhook: customer.subscription.created
- [ ] Webhook: customer.subscription.updated
- [ ] Webhook: customer.subscription.deleted
- [ ] Validar assinatura do webhook

### Upgrade/Downgrade
- [ ] Endpoint para upgrade de plano
- [ ] Endpoint para downgrade de plano
- [ ] Cálculo pro-rata
- [ ] Confirmação por email

### Endpoints Pagamentos
```bash
- [ ] GET  /api/v1/billing/plans
- [ ] POST /api/v1/billing/checkout
- [ ] POST /api/v1/billing/upgrade
- [ ] POST /api/v1/billing/cancel
- [ ] POST /api/v1/webhooks/stripe
```

### Testes Semana 10-11
- [ ] Listar planos disponíveis
- [ ] Criar checkout session
- [ ] Simular pagamento bem-sucedido
- [ ] Verificar tenant mudou de plano
- [ ] Testar limite de bookings
- [ ] Cancelar assinatura

---

## 📅 SEMANA 12+: Features Avançadas

### Múltiplos Profissionais
- [ ] Criar Employee entity
- [ ] Associar serviços a profissionais
- [ ] Agenda separada por profissional
- [ ] Cliente escolhe profissional (opcional)

### Analytics Avançado
- [ ] Taxa de conversão (visitas → agendamentos)
- [ ] Horários de pico
- [ ] Serviços mais rentáveis
- [ ] Clientes recorrentes
- [ ] NPS (Net Promoter Score)

### Integração Google Calendar
- [ ] OAuth com Google
- [ ] Sincronizar agendamentos
- [ ] Bloquear horários do Google no sistema
- [ ] Criar evento no Google ao confirmar

### App Mobile (React Native) - opcional
- [ ] Setup React Native
- [ ] Telas: Login, Dashboard, Agendamentos
- [ ] Push notifications
- [ ] Deploy na Play Store / App Store

### Personalização
- [ ] Upload de logo
- [ ] Escolher cores do tema
- [ ] Personalizar emails
- [ ] Custom domain (myshop.com)

### Sistema de Avaliações
- [ ] Criar Review entity
- [ ] Permitir cliente avaliar após serviço
- [ ] Exibir média de avaliações
- [ ] Moderar reviews

### Pagamento Online Antecipado
- [ ] Integrar Stripe Checkout
- [ ] Permitir pagamento ao agendar
- [ ] Política de cancelamento com estorno
- [ ] Comprovante de pagamento

---

## 🧪 TESTES GERAIS

### Testes Unitários
- [ ] Testes para entities
- [ ] Testes para repositories
- [ ] Testes para use cases
- [ ] Coverage mínimo de 70%

### Testes de Integração
- [ ] Testes de endpoints
- [ ] Testes com banco de dados
- [ ] Testes de autenticação
- [ ] Testes de autorização

### Testes E2E
- [ ] Fluxo completo: Cadastro → Agendamento → Confirmação
- [ ] Fluxo de onboarding
- [ ] Fluxo de pagamento
- [ ] Chat com IA

---

## 🚀 DEPLOY

### Preparação para Produção
- [ ] Criar Dockerfile otimizado
- [ ] Multi-stage build
- [ ] Configurar health checks
- [ ] Configurar graceful shutdown

### CI/CD
- [ ] GitHub Actions para testes
- [ ] GitHub Actions para deploy
- [ ] Rodar testes em cada PR
- [ ] Deploy automático na main

### Escolher Hosting
- [ ] Railway (recomendado)
- [ ] DigitalOcean
- [ ] AWS
- [ ] Google Cloud

### Database em Produção
- [ ] MongoDB Atlas
- [ ] Backup automático
- [ ] Réplicas
- [ ] Monitoramento

### Monitoramento
- [ ] Sentry para error tracking
- [ ] Logs centralizados
- [ ] Métricas de performance
- [ ] Alertas

### Segurança
- [ ] Rate limiting
- [ ] CORS configurado
- [ ] Headers de segurança
- [ ] SSL/TLS
- [ ] Sanitização de inputs
- [ ] Proteção contra SQL injection
- [ ] Proteção contra XSS

---

## 📚 DOCUMENTAÇÃO

- [ ] README atualizado
- [ ] Swagger/OpenAPI
- [ ] Guia de contribuição
- [ ] Guia de deploy
- [ ] Postman collection
- [ ] Diagramas de arquitetura

---

## 🎯 MVP MÍNIMO (3-4 semanas)

Se quiser lançar rápido, foque nestes itens essenciais:

- [x] Multi-tenant básico
- [x] CRUD de serviços
- [x] Agendamento público
- [x] Gestão de agendamentos
- [x] Dashboard simples
- [x] Notificação por email
- [ ] Deploy em produção

**Com isso você já pode validar a ideia com clientes reais!**

---

## 💡 DICAS

1. **Commit frequentemente**: A cada feature completa
2. **Teste antes de marcar**: Não marque como feito se não testou
3. **Documente problemas**: Anote bugs e ideias para depois
4. **Peça feedback cedo**: Mostre para potenciais clientes logo
5. **Não persiga perfeição**: MVP funcional > código perfeito
6. **Priorize value**: Faça primeiro o que traz mais valor
7. **Automatize**: Testes automatizados salvam tempo

---

**Boa sorte!** 🚀

Dúvidas? Adicione comentários aqui ou abra issues no repo.
