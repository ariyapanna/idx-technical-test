import { useState, useCallback, useEffect, useRef } from 'react';
import { App } from 'antd';
import { todoService, categoryService } from '../../../services/api';
import type { Todo, Category } from '../../../types';

export const useTodoManagement = () => {
    const { message } = App.useApp();
    const [todos, setTodos] = useState<Todo[]>([]);
    const [categories, setCategories] = useState<Category[]>([]);
    const [loading, setLoading] = useState(false);
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        total: 0,
    });
    const [search, setSearch] = useState("");
    const [sort, setSort] = useState<{ field?: string, order?: 'asc' | 'desc' }>({
        field: 'created_at',
        order: 'desc'
    });

    const [isModalVisible, setIsModalVisible] = useState(false);
    const [editingTodo, setEditingTodo] = useState<Todo | null>(null);
    const [detailTodo, setDetailTodo] = useState<Todo | null>(null);
    const [isDetailVisible, setIsDetailVisible] = useState(false);

    const searchTimeout = useRef<ReturnType<typeof setTimeout> | null>(null);

    const fetchTodos = useCallback(async (page = 1, pageSize = 10, searchText = "", sortField?: string, sortOrder?: 'asc' | 'desc') => {
        setLoading(true);
        try {
            const response = await todoService.list({
                page,
                limit: pageSize,
                search: searchText,
                sort_by: sortField,
                sort_order: sortOrder
            });
            if (response.success && response.data) {
                setTodos(response.data);
                if (response.pagination) {
                    setPagination({
                        current: response.pagination.current_page,
                        pageSize: response.pagination.per_page,
                        total: response.pagination.total,
                    });
                }
            }
        } catch (error: any) {
            message.error(error.message || "Failed to fetch todos");
        } finally {
            setLoading(false);
        }
    }, [message]);

    const fetchCategories = useCallback(async () => {
        try {
            const response = await categoryService.list();
            if (response.success) setCategories(response.data);
        } catch (error: any) {
            message.error(error.message || "Failed to fetch categories");
        }
    }, [message]);

    useEffect(() => {
        fetchTodos(pagination.current, pagination.pageSize, search, sort.field, sort.order);
        fetchCategories();
    }, [fetchTodos, fetchCategories]);

    const handleSearch = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const value = e.target.value;
        setSearch(value);

        if (searchTimeout.current) clearTimeout(searchTimeout.current);

        searchTimeout.current = setTimeout(() => {
            fetchTodos(1, pagination.pageSize, value, sort.field, sort.order);
        }, 500);
    };

    const handleTableChange = (p: any, _filters: any, sorter: any) => {
        const newSort = sorter.field ? {
            field: sorter.field,
            order: sorter.order === 'descend' ? 'desc' : 'asc'
        } as const : { field: 'created_at', order: 'desc' as const };

        setSort(newSort);
        fetchTodos(p.current, p.pageSize, search, newSort.field, newSort.order);
    };

    const handleSortChange = (value: 'asc' | 'desc') => {
        const newSort = { field: 'created_at', order: value };
        setSort(newSort);
        fetchTodos(1, pagination.pageSize, search, newSort.field, newSort.order);
    };

    const showModal = (todo?: Todo) => {
        fetchCategories();
        setEditingTodo(todo || null);
        setIsModalVisible(true);
    };

    const hideModal = () => {
        setIsModalVisible(false);
        setEditingTodo(null);
    };

    const showDetail = (todo: Todo) => {
        setDetailTodo(todo);
        setIsDetailVisible(true);
    };

    const hideDetail = () => {
        setIsDetailVisible(false);
    };

    const handleDelete = async (id: number) => {
        try {
            const res = await todoService.delete(id);
            if (res.success) {
                message.success("Task deleted");
                fetchTodos(pagination.current, pagination.pageSize, search, sort.field, sort.order);
            }
        } catch (error: any) {
            message.error(error.message || "Delete failed");
        }
    };

    const handleToggle = async (id: number) => {
        try {
            const res = await todoService.toggleComplete(id);
            if (res.success) {
                message.success("Status updated");
                fetchTodos(pagination.current, pagination.pageSize, search, sort.field, sort.order);
            }
        } catch (error: any) {
            message.error(error.message || "Update failed");
        }
    };

    const handleFormSubmit = async (values: any) => {
        try {
            const data = {
                ...values,
                due_date: values.due_date ? values.due_date.format("YYYY-MM-DD") : undefined,
            };

            const response = editingTodo
                ? await todoService.update(editingTodo.id, { ...data, completed: editingTodo.completed })
                : await todoService.create(data);

            if (response.success) {
                message.success(editingTodo ? "Updated successfully" : "Created successfully");
                hideModal();
                fetchTodos(pagination.current, pagination.pageSize, search, sort.field, sort.order);
                return true;
            }
        } catch (error: any) {
            message.error(error.message || "Operation failed");
        }
        return false;
    };

    return {
        todos,
        categories,
        loading,
        pagination,
        search,
        sort,
        isModalVisible,
        editingTodo,
        detailTodo,
        isDetailVisible,
        handleSearch,
        handleTableChange,
        handleSortChange,
        showModal,
        hideModal,
        showDetail,
        hideDetail,
        handleDelete,
        handleToggle,
        handleFormSubmit,
    };
};
