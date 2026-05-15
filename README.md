# IDX Technical Test - Todo Application (Ariya Panna)

A robust, full-stack Todo Management System built with **Go (Clean Architecture)** and **React (Ant Design)**. This application features category management, task prioritization, due dates, and a responsive UI.

## 🚀 Features
- **Todo Management**: Create, Read, Update, Delete tasks.
- **Category Management**: Organize tasks with custom categories and colors.
- **Advanced Filtering & Sorting**: Search tasks by title, filter by status, and sort by priority or due date.
- **Responsive Design**: Seamless experience across Desktop, Tablet, and Mobile.
- **Robust Backend**: Built with Clean Architecture principles for maintainability and scalability.

---

## 🛠️ Setup and Installation

### Option 1: Running with Docker (Recommended)
From the root directory, simply run:
```bash
docker compose up -d --build
```
This will start the database, backend, and frontend automatically.
- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend API: [http://localhost:8080/api/v1](http://localhost:8080/api/v1)

### Option 2: Running Locally (Without Docker)

#### Prerequisites
- **PostgreSQL 15+** installed and running.
- **Go 1.22+** installed.
- **Node.js 18+** and **npm** installed.

#### 1. Database Setup
Create a PostgreSQL database named `idx_todo_app`.

#### 2. Backend Setup
1. Go to the backend directory:
   ```bash
   cd backend
   ```
2. Create a `.env` file (copy from .env.example):
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_postgres_user
   DB_PASS=your_postgres_password
   DB_NAME=idx_todo_app
   DB_SSL_MODE=disable
   SERVER_PORT=8080
   ```
3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```
   *The migrations will run automatically upon startup.*

#### 3. Frontend Setup
1. Go to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Create a `.env` file (copy from .env.example):
   ```env
   VITE_API_BASE_URL=http://localhost:8080/api/v1
   ```
4. Run the development server:
   ```bash
   npm run dev
   ```
   Access the app at [http://localhost:5173](http://localhost:5173).


---

## 🧪 Running Tests

### Backend Unit Tests
To run the backend unit tests, ensure you have Go installed locally or run inside the container:
```bash
cd backend
go test ./internal/usecase/...
```

---

## 📖 API Documentation

### Categories
- `GET /api/v1/categories` - List all categories.
- `POST /api/v1/categories` - Create a new category.
- `PUT /api/v1/categories/:id` - Update a category.
- `DELETE /api/v1/categories/:id` - Delete a category.

### Todos
- `GET /api/v1/todos` - List todos (supports `page`, `limit`, `search`, `sort_by`, `sort_order`).
- `POST /api/v1/todos` - Create a new task.
- `PUT /api/v1/todos/:id` - Update a task.
- `DELETE /api/v1/todos/:id` - Delete a task.
- `PATCH /api/v1/todos/:id/toggle` - Toggle task completion status.

---

## 🧠 Technical Questions

### Database Design
1. **What database tables did you create and why?**
   - `categories`: Stores category metadata (name, color). This allows tasks to be grouped visually.
   - `todos`: The core table storing task details. It has a foreign key `category_id` referencing the `categories` table.
   - **Structure Choice**: I utilized a relational normalized structure by separating `categories` and `todos` into distinct tables. This avoids data redundancy (One-to-Many relationship), where a single category can be reused across multiple tasks. This ensures that category attributes like colors and names are managed centrally, making the system more maintainable and consistent.

2. **How did you handle pagination and filtering?**
   - **Filtering**: Implemented using SQL `WHERE` clauses (via GORM). For searching, I used the `ILIKE` operator (case-insensitive) on the title.
   - **Pagination**: Handled efficiently using `LIMIT` and `OFFSET` in the SQL queries.
   - **Indexes**: Added indexes on `title` (for search), `completed` (for filtering), and `category_id` (for joins) to optimize query performance as the data grows.

### Technical Decisions
1. **How did you implement responsive design?**
   - **Ant Design Components**: Used `Row`, `Col`, and `Space` with responsive props. The `Table` component uses `scroll={{ x: 800 }}` to handle small screens gracefully.
   - **Custom Breakpoints**: Primarily targeted Mobile (<576px) and Desktop (>992px) using media queries for layout stacking (e.g., search bar and buttons stacking vertically on mobile).

2. **How did you structure your React components?**
   - **Feature-based Structure**: Divided the app into `features/todo-management` and `features/category-management`.
   - **Custom Hooks**: Business logic is decoupled into hooks like `useTodoManagement.ts`, making the components clean and testable.
   - **State Management**: Used React `useState` and `useCallback` for local state. Fetching is triggered via `useEffect` and manually upon actions (like opening a modal).

3. **What backend architecture did you choose?**
   - **Clean Architecture & Layered Architecture**: Organized into:
     - `domain`: Core entities and repository interfaces.
     - `usecase`: Business logic layer.
     - `delivery/http`: HTTP handlers and routing.
     - `infrastructure`: Database & external services implementation (GORM).
   - **Error Handling**: Implemented a custom `apperror` package to standardize HTTP error responses (404 for not found, 400 for validation).

4. **How did you handle data validation?**
   - **Both sides**: 
     - **Frontend**: Used Ant Design Form validation for immediate feedback.
     - **Backend**: Implemented validation in the `entity` layer to ensure business rules are strictly enforced (e.g., title cannot be empty).

### Testing & Quality
1. **What did you choose to unit test and why?**
   - **UseCase Layer**: Tested `TodoUsecase` because it contains the core business logic (date parsing, mapping, validation checks).
   - **Edge Cases**: Specifically tested invalid date formats and empty lists to ensure robust error handling.

2. **How did you structure your tests?**
   - **Interface-based Mocking**: Following Clean Architecture principles, I used interfaces to decouple the UseCase layer from the Persistence layer. This allows us to mock the Repository and test business logic in isolation without a real database.
   - **External Package Testing**: I used the `package usecase_test` suffix for test files. This is a Go best practice that treats the test as an external consumer of the package, ensuring we only test the exported (public) API and preventing circular dependencies.
   - **Table-Driven Tests**: Used Go's idiomatic table-driven test pattern to handle multiple test scenarios (success/failure) in a single test function, making the test suite clean and easily extendable.

3. **Future Improvements**
   - **Technical Debt**: Replace the custom SQL runner with an industry-standard migration tool like `golang-migrate` for better version tracking.
   - **Refactoring**: Implement **caching using Redis** for frequently accessed todo lists to optimize performance and reduce database load during high traffic.
   - **Features**: Implement User Authentication (JWT) and a more advanced tagging system.

---
*Created for IDX Technical Test - 2026*