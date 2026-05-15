import { useEffect } from "react";
import dayjs from "dayjs";
import { colors, spacing } from "../../../theme/design-system";

import type { Todo, Category, Priority } from "../../../types";

import { useTodoManagement } from "../hooks/useTodoManagement";

import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    CheckOutlined,
    UndoOutlined
} from "@ant-design/icons";

import Card from "../../../components/ui/antdesign/Card";
import Table from "../../../components/ui/antdesign/Table";
import Button from "../../../components/ui/antdesign/Button";
import Input from "../../../components/ui/antdesign/Input";
import Tag from "../../../components/ui/antdesign/Tag";
import Select from "../../../components/ui/antdesign/Select";
import Modal from "../../../components/ui/antdesign/Modal";
import Form from "../../../components/ui/antdesign/Form";
import DatePicker from "../../../components/ui/antdesign/DatePicker";
import Space from "../../../components/ui/antdesign/Space";
import Tooltip from "../../../components/ui/antdesign/Tooltip";
import Popconfirm from "../../../components/ui/antdesign/Popconfirm";
import { Text } from "../../../components/ui/antdesign/Typography";

export default function TodoManagementPage() {
    const [form] = Form.useForm();
    const {
        todos,
        categories,
        loading,
        pagination,
        search,
        isModalVisible,
        editingTodo,
        detailTodo,
        isDetailVisible,
        handleSearch,
        handleTableChange,
        showModal,
        hideModal,
        showDetail,
        hideDetail,
        handleDelete,
        handleToggle,
        handleFormSubmit,
    } = useTodoManagement();

    useEffect(() => {
        if (isModalVisible) {
            if (editingTodo) {
                form.setFieldsValue({
                    ...editingTodo,
                    category_id: editingTodo.category.id,
                    due_date: editingTodo.due_date ? dayjs(editingTodo.due_date) : null,
                });
            } else {
                form.resetFields();
            }
        }
    }, [isModalVisible, editingTodo, form]);

    const onFinish = async () => {
        const values = await form.validateFields();
        await handleFormSubmit(values);
    };

    const columns = [
        {
            title: 'Status',
            dataIndex: 'completed',
            key: 'status',
            width: 100,
            sorter: true,
            render: (completed: boolean) => (
                completed ? <Tag color="success">Done</Tag> : <Tag color="warning">Pending</Tag>
            ),
        },
        {
            title: 'Title',
            dataIndex: 'title',
            key: 'title',
            sorter: true,
            render: (text: string, record: Todo) => (
                <Text delete={record.completed} strong style={{ color: record.completed ? colors.neutral.muted : colors.neutral.title }}>
                    {text}
                </Text>
            ),
        },
        {
            title: 'Category',
            dataIndex: 'category',
            key: 'category',
            render: (cat: Category | null) => cat ? (
                <Tag color={cat.color}>{cat.name}</Tag>
            ) : '-',
        },

        {
            title: 'Priority',
            dataIndex: 'priority',
            key: 'priority',
            sorter: true,
            render: (p: Priority) => {
                const color = p === 'high' ? 'red' : p === 'medium' ? 'orange' : 'blue';
                return <Tag color={color}>{p.toUpperCase()}</Tag>;
            },
        },
        {
            title: 'Due Date',
            dataIndex: 'due_date',
            key: 'due_date',
            sorter: true,
            render: (d: string) => d ? (
                <Text style={{ fontSize: '13px', color: colors.neutral.secondary }}>
                    {dayjs(d).format("DD MMM YYYY")}
                </Text>
            ) : '-',
        },

        {
            title: 'Action',
            key: 'action',
            align: 'right' as const,
            render: (_: any, record: Todo) => (
                <Space size="small" onClick={(e) => e.stopPropagation()}>
                    <Tooltip title={record.completed ? "Mark as Incomplete" : "Mark as Complete"}>
                        <Button
                            type="text"
                            size="small"
                            icon={record.completed ? <UndoOutlined /> : <CheckOutlined />}
                            onClick={() => handleToggle(record.id)}
                            style={{ color: record.completed ? colors.warning : colors.success }}
                        />
                    </Tooltip>
                    <Button type="text" size="small" icon={<EditOutlined />} onClick={() => showModal(record)} />
                    <Popconfirm title="Delete this task?" onConfirm={() => handleDelete(record.id)} okText="Yes" cancelText="No">
                        <Button type="text" size="small" danger icon={<DeleteOutlined />} />
                    </Popconfirm>
                </Space>
            ),
        },
    ];

    return (
        <div className="fade-in">
            <Card title="What you need to do?">
                <Space orientation="vertical" size="middle" style={{ width: '100%' }}>
                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', gap: spacing.md }} className="mobile-stack">
                        <div style={{ display: 'flex', gap: spacing.md }} className="mobile-stack mobile-full-width">
                            <Input
                                placeholder="Search tasks..."
                                allowClear
                                value={search}
                                onChange={handleSearch}
                                className="mobile-full-width"
                                style={{ width: 300 }}
                            />
                        </div>
                        <Button type="primary" icon={<PlusOutlined />} onClick={() => showModal()}>
                            Create Task
                        </Button>
                    </div>


                    <Table
                        dataSource={todos}
                        columns={columns}
                        rowKey="id"
                        loading={loading}
                        scroll={{ x: 800 }}
                        onRow={(record) => ({
                            onClick: () => showDetail(record),
                            style: { cursor: 'pointer' }
                        })}
                        pagination={{
                            ...pagination,
                            showSizeChanger: true,
                            pageSizeOptions: ['5', '10', '20', '50'],
                            placement: ['bottomCenter'] as any,
                            showTotal: (total) => `Total ${total} items`,
                        }}
                        onChange={handleTableChange}
                    />
                </Space>
            </Card>

            <Modal
                title="Task Details"
                open={isDetailVisible}
                onCancel={hideDetail}
                footer={[
                    <Button key="close" onClick={hideDetail}>Close</Button>,
                    <Button key="edit" type="primary" onClick={() => {
                        hideDetail();
                        if (detailTodo) showModal(detailTodo);
                    }}>Edit Task</Button>
                ]}
            >
                {detailTodo && (
                    <Space orientation="vertical" size="middle" style={{ width: '100%', paddingTop: spacing.md }}>
                        <div>
                            <Text type="secondary" style={{ display: 'block', marginBottom: '4px' }}>Title</Text>
                            <div style={{ fontSize: '18px', fontWeight: 600 }}>{detailTodo.title}</div>
                        </div>
                        <div>
                            <Text type="secondary" style={{ display: 'block', marginBottom: '4px' }}>Description</Text>
                            <div style={{ whiteSpace: 'pre-wrap' }}>{detailTodo.description || 'No description provided'}</div>
                        </div>
                        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: spacing.md }}>
                            <div>
                                <Text type="secondary" style={{ display: 'block', marginBottom: '4px' }}>Category</Text>
                                <div>
                                    {detailTodo.category ? (
                                        <Tag color={detailTodo.category.color}>{detailTodo.category.name}</Tag>
                                    ) : (
                                        <Text type="secondary">None</Text>
                                    )}
                                </div>
                            </div>

                            <div>
                                <Text type="secondary" style={{ display: 'block', marginBottom: '4px' }}>Priority</Text>
                                <div><Tag color={detailTodo.priority === 'high' ? 'red' : detailTodo.priority === 'medium' ? 'orange' : 'blue'}>{detailTodo.priority.toUpperCase()}</Tag></div>
                            </div>
                        </div>
                    </Space>
                )}
            </Modal>

            <Modal
                title={editingTodo ? "Update Task" : "Create Task"}
                open={isModalVisible}
                onOk={onFinish}
                onCancel={hideModal}
                okText={editingTodo ? "Save Changes" : "Create"}
                destroyOnHidden
            >
                <Form form={form} layout="vertical" initialValues={{ priority: 'medium' }} style={{ marginTop: spacing.md }}>
                    <Form.Item name="title" label="Task Title" rules={[{ required: true, message: 'Please enter a title' }]}>
                        <Input placeholder="What needs to be done?" />
                    </Form.Item>
                    <Form.Item name="description" label="Description">
                        <Input.TextArea placeholder="Add more details..." rows={3} />
                    </Form.Item>
                    <Form.Item name="category_id" label="Category" rules={[{ required: true }]}>
                        <Select placeholder="Select a category">
                            {categories.map(c => <Select.Option key={c.id} value={c.id}>{c.name}</Select.Option>)}
                        </Select>
                    </Form.Item>
                    <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: spacing.md }}>
                        <Form.Item name="priority" label="Priority" rules={[{ required: true }]}>
                            <Select>
                                <Select.Option value="low">Low</Select.Option>
                                <Select.Option value="medium">Medium</Select.Option>
                                <Select.Option value="high">High</Select.Option>
                            </Select>
                        </Form.Item>
                        <Form.Item name="due_date" label="Due Date">
                            <DatePicker style={{ width: '100%' }} />
                        </Form.Item>
                    </div>
                </Form>
            </Modal>
        </div>
    );
}