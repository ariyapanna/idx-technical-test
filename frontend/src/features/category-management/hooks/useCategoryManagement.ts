import { useState, useCallback, useEffect } from 'react';
import { App } from 'antd';
import { categoryService } from '../../../services/api';
import type { Category } from '../../../types';

export const useCategoryManagement = () => {
    const { message } = App.useApp();
    const [categories, setCategories] = useState<Category[]>([]);
    const [loading, setLoading] = useState(false);
    
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [editingCategory, setEditingCategory] = useState<Category | null>(null);

    const fetchCategories = useCallback(async () => {
        setLoading(true);
        try {
            const response = await categoryService.list();
            if (response.success) setCategories(response.data);
        } catch (error: any) {
            message.error(error.message || "Failed to fetch categories");
        } finally {
            setLoading(false);
        }
    }, [message]);

    useEffect(() => {
        fetchCategories();
    }, [fetchCategories]);

    const showModal = (category?: Category) => {
        setEditingCategory(category || null);
        setIsModalVisible(true);
    };

    const hideModal = () => {
        setIsModalVisible(false);
        setEditingCategory(null);
    };

    const handleFormSubmit = async (values: any) => {
        try {
            const color = typeof values.color === 'string' ? values.color : values.color.toHexString();
            const data = { ...values, color };

            const response = editingCategory
                ? await categoryService.update(editingCategory.id, data)
                : await categoryService.create(data);

            if (response.success) {
                message.success("Category saved successfully");
                hideModal();
                fetchCategories();
                return true;
            }
        } catch (error: any) {
            message.error(error.message || "Operation failed");
        }
        return false;
    };

    const handleDelete = async (id: number) => {
        try {
            const res = await categoryService.delete(id);
            if (res.success) {
                message.success("Category deleted");
                fetchCategories();
            }
        } catch (error: any) {
            message.error(error.message || "Delete failed");
        }
    };

    return {
        categories,
        loading,
        isModalVisible,
        editingCategory,
        showModal,
        hideModal,
        handleFormSubmit,
        handleDelete,
        fetchCategories,
    };
};
