package models

import (
	"time"
)

// TaskStatus describes the current status of a specific task
type TaskStatus string

// TaskStatus constants, includes pending, in progress, completed and cancelled
const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

// Task represents a task in the system
type Task struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     *time.Time `json:"due_time,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TaskCreate struct that represents the data required to create a new task
type TaskCreate struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     *time.Time `json:"due_time,omitempty"`
}

// TaskUpdate represents the information needed to update a task
type TaskUpdate struct {
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	DueDate     *time.Time  `json:"due_time,omitempty"`
}

// TaskFilter represents the filters that can be applied when looking for a task
type TaskFilter struct {
	Status     TaskStatus `json:"status"`
	DueBefore  *time.Time `json:"due_before,omitempty"`
	DueAfter   *time.Time `json:"due_after,omitempty"`
	SearchTerm *string    `json:"search_term,omitempty"`
	Page       int        `json:"page,omitempty"`
	PageSize   int        `json:"page_size,omitempty"`
}
