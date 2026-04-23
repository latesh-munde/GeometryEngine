# 📐 GeomEngine

A modular, concurrency-aware geometry processing engine written in Go.  
Designed with **clean separation of concerns**, **scalability**, and **testability**, this project demonstrates how to structure a production-grade system around computational geometry and concurrent execution.

---

# 🧠 Core Design Philosophy

The system is built on strict architectural boundaries:

- Geometry is **pure and deterministic**
- Execution is **stateless and isolated**
- Concurrency is **controlled and bounded**
- Components are **loosely coupled and composable**

---

# 📁 Project Structure

```
geomengine/
│
├── batch/
│ ├── processor.go → Coordinates batch execution
│ ├── worker.go → Worker logic for task execution
│ ├── task.go → Task definition for batch processing
│ ├── result.go → Result structure for processed tasks
│ ├── processor_test.go → Tests for batch processor
│
├── cmd/
│ └── runner/
│ └── main.go → Application entry point
│
├── dto/
│ └── task_json.go → JSON input structure for tasks
│
├── engine/
│ ├── engineErr.go → Engine-level error definitions
│ ├── executor.go → Executes operations on a single shape
│ ├── executor_test.go → Tests for execution logic
│ ├── operations.go → Supported geometric operations
│
├── errors/
│ └── geometryErr.go → Geometry-specific validation errors
│
├── geometry/
│ ├── shape.go → Shape interface definition
│ ├── circle.go → Circle implementation
│ ├── rectangle.go → Rectangle implementation
│ ├── polygon.go → Polygon implementation
│ ├── line.go → Line implementation
│ ├── point.go → Point implementation
│ ├── bbox.go → Bounding box logic
│ ├── \*\_test.go → Unit tests for shapes
│
├── internal/
│ ├── builder/
│ │ └── shape_builder.go → Converts DTO → geometry objects
│ │
│ ├── generator/
│ │ ├── task_generator.go → Creates structured tasks
│ │ └── random_task.go → Generates random workloads
│ │
│ ├── runner/
│ │ ├── sequential.go → Sequential execution strategy
│ │ ├── concurrent.go → Concurrent execution strategy
│ │ ├── streaming.go → Streaming execution model
│ │ └── helpers.go → Shared runner utilities
│ │
│ ├── summary/
│ │ └── summary_writer.go → Aggregates and prints results
│
├── results/ → Output storage / aggregation
│
├── go.mod
└── README.md
```

---

# 📦 Module Breakdown

## 1. geometry — Geometry Domain

Defines geometric primitives and their behavior.

### Responsibilities

- Encapsulates geometry concepts cleanly
- Each shape is self-contained
- Implements a common `Shape` interface

### Rules

- One file per shape
- No concurrency
- No batch logic
- No logging or printing

### Each Shape

- Stores its own data
- Implements geometric operations (area, bounds, etc.)

---

## 2. errors — Geometry Validation & Errors

Handles domain-specific validation failures.

### Purpose

Ensures invalid geometry fails clearly and predictably.

### Examples

- Invalid radius
- Degenerate polygon
- Zero-length line

---

## 3. engine — Single-Shape Execution Layer

Executes operations on a single shape.

### Responsibilities

- Accepts one shape
- Executes one operation
- Returns result or error

### Constraints

- No batch awareness
- No loops over collections
- No goroutines

### Benefit

- Reusable and composable core logic
- Easy to test and extend

---

## 4. tests — Geometry & Engine Testing

Ensures correctness and robustness.

### Focus Areas

- Area calculations
- Bounding boxes
- Edge cases and invalid inputs

---

## 5. batch — Concurrent Processing Layer

Implements worker pool for processing multiple tasks.

### Responsibilities

- Task distribution
- Worker lifecycle management
- Result collection

### Rules

- No geometry logic
- Calls engine only
- No shared mutable state

### Guarantees

- All tasks are processed
- No lost results
- Proper error propagation

---

## 6. benchmarks — Performance Evaluation

Measures system efficiency.

### Includes

- Single-thread execution
- Batch execution
- Worker scaling performance

---

## 🔧 internal — Supporting Infrastructure

Not exposed publicly; used for orchestration.

### builder/

- Converts DTO input into geometry objects

### generator/

- Generates tasks (random or structured workloads)

### runner/

Execution strategies:

- Sequential
- Concurrent
- Streaming

### summary/

- Aggregates and formats final results

---

## 🧾 dto — Data Transfer Objects

Defines structured input/output formats (e.g., JSON tasks).

---

## 📊 results — Output Handling

Stores or organizes processed results.

---

## 🚀 cmd/runner — Entry Point

Wires all components together:

- Reads input
- Builds shapes
- Executes tasks
- Outputs results

---

# ⚙️ Execution Flow

Input → DTO → Builder → Engine → Batch Runner → Results → Summary

---

# 🧪 Testing Strategy

- Unit tests per shape
- Engine-level validation
- Edge case coverage
- Deterministic outputs

---

# 📈 Performance Strategy

- Bounded concurrency via worker pools
- Channel-based communication
- No shared state → avoids race conditions
- Scales with worker count

---

# 🎯 Key Highlights

- Clean architecture (domain → engine → concurrency)
- Strict separation of concerns
- Fully testable components
- Production-ready concurrency model
- Easily extensible for new shapes and operations

---

# 🧭 Future Improvements

- Add more geometric operations
- Extend to 3D geometry
- Add persistent storage
- Expose as REST/gRPC service
- Add visualization layer

---

# 🏁 Getting Started

### 1. Clone Repository

```
git clone https://github.com/latesh-munde/GeometryEngine.git
cd geom-engine
```

### 2. Run Server

```bash
go mod tidy
go run cmd/runner/main.go
```

---
