package task

type Task struct {
    ID          uint32  `gorm:"primaryKey"`
    UserID      uint32  `gorm:"column:user_id"` // !!! Важно: убедись, что имя колонки указано правильно
    Title       string
    Description string
    IsDone      bool
}