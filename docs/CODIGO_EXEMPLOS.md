# 🛠️ Exemplos de Código - Início do Projeto

## 📁 Estrutura de Pastas Completa

```
api-rest-crud-golang/
├── cmd/
│   ├── api/
│   │   └── main.go                 # Entry point da API
│   ├── worker/
│   │   └── main.go                 # Background jobs
│   └── migrations/
│       └── main.go                 # Migrations
│
├── internal/
│   ├── domain/                     # Entities (regras de negócio)
│   │   ├── tenant/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── service/                # Serviços oferecidos
│   │   │   ├── entity.go
│   │   │   └── repository.go
│   │   ├── booking/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── chat/
│   │       ├── entity.go
│   │       └── repository.go
│   │
│   ├── application/                # Use cases
│   │   ├── tenant/
│   │   │   ├── create_tenant.go
│   │   │   └── update_settings.go
│   │   ├── booking/
│   │   │   ├── create_booking.go
│   │   │   ├── get_availability.go
│   │   │   └── cancel_booking.go
│   │   └── chat/
│   │       └── process_message.go
│   │
│   ├── infrastructure/             # External dependencies
│   │   ├── database/
│   │   │   ├── mongodb.go
│   │   │   └── redis.go
│   │   ├── ai/
│   │   │   └── openai_client.go
│   │   ├── notification/
│   │   │   ├── email.go
│   │   │   └── sms.go
│   │   └── queue/
│   │       └── redis_queue.go
│   │
│   └── interfaces/                 # HTTP, WebSocket, etc
│       ├── http/
│       │   ├── routes/
│       │   │   ├── tenant.go
│       │   │   ├── booking.go
│       │   │   ├── service.go
│       │   │   └── chat.go
│       │   ├── middleware/
│       │   │   ├── auth.go
│       │   │   ├── tenant.go
│       │   │   └── rate_limit.go
│       │   └── handlers/
│       │       ├── tenant_handler.go
│       │       ├── booking_handler.go
│       │       └── chat_handler.go
│       └── websocket/
│           └── chat_ws.go
│
├── pkg/                            # Shared packages
│   ├── jwt/
│   │   └── jwt.go
│   ├── logger/
│   │   └── logger.go
│   ├── validator/
│   │   └── validator.go
│   └── utils/
│       ├── slug.go
│       └── time.go
│
├── config/
│   └── config.go
│
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## 1️⃣ DOMAIN LAYER - Entidades

### internal/domain/tenant/entity.go
```go
package tenant

import (
    "time"
)

type Tenant struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Name      string    `json:"name" bson:"name" binding:"required"`
    Slug      string    `json:"slug" bson:"slug" binding:"required"`
    Email     string    `json:"email" bson:"email" binding:"required,email"`
    Phone     string    `json:"phone" bson:"phone"`
    Plan      string    `json:"plan" bson:"plan"` // free, basic, pro
    Status    string    `json:"status" bson:"status"` // active, suspended, cancelled
    Settings  Settings  `json:"settings" bson:"settings"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Settings struct {
    BusinessType      string   `json:"business_type" bson:"business_type"` // barbershop, salon, clinic
    WorkingDays       []string `json:"working_days" bson:"working_days"`
    WorkingHours      string   `json:"working_hours" bson:"working_hours"` // "09:00-18:00"
    BookingInterval   int      `json:"booking_interval" bson:"booking_interval"` // minutes
    MaxAdvanceDays    int      `json:"max_advance_days" bson:"max_advance_days"`
    Timezone          string   `json:"timezone" bson:"timezone"`
    AutoConfirm       bool     `json:"auto_confirm" bson:"auto_confirm"`
    RequirePhone      bool     `json:"require_phone" bson:"require_phone"`
    AllowCancellation bool     `json:"allow_cancellation" bson:"allow_cancellation"`
}

// Métodos de negócio
func (t *Tenant) IsActive() bool {
    return t.Status == "active"
}

func (t *Tenant) CanCreateService() bool {
    if t.Plan == "free" {
        return false // Pode ter lógica de limite aqui
    }
    return t.IsActive()
}
```

### internal/domain/user/entity.go
```go
package user

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    TenantID  string    `json:"tenant_id" bson:"tenant_id" binding:"required"`
    Email     string    `json:"email" bson:"email" binding:"required,email"`
    Password  string    `json:"-" bson:"password"`
    Role      string    `json:"role" bson:"role"` // owner, employee, customer
    Name      string    `json:"name" bson:"name" binding:"required"`
    Phone     string    `json:"phone" bson:"phone"`
    IsActive  bool      `json:"is_active" bson:"is_active"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *User) HashPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return err
    }
    u.Password = string(bytes)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

func (u *User) IsOwner() bool {
    return u.Role == "owner"
}

func (u *User) CanManageBookings() bool {
    return u.Role == "owner" || u.Role == "employee"
}
```

### internal/domain/service/entity.go
```go
package service

import "time"

type Service struct {
    ID          string    `json:"id" bson:"_id,omitempty"`
    TenantID    string    `json:"tenant_id" bson:"tenant_id" binding:"required"`
    Name        string    `json:"name" bson:"name" binding:"required"`
    Description string    `json:"description" bson:"description"`
    Duration    int       `json:"duration" bson:"duration" binding:"required,min=15"` // minutes
    Price       float64   `json:"price" bson:"price" binding:"required,min=0"`
    Category    string    `json:"category" bson:"category"`
    IsActive    bool      `json:"is_active" bson:"is_active"`
    ImageURL    string    `json:"image_url" bson:"image_url"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (s *Service) FormattedPrice() string {
    return fmt.Sprintf("R$ %.2f", s.Price)
}
```

### internal/domain/booking/entity.go
```go
package booking

import (
    "time"
)

type Booking struct {
    ID            string    `json:"id" bson:"_id,omitempty"`
    TenantID      string    `json:"tenant_id" bson:"tenant_id" binding:"required"`
    ServiceID     string    `json:"service_id" bson:"service_id" binding:"required"`
    CustomerID    string    `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
    
    // Scheduling info
    Date          string    `json:"date" bson:"date" binding:"required"` // "2024-03-15"
    TimeSlot      string    `json:"time_slot" bson:"time_slot" binding:"required"` // "14:00"
    Duration      int       `json:"duration" bson:"duration"` // minutes
    EndTime       string    `json:"end_time" bson:"end_time"` // calculated
    
    // Status
    Status        string    `json:"status" bson:"status"` // pending, confirmed, cancelled, completed, no_show
    
    // Customer info
    CustomerName  string    `json:"customer_name" bson:"customer_name" binding:"required"`
    CustomerPhone string    `json:"customer_phone" bson:"customer_phone" binding:"required"`
    CustomerEmail string    `json:"customer_email" bson:"customer_email"`
    
    // Additional
    Notes         string    `json:"notes" bson:"notes"`
    CancelReason  string    `json:"cancel_reason,omitempty" bson:"cancel_reason,omitempty"`
    
    // Metadata
    CreatedAt     time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
    ConfirmedAt   *time.Time `json:"confirmed_at,omitempty" bson:"confirmed_at,omitempty"`
    CancelledAt   *time.Time `json:"cancelled_at,omitempty" bson:"cancelled_at,omitempty"`
}

func (b *Booking) CanBeCancelled() bool {
    return b.Status == "pending" || b.Status == "confirmed"
}

func (b *Booking) IsUpcoming() bool {
    // Lógica para verificar se está no futuro
    bookingDateTime := parseDateTime(b.Date, b.TimeSlot)
    return bookingDateTime.After(time.Now())
}
```

---

## 2️⃣ REPOSITORY LAYER

### internal/domain/tenant/repository.go
```go
package tenant

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
    Create(ctx context.Context, tenant *Tenant) error
    FindByID(ctx context.Context, id string) (*Tenant, error)
    FindBySlug(ctx context.Context, slug string) (*Tenant, error)
    Update(ctx context.Context, tenant *Tenant) error
    Delete(ctx context.Context, id string) error
}

type MongoRepository struct {
    collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) Repository {
    return &MongoRepository{
        collection: db.Collection("tenants"),
    }
}

func (r *MongoRepository) Create(ctx context.Context, tenant *Tenant) error {
    _, err := r.collection.InsertOne(ctx, tenant)
    return err
}

func (r *MongoRepository) FindByID(ctx context.Context, id string) (*Tenant, error) {
    var tenant Tenant
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&tenant)
    if err != nil {
        return nil, err
    }
    return &tenant, nil
}

func (r *MongoRepository) FindBySlug(ctx context.Context, slug string) (*Tenant, error) {
    var tenant Tenant
    err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&tenant)
    if err != nil {
        return nil, err
    }
    return &tenant, nil
}

func (r *MongoRepository) Update(ctx context.Context, tenant *Tenant) error {
    filter := bson.M{"_id": tenant.ID}
    update := bson.M{"$set": tenant}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *MongoRepository) Delete(ctx context.Context, id string) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
```

### internal/domain/booking/repository.go
```go
package booking

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
    Create(ctx context.Context, booking *Booking) error
    FindByID(ctx context.Context, id string) (*Booking, error)
    FindByTenant(ctx context.Context, tenantID string, filters map[string]interface{}) ([]*Booking, error)
    FindByDateRange(ctx context.Context, tenantID, startDate, endDate string) ([]*Booking, error)
    Update(ctx context.Context, booking *Booking) error
}

type MongoRepository struct {
    collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) Repository {
    return &MongoRepository{
        collection: db.Collection("bookings"),
    }
}

func (r *MongoRepository) FindByDateRange(
    ctx context.Context,
    tenantID, startDate, endDate string,
) ([]*Booking, error) {
    filter := bson.M{
        "tenant_id": tenantID,
        "date": bson.M{
            "$gte": startDate,
            "$lte": endDate,
        },
        "status": bson.M{
            "$in": []string{"confirmed", "pending"},
        },
    }
    
    cursor, err := r.collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    
    var bookings []*Booking
    if err = cursor.All(ctx, &bookings); err != nil {
        return nil, err
    }
    
    return bookings, nil
}
```

---

## 3️⃣ APPLICATION LAYER - Use Cases

### internal/application/booking/get_availability.go
```go
package booking

import (
    "context"
    "fmt"
    "time"
    "your-app/internal/domain/booking"
    "your-app/internal/domain/tenant"
)

type GetAvailabilityUseCase struct {
    tenantRepo  tenant.Repository
    bookingRepo booking.Repository
}

type AvailabilityRequest struct {
    TenantSlug string
    Date       string
    ServiceID  string
}

type TimeSlot struct {
    Time      string `json:"time"`
    Available bool   `json:"available"`
}

func (uc *GetAvailabilityUseCase) Execute(
    ctx context.Context,
    req AvailabilityRequest,
) ([]TimeSlot, error) {
    // 1. Buscar tenant
    tenant, err := uc.tenantRepo.FindBySlug(ctx, req.TenantSlug)
    if err != nil {
        return nil, fmt.Errorf("tenant not found")
    }
    
    // 2. Validar se data está no range permitido
    requestDate, _ := time.Parse("2006-01-02", req.Date)
    maxDate := time.Now().AddDate(0, 0, tenant.Settings.MaxAdvanceDays)
    
    if requestDate.After(maxDate) {
        return nil, fmt.Errorf("date too far in advance")
    }
    
    // 3. Gerar slots baseado nas configurações
    slots := uc.generateTimeSlots(tenant.Settings)
    
    // 4. Buscar agendamentos existentes
    bookings, err := uc.bookingRepo.FindByDateRange(ctx, tenant.ID, req.Date, req.Date)
    if err != nil {
        return nil, err
    }
    
    // 5. Marcar slots ocupados
    for i := range slots {
        slots[i].Available = !uc.isSlotOccupied(slots[i].Time, bookings)
    }
    
    return slots, nil
}

func (uc *GetAvailabilityUseCase) generateTimeSlots(settings tenant.Settings) []TimeSlot {
    // Parse working hours: "09:00-18:00"
    // Gerar slots a cada X minutos (settings.BookingInterval)
    // Retornar lista de TimeSlot
    
    var slots []TimeSlot
    // Implementação aqui...
    return slots
}

func (uc *GetAvailabilityUseCase) isSlotOccupied(timeSlot string, bookings []*booking.Booking) bool {
    for _, b := range bookings {
        if b.TimeSlot == timeSlot && (b.Status == "confirmed" || b.Status == "pending") {
            return true
        }
    }
    return false
}
```

### internal/application/booking/create_booking.go
```go
package booking

import (
    "context"
    "fmt"
    "time"
    "github.com/google/uuid"
    "your-app/internal/domain/booking"
    "your-app/internal/domain/tenant"
    "your-app/internal/domain/service"
)

type CreateBookingUseCase struct {
    tenantRepo  tenant.Repository
    serviceRepo service.Repository
    bookingRepo booking.Repository
}

type CreateBookingRequest struct {
    TenantSlug    string  `json:"tenant_slug" binding:"required"`
    ServiceID     string  `json:"service_id" binding:"required"`
    Date          string  `json:"date" binding:"required"`
    TimeSlot      string  `json:"time_slot" binding:"required"`
    CustomerName  string  `json:"customer_name" binding:"required"`
    CustomerPhone string  `json:"customer_phone" binding:"required"`
    CustomerEmail string  `json:"customer_email"`
    Notes         string  `json:"notes"`
}

func (uc *CreateBookingUseCase) Execute(
    ctx context.Context,
    req CreateBookingRequest,
) (*booking.Booking, error) {
    
    // 1. Buscar tenant
    tenant, err := uc.tenantRepo.FindBySlug(ctx, req.TenantSlug)
    if err != nil {
        return nil, fmt.Errorf("tenant not found")
    }
    
    if !tenant.IsActive() {
        return nil, fmt.Errorf("tenant is not active")
    }
    
    // 2. Buscar serviço
    svc, err := uc.serviceRepo.FindByID(ctx, req.ServiceID)
    if err != nil {
        return nil, fmt.Errorf("service not found")
    }
    
    // 3. Validar disponibilidade
    bookings, err := uc.bookingRepo.FindByDateRange(ctx, tenant.ID, req.Date, req.Date)
    if err != nil {
        return nil, err
    }
    
    if uc.isTimeSlotOccupied(req.TimeSlot, bookings) {
        return nil, fmt.Errorf("time slot not available")
    }
    
    // 4. Calcular end_time
    endTime := uc.calculateEndTime(req.TimeSlot, svc.Duration)
    
    // 5. Criar booking
    newBooking := &booking.Booking{
        ID:            uuid.New().String(),
        TenantID:      tenant.ID,
        ServiceID:     req.ServiceID,
        Date:          req.Date,
        TimeSlot:      req.TimeSlot,
        Duration:      svc.Duration,
        EndTime:       endTime,
        Status:        uc.getInitialStatus(tenant.Settings.AutoConfirm),
        CustomerName:  req.CustomerName,
        CustomerPhone: req.CustomerPhone,
        CustomerEmail: req.CustomerEmail,
        Notes:         req.Notes,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }
    
    // 6. Salvar
    if err := uc.bookingRepo.Create(ctx, newBooking); err != nil {
        return nil, err
    }
    
    // 7. Enviar notificação (async - pode usar queue)
    go uc.sendBookingNotification(newBooking)
    
    return newBooking, nil
}

func (uc *CreateBookingUseCase) getInitialStatus(autoConfirm bool) string {
    if autoConfirm {
        return "confirmed"
    }
    return "pending"
}

func (uc *CreateBookingUseCase) calculateEndTime(startTime string, duration int) string {
    // Parse "14:00" -> add duration -> return "14:30"
    t, _ := time.Parse("15:04", startTime)
    endTime := t.Add(time.Duration(duration) * time.Minute)
    return endTime.Format("15:04")
}
```

---

## 4️⃣ HTTP LAYER - Handlers

### internal/interfaces/http/handlers/booking_handler.go
```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "your-app/internal/application/booking"
)

type BookingHandler struct {
    getAvailability *booking.GetAvailabilityUseCase
    createBooking   *booking.CreateBookingUseCase
}

func NewBookingHandler(
    getAvailability *booking.GetAvailabilityUseCase,
    createBooking *booking.CreateBookingUseCase,
) *BookingHandler {
    return &BookingHandler{
        getAvailability: getAvailability,
        createBooking:   createBooking,
    }
}

// GET /api/v1/booking/:slug/availability?date=2024-03-15&service_id=xxx
func (h *BookingHandler) GetAvailability(c *gin.Context) {
    slug := c.Param("slug")
    date := c.Query("date")
    serviceID := c.Query("service_id")
    
    if date == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "date is required",
        })
        return
    }
    
    req := booking.AvailabilityRequest{
        TenantSlug: slug,
        Date:       date,
        ServiceID:  serviceID,
    }
    
    slots, err := h.getAvailability.Execute(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "date":  date,
        "slots": slots,
    })
}

// POST /api/v1/booking/:slug/appointments
func (h *BookingHandler) CreateBooking(c *gin.Context) {
    slug := c.Param("slug")
    
    var req booking.CreateBookingRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    req.TenantSlug = slug
    
    newBooking, err := h.createBooking.Execute(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "Booking created successfully",
        "booking": newBooking,
    })
}
```

---

## 5️⃣ ROUTES

### internal/interfaces/http/routes/booking.go
```go
package routes

import (
    "github.com/gin-gonic/gin"
    "your-app/internal/interfaces/http/handlers"
)

func SetupBookingRoutes(r *gin.Engine, handler *handlers.BookingHandler) {
    
    // Rotas públicas (por slug)
    public := r.Group("/api/v1/booking/:slug")
    {
        public.GET("/info", handler.GetTenantInfo)
        public.GET("/services", handler.GetServices)
        public.GET("/availability", handler.GetAvailability)
        public.POST("/appointments", handler.CreateBooking)
    }
    
    // Rotas protegidas (gestão)
    private := r.Group("/api/v1/appointments")
    private.Use(AuthMiddleware()) // Requer JWT
    private.Use(TenantMiddleware()) // Injeta tenant_id
    {
        private.GET("", handler.ListBookings)
        private.GET("/:id", handler.GetBooking)
        private.PATCH("/:id/confirm", handler.ConfirmBooking)
        private.PATCH("/:id/cancel", handler.CancelBooking)
        private.PATCH("/:id/complete", handler.CompleteBooking)
    }
}
```

---

## 6️⃣ INTEGRAÇÃO COM OPENAI

### internal/infrastructure/ai/openai_client.go
```go
package ai

import (
    "context"
    "fmt"
    openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
    client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
    return &OpenAIClient{
        client: openai.NewClient(apiKey),
    }
}

type TenantContext struct {
    Name         string
    BusinessType string
    Services     []ServiceInfo
    WorkingHours string
}

type ServiceInfo struct {
    Name  string
    Price float64
}

func (c *OpenAIClient) ChatWithContext(
    ctx context.Context,
    messages []openai.ChatCompletionMessage,
    tenantCtx TenantContext,
) (string, error) {
    
    // Montar system prompt com contexto
    systemPrompt := c.buildSystemPrompt(tenantCtx)
    
    // Adicionar system message
    allMessages := append(
        []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleSystem,
                Content: systemPrompt,
            },
        },
        messages...,
    )
    
    resp, err := c.client.CreateChatCompletion(
        ctx,
        openai.ChatCompletionRequest{
            Model:    openai.GPT4TurboPreview,
            Messages: allMessages,
            Temperature: 0.7,
        },
    )
    
    if err != nil {
        return "", err
    }
    
    return resp.Choices[0].Message.Content, nil
}

func (c *OpenAIClient) buildSystemPrompt(ctx TenantContext) string {
    servicesStr := ""
    for _, s := range ctx.Services {
        servicesStr += fmt.Sprintf("- %s: R$ %.2f\n", s.Name, s.Price)
    }
    
    return fmt.Sprintf(`
Você é um assistente virtual para %s, um estabelecimento do tipo %s.

INFORMAÇÕES DO ESTABELECIMENTO:
Horário de funcionamento: %s

SERVIÇOS DISPONÍVEIS:
%s

SUAS FUNÇÕES:
1. Responder perguntas sobre serviços, preços e horários
2. Verificar disponibilidade de horários
3. Auxiliar no processo de agendamento
4. Ser cordial, profissional e objetivo

REGRAS:
- Sempre confirme os dados antes de criar um agendamento
- Se não souber algo, seja honesto e sugira contato direto
- Use linguagem natural e brasileira
- Mantenha respostas concisas
`, ctx.Name, ctx.BusinessType, ctx.WorkingHours, servicesStr)
}
```

---

## 7️⃣ CONFIGURAÇÃO

### config/config.go
```go
package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Port            string
    MongoURI        string
    DatabaseName    string
    RedisURL        string
    JWTSecret       string
    OpenAIKey       string
    SendGridKey     string
    TwilioSID       string
    TwilioToken     string
    Environment     string
}

func Load() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        // Ignore error in production
    }
    
    return &Config{
        Port:         getEnv("PORT", "8080"),
        MongoURI:     getEnv("MONGODB_URI", "mongodb://localhost:27017"),
        DatabaseName: getEnv("DATABASE_NAME", "scheduler_saas"),
        RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379"),
        JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
        OpenAIKey:    getEnv("OPENAI_API_KEY", ""),
        SendGridKey:  getEnv("SENDGRID_API_KEY", ""),
        Environment:  getEnv("ENVIRONMENT", "development"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
```

### .env.example
```bash
# Server
PORT=8080
ENVIRONMENT=development

# Database
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=scheduler_saas

# Cache
REDIS_URL=redis://localhost:6379

# JWT
JWT_SECRET=your-super-secret-key-change-in-production

# OpenAI
OPENAI_API_KEY=sk-your-openai-key

# Email (SendGrid)
SENDGRID_API_KEY=your-sendgrid-key
FROM_EMAIL=noreply@yourapp.com

# SMS (Twilio)
TWILIO_ACCOUNT_SID=your-twilio-sid
TWILIO_AUTH_TOKEN=your-twilio-token
TWILIO_PHONE_NUMBER=+15555555555

# Frontend URL
FRONTEND_URL=http://localhost:3000

# Cors
ALLOWED_ORIGINS=http://localhost:3000,https://yourapp.com
```

---

## 8️⃣ DOCKER SETUP

### docker-compose.yml
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
      - DATABASE_NAME=scheduler_saas
      - JWT_SECRET=${JWT_SECRET}
      - OPENAI_API_KEY=${OPENAI_API_KEY}
    depends_on:
      - mongo
      - redis
    volumes:
      - .:/app
    command: go run cmd/api/main.go

  mongo:
    image: mongo:7
    ports:
      - "27017:27017"
    volumes:
      - mongo_data:/data/db
    environment:
      - MONGO_INITDB_DATABASE=scheduler_saas

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"

  mongo-express:
    image: mongo-express
    ports:
      - "8081:8081"
    environment:
      - ME_CONFIG_MONGODB_URL=mongodb://mongo:27017
      - ME_CONFIG_BASICAUTH_USERNAME=admin
      - ME_CONFIG_BASICAUTH_PASSWORD=admin
    depends_on:
      - mongo

volumes:
  mongo_data:
```

### Dockerfile
```dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN go build -o main cmd/api/main.go

# Run
CMD ["./main"]
```

---

## 9️⃣ MAIN ENTRY POINT

### cmd/api/main.go
```go
package main

import (
    "context"
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    
    "your-app/config"
    "your-app/internal/domain/tenant"
    "your-app/internal/domain/booking"
    bookingApp "your-app/internal/application/booking"
    "your-app/internal/interfaces/http/handlers"
    "your-app/internal/interfaces/http/routes"
    "your-app/internal/infrastructure/ai"
)

func main() {
    // Load config
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
    
    // Connect to MongoDB
    client, err := mongo.Connect(
        context.Background(),
        options.Client().ApplyURI(cfg.MongoURI),
    )
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }
    defer client.Disconnect(context.Background())
    
    db := client.Database(cfg.DatabaseName)
    
    // Initialize repositories
    tenantRepo := tenant.NewMongoRepository(db)
    bookingRepo := booking.NewMongoRepository(db)
    
    // Initialize use cases
    getAvailability := &bookingApp.GetAvailabilityUseCase{
        tenantRepo:  tenantRepo,
        bookingRepo: bookingRepo,
    }
    
    createBooking := &bookingApp.CreateBookingUseCase{
        tenantRepo:  tenantRepo,
        bookingRepo: bookingRepo,
    }
    
    // Initialize OpenAI
    aiClient := ai.NewOpenAIClient(cfg.OpenAIKey)
    
    // Initialize handlers
    bookingHandler := handlers.NewBookingHandler(
        getAvailability,
        createBooking,
    )
    
    // Setup Gin
    r := gin.Default()
    
    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    // Setup routes
    routes.SetupBookingRoutes(r, bookingHandler)
    
    // Start server
    addr := fmt.Sprintf(":%s", cfg.Port)
    log.Printf("Server running on %s", addr)
    r.Run(addr)
}
```

---

## 🚀 Como começar

### 1. Clone e configure
```bash
git checkout -b scheduler-saas
cp .env.example .env
# Edite .env com suas chaves
```

### 2. Instale dependências
```bash
go mod init your-app
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get github.com/google/uuid
go get github.com/joho/godotenv
go get github.com/sashabaranov/go-openai
go get golang.org/x/crypto/bcrypt
```

### 3. Rode com Docker
```bash
docker-compose up -d
```

### 4. Teste a API
```bash
# Health check
curl http://localhost:8080/health

# Criar tenant
curl -X POST http://localhost:8080/api/v1/tenants/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Barbearia do João",
    "email": "joao@barbearia.com",
    "password": "senha123"
  }'

# Ver disponibilidade
curl http://localhost:8080/api/v1/booking/barbearia-joao/availability?date=2024-03-15
```

---

## ✅ Próximos passos

1. Implementar autenticação JWT completa
2. Criar endpoints de gestão de serviços
3. Adicionar dashboard com métricas
4. Implementar chat com IA
5. Sistema de notificações
6. Testes unitários
7. Deploy em produção

**Boa sorte!** 🚀
