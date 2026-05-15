import { useEffect } from "react";
import { ColorPicker } from "antd";
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined
} from "@ant-design/icons";
import { colors, spacing } from "../../../theme/design-system";
import type { Category } from "../../../types";

// Hooks
import { useCategoryManagement } from "../hooks/useCategoryManagement";

// Custom UI Wrappers
import Card from "../../../components/ui/antdesign/Card";
import Button from "../../../components/ui/antdesign/Button";
import Table from "../../../components/ui/antdesign/Table";
import Input from "../../../components/ui/antdesign/Input";
import Tag from "../../../components/ui/antdesign/Tag";
import Modal from "../../../components/ui/antdesign/Modal";
import Form from "../../../components/ui/antdesign/Form";
import Popconfirm from "../../../components/ui/antdesign/Popconfirm";
import Space from "../../../components/ui/antdesign/Space";

export default function CategoryManagementPage() {
    const [form] = Form.useForm();
    const {
        categories,
        loading,
        isModalVisible,
        editingCategory,
        showModal,
        hideModal,
        handleFormSubmit,
        handleDelete,
    } = useCategoryManagement();

    useEffect(() => {
        if (isModalVisible) {
            if (editingCategory) {
                form.setFieldsValue(editingCategory);
            } else {
                form.resetFields();
                form.setFieldsValue({ color: colors.primary });
            }
        }
    }, [isModalVisible, editingCategory, form]);

    const onFinish = async () => {
        const values = await form.validateFields();
        await handleFormSubmit(values);
    };

    const columns = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
            render: (text: string, record: Category) => (
                <Tag color={record.color}>{text}</Tag>
            ),
        },
        {
            title: 'Color Code',
            dataIndex: 'color',
            key: 'color',
            render: (color: string) => <code style={{ color: colors.neutral.secondary }}>{color.toUpperCase()}</code>,
        },
        {
            title: 'Action',
            key: 'action',
            align: 'right' as const,
            render: (_: any, record: Category) => (
                <Space>
                    <Button type="text" size="small" icon={<EditOutlined />} onClick={() => showModal(record)} />
                    <Popconfirm title="Delete this category?" onConfirm={() => handleDelete(record.id)} okText="Yes" cancelText="No">
                        <Button type="text" size="small" danger icon={<DeleteOutlined />} />
                    </Popconfirm>
                </Space>
            ),
        },
    ];

    return (
        <div className="fade-in">
            <Card
                title="Category Management"
                extra={
                    <Button type="primary" icon={<PlusOutlined />} onClick={() => showModal()}>
                        Add Category
                    </Button>
                }
            >
                <Table
                    dataSource={categories}
                    columns={columns}
                    rowKey="id"
                    loading={loading}
                    pagination={false}
                    size="middle"
                    scroll={{ x: 500 }}
                />
            </Card>

            <Modal
                title={editingCategory ? "Update Category" : "New Category"}
                open={isModalVisible}
                onOk={onFinish}
                onCancel={hideModal}
                okText={editingCategory ? "Save Changes" : "Create"}
                destroyOnHidden
            >
                <Form form={form} layout="vertical" style={{ marginTop: spacing.md }}>
                    <Form.Item name="name" label="Category Name" rules={[{ required: true, message: 'Please enter a name' }]}>
                        <Input placeholder="Work, Personal, etc." />
                    </Form.Item>
                    <Form.Item name="color" label="Theme Color" rules={[{ required: true }]}>
                        <ColorPicker showText />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
    );
}
