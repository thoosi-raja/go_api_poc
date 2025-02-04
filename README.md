# To-Do Tasking Project

## Overview
This is a **To-Do Tasking Project** that allows users to create tasks with a title and description. Users can update the task status accordingly. If a task is not marked as "completed" within **1 hour**, it is automatically invalidated, and its status is changed to **"critical"**.

## API Endpoints

### **1. Get All Tasks**  
**Method:** `GET`  
**URL:** `http://localhost:8080/tasks`  
**Description:** Fetches all tasks from the database.  
**Response:**
```json
[
    {
        "Id": 1,
        "Title": "Sample Task",
        "Description": "This is a sample task description",
        "Status": "pending",
        "Created_on": "2024-02-03T12:00:00Z"
    }
]
```

---
### **2. Get Task**  
**Method:** `GET`  
**URL:** `http://localhost:8080/tasks/{taskId}`  
**Description:** Fetches a specific task using its `taskId`.  
**Response:**
```json
{
    "Id": 1,
    "Title": "Sample Task",
    "Description": "This is a sample task description",
    "Status": "pending",
    "Created_on": "2024-02-03T12:00:00Z"
}
```

---
### **3. Create Task**  
**Method:** `POST`  
**URL:** `http://localhost:8080/tasks`  
**Description:** Creates a new task in the database.  
**Request Body:**
```json
{
    "Title": "New Task",
    "Description": "Task details",
    "Status": "pending"
}
```
**Response:**
```json
{
    "id": 2,
    "Title": "New Task",
    "Description": "Task details",
    "Status": "pending",
    "CreatedOn": "2024-02-03T13:00:00Z"
}
```

---
### **4. Update Task**  
**Method:** `PUT`  
**URL:** `http://localhost:8080/tasks/{taskId}`  
**Description:** Updates a specific task.  
**Request Body:**
```json
{
    "Id": 2,
    "Title": "Updated Task",
    "Description": "Updated task details",
    "Status": "in_progress"
}
```
**Response:**
```json
{
    "id": 2,
    "Title": "Updated Task",
    "Description": "Updated task details",
    "Status": "in_progress",
    "CreatedOn": "2024-02-03T13:00:00Z"
}
```

---
### **5. Delete Task**  
**Method:** `DELETE`  
**URL:** `http://localhost:8080/tasks/{taskId}`  
**Description:** Deletes a specific task from the database.

---

## Automatic Task Invalidation (Cron Job)
- A **cron job** runs every minute to check for tasks older than **1 hour** that are **not marked as "completed"**.
- If a task meets this condition, its status is updated to **"critical"**.

---

