import type {
  ApiResponse,
  ApiPaginationResponse,
  Todo,
  Category,
  CreateTodoRequest,
  UpdateTodoRequest,
  CreateCategoryRequest,
  UpdateCategoryRequest
} from '../types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';


async function handleResponse<T>(response: Response): Promise<T> {
  const isJson = response.headers.get('content-type')?.includes('application/json');
  const data = isJson ? await response.json() : null;

  if (!response.ok) {
    const error = (data && data.message) || response.statusText;
    throw new Error(error);
  }

  return data as T;
}

export const todoService = {
  async list(params: { page?: number; limit?: number; search?: string; sort_by?: string; sort_order?: string }) {
    const cleanParams = Object.fromEntries(
      Object.entries(params).filter(([_, v]) => v != null && v !== '')
    );
    const query = new URLSearchParams(cleanParams as any).toString();
    const response = await fetch(`${API_BASE_URL}/todos?${query}`);
    return handleResponse<ApiPaginationResponse<Todo[]>>(response);
  },

  async create(data: CreateTodoRequest) {
    const response = await fetch(`${API_BASE_URL}/todos`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return handleResponse<ApiResponse<Todo>>(response);
  },

  async update(id: number, data: UpdateTodoRequest) {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return handleResponse<ApiResponse<Todo>>(response);
  },

  async delete(id: number) {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
      method: 'DELETE',
    });
    return handleResponse<ApiResponse<null>>(response);
  },

  async toggleComplete(id: number) {
    const response = await fetch(`${API_BASE_URL}/todos/${id}/complete`, {
      method: 'PATCH',
    });
    return handleResponse<ApiResponse<Todo>>(response);
  },
};

export const categoryService = {
  async list() {
    const response = await fetch(`${API_BASE_URL}/categories`);
    return handleResponse<ApiResponse<Category[]>>(response);
  },

  async create(data: CreateCategoryRequest) {
    const response = await fetch(`${API_BASE_URL}/categories`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return handleResponse<ApiResponse<Category>>(response);
  },

  async update(id: number, data: UpdateCategoryRequest) {
    const response = await fetch(`${API_BASE_URL}/categories/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return handleResponse<ApiResponse<Category>>(response);
  },

  async delete(id: number) {
    const response = await fetch(`${API_BASE_URL}/categories/${id}`, {
      method: 'DELETE',
    });
    return handleResponse<ApiResponse<null>>(response);
  },
};

