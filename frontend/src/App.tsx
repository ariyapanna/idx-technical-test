import {
  AppstoreOutlined,
  DiffOutlined,
} from '@ant-design/icons';

import CategoryManagementPage from './features/category-management/pages/CategoryManagementPage';

import Header from './components/layout/Header';
import Tabs from './components/ui/antdesign/Tabs';
import TodoManagementPage from './features/todo-management/pages/TodoManagementPage';

function App() {
  const items = [
    {
      key: '1',
      label: 'Todo Management',
      icon: <DiffOutlined />,
      children: <TodoManagementPage />
    },
    {
      key: '2',
      label: 'Category Management',
      icon: <AppstoreOutlined />,
      children: <CategoryManagementPage />,
    },
  ];

  return (
    <>
      <Header />
      <div style={{ padding: '16px', maxWidth: '1200px', margin: '0 auto' }}>
        <Tabs items={items} defaultActiveKey="1" />
      </div>
    </>
  );
}

export default App;
