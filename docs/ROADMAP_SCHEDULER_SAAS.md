# 🚀 Roadmap - SaaS de Agendamento Inteligente

## 📋 Índice
1. [Visão Geral do Projeto](#visão-geral)
2. [Arquitetura Atual vs Nova](#arquitetura)
3. [Plano de Desenvolvimento por Fases](#fases)
4. [Stack Tecnológica](#stack)
5. [Database Schema](#database)
6. [APIs e Endpoints](#apis)
7. [Integrações com IA](#ia)
8. [Deploy e Infraestrutura](#deploy)

---

## 🎯 Visão Geral do Projeto

### O que você tem agora:
- ✅ CRUD de usuários
- ✅ Autenticação JWT
- ✅ API REST básica em Go
- ✅ Estrutura de projeto organizada

### O que vai construir:
**SaaS Multi-tenant de Agendamento com IA**

Um sistema onde:
- 🏢 **Empresas** criam contas e gerenciam suas agendas
- 👥 **Clientes finais** agendam serviços online
- 🤖 **IA** responde dúvidas e facilita agendamentos
- 💰 **Modelo de negócio**: Mensalidade (R$ 19-99/mês)

---

## 🏗️ Arquitetura Atual vs Nova

### Estrutura Atual (CRUD Simples)
```
api-rest-crud-golang/
├── configuration/
├── src/
│   ├── controller/
│   ├── model/
│   └── ...
├── main.go
└── init_dependencies.go
```

### Nova Estrutura (Multi-tenant SaaS)
```
api-rest-crud-golang/
├── cmd/
│   ├── api/           # API principal
│   ├── worker/        # Jobs em background
│   └── migrations/    # Migrações de DB
├── internal/
│   ├── domain/        # Entidades de negócio
│   │   ├── tenant/    # Empresas (multi-tenant)
│   │   ├── user/      # Usuários (clientes + donos)
│   │   ├── service/   # Serviços oferecidos
│   │   ├── booking/   # Agendamentos
│   │   └── chat/      # Mensagens com IA
│   ├── application/   # Casos de uso
│   ├── infrastructure/
│   │   ├── database/
│   │   ├── cache/     # Redis
│   │   ├── queue/     # Jobs
│   │   └── ai/        # OpenAI integration
│   └── interfaces/
│       ├── http/      # REST API
│       └── websocket/ # Chat em tempo real
├── pkg/               # Código reutilizável
│   ├── jwt/
│   ├── logger/
│   └── validator/
├── web/               # Frontend (opcional)
└── docker-compose.yml
```

---

## 📅 Plano de Desenvolvimento por Fases

### **FASE 1: Fundação Multi-tenant (Semanas 1-2)**

#### Objetivos:
- Transformar sistema single-user em multi-tenant
- Separar dados por empresa (tenant)
- Implementar onboarding de empresas

#### Tarefas:

**1.1 - Criar modelo de Tenant (Empresa)**
```go
// internal/domain/tenant/tenant.go
type Tenant struct {
    ID          string    `json:"id" bson:"_id"`
    Name        string    `json:"name" bson:"name"`
    Slug        string    `json:"slug" bson:"slug"` // barbearia-joao
    Email       string    `json:"email" bson:"email"`
    Phone       string    `json:"phone" bson:"phone"`
    Plan        string    `json:"plan" bson:"plan"` // free, basic, pro
    Status      string    `json:"status" bson:"status"` // active, suspended
    Settings    Settings  `json:"settings" bson:"settings"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type Settings struct {
    BusinessType    string   `json:"business_type"` // salon, barbershop, clinic
    WorkingDays     []string `json:"working_days"`  // ["mon", "tue", "wed"]
    WorkingHours    string   `json:"working_hours"` // "09:00-18:00"
    BookingInterval int      `json:"booking_interval"` // 30 minutos
    MaxAdvanceDays  int      `json:"max_advance_days"` // 30 dias
}
```

**1.2 - Modificar User model para incluir tenant**
```go
// internal/domain/user/user.go
type User struct {
    ID        string    `json:"id" bson:"_id"`
    TenantID  string    `json:"tenant_id" bson:"tenant_id"` // NOVO
    Email     string    `json:"email" bson:"email"`
    Password  string    `json:"-" bson:"password"`
    Role      string    `json:"role" bson:"role"` // owner, employee, customer
    Name      string    `json:"name" bson:"name"`
    Phone     string    `json:"phone" bson:"phone"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
```

**1.3 - Middleware de isolamento de tenant**
```go
// internal/interfaces/http/middleware/tenant.go
func TenantIsolation() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        
        // Buscar tenant do usuário
        user, err := userRepo.FindByID(userID)
        if err != nil {
            c.JSON(401, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }
        
        // Injetar tenant_id no contexto
        c.Set("tenant_id", user.TenantID)
        c.Next()
    }
}
```

**1.4 - Endpoints de Tenant**
```
POST   /api/v1/tenants/register    # Cadastro de nova empresa
GET    /api/v1/tenants/me           # Dados da empresa logada
PATCH  /api/v1/tenants/settings     # Atualizar configurações
```

---

### **FASE 2: Sistema de Serviços e Agenda (Semanas 3-4)**

#### Objetivos:
- Criar catálogo de serviços
- Implementar disponibilidade de horários
- Sistema de agendamento básico

#### Tarefas:

**2.1 - Modelo de Serviços**
```go
// internal/domain/service/service.go
type Service struct {
    ID          string    `json:"id" bson:"_id"`
    TenantID    string    `json:"tenant_id" bson:"tenant_id"`
    Name        string    `json:"name" bson:"name"` // "Corte masculino"
    Description string    `json:"description" bson:"description"`
    Duration    int       `json:"duration" bson:"duration"` // 30 minutos
    Price       float64   `json:"price" bson:"price"`
    IsActive    bool      `json:"is_active" bson:"is_active"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
```

**2.2 - Modelo de Agendamentos (Bookings)**
```go
// internal/domain/booking/booking.go
type Booking struct {
    ID          string    `json:"id" bson:"_id"`
    TenantID    string    `json:"tenant_id" bson:"tenant_id"`
    ServiceID   string    `json:"service_id" bson:"service_id"`
    CustomerID  string    `json:"customer_id" bson:"customer_id"`
    
    // Dados do agendamento
    Date        string    `json:"date" bson:"date"` // "2024-03-15"
    TimeSlot    string    `json:"time_slot" bson:"time_slot"` // "14:00"
    Duration    int       `json:"duration" bson:"duration"` // 30
    
    // Status
    Status      string    `json:"status" bson:"status"` // pending, confirmed, cancelled
    
    // Informações adicionais
    CustomerName  string  `json:"customer_name" bson:"customer_name"`
    CustomerPhone string  `json:"customer_phone" bson:"customer_phone"`
    Notes         string  `json:"notes" bson:"notes"`
    
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
```

**2.3 - Lógica de disponibilidade**
```go
// internal/application/booking/availability.go
type AvailabilityService struct {
    bookingRepo repository.BookingRepository
}

func (s *AvailabilityService) GetAvailableSlots(
    tenantID string, 
    date string,
) ([]TimeSlot, error) {
    // 1. Buscar configurações do tenant
    // 2. Gerar slots baseado em working_hours e interval
    // 3. Buscar agendamentos existentes
    // 4. Remover slots ocupados
    // 5. Retornar slots disponíveis
}
```

**2.4 - Endpoints de Serviços e Agendamentos**
```
# Serviços (apenas tenant owner)
POST   /api/v1/services           # Criar serviço
GET    /api/v1/services           # Listar serviços
PATCH  /api/v1/services/:id       # Editar serviço
DELETE /api/v1/services/:id       # Deletar serviço

# Agendamentos (público - por slug)
GET    /api/v1/booking/:slug/services           # Ver serviços disponíveis
GET    /api/v1/booking/:slug/availability       # Ver horários
POST   /api/v1/booking/:slug/appointments       # Criar agendamento

# Gestão de agendamentos (apenas tenant)
GET    /api/v1/appointments                     # Listar agendamentos
GET    /api/v1/appointments/:id                 # Ver detalhes
PATCH  /api/v1/appointments/:id/confirm         # Confirmar
PATCH  /api/v1/appointments/:id/cancel          # Cancelar
```

---

### **FASE 3: Dashboard e Métricas (Semana 5)**

#### Objetivos:
- Dashboard com estatísticas
- Visualização de agenda
- Relatórios básicos

**3.1 - Modelo de Métricas**
```go
type DashboardMetrics struct {
    TotalBookings       int     `json:"total_bookings"`
    ConfirmedBookings   int     `json:"confirmed_bookings"`
    CancelledBookings   int     `json:"cancelled_bookings"`
    Revenue             float64 `json:"revenue"`
    UpcomingBookings    int     `json:"upcoming_bookings"`
    TopServices         []ServiceMetric `json:"top_services"`
}
```

**3.2 - Endpoints**
```
GET /api/v1/dashboard/metrics        # Métricas gerais
GET /api/v1/dashboard/calendar       # Visão de calendário
GET /api/v1/dashboard/reports        # Relatórios
```

---

### **FASE 4: Chat com IA (Semanas 6-7)**

#### Objetivos:
- Integrar OpenAI GPT
- Chat que responde perguntas
- Permite agendamento via conversa

**4.1 - Modelo de Conversa**
```go
// internal/domain/chat/conversation.go
type Conversation struct {
    ID          string    `json:"id" bson:"_id"`
    TenantID    string    `json:"tenant_id" bson:"tenant_id"`
    CustomerID  string    `json:"customer_id" bson:"customer_id"`
    Messages    []Message `json:"messages" bson:"messages"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type Message struct {
    ID        string    `json:"id"`
    Role      string    `json:"role"` // user, assistant
    Content   string    `json:"content"`
    Timestamp time.Time `json:"timestamp"`
}
```

**4.2 - Integração OpenAI**
```go
// internal/infrastructure/ai/openai_client.go
type OpenAIClient struct {
    apiKey string
    client *openai.Client
}

func (c *OpenAIClient) Chat(
    conversationHistory []Message,
    tenantContext TenantContext,
) (string, error) {
    // Montar prompt com contexto do negócio
    systemPrompt := fmt.Sprintf(`
        Você é um assistente virtual para %s.
        Serviços disponíveis: %v
        Horário de funcionamento: %s
        
        Você pode:
        - Responder sobre serviços e preços
        - Verificar disponibilidade
        - Ajudar a agendar
    `, tenantContext.Name, tenantContext.Services, tenantContext.Hours)
    
    // Chamar OpenAI API
    // Retornar resposta
}
```

**4.3 - Endpoints de Chat**
```
POST /api/v1/chat/:slug/conversations      # Iniciar conversa
POST /api/v1/chat/:slug/messages           # Enviar mensagem
GET  /api/v1/chat/:slug/conversations/:id  # Histórico
```

**4.4 - WebSocket para tempo real (opcional)**
```go
// internal/interfaces/websocket/chat_handler.go
func HandleChatWebSocket(c *gin.Context) {
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    // Implementar comunicação bidirecional
}
```

---

### **FASE 5: Autenticação Pública e Onboarding (Semana 8)**

#### Objetivos:
- Permitir que clientes agendem sem cadastro
- Fluxo de onboarding para novas empresas
- Landing page pública

**5.1 - Agendamento sem login**
- Cliente informa apenas nome e telefone
- Recebe confirmação por SMS/Email
- Pode cancelar com link único

**5.2 - Onboarding de Empresas**
```
Passo 1: Dados básicos (nome, email, senha)
Passo 2: Tipo de negócio e configurações
Passo 3: Adicionar primeiro serviço
Passo 4: Configurar horários
Passo 5: URL pública gerada (exemplo.agendeja.com.br)
```

---

### **FASE 6: Notificações e Lembretes (Semana 9)**

**6.1 - Sistema de notificações**
```go
type Notification struct {
    Type      string // email, sms, whatsapp
    To        string
    Template  string
    Data      map[string]interface{}
}
```

**6.2 - Casos de uso:**
- Confirmação de agendamento (imediato)
- Lembrete 24h antes
- Lembrete 1h antes
- Cancelamento
- Reagendamento

**6.3 - Integrations:**
- SendGrid (email)
- Twilio (SMS)
- WhatsApp Business API

---

### **FASE 7: Planos e Pagamentos (Semanas 10-11)**

**7.1 - Modelos de plano**
```go
type Plan struct {
    ID              string
    Name            string  // "Free", "Basic", "Pro"
    Price           float64
    MaxBookings     int     // Limite mensal
    MaxServices     int
    HasAI           bool
    HasWhatsApp     bool
    HasCustomDomain bool
}
```

**7.2 - Integração com Stripe/Mercado Pago**
```
- Checkout de plano
- Webhooks de pagamento
- Upgrade/downgrade
- Cancelamento
```

---

### **FASE 8: Recursos Avançados (Semanas 12+)**

**8.1 - Funcionalidades extras:**
- ✨ Múltiplos profissionais (barbeiro A, B, C)
- 📊 Analytics avançado
- 🔗 Integração com Google Calendar
- 📱 App mobile (React Native)
- 🎨 Personalização de cores/logo
- 📧 Email marketing
- ⭐ Sistema de avaliações
- 💳 Pagamento online antecipado

---

## 🛠️ Stack Tecnológica Recomendada

### Backend (já tem):
- **Go** (Gin framework)
- **MongoDB** (principal)
- **JWT** para auth

### Adicionar:
- **Redis** - Cache e filas
- **PostgreSQL** - Alternativa (melhor para relatórios)
- **RabbitMQ/Redis Queue** - Jobs assíncronos
- **OpenAI API** - Chat com IA
- **Twilio/SendGrid** - Notificações

### Frontend:
- **React** ou **Next.js** - Dashboard admin
- **TailwindCSS** - Estilização
- **shadcn/ui** - Componentes

### DevOps:
- **Docker** - Containerização
- **Docker Compose** - Dev environment
- **GitHub Actions** - CI/CD
- **Railway/Render** - Hosting inicial
- **Vercel** - Frontend

---

## 💾 Database Schema (MongoDB)

### Collections:

```javascript
// tenants
{
  _id: ObjectId,
  name: "Barbearia do João",
  slug: "barbearia-joao",
  email: "contato@barbearia.com",
  plan: "basic",
  status: "active",
  settings: {
    business_type: "barbershop",
    working_days: ["mon", "tue", "wed", "thu", "fri"],
    working_hours: "09:00-18:00",
    booking_interval: 30,
    timezone: "America/Sao_Paulo"
  },
  created_at: ISODate(),
  updated_at: ISODate()
}

// users
{
  _id: ObjectId,
  tenant_id: ObjectId,
  email: "joao@email.com",
  password: "hashed",
  role: "owner", // owner, employee, customer
  name: "João Silva",
  phone: "+5511999999999",
  created_at: ISODate()
}

// services
{
  _id: ObjectId,
  tenant_id: ObjectId,
  name: "Corte Masculino",
  description: "Corte completo com máquina e tesoura",
  duration: 30, // minutos
  price: 35.00,
  is_active: true,
  created_at: ISODate()
}

// bookings
{
  _id: ObjectId,
  tenant_id: ObjectId,
  service_id: ObjectId,
  customer_id: ObjectId, // null se agendou sem cadastro
  date: "2024-03-15",
  time_slot: "14:00",
  duration: 30,
  status: "confirmed", // pending, confirmed, cancelled, completed
  customer_name: "Maria Santos",
  customer_phone: "+5511988888888",
  customer_email: "maria@email.com",
  notes: "Preferência por barbeiro João",
  created_at: ISODate(),
  updated_at: ISODate()
}

// conversations
{
  _id: ObjectId,
  tenant_id: ObjectId,
  customer_phone: "+5511988888888",
  messages: [
    {
      id: "uuid",
      role: "user",
      content: "Tem horário hoje?",
      timestamp: ISODate()
    },
    {
      id: "uuid",
      role: "assistant",
      content: "Sim! Temos disponível às 14:00 e 16:00.",
      timestamp: ISODate()
    }
  ],
  created_at: ISODate(),
  updated_at: ISODate()
}
```

---

## 🔌 APIs e Endpoints Completos

```
# AUTENTICAÇÃO
POST   /api/v1/auth/register          # Cadastro de empresa
POST   /api/v1/auth/login             # Login
POST   /api/v1/auth/refresh           # Refresh token

# TENANT (requer auth)
GET    /api/v1/tenants/me             # Dados da empresa
PATCH  /api/v1/tenants/settings       # Atualizar config
PATCH  /api/v1/tenants/plan           # Mudar plano

# SERVIÇOS (requer auth + tenant)
POST   /api/v1/services               # Criar
GET    /api/v1/services               # Listar
GET    /api/v1/services/:id           # Ver
PATCH  /api/v1/services/:id           # Editar
DELETE /api/v1/services/:id           # Deletar

# AGENDAMENTOS - Gestão (requer auth)
GET    /api/v1/appointments           # Listar todos
GET    /api/v1/appointments/:id       # Ver detalhes
PATCH  /api/v1/appointments/:id/confirm
PATCH  /api/v1/appointments/:id/cancel
PATCH  /api/v1/appointments/:id/complete

# AGENDAMENTOS - Público (por slug)
GET    /api/v1/booking/:slug/info             # Info do negócio
GET    /api/v1/booking/:slug/services         # Serviços
GET    /api/v1/booking/:slug/availability     # Horários disponíveis
POST   /api/v1/booking/:slug/appointments     # Criar agendamento

# CHAT IA
POST   /api/v1/chat/:slug/conversations       # Iniciar
POST   /api/v1/chat/:slug/messages            # Enviar msg
GET    /api/v1/chat/:slug/conversations/:id   # Histórico

# DASHBOARD
GET    /api/v1/dashboard/metrics              # KPIs
GET    /api/v1/dashboard/calendar             # Calendário
GET    /api/v1/dashboard/revenue              # Faturamento

# NOTIFICAÇÕES
POST   /api/v1/notifications/send             # Enviar (admin)
GET    /api/v1/notifications/templates        # Templates

# WEBHOOKS (pagamentos)
POST   /api/v1/webhooks/stripe
POST   /api/v1/webhooks/mercadopago
```

---

## 🤖 Integrações com IA

### OpenAI - Casos de uso:

**1. Chat de Atendimento**
```go
systemPrompt := `
Você é assistente virtual da {business_name}.

INFORMAÇÕES DO NEGÓCIO:
- Tipo: {business_type}
- Serviços: {services_list}
- Horários: {working_hours}
- Dias: {working_days}

SUAS FUNÇÕES:
1. Responder sobre serviços e preços
2. Verificar disponibilidade de horários
3. Ajudar no agendamento
4. Ser cordial e profissional

REGRAS:
- Sempre confirme dados antes de agendar
- Se não souber, peça para falar com atendente
- Use linguagem natural brasileira
`
```

**2. Function Calling para agendar**
```go
functions := []openai.Function{
    {
        Name: "check_availability",
        Description: "Verifica horários disponíveis",
        Parameters: {
            "date": "string",
            "service_id": "string",
        },
    },
    {
        Name: "create_booking",
        Description: "Cria um agendamento",
        Parameters: {
            "date": "string",
            "time": "string",
            "service_id": "string",
            "customer_name": "string",
            "customer_phone": "string",
        },
    },
}
```

**3. Análise de Sentimento (futuro)**
- Identificar clientes insatisfeitos
- Priorizar atendimentos urgentes
- Melhorar scripts de atendimento

---

## 🚀 Deploy e Infraestrutura

### Opção 1: Railway (recomendado para MVP)
```yaml
# railway.toml
[build]
builder = "nixpacks"

[deploy]
startCommand = "go run main.go"
healthcheckPath = "/health"
restartPolicyType = "on-failure"
```

**Vantagens:**
- Deploy automático via Git
- MongoDB integrado
- Redis disponível
- SSL grátis
- ~$5-20/mês

### Opção 2: Docker Compose (local/VPS)
```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - MONGODB_URI=mongodb://mongo:27017
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=${JWT_SECRET}
      - OPENAI_API_KEY=${OPENAI_API_KEY}
    depends_on:
      - mongo
      - redis

  mongo:
    image: mongo:7
    volumes:
      - mongo_data:/data/db
    ports:
      - "27017:27017"

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"

  worker:
    build: .
    command: go run cmd/worker/main.go
    depends_on:
      - mongo
      - redis

volumes:
  mongo_data:
```

### Opção 3: Cloud (produção)
- **Backend**: AWS ECS / Google Cloud Run
- **Database**: MongoDB Atlas
- **Cache**: AWS ElastiCache / Redis Cloud
- **Storage**: AWS S3
- **CDN**: CloudFlare

---

## 📊 Métricas de Sucesso

### KPIs para acompanhar:
- 📈 **MRR** (Monthly Recurring Revenue)
- 👥 **Tenants ativos**
- 📅 **Agendamentos por mês**
- 💬 **Conversas com IA**
- ⭐ **Churn rate**
- 💵 **Lifetime Value (LTV)**

---

## ⚡ Quick Wins (features rápidas de implementar)

1. **Página pública de agendamento** (2-3 dias)
2. **Notificação por email** (1 dia)
3. **Dashboard básico** (2 dias)
4. **Chat com IA** (3-4 dias com OpenAI)
5. **Exportar relatórios CSV** (1 dia)

---

## 🎯 Próximos Passos Imediatos

### Semana 1-2: Fundação
1. ✅ Criar branch `scheduler-saas`
2. ✅ Atualizar estrutura de pastas
3. ✅ Implementar modelo Tenant
4. ✅ Migrar auth para suportar multi-tenant
5. ✅ Criar endpoints de tenant

### Você deveria começar com:
```bash
# 1. Criar migrations
mkdir -p cmd/migrations

# 2. Atualizar dependências
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/sashabaranov/go-openai

# 3. Configurar variáveis de ambiente
cp .env.example .env

# 4. Rodar com Docker
docker-compose up -d
```

---

## 📚 Recursos Úteis

### Documentação:
- [Gin Framework](https://gin-gonic.com/docs/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [OpenAI Go SDK](https://github.com/sashabaranov/go-openai)
- [JWT Go](https://github.com/golang-jwt/jwt)

### Inspirações:
- [Cal.com](https://cal.com) - Agendamento open source
- [Calendly](https://calendly.com) - Líder de mercado
- [Acuity Scheduling](https://acuityscheduling.com)

---

## 💡 Dicas Importantes

1. **Comece simples**: MVP funcional > features complexas
2. **Multi-tenant desde o início**: Mais difícil refatorar depois
3. **Valide com usuários reais**: Encontre 3-5 barbearias/salões para testar
4. **Documentar APIs**: Use Swagger/OpenAPI
5. **Testes**: Pelo menos nos casos de uso críticos
6. **Segurança**: 
   - Rate limiting
   - Validação de inputs
   - CORS configurado
   - Secrets em variáveis de ambiente

---

## 🎁 Bonus: Monetização

### Modelo de Pricing sugerido:

**Free Tier** (para tração inicial)
- 50 agendamentos/mês
- 1 serviço
- Sem IA
- Branding "Powered by YourApp"

**Basic - R$ 29/mês**
- 200 agendamentos/mês
- Serviços ilimitados
- Chat IA básico
- Notificações por email
- Sem branding

**Pro - R$ 79/mês**
- Agendamentos ilimitados
- IA avançada
- WhatsApp integration
- Analytics
- Custom domain
- API access

**Enterprise - R$ 199/mês**
- Tudo do Pro
- Múltiplos profissionais
- White label
- Suporte prioritário

---

## 🚀 Boa sorte com o projeto!

Lembre-se: **feito é melhor que perfeito**. 

Comece pequeno, valide com usuários, e evolua baseado em feedback real!

---

**Próximo arquivo**: Vou criar o schema detalhado do banco de dados e exemplos de código para começar.
