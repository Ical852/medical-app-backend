package models

import _ "github.com/jinzhu/gorm"

type User struct {
    ID       uint   `gorm:"primary_key"`
    Name     string `gorm:"size:100"`
    Email    string `gorm:"unique;size:100"`
    Password string `gorm:"size:100"`
    Role     string `gorm:"size:20"`
}

type Doctor struct {
    ID            uint   `gorm:"primary_key"`
    UserID        uint   `gorm:"index"`
    Specialization string `gorm:"size:100"`
}

type Medicine struct {
    ID          uint    `gorm:"primary_key"`
    Name        string  `gorm:"size:100"`
    Description string  `gorm:"size:255"`
    Price       float64
    Stock       int
}

type Chat struct {
    ID        uint `gorm:"primary_key"`
    UserID    uint `gorm:"index"`
    DoctorID  uint `gorm:"index"`
    CreatedAt int64
}

type Message struct {
    ID        uint `gorm:"primary_key"`
    ChatID    uint `gorm:"index"`
    SenderID  uint `gorm:"index"`
    Message   string `gorm:"size:255"`
    SentAt    int64
}

type Order struct {
    ID         uint `gorm:"primary_key"`
    UserID     uint `gorm:"index"`
    TotalAmount float64
    Status     string `gorm:"size:20"`
    CreatedAt  int64
}

type OrderItem struct {
    ID         uint `gorm:"primary_key"`
    OrderID    uint `gorm:"index"`
    MedicineID uint `gorm:"index"`
    Quantity   int
    Price      float64
}
