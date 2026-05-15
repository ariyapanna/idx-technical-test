export interface Category {
  id: number;
  name: string;
  color: string;
  created_at: string;
}

export type Priority = 'low' | 'medium' | 'high';

export interface Todo {
  id: number;
  category: Category;
  title: string;
  description: string;
  priority: Priority;
  due_date: string | null;
  completed: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateTodoRequest {
  category_id: number;
  title: string;
  description?: string;
  priority: string;
  due_date?: string;
}

export interface UpdateTodoRequest {
  category_id: number;
  title: string;
  description?: string;
  priority: string;
  due_date?: string;
  completed: boolean;
}

export interface CreateCategoryRequest {
  name: string;
  color: string;
}

export interface UpdateCategoryRequest {
  name: string;
  color: string;
}

export interface PaginationInfo {
  current_page: number;
  per_page: number;
  total: number;
  total_pages: number;
}

export interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: T;
}

export interface ApiPaginationResponse<T> extends ApiResponse<T> {
  pagination: PaginationInfo;
}

